import os
import re
from bs4 import BeautifulSoup
import html2text

def ensure_directory(directory):
    if not os.path.exists(directory):
        os.makedirs(directory)

def process_html_file(input_path, output_path):
    with open(input_path, 'r', encoding='utf-8') as f:
        html = f.read()
    
    # Parse the HTML
    soup = BeautifulSoup(html, 'html.parser')
    
    # Remove header with logo and title
    logo = soup.find('a', id='logo')
    if logo:
        logo.decompose()
    
    # Remove sidenav links
    sidenav = soup.find('nav')
    if sidenav:
        sidenav.decompose()
    
    # Remove badges from usage section
    usage_section = soup.find('h3', id='usage')
    if usage_section:
        badges = usage_section.find_next('p', class_='spectrum-Body')
        if badges and badges.find('img'):
            badges.decompose()
    
    # Find all tabs
    tabs_elements = soup.find_all('sp-tab')
    for tab in tabs_elements:
        if 'value' in tab.attrs and tab['value'] == 'changelog':
            tab.decompose()
    
    # Find and remove the changelog panel
    changelog_panel = soup.find('sp-tab-panel', {'value': 'changelog'})
    if changelog_panel:
        changelog_panel.decompose()
    else:
        print(f"No changelog panel found in {input_path}")
    
    # Extract tables and save their HTML version before converting to markdown
    tables = {}
    table_count = 0
    
    # Find API section to handle tables specially
    api_section = soup.find('sp-tab-panel', {'value': 'api'})
    
    for table in soup.find_all('sp-table'):
        # Convert sp-table to standard HTML table
        html_table = '<table>\n'
        
        # Find the heading above the table
        table_heading = None
        prev_element = table.previous_element
        while prev_element and not table_heading:
            if prev_element.name in ['h2', 'h3', 'h4'] and prev_element.get_text().strip():
                table_heading = prev_element.get_text().strip()
            prev_element = prev_element.previous_element
        
        # Add hard-coded headers based on section
        if table_heading and "Attributes and Properties" in table_heading:
            # Property table - add headers manually
            html_table += '  <thead>\n'
            html_table += '    <tr>\n'
            html_table += '      <th>Property</th>\n'
            html_table += '      <th>Attribute</th>\n'
            html_table += '      <th>Type</th>\n'
            html_table += '      <th>Default</th>\n'
            html_table += '      <th>Description</th>\n'
            html_table += '    </tr>\n'
            html_table += '  </thead>\n'
        elif table_heading and "Slots" in table_heading:
            # Slots table - add headers manually
            html_table += '  <thead>\n'
            html_table += '    <tr>\n'
            html_table += '      <th>Name</th>\n'
            html_table += '      <th>Description</th>\n'
            html_table += '    </tr>\n'
            html_table += '  </thead>\n'
        else:
            # Process table head normally
            thead = table.find('sp-table-head')
            if thead:
                html_table += '  <thead>\n'
                
                # Process table head rows
                head_rows = thead.find_all('sp-table-row')
                for row in head_rows:
                    html_table += '    <tr>\n'
                    
                    # Process head cells
                    cells = row.find_all('sp-table-head-cell')
                    for cell in cells:
                        cell_text = cell.get_text().strip()
                        html_table += f'      <th>{cell_text}</th>\n'
                    
                    html_table += '    </tr>\n'
                
                html_table += '  </thead>\n'
        
        # Process table body
        tbody = table.find('sp-table-body')
        if tbody:
            html_table += '  <tbody>\n'
            
            # Process body rows
            body_rows = tbody.find_all('sp-table-row')
            for row in body_rows:
                html_table += '    <tr>\n'
                
                # Process body cells
                cells = row.find_all('sp-table-cell')
                for cell in cells:
                    # Get cell content, preserving code elements
                    cell_content = ''
                    for element in cell.contents:
                        if hasattr(element, 'name') and element.name == 'code':
                            cell_content += f'<code>{element.get_text().strip()}</code>'
                        elif isinstance(element, str):
                            cell_content += element.strip()
                    
                    html_table += f'      <td>{cell_content}</td>\n'
                
                html_table += '    </tr>\n'
            
            html_table += '  </tbody>\n'
        
        html_table += '</table>'
        
        # Create a unique placeholder
        placeholder = f"TABLE_PLACEHOLDER_{table_count}"
        table_count += 1
        
        # Store the HTML table
        tables[placeholder] = html_table
        
        # Replace the table with a placeholder
        placeholder_tag = soup.new_tag('div')
        placeholder_tag.string = placeholder
        table.replace_with(placeholder_tag)
    
    # Get main content after processing
    main = soup.find('main')
    if not main:
        print(f"No main content found in {input_path}")
        return
    
    # Convert to markdown
    h = html2text.HTML2Text()
    h.ignore_links = False
    h.ignore_images = True  # Ignore images to remove badges
    h.body_width = 0  # Don't wrap lines
    h.ignore_emphasis = False
    h.ignore_tables = False
    h.single_line_break = True  # Preserve line breaks
    h.strong_mark = '**'  # Use ** for bold
    h.em_mark = '*'  # Use * for italic
    h.ul_item_mark = '- '  # Use - for unordered lists
    h.emphasis_mark = '*'  # Use * for emphasis
    
    # Convert the content to markdown
    markdown = h.handle(str(main))
    
    # Clean up the markdown
    markdown = re.sub(r'\n{3,}', '\n\n', markdown)  # Remove excessive newlines
    markdown = re.sub(r'#\s+', '# ', markdown)  # Fix heading spacing
    
    # Replace placeholders with HTML tables
    for placeholder, html_table in tables.items():
        markdown = markdown.replace(placeholder, html_table)
    
    # Fix code formatting - wrap all code elements in backticks
    markdown = re.sub(r'<code>(.*?)</code>', r'`\1`', markdown)
    
    # Remove "Section titled" text
    markdown = re.sub(r'#Section titled.*?\n', '', markdown)
    
    # Remove "Overview API Changelog" navigation text
    markdown = re.sub(r'Overview\s+API\s+Changelog', '', markdown)
    
    # Clean up empty sections
    markdown = re.sub(r'\n\s*\n\s*\n', '\n\n', markdown)
    
    # Write the markdown file
    with open(output_path, 'w', encoding='utf-8') as f:
        f.write(markdown)
    
    print(f"Processed: {output_path}")

def main():
    input_dir = "html"
    output_dir = "md"
    ensure_directory(output_dir)
    
    for filename in os.listdir(input_dir):
        if filename.endswith('.html'):
            input_path = os.path.join(input_dir, filename)
            output_path = os.path.join(output_dir, filename.replace('.html', '.md'))
            process_html_file(input_path, output_path)

if __name__ == "__main__":
    main() 