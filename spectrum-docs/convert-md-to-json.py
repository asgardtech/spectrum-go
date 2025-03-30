#!/usr/bin/env python3
"""
Script to convert Spectrum component markdown files to JSON schema files.
This parses the markdown files in spectrum-docs/md/ and creates corresponding JSON files
in spectrum-docs/json/ following the component schema structure.
"""

import os
import re
import json
import glob
from pathlib import Path

# Directory paths
MD_DIR = Path("md")
JSON_DIR = Path("json")
SCHEMA_FILE = Path("component-schema.json")

# Helper functions to extract information from markdown
def extract_tag_name(content):
    """Extract the tag name from markdown content."""
    match = re.search(r'# (sp-\w[\w-]*)', content)
    if match:
        return match.group(1)
    return None

def extract_description(content):
    """Extract the component description."""
    # Look for description after the Overview heading
    match = re.search(r'## Overview\s+([^\n]+)', content)
    if match:
        return match.group(1)
    return ""

def extract_attributes(content):
    """Extract component attributes from API table."""
    attributes = []
    
    # Find the attributes table
    attr_table_match = re.search(r'### Attributes and Properties.*?<table>.*?<tbody>(.*?)</tbody>.*?</table>', 
                               content, re.DOTALL)
    
    if not attr_table_match:
        return attributes
    
    table_content = attr_table_match.group(1)
    
    # Parse each row in the table
    rows = re.findall(r'<tr>(.*?)</tr>', table_content, re.DOTALL)
    for row in rows:
        cols = re.findall(r'<td>(.*?)</td>', row, re.DOTALL)
        if len(cols) >= 5:
            # Clean up property name, removing backticks
            prop_name = cols[1].replace('`', '').strip()
            
            # Try to determine the type
            type_text = cols[2].replace('`', '').strip()
            
            # Default value (may be empty)
            default_value = cols[3].replace('`', '').strip()
            if default_value == "``":
                default_value = ""
            
            # Description
            description = cols[4].strip()
            
            # Determine attribute type based on the type text
            attr_type = "string"  # Default
            enum_values = []
            
            if "boolean" in type_text.lower():
                attr_type = "boolean"
                if default_value.lower() == "false":
                    default_value = False
                elif default_value.lower() == "true":
                    default_value = True
                else:
                    default_value = None
            elif "number" in type_text.lower():
                attr_type = "number"
                if default_value:
                    try:
                        default_value = float(default_value)
                        if default_value.is_integer():
                            default_value = int(default_value)
                    except ValueError:
                        default_value = None
            elif "|" in type_text:
                # Likely an enum type
                attr_type = "enum"
                enum_values = [v.strip().replace("'", "").replace('"', '') 
                              for v in re.findall(r"'([^']*)'|\"([^\"]*)\"", type_text)]
                if not enum_values:
                    enum_values = [v.strip() for v in type_text.split("|") if v.strip() and "undefined" not in v]
            
            attribute = {
                "name": prop_name,
                "type": attr_type,
                "description": description
            }
            
            if enum_values:
                attribute["values"] = enum_values
            
            if default_value not in ("", None):
                attribute["default"] = default_value
                
            attributes.append(attribute)
    
    return attributes

def extract_slots(content):
    """Extract component slots."""
    slots = []
    
    # Find the slots table
    slots_table_match = re.search(r'### Slots.*?<table>.*?<tbody>(.*?)</tbody>.*?</table>', 
                                content, re.DOTALL)
    
    if not slots_table_match:
        return slots
    
    table_content = slots_table_match.group(1)
    
    # Parse each row in the table
    rows = re.findall(r'<tr>(.*?)</tr>', table_content, re.DOTALL)
    for row in rows:
        cols = re.findall(r'<td>(.*?)</td>', row, re.DOTALL)
        if len(cols) >= 2:
            name = cols[0].replace('`', '').strip()
            description = cols[1].strip()
            
            # Handle default slot
            if "default slot" in name.lower():
                name = "default"
            
            slots.append({
                "name": name,
                "description": description
            })
    
    return slots

def extract_events(content):
    """Extract component events."""
    # This would need to be expanded for more thorough event extraction
    # For now, just create a click event as a default for most components
    if "button" in content.lower() or "click" in content.lower():
        return [{
            "name": "click",
            "description": "Fired when the component is clicked",
            "bubbles": True,
            "cancelable": False
        }]
    return []

def extract_accessibility(content):
    """Extract accessibility information."""
    accessibility = {
        "role": "",
        "ariaAttributes": [],
        "keyboardInteractions": []
    }
    
    # Try to find accessibility section
    accessibility_section = re.search(r'### Accessibility(.*?)(?:^##|\Z)', 
                                    content, re.DOTALL | re.MULTILINE)
    
    if accessibility_section:
        section_content = accessibility_section.group(1)
        
        # Look for ARIA role
        role_match = re.search(r'role=["\']([\w-]+)["\']', section_content)
        if role_match:
            accessibility["role"] = role_match.group(1)
        
        # Look for keyboard interactions
        key_patterns = [
            (r'(?:press|key|keyboard).*?(Enter|Space|Arrow\w+|Tab|Escape|Esc)', "Activates or navigates the component"),
            (r'(Enter|Space|Arrow\w+|Tab|Escape|Esc).*?(?:press|key|keyboard)', "Activates or navigates the component")
        ]
        
        for pattern, default_desc in key_patterns:
            matches = re.finditer(pattern, section_content, re.IGNORECASE)
            for match in matches:
                key = match.group(1)
                if not any(ki["key"] == key for ki in accessibility["keyboardInteractions"]):
                    accessibility["keyboardInteractions"].append({
                        "key": key,
                        "description": default_desc
                    })
    
    # Default aria attributes based on likely properties
    if "label" in content.lower():
        accessibility["ariaAttributes"].append({
            "name": "aria-label",
            "description": "Provides an accessible name for the component"
        })
    
    if "disabled" in content.lower():
        accessibility["ariaAttributes"].append({
            "name": "aria-disabled",
            "description": "Indicates that the component is disabled"
        })
    
    return accessibility

def extract_examples(content):
    """Extract usage examples."""
    examples = []
    
    # Look for code examples
    code_blocks = re.finditer(r'(#+.*?)\s+```(html)?\s+(.*?)```', content, re.DOTALL)
    
    for i, match in enumerate(code_blocks):
        title = match.group(1).strip('#').strip() or f"Example {i+1}"
        code = match.group(3).strip()
        
        examples.append({
            "title": title,
            "description": "",
            "code": code
        })
    
    # If no code blocks with backticks, try looking for indented examples
    if not examples:
        indented_blocks = re.finditer(r'(?:^|\n)(\w+.*?)\n\s{4,}(.*?)(?=\n\S|$)', content, re.DOTALL | re.MULTILINE)
        for i, match in enumerate(indented_blocks):
            title = match.group(1).strip() or f"Example {i+1}"
            code = '\n'.join(line.strip() for line in match.group(2).strip().split('\n'))
            
            if "<sp-" in code:  # Only add if it seems to have component tags
                examples.append({
                    "title": title,
                    "description": "",
                    "code": code
                })
    
    return examples

def infer_component_name(file_name):
    """Convert filename to component name in PascalCase."""
    # Remove file extension
    base_name = os.path.splitext(os.path.basename(file_name))[0]
    
    # Handle special cases for component names
    name_parts = base_name.split('-')
    pascal_case = ''.join(part.capitalize() for part in name_parts)
    
    # Special handling for some components
    if pascal_case == "HelpTextMixin":
        pascal_case = "HelpTextMixin"
    elif pascal_case == "IconsUi":
        pascal_case = "IconsUI"
    
    return pascal_case

def infer_category(file_name, content):
    """Infer the component category based on file name or content."""
    base_name = os.path.splitext(os.path.basename(file_name))[0]
    
    # Define category mappings
    category_mappings = {
        "button": "Action",
        "action": "Action",
        "link": "Action",
        "tag": "Content",
        "badge": "Content",
        "icon": "Content",
        "field": "Form",
        "input": "Form",
        "checkbox": "Form",
        "radio": "Form",
        "switch": "Form",
        "textfield": "Form",
        "textarea": "Form",
        "number-field": "Form",
        "color": "Color",
        "slider": "Form",
        "dialog": "Overlay",
        "popover": "Overlay",
        "overlay": "Overlay",
        "toast": "Overlay",
        "tooltip": "Overlay",
        "menu": "Navigation",
        "breadcrumb": "Navigation",
        "sidenav": "Navigation",
        "tab": "Navigation",
        "progress": "Feedback",
        "alert": "Feedback",
        "status": "Feedback",
        "meter": "Feedback"
    }
    
    # Check for category matches
    for key, category in category_mappings.items():
        if key in base_name:
            return category
    
    # Look for clues in content
    if "form" in content.lower() or "input" in content.lower():
        return "Form"
    elif "action" in content.lower() or "click" in content.lower():
        return "Action"
    elif "navigation" in content.lower() or "nav" in content.lower():
        return "Navigation"
    elif "overlay" in content.lower() or "modal" in content.lower():
        return "Overlay"
    
    # Default category
    return "UI Component"

def infer_dependencies(content, component_name):
    """Infer component dependencies based on content."""
    dependencies = []
    
    # Look for references to other components
    potential_deps = re.findall(r'<sp-([\w-]+)', content)
    for dep in potential_deps:
        # Convert kebab-case to PascalCase
        dep_parts = dep.split('-')
        dep_name = ''.join(part.capitalize() for part in dep_parts)
        
        # Don't add self as dependency
        if dep_name != component_name and dep_name not in dependencies:
            dependencies.append(dep_name)
    
    return dependencies

def process_markdown_file(file_path):
    """Process a markdown file and convert it to component schema JSON."""
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            content = f.read()
        
        file_name = os.path.basename(file_path)
        component_name = infer_component_name(file_name)
        tag_name = extract_tag_name(content) or f"sp-{file_name.replace('.md', '')}"
        
        component_json = {
            "name": component_name,
            "tagName": tag_name,
            "description": extract_description(content),
            "category": infer_category(file_name, content),
            "attributes": extract_attributes(content),
            "slots": extract_slots(content),
            "events": extract_events(content),
            "accessibility": extract_accessibility(content),
            "examples": extract_examples(content),
            "implementation": {
                "goStructName": f"spectrum{component_name}",
                "constructorName": component_name,
                "dependsOn": infer_dependencies(content, component_name)
            }
        }
        
        # Determine output path
        output_file = os.path.join(JSON_DIR, file_name.replace('.md', '.json'))
        
        # Write the JSON file
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(component_json, f, indent=2)
            
        return True, f"Processed {file_name} -> {os.path.basename(output_file)}"
    
    except Exception as e:
        return False, f"Error processing {file_name}: {str(e)}"

def main():
    """Main function to process all markdown files."""
    # Ensure output directory exists
    os.makedirs(JSON_DIR, exist_ok=True)
    
    # Get all markdown files
    md_files = glob.glob(os.path.join(MD_DIR, "*.md"))
    
    results = []
    for file_path in md_files:
        success, message = process_markdown_file(file_path)
        results.append((success, message))
    
    # Print summary
    print(f"\nProcessed {len(md_files)} files:")
    success_count = sum(1 for r in results if r[0])
    print(f"Success: {success_count}")
    print(f"Failed: {len(results) - success_count}")
    
    if len(results) - success_count > 0:
        print("\nErrors:")
        for success, message in results:
            if not success:
                print(f"- {message}")

if __name__ == "__main__":
    main()