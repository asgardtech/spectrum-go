import os
import requests
import re
from bs4 import BeautifulSoup
from urllib.parse import urljoin

def ensure_directory(directory):
    if not os.path.exists(directory):
        os.makedirs(directory)

def download_page(url, output_path):
    try:
        response = requests.get(url)
        response.raise_for_status()
        with open(output_path, 'w', encoding='utf-8') as f:
            f.write(response.text)
        print(f"Downloaded: {output_path}")
        return True
    except Exception as e:
        print(f"Error downloading {url}: {str(e)}")
        return False

def main():
    base_url = "https://opensource.adobe.com/spectrum-web-components/index.html"
    output_dir = "html"
    ensure_directory(output_dir)

    # Download the main index page
    print("Downloading main index page...")
    response = requests.get(base_url)
    response.raise_for_status()
    soup = BeautifulSoup(response.text, 'html.parser')

    # Components that don't have API pages
    skip_components = {
        'trigger-directive',
        'imperative-api',
        'slottable-request',
        'directive',
        'icons',
        'iconset'
    }
    
    # Find and download component API pages
    print("Processing component links...")
    component_links = soup.find_all('a', href=lambda x: x and '/components/' in x)
    for link in component_links:
        component_url = urljoin(base_url, link['href'])
        if not component_url.endswith('/'):
            component_url += '/'
        
        # Extract component name from URL
        component_name = component_url.split('/components/')[-1].rstrip('/')
        
        # Skip components that don't have API pages
        if component_name in skip_components:
            print(f"Skipping {component_name} - no API available")
            continue
        
        # Construct API URL
        api_url = f"{component_url}api/"
        
        # Download API page
        output_path = os.path.join(output_dir, f"{component_name}.html")
        download_page(api_url, output_path)
    
    # Parse the HTML to find all tool links
    print("\nFinding and processing all tool links...")
    
    # Find all links in the page
    all_links = soup.find_all('a', href=True)
    
    # Keep track of processed tools to avoid duplicates
    processed_tools = set()
    
    # Match links with /tools/ in the URL
    tool_pattern = re.compile(r'/spectrum-web-components/tools/([^/]+)')
    
    for link in all_links:
        href = link['href']
        match = tool_pattern.search(href)
        
        if match:
            tool_name = match.group(1)
            
            # Skip if we've already processed this tool
            if tool_name in processed_tools:
                continue
            
            # Mark as processed
            processed_tools.add(tool_name)
            
            # Construct full tool URL
            tool_url = urljoin(base_url, href)
            
            # Ensure URL ends with a slash
            if not tool_url.endswith('/'):
                tool_url += '/'
            
            # Remove any API suffix if present
            if tool_url.endswith('/api/'):
                tool_url = tool_url[:-4]
            
            print(f"Found tool: {tool_name} at {tool_url}")
            
            # Download the tool page
            output_path = os.path.join(output_dir, f"{tool_name}.html")
            if download_page(tool_url, output_path):
                print(f"Successfully downloaded {tool_name}")

if __name__ == "__main__":
    main()
