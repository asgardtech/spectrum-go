import os
import requests
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import time

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
    components_output_dir = "html/components"
    tools_output_dir = "html/tools"
    ensure_directory(components_output_dir)
    ensure_directory(tools_output_dir)

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
        output_path = os.path.join(components_output_dir, f"{component_name}.html")
        download_page(api_url, output_path)
    
    # Find and download tools API pages
    print("\nProcessing tools links...")
    tools_links = soup.find_all('a', href=lambda x: x and '/tools/' in x)
    
    # If no tools links found directly, try to find a tools section
    if not tools_links:
        print("No direct tools links found, trying to find tools section...")
        tools_section = soup.find('section', id='tools')
        if tools_section:
            tools_links = tools_section.find_all('a')
    
    # If still no links, manually add known tools URLs
    if not tools_links:
        print("Adding known tools links manually...")
        known_tools = [
            {"href": "/spectrum-web-components/tools/shared/", "text": "Shared"},
            {"href": "/spectrum-web-components/tools/reactive-controllers/", "text": "Reactive Controllers"},
            {"href": "/spectrum-web-components/tools/base/", "text": "Base"},
            {"href": "/spectrum-web-components/tools/localization/", "text": "Localization"}
        ]
        
        for tool in known_tools:
            tool_url = urljoin(base_url, tool["href"])
            if not tool_url.endswith('/'):
                tool_url += '/'
            
            # Extract tool name
            tool_name = tool["text"].lower().replace(' ', '-')
            
            # Construct API URL
            api_url = f"{tool_url}api/"
            
            # Download API page
            output_path = os.path.join(tools_output_dir, f"{tool_name}.html")
            print(f"Trying known tool: {api_url}")
            if download_page(api_url, output_path):
                print(f"Successfully downloaded {tool_name}")
    else:
        # Process found tools links
        for link in tools_links:
            tool_url = urljoin(base_url, link['href'])
            if not tool_url.endswith('/'):
                tool_url += '/'
            
            # Extract tool name from URL or link text
            if '/tools/' in tool_url:
                tool_name = tool_url.split('/tools/')[-1].rstrip('/')
                if '/' in tool_name:
                    tool_name = tool_name.split('/')[0]  # Take only the first part
            else:
                # Use link text if URL structure is different
                tool_name = link.get_text().strip().lower().replace(' ', '-')
            
            # Construct API URL
            api_url = f"{tool_url}api/"
            
            # Download API page
            output_path = os.path.join(tools_output_dir, f"{tool_name}.html")
            if download_page(api_url, output_path):
                print(f"Successfully downloaded {tool_name}")
    
    # Direct download of specific known tools
    print("\nDirectly downloading specific known tools...")
    specific_tools = {
        "shared": "https://opensource.adobe.com/spectrum-web-components/tools/shared/api/",
        "reactive-controllers": "https://opensource.adobe.com/spectrum-web-components/tools/reactive-controllers/api/",
        "base": "https://opensource.adobe.com/spectrum-web-components/tools/base/api/",
        "localization": "https://opensource.adobe.com/spectrum-web-components/tools/localization/api/"
    }
    
    for tool_name, api_url in specific_tools.items():
        output_path = os.path.join(tools_output_dir, f"{tool_name}.html")
        if download_page(api_url, output_path):
            print(f"Successfully downloaded {tool_name} (direct)")

if __name__ == "__main__":
    main()
