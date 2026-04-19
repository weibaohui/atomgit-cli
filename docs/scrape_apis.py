#!/usr/bin/env python3
"""
Scrape GitCode API documentation and organize by category.
"""

import os
import re
import json
import asyncio
from pathlib import Path
from typing import Optional
from playwright.async_api import async_playwright

BASE_URL = "https://docs.atomgit.com/docs/apis"
API_LIST_PATH = Path(__file__).parent.parent / "docs" / "apis_list.md"
DOCS_DIR = Path(__file__).parent.parent / "docs"

CATEGORIES = {
    "repos": ["repos-owner-repo", "repos-owner", "org-org-repo"],
    "branch": ["repos-owner-repo-branches", "repos-owner-repo-protect-branches"],
    "issue": ["repos-owner-repo-issues", "repos-owner-issues", "search-issues", "enterprises-enterprise-issues"],
    "pull_request": ["repos-owner-repo-pulls", "users-merge-requests"],
    "commit": ["repos-owner-repo-commits", "repos-owner-repo-comments", "owner-repo-repository-commit"],
    "tag": ["repos-owner-repo-tags", "repos-owner-repo-protected-tags"],
    "label": ["repos-owner-repo-labels", "enterprises-enterprise-labels"],
    "milestone": ["repos-owner-repo-milestones", "enterprise-enterprise-id-milestones"],
    "user": ["users-", "user-", "emails"],
    "org": ["orgs-", "org-org", "org-owner", "enterprises-enterprise-members"],
    "hook": ["repos-owner-repo-hooks"],
    "release": ["repos-owner-repo-releases", "repos-owner-repo-releases-attach"],
    "enterprise": ["enterprises-", "enterprise-enterprise"],
    "dashboard": ["kanban", "org-owner-kanban"],
    "ai": ["similarity", "audio-transcriptions", "detect-yolo", "video-generate", "video-status", "audio-classification"],
    "search": ["search-repositories"],
}

def get_category_from_url(url: str) -> str:
    """Determine category from API URL."""
    url_lower = url.lower()

    # Special cases
    if "kanban" in url_lower:
        return "dashboard"
    if "similarity" in url_lower or "audio" in url_lower or "detect-yolo" in url_lower or "video" in url_lower:
        return "ai"
    if "search-issues" in url_lower or "search-repositories" in url_lower:
        return "search"
    if "hooks" in url_lower:
        return "hook"
    if "releases" in url_lower or "release" in url_lower:
        return "release"
    if "enterprises" in url_lower or "enterprise-enterprise" in url_lower:
        return "enterprise"
    if "orgs-org" in url_lower or "org-org" in url_lower or "org-owner" in url_lower:
        return "org"
    if "milestones" in url_lower:
        return "milestone"
    if "labels" in url_lower:
        return "label"
    if "tags" in url_lower or "protected-tags" in url_lower:
        return "tag"
    if "commits" in url_lower or "comments" in url_lower or "commit-statistics" in url_lower:
        return "commit"
    if "pulls" in url_lower or "merge-requests" in url_lower:
        return "pull_request"
    if "issues" in url_lower:
        return "issue"
    if "users" in url_lower or "user-" in url_lower or "emails" in url_lower:
        return "user"
    if "repos" in url_lower or "fork" in url_lower:
        return "repos"
    if "branch" in url_lower:
        return "branch"

    return "repos"

def extract_api_info(page_content: dict) -> dict:
    """Extract API information from page content."""
    info = {
        "title": "",
        "method": "",
        "endpoint": "",
        "path_params": [],
        "query_params": [],
        "body_params": [],
        "request_example": "",
        "response_schema": "",
        "response_example": "",
    }

    # This is a simplified extractor - in reality we'd need to parse the DOM more carefully
    return info

async def scrape_api_page(page, url: str) -> dict:
    """Scrape a single API page."""
    try:
        await page.goto(url, wait_until="networkidle", timeout=30000)
        await page.wait_for_timeout(1000)  # Give time for dynamic content

        # Get page title
        title = await page.title()

        # Get main content
        content = await page.content()

        # Extract using regex patterns
        method_match = re.search(r'<(?:span|div|generic)[^>]*>(GET|POST|PUT|PATCH|DELETE)</', content, re.IGNORECASE)
        method = method_match.group(1) if method_match else "GET"

        endpoint_match = re.search(r'https://api\.atomgit\.com/api/[v\d]+/[^\s<>"\']+', content)
        endpoint = endpoint_match.group(0) if endpoint_match else ""

        # Try to get JSON content
        try:
            json_content = await page.evaluate("""() => {
                const schemaEl = document.querySelector('.tabpanel[data-tabpanel]');
                if (schemaEl) {
                    const codeEl = schemaEl.querySelector('code');
                    if (codeEl) return codeEl.textContent;
                }
                // Try to find response schema
                const respEl = document.querySelector('[data-testid="response-schema"]');
                if (respEl) return respEl.textContent;
                return '';
            }""")
        except:
            json_content = ""

        return {
            "url": url,
            "title": title,
            "method": method,
            "endpoint": endpoint,
            "raw_content": content[:5000],  # First 5000 chars for debugging
            "json_content": json_content,
        }
    except Exception as e:
        return {"url": url, "error": str(e)}

async def main():
    """Main scraping function."""
    # Read API list
    with open(API_LIST_PATH, "r") as f:
        api_list_content = f.read()

    # Extract URLs
    url_pattern = re.compile(r'https://docs\.atomgit\.com/docs/apis/[a-zA-Z0-9_-]+')
    urls = url_pattern.findall(api_list_content)
    unique_urls = list(set(urls))

    print(f"Found {len(unique_urls)} unique API URLs")

    # Create category directories
    for category in CATEGORIES.keys():
        cat_dir = DOCS_DIR / category
        cat_dir.mkdir(exist_ok=True)

    # Also create search category if needed
    (DOCS_DIR / "search").mkdir(exist_ok=True)

    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=True)
        page = await browser.new_page()

        results = []
        for i, url in enumerate(unique_urls[:]):  # Process all
            print(f"[{i+1}/{len(unique_urls)}] Scraping: {url}")
            result = await scrape_api_page(page, url)
            results.append(result)

            # Rate limiting
            await asyncio.sleep(0.5)

        await browser.close()

    # Save raw results
    with open(DOCS_DIR / "scrape_results.json", "w") as f:
        json.dump(results, f, indent=2)

    print(f"\nScraping complete. Results saved to {DOCS_DIR / 'scrape_results.json'}")
    print(f"Total APIs scraped: {len(results)}")

if __name__ == "__main__":
    asyncio.run(main())
