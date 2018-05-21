#!/usr/bin/python3
from unittest import TestCase


class TestVertex(TestCase):

    def test_using_lambda_in_vertices(self):
        vertices = [1, 2, 3, 4, 5]
        coll = list(filter(lambda x: x % 2 == 0, vertices))
        self.assertEqual(2, len(coll))
