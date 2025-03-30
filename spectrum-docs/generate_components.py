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
    
    # Special handling for prepositions and short words
    special_words = {
        'in': 'In',
        'at': 'At',
        'by': 'By',
        'of': 'Of',
        'to': 'To',
        'on': 'On',
        'for': 'For',
        'is': 'Is',
        'has': 'Has',
        'with': 'With',
        'from': 'From',
        'into': 'Into',
        'over': 'Over',
        'under': 'Under',
        'after': 'After',
        'before': 'Before',
        'between': 'Between',
        'through': 'Through',
        'during': 'During',
        'within': 'Within',
        'without': 'Without',
        'per': 'Per',
        'via': 'Via'
    }
    
    # First word is lowercase, rest are title case
    result = components[0]
    for i in range(1, len(components)):
        word = components[i]
        # Check if it's a special word that needs consistent casing
        if word.lower() in special_words:
            result += special_words[word.lower()]
        else:
            result += word.title()
    
    return result

def get_go_type(attr_type, item_type=None):
    """Convert JSON type to Go type."""
    type_mapping = {
        "string": "string",
        "number": "float64",
        "integer": "int",
        "boolean": "bool",
        "object": "any",
        "array": "[]any"
    }
    
    # Handle arrays with specific item types
    if attr_type == "array" and item_type:
        if item_type == "string":
            return "[]string"
        elif item_type == "number":
            return "[]float64"
        elif item_type == "integer":
            return "[]int"
        elif item_type == "boolean":
            return "[]bool"
    
    return type_mapping.get(attr_type, "string")

def process_component(component):
    """Process component data for template rendering."""
    # Enhance the component data with additional information for templates
    processed = component.copy()
    
    # Track slot/attribute names to detect duplicates
    method_names = set()
    
    # Go reserved keywords that need special parameter names
    reserved_keywords = {
        'type', 'for', 'range', 'map', 'select', 'go', 'interface', 'switch', 
        'case', 'default', 'func', 'defer', 'package', 'import', 'break', 'chan', 
        'const', 'continue', 'else', 'fallthrough', 'goto', 'if', 'return', 
        'struct', 'var'
    }
    
    # Special handling for words like "in", "on" in property names
    special_words = {
        'in': 'In',
        'at': 'At',
        'by': 'By',
        'of': 'Of',
        'to': 'To',
        'on': 'On',
        'for': 'For',
        'is': 'Is',
        'has': 'Has',
        'with': 'With',
        'from': 'From',
        'into': 'Into',
        'over': 'Over',
        'under': 'Under',
        'after': 'After',
        'before': 'Before',
        'between': 'Between',
        'through': 'Through',
        'during': 'During',
        'within': 'Within',
        'without': 'Without',
        'per': 'Per',
        'via': 'Via'
    }
    
    # Process attributes for Go types
    for attr in processed.get('attributes', []):
        # Convert kebab-case to camelCase for the Go name with special handling for words like "in"
        parts = attr['name'].split('-')
        special_parts = []
        
        # First part (capitalize if it's a special word and not the first part)
        if parts[0].lower() in special_words:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
        else:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
            
        # Rest of the parts
        for part in parts[1:]:
            if part.lower() in special_words:
                special_parts.append(special_words[part.lower()])
            else:
                special_parts.append(part.title())
        
        attr_go_name = ''.join(special_parts)
        attr['go_name'] = attr_go_name[0].upper() + attr_go_name[1:]  # Capitalize first letter for method name
        attr['prop_name'] = f"Prop{attr['go_name']}"
        
        # Add parameter_name for function parameters (without hyphens)
        attr['parameter_name'] = kebab_to_camel(attr['name'])
        
        # Handle Go reserved keywords in parameter names
        if attr['parameter_name'] in reserved_keywords:
            attr['parameter_name'] = f"{attr['parameter_name']}Value"
        
        # Track method names to detect duplicates with slots
        method_names.add(attr['go_name'])
        
        # Ensure values key is accessible properly
        if attr.get('type') == 'enum':
            # Make sure 'values' is accessed via dict method not attribute
            if 'values' in attr:
                attr['values_list'] = attr['values']
            else:
                attr['values_list'] = []
                
            # Make the enum type name unique by combining component name and attribute name
            attr['go_type'] = f"{processed['name']}{attr['go_name']}"
        else:
            # Handle array types specifically and determine array item type
            if attr.get('type') == 'array':
                # Default to string for arrays if not specified
                item_type = 'string'
                
                # Check if there's items property with type
                if 'items' in attr and 'type' in attr['items']:
                    item_type = attr['items']['type']
                
                # Special case for modifierKeys which should be string array
                if attr['name'] == 'modifierKeys':
                    item_type = 'string'
                    
                attr['go_type'] = get_go_type(attr.get('type', 'string'), item_type)
                
                # Handle array default values
                if 'default' in attr and attr['default'] == '[]':
                    # Replace empty array defaults with the proper init value
                    if attr['go_type'] == '[]string':
                        attr['default'] = '[]string{}'
                    elif attr['go_type'] == '[]float64':
                        attr['default'] = '[]float64{}'
                    elif attr['go_type'] == '[]int':
                        attr['default'] = '[]int{}'
                    elif attr['go_type'] == '[]bool':
                        attr['default'] = '[]bool{}'
                    else:
                        attr['default'] = '[]any{}'
            else:
                attr['go_type'] = get_go_type(attr.get('type', 'string'))
            
        # Zero value for the type
        if attr.get('type') == 'string':
            attr['zero_value'] = '""'
        elif attr.get('type') == 'boolean':
            attr['zero_value'] = 'false'
        elif attr.get('type') in ('number', 'integer'):
            attr['zero_value'] = '0'
            # Set a proper default value for numeric types if not specified
            if 'default' not in attr or attr['default'] is None:
                attr['default'] = 0
        elif attr.get('type') == 'array':
            attr['zero_value'] = 'nil'
            # For array types, store a proper initialization string
            if attr['go_type'] == '[]string':
                attr['init_value'] = '[]string{}'
            elif attr['go_type'] == '[]float64':
                attr['init_value'] = '[]float64{}'
            elif attr['go_type'] == '[]int':
                attr['init_value'] = '[]int{}'
            elif attr['go_type'] == '[]bool':
                attr['init_value'] = '[]bool{}'
            else:
                attr['init_value'] = '[]any{}'
        elif attr.get('type') == 'object':
            attr['zero_value'] = 'nil'
            # For object types, ensure proper initialization
            attr['init_value'] = 'map[string]any{}'
            # Set a proper default for empty objects
            if 'default' in attr and (attr['default'] == '{}' or attr['default'] == {}):
                attr['default'] = 'map[string]any{}'
        else:
            attr['zero_value'] = 'nil'
    
    # Process slots
    has_default_slot = False
    default_slot_allowed_elements = []
    
    for slot in processed.get('slots', []):
        # Check for default slot
        if slot['name'] == 'default':
            has_default_slot = True
            if 'allowedElements' in slot:
                default_slot_allowed_elements = slot['allowedElements']
        
        # Convert kebab-case to camelCase for the Go name
        parts = slot['name'].split('-')
        special_parts = []
        
        # First part (capitalize if it's a special word and not the first part)
        if parts[0].lower() in special_words:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
        else:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
            
        # Rest of the parts
        for part in parts[1:]:
            if part.lower() in special_words:
                special_parts.append(special_words[part.lower()])
            else:
                special_parts.append(part.title())
        
        slot_go_name = ''.join(special_parts)
        slot['go_name'] = slot_go_name[0].upper() + slot_go_name[1:]  # Capitalize first letter for slot method name
        slot['prop_name'] = f"Prop{slot['go_name']}Slot"
        slot['parameter_name'] = kebab_to_camel(slot['name'])
        
        # Handle Go reserved keywords in parameter names
        if slot['parameter_name'] in reserved_keywords:
            slot['parameter_name'] = f"{slot['parameter_name']}Value"
        
        # Check for duplicate method names with attributes
        if slot['go_name'] in method_names:
            # Mark this slot as having a conflicting name with an attribute
            slot['has_name_conflict'] = True
            # Modify the go_name to avoid conflict
            slot['go_name'] = f"{slot['go_name']}Content"
        else:
            slot['has_name_conflict'] = False
            method_names.add(slot['go_name'])
    
    # Add special Text property for default slot without allowedElements
    if has_default_slot and not default_slot_allowed_elements:
        processed['has_text_prop'] = True
    else:
        processed['has_text_prop'] = False
        
    # Add default slot allowed elements info
    processed['default_slot_allowed_elements'] = default_slot_allowed_elements
    
    # Process events
    for event in processed.get('events', []):
        # Convert kebab-case to camelCase for the Go name
        parts = event['name'].split('-')
        special_parts = []
        
        # First part (capitalize if it's a special word and not the first part)
        if parts[0].lower() in special_words:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
        else:
            special_parts.append(parts[0].lower())  # Keep first word lowercase
            
        # Rest of the parts
        for part in parts[1:]:
            if part.lower() in special_words:
                special_parts.append(special_words[part.lower()])
            else:
                special_parts.append(part.title())
        
        event_go_name = ''.join(special_parts)
        event['go_name'] = event_go_name[0].upper() + event_go_name[1:]  # Capitalize first letter for event method name
        event['prop_name'] = f"PropOn{event['go_name']}"
        event['parameter_name'] = kebab_to_camel(event['name'])
        
        # Handle Go reserved keywords in parameter names
        if event['parameter_name'] in reserved_keywords:
            event['parameter_name'] = f"{event['parameter_name']}Value"
    
    # Set struct name
    processed['struct_name'] = f"spectrum{processed['name']}"
    
    return processed

def main():
    input_dir = "json"
    output_dir = "../"
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
        
        # Replace hyphens with underscores and add sp_ prefix
        output_filename = f"sp_{component_name.replace('-', '_')}.go"
        output_path = os.path.join(output_dir, output_filename)
        
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