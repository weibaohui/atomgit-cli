const { chromium } = require('playwright');
const fs = require('fs');
const path = require('path');
const https = require('https');
const http = require('http');

// Configuration
const BASE_URL = 'https://docs.atomgit.com/docs/apis';
const API_LIST_PATH = path.join(__dirname, 'apis_list.md');
const DOCS_DIR = __dirname;

// Category mapping
const getCategory = (url) => {
  const urlLower = url.toLowerCase();
  if (urlLower.includes('kanban')) return 'dashboard';
  if (urlLower.includes('similarity') || urlLower.includes('audio-transcriptions') ||
      urlLower.includes('detect-yolo') || urlLower.includes('video-generate') ||
      urlLower.includes('video-status') || urlLower.includes('audio-classification')) return 'ai';
  if (urlLower.includes('search-issues') || urlLower.includes('search-repositories')) return 'search';
  if (urlLower.includes('hooks')) return 'hook';
  if (urlLower.includes('releases') || urlLower.includes('release')) return 'release';
  if (urlLower.includes('enterprises') || urlLower.includes('enterprise-enterprise')) return 'enterprise';
  if (urlLower.includes('orgs-org') || urlLower.includes('org-org') || urlLower.includes('org-owner')) return 'org';
  if (urlLower.includes('milestones')) return 'milestone';
  if (urlLower.includes('labels')) return 'label';
  if (urlLower.includes('tags') || urlLower.includes('protected-tags')) return 'tag';
  if (urlLower.includes('commits') || urlLower.includes('comments') || urlLower.includes('commit-statistics')) return 'commit';
  if (urlLower.includes('pulls') || urlLower.includes('merge-requests')) return 'pull_request';
  if (urlLower.includes('issues')) return 'issue';
  if (urlLower.includes('users') || urlLower.includes('user-') || urlLower.includes('emails')) return 'user';
  if (urlLower.includes('repos') || urlLower.includes('fork')) return 'repos';
  if (urlLower.includes('branch')) return 'branch';
  return 'repos';
};

// Get filename from URL
const getFilename = (url) => {
  return url.replace('https://docs.atomgit.com/docs/apis/', '');
};

// Extract URLs from apis_list.md
const extractUrls = () => {
  const content = fs.readFileSync(API_LIST_PATH, 'utf-8');
  const urlRegex = /https:\/\/docs\.atomgit\.com\/docs\/apis\/[a-zA-Z0-9_-]+/g;
  const urls = content.match(urlRegex) || [];
  return [...new Set(urls)];
};

// Scrape a page
const scrapePage = async (page, url) => {
  console.log(`Scraping: ${url}`);

  try {
    await page.goto(url, { waitUntil: 'networkidle', timeout: 30000 });
    await page.waitForTimeout(1000);

    // Get title
    const title = await page.title();

    // Get h1
    const h1 = await page.$eval('h1', el => el.textContent).catch(() => '');

    // Get method and endpoint from h2
    const h2Content = await page.$eval('h2', el => el.textContent).catch(() => '');
    const methodMatch = h2Content.match(/GET|POST|PUT|PATCH|DELETE/);
    const method = methodMatch ? methodMatch[0] : 'GET';
    const endpoint = h2Content.replace(/GET|POST|PUT|PATCH|DELETE/, '').trim();

    // Extract parameters by examining the DOM structure
    const pathParams = [];
    const queryParams = [];
    const bodyParams = [];

    // Get all strong elements (parameter names)
    const strongElements = await page.$$('strong');
    for (const strong of strongElements) {
      const text = await strong.textContent();
      // Check if this looks like a parameter name (not a heading)
      if (text && !['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'Schema', 'Example'].includes(text)) {
        // Get parent/sibling context
        const parent = await strong.evaluateHandle(el => el.parentElement);
        const parentText = await parent.textContent();

        if (parentText.includes('Path') || parentText.includes('path')) {
          const desc = parentText.replace(text, '').replace(/string|number|required|必填/gi, '').trim();
          pathParams.push({ name: text, type: 'string', required: true, description: desc });
        } else if (parentText.includes('Query') || parentText.includes('query')) {
          const desc = parentText.replace(text, '').replace(/string|number|required|必填/gi, '').trim();
          queryParams.push({ name: text, type: 'string', required: true, description: desc });
        }
      }
    }

    // Get response example
    const responseExample = await page.evaluate(() => {
      // Try to find code blocks with JSON
      const codeBlocks = document.querySelectorAll('code');
      for (const code of codeBlocks) {
        const text = code.textContent;
        if (text && text.includes('{') && text.includes('}')) {
          return text;
        }
      }
      return '';
    });

    // Get curl example
    const curlExample = await page.evaluate(() => {
      const tabs = document.querySelectorAll('[role="tab"]');
      for (const tab of tabs) {
        if (tab.textContent.toLowerCase().includes('curl')) {
          const tabPanel = tab.closest('[role="tabpanel"]');
          if (tabPanel) {
            const code = tabPanel.querySelector('code');
            if (code) return code.textContent;
          }
        }
      }
      return '';
    });

    return {
      url,
      title: h1,
      method,
      endpoint,
      pathParams,
      queryParams,
      bodyParams,
      responseExample,
      curlExample,
      rawTitle: title
    };
  } catch (error) {
    console.error(`Error scraping ${url}: ${error.message}`);
    return { url, error: error.message };
  }
};

// Main function
const main = async () => {
  // Extract URLs
  const urls = extractUrls();
  console.log(`Found ${urls.length} unique API URLs`);

  // Create category directories
  const categories = ['repos', 'branch', 'issue', 'pull_request', 'commit', 'tag', 'label',
                      'milestone', 'user', 'org', 'hook', 'release', 'enterprise', 'dashboard', 'ai', 'search'];
  for (const cat of categories) {
    const catDir = path.join(DOCS_DIR, cat);
    if (!fs.existsSync(catDir)) {
      fs.mkdirSync(catDir, { recursive: true });
    }
  }

  // Launch browser
  const browser = await chromium.launch({ headless: true });
  const page = await browser.newPage();

  // Scrape each URL
  const results = [];
  for (let i = 0; i < urls.length; i++) {
    const url = urls[i];
    console.log(`[${i+1}/${urls.length}] ${url}`);
    const result = await scrapePage(page, url);
    results.push(result);

    // Rate limiting
    await new Promise(r => setTimeout(r, 500));
  }

  await browser.close();

  // Save results
  fs.writeFileSync(path.join(DOCS_DIR, 'scrape_results.json'), JSON.stringify(results, null, 2));
  console.log(`\nScraping complete. ${results.length} APIs processed.`);
  console.log(`Results saved to ${path.join(DOCS_DIR, 'scrape_results.json')}`);
};

main().catch(console.error);
