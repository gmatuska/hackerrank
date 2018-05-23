from unittest import TestCase

from edge import Edge


class TestGraph(TestCase):
    def test_get_vertex_edges(self):
        """take a union of the edges' vertices"""
        self.assertEqual({1, 2}, Edge(1, 2).vertices())
        self.assertEqual({1, 2, 3}, Edge(1, 2).vertices().union(Edge(1, 3).vertices()))
