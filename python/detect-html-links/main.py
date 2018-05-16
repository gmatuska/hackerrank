#!/bin/python3
import re

num_lines = int(input())
search_pattern = r'<a\shref="(.*?)"(?:>|(?:\stitle=".*?>))(.*?)<'
for i in range(num_lines):
    match = re.findall(search_pattern, input())
    if(bool(match)):
        for link, title in match:
            print(link + ',' + title.strip())


