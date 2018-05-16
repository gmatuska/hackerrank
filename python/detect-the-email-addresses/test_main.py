from unittest import TestCase
import re


class TestMain(TestCase):
    def test_main(self):
        pattern = '(\w+@[\w.]+)'
        lines = 'Mail us at hackers@hackerrank.com to chat more. Or you can write to us at interviewstreet@hackerrank.com.in.out!'
        matches = re.findall(pattern, lines)
        self.assertEqual(2, len(matches))
