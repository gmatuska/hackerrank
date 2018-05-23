from unittest import TestCase

from graph import Graph


class TestGraph(TestCase):
    def test_get_vertex_edges(self):
        """take a union of the edges' vertices"""
        self.assertEqual({1, 2, 3}, {1, 2}.union({1, 3}))
        edge_set = [{1, 2}, {1, 3}, {1, 4}, {1, 5}]
        result_set = set().union(*edge_set)
        self.assertEqual({1, 2, 3, 4, 5}, result_set)
        self.assertEqual({2, 3, 4, 5}, result_set.symmetric_difference({1}))
        graph = Graph(edge_set)
        self.assertEqual({2, 3, 4, 5}, graph.get_vertex_edge_vertices(1))
