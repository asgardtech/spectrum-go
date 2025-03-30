#!/usr/bin/env python3

import json
import os
from jinja2 import Environment, FileSystemLoader
import re

def camel_to_kebab(s):
    """Convert CamelCase to kebab-case."""
    return re.sub(r'(?<!^)(?=[A-Z])', '-', s).lower()

def kebab_to_camel(s):
    """Convert kebab-case to camelCase."""
    components = s.split('-')
    return components[0] + ''.join(x.title() for x in components[1:])

def get_go_type(attr_type):
    """Convert JSON type to Go type."""
    type_mapping = {
        "string": "string",
        "number": "float64",
        "integer": "int",
        "boolean": "bool",
        "object": "interface{}",
        "array": "[]interface{}"
    }
    return type_mapping.get(attr_type, "string")

def process_component(component):
    """Process component data for template rendering."""
    # Enhance the component data with additional information for templates
    processed = component.copy()
    
    # Process attributes for Go types
    for attr in processed.get('attributes', []):
        attr['go_name'] = attr['name'].title().replace('-', '')
        attr['prop_name'] = f"Prop{attr['go_name']}"
        
        if attr.get('type') == 'enum' and attr.get('values'):
            attr['go_type'] = f"{processed['name']}{attr['go_name']}"
        else:
            attr['go_type'] = get_go_type(attr.get('type', 'string'))
            
        # Zero value for the type
        if attr.get('type') == 'string':
            attr['zero_value'] = '""'
        elif attr.get('type') == 'boolean':
            attr['zero_value'] = 'false'
        elif attr.get('type') in ('number', 'integer'):
            attr['zero_value'] = '0'
        else:
            attr['zero_value'] = 'nil'
    
    # Process slots
    for slot in processed.get('slots', []):
        slot['go_name'] = slot['name'].title().replace('-', '')
        slot['prop_name'] = f"Prop{slot['go_name']}Slot"
    
    # Process events
    for event in processed.get('events', []):
        event['go_name'] = event['name'].title().replace('-', '')
        event['prop_name'] = f"PropOn{event['go_name']}"
    
    # Set struct name
    processed['struct_name'] = f"spectrum{processed['name']}"
    
    return processed

def main():
    input_dir = "spectrum-docs/json"
    output_dir = "generated"
    template_dir = "templates"
    
    # Create output directory if it doesn't exist
    os.makedirs(output_dir, exist_ok=True)
    
    # Set up Jinja2 environment
    env = Environment(
        loader=FileSystemLoader(template_dir),
        trim_blocks=True,
        lstrip_blocks=True
    )
    
    # Load the template
    template = env.get_template("component.go.j2")
    
    # Process each JSON file
    for filename in os.listdir(input_dir):
        if not filename.endswith(".json"):
            continue
            
        input_path = os.path.join(input_dir, filename)
        component_name = os.path.splitext(filename)[0]
        output_path = os.path.join(output_dir, f"{component_name}.go")
        
        with open(input_path, 'r') as f:
            component = json.load(f)
        
        # Process component data
        processed_component = process_component(component)
        
        # Render template
        go_code = template.render(component=processed_component)
        
        # Write to output file
        with open(output_path, 'w') as f:
            f.write(go_code)
        
        print(f"Generated {output_path}")

if __name__ == "__main__":
    main() 