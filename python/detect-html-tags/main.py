#!/bin/python3

import re

def detect_html_tags():
    d = {}
    num_lines = int(input())
    search_pattern = '<\s*([a-z0-9]+)\s*(?=>)*'
    for i in range(num_lines):
        match = re.findall(search_pattern,input())
        if bool(match):
            for tag in match:
                d[tag] = tag
    output = (';'.join('%s' % (key) for (key,_) in sorted(d.items())))
    print(output)

detect_html_tags()

