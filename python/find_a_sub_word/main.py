import re


class SubWordFinder:

    @staticmethod
    def get_match_count(sentence, subword):
        return len(re.findall('(?<=\w)%s(?=\w)' % subword, sentence))


def main():
    finder = SubWordFinder()
    n = int(input())
    sentence = "\n".join(input() for _ in range(n))
    t = int(input())
    for _ in range(t):
        print(finder.get_match_count(sentence,input().strip()))

main()
