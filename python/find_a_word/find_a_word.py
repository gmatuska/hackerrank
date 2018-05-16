#!/bin/python3
import re


class WordFinder:

    @staticmethod
    def find_word(sentences, word):
        counter = 0
        pattern = '(^|(?<=\W)){}((?=\W)|$)'.format(word)
        for s in sentences:
            counter += len(re.findall(pattern,s))
        return counter


def main():
    finder = WordFinder()
    results = []
    sentences = []
    n = int(input())
    for _ in range(n):
        sentences.append(input())
    t = int(input())
    for _ in range(t):
        results.append(finder.find_word(sentences, input()))
    print(*results, sep='\n')

main()
