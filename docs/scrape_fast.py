#!/usr/bin/env python3
"""Fast concurrent API doc scraper for GitCode/AtomGit."""

import os
import re
import json
import asyncio
from pathlib import Path
from playwright.async_api import async_playwright

BASE_URL = "https://docs.atomgit.com/docs/apis"
API_BASE = "https://api.atomgit.com/api/v5"
DOCS_DIR = Path(__file__).parent / "docs"

CATEGORIES = {
    "repos": "repos",
    "branch": "branch",
    "issue": "issue",
    "pull_request": "pull_request",
    "commit": "commit",
    "tag": "tag",
    "label": "label",
    "milestone": "milestone",
    "user": "user",
    "org": "org",
    "hook": "hook",
    "release": "release",
    "enterprise": "enterprise",
    "dashboard": "dashboard",
    "ai": "ai",
    "search": "search",
}

def get_category(url: str) -> str:
    url_lower = url.lower()
    if "kanban" in url_lower: return "dashboard"
    if any(x in url_lower for x in ["similarity", "audio-transcriptions", "detect-yolo", "video-generate", "video-status", "audio-classification"]): return "ai"
    if "search-issues" in url_lower or "search-repositories" in url_lower: return "search"
    if "hooks" in url_lower: return "hook"
    if "releases" in url_lower or "release" in url_lower: return "release"
    if "enterprises" in url_lower or "enterprise-enterprise" in url_lower: return "enterprise"
    if "orgs-org" in url_lower or "org-org" in url_lower or "org-owner" in url_lower: return "org"
    if "milestones" in url_lower: return "milestone"
    if "labels" in url_lower: return "label"
    if "tags" in url_lower or "protected-tags" in url_lower: return "tag"
    if "commits" in url_lower or "comments" in url_lower: return "commit"
    if "pulls" in url_lower or "merge-requests" in url_lower: return "pull_request"
    if "issues" in url_lower: return "issue"
    if "users" in url_lower or "user-" in url_lower or "emails" in url_lower: return "user"
    if "repos" in url_lower or "fork" in url_lower: return "repos"
    if "branch" in url_lower: return "branch"
    return "repos"

def extract_info(content: str, url: str) -> dict:
    info = {
        "title": "",
        "method": "GET",
        "endpoint": "",
        "path_params": [],
        "query_params": [("access_token", "string", True, "用户授权码")],
        "body_params": [],
        "request_example": "",
        "response_schema": "",
    }

    # Extract title
    title_match = re.search(r'heading[^>]*level=1[^>]*ref=e\d+[^>]*>\s*<generic[^>]*>([^<]+)', content)
    if not title_match:
        title_match = re.search(r'<h1[^>]*>([^<]+)</h1>', content)
    if title_match:
        info["title"] = title_match.group(1).strip()

    # Extract method
    method_match = re.search(r'<(?:span|generic)[^>]*>(GET|POST|PUT|PATCH|DELETE)</', content, re.I)
    if method_match:
        info["method"] = method_match.group(1).upper()

    # Extract endpoint
    endpoint_match = re.search(r'https://api\.atomgit\.com(/api/v\d+/[^\s<>"\']+)', content)
    if endpoint_match:
        info["endpoint"] = endpoint_match.group(1)

    # Extract parameters from Schema section
    param_pattern = re.compile(r'<strong[^>]*>([^<]+)</strong>\s*<generic[^>]*>([^<]+)</generic>\s*<generic[^>]*>(required|optional)</generic>', re.DOTALL)
    for match in param_pattern.finditer(content):
        param_name = match.group(1).strip()
        param_type = match.group(2).strip()
        param_required = match.group(3).strip() == "required"
        if param_name == "access_token":
            continue
        if ":" in param_name or "/" in param_name:
            info["path_params"].append((param_name, param_type, param_required, ""))
        elif param_name:
            info["query_params"].append((param_name, param_type, param_required, ""))

    # Extract body params - look for JSON object properties
    body_pattern = re.compile(r'<code[^>]*>\s*"?(\w+)"?\s*</code>\s*<generic[^>]*>([^<]+)</generic>\s*<generic[^>]*>(required|optional)</generic>', re.DOTALL)
    for match in body_pattern.finditer(content):
        param_name = match.group(1).strip()
        param_type = match.group(2).strip()
        param_required = match.group(3).strip() == "required"
        if param_name and param_name not in [p[0] for p in info["path_params"] + info["query_params"] + info["body_params"]]:
            info["body_params"].append((param_name, param_type, param_required, ""))

    # Build request example
    endpoint = info["endpoint"]
    if info["method"] in ["POST", "PUT", "PATCH"]:
        info["request_example"] = f'''curl -X {info["method"]} "https://api.atomgit.com{endpoint}" \\
  -H "Authorization: token $ATOMGIT_TOKEN" \\
  -H "Content-Type: application/json" \\
  -d '{{"param": "value"}}'
'''
    else:
        params = "?".join([f"{p[0]}=$ATOMGIT_TOKEN" for p in info["query_params"]]) if info["query_params"] else "access_token=$ATOMGIT_TOKEN"
        info["request_example"] = f'curl "https://api.atomgit.com{endpoint}?{params}" \\\n  -H "Authorization: token $ATOMGIT_TOKEN"'

    return info

def create_markdown(api_id: str, info: dict, url: str) -> str:
    method = info["method"]
    endpoint = info["endpoint"]
    title = info["title"] or api_id.replace("-", " ").title()

    md = f"# {title}\n\n## 基本信息\n\n| 项目 | 值 |\n|------|-----|\n| **HTTP Method** | {method} |\n| **Endpoint** | `https://api.atomgit.com{endpoint}` |\n| **文档链接** | {url} |\n\n"

    if info["path_params"]:
        md += "## Path Parameters\n\n| 参数名 | 类型 | 必填 | 描述 |\n|--------|------|------|------|\n"
        for p in info["path_params"]:
            md += f"| {p[0]} | {p[1]} | {'是' if p[2] else '否'} | {p[3]} |\n"

    if info["query_params"]:
        md += "\n## Query Parameters\n\n| 参数名 | 类型 | 必填 | 描述 |\n|--------|------|------|------|\n"
        for p in info["query_params"]:
            md += f"| {p[0]} | {p[1]} | {'是' if p[2] else '否'} | {p[3]} |\n"

    if info["body_params"]:
        md += "\n## Body Parameters\n\n| 参数名 | 类型 | 必填 | 描述 |\n|--------|------|------|------|\n"
        for p in info["body_params"]:
            md += f"| {p[0]} | {p[1]} | {'是' if p[2] else '否'} | {p[3]} |\n"

    md += f"\n## 请求示例\n\n```bash\n{info['request_example']}\n```\n\n## 相关 CLI 命令\n\n参见 [API列表](/docs/apis_list.md)\n"

    return md

async def scrape_page(page, url: str, semaphore: asyncio.Semaphore) -> dict:
    async with semaphore:
        try:
            await page.goto(url, wait_until="domcontentloaded", timeout=15000)
            await page.wait_for_timeout(500)
            content = await page.content()
            api_path = url.replace(f"{BASE_URL}/", "")
            return {"url": url, "content": content, "api_path": api_path, "success": True}
        except Exception as e:
            return {"url": url, "error": str(e), "success": False}

async def main():
    # Create dirs
    for cat in CATEGORIES.values():
        (DOCS_DIR / cat).mkdir(parents=True, exist_ok=True)

    # Read API list and extract URLs
    api_list_path = DOCS_DIR / "apis_list.md"
    with open(api_list_path) as f:
        content = f.read()

    urls = re.findall(r'https://docs\.atomgit\.com/docs/apis/[a-zA-Z0-9_-]+', content)
    urls = list(dict.fromkeys(urls))  # dedupe preserve order
    print(f"Found {len(urls)} APIs")

    semaphore = asyncio.Semaphore(5)  # 5 concurrent
    results = []

    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=True)
        page = await browser.new_page()

        tasks = [scrape_page(page, url, semaphore) for url in urls]
        for i, coro in enumerate(asyncio.as_completed(tasks), 1):
            result = await coro
            status = "✓" if result.get("success") else "✗"
            print(f"[{i}/{len(urls)}] {status} {result.get('url', '?')}")
            if result.get("success"):
                results.append(result)

        await browser.close()

    # Process results
    created = 0
    for result in results:
        content = result["content"]
        url = result["url"]
        api_path = result["api_path"]

        category = get_category(url)
        info = extract_info(content, url)

        filename = api_path + ".md"
        filepath = DOCS_DIR / category / filename

        md = create_markdown(api_path, info, url)

        with open(filepath, "w") as f:
            f.write(md)
        created += 1

    print(f"\n✓ Created {created} API docs in go/docs/")

if __name__ == "__main__":
    asyncio.run(main())