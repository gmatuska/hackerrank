from unittest import TestCase
from find_a_word import WordFinder


class TestWordFinder(TestCase):

    def test_find_word(self):
        finder = WordFinder()
        sentence = "foo bar (foo) bar foo-bar foo_bar foo'bar bar-foo bar, foo."
        word = 'foo'
        self.assertEqual(6, finder.find_word(sentence,word))
