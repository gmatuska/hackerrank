import re


class SubWordFinder:

    def get_match_count(self, sentence, subword):
        pattern = ''.join(['(?<=\w)(', subword, ')(?=\w)'])
        return len(re.findall(pattern, sentence))
