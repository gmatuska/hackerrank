from unittest import TestCase
from find_a_sub_word import SubWordFinder


class TestFind_Sub_Word(TestCase):

    def test_find_sub_word(self):
        sentence = 'existing pessimist optimist this is'
        wordFinder = SubWordFinder()
        subword = 'is'
        self.assertEqual(wordFinder.get_match_count(sentence, subword), 3)
