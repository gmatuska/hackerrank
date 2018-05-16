#!/bin/python3

import re


class AlienUsername:

    @staticmethod
    def is_valid_alien_name(name):
        return len(re.findall('^[_.][0-9]+[a-zA-Z]*_?$', name)) > 0


def main():
    results =[]
    n = int(input())
    alien = AlienUsername()
    for _ in range(n):
        if alien.is_valid_alien_name(input().strip()):
            results.append('VALID')
        else:
            results.append('INVALID')
    print(*results, sep='\n')

main()
