from unittest import TestCase

from graph import Graph


class TestGraph(TestCase):
    def test_get_vertex_edge_vertices(self):
        edge_set = [{1, 2}, {1, 3}, {1, 4}, {1, 5}]
        result_set = set().union(*edge_set)
        self.assertEqual({1, 2, 3, 4, 5}, result_set)
        self.assertEqual({2, 3, 4, 5}, result_set.symmetric_difference({1}))
        graph = Graph(edge_set)
        self.assertEqual({2, 3, 4, 5}, graph.get_vertex_edge_vertices(1))

    def test_add_edge(self):
        edges = [{0, 1}, {1, 2}, {2, 2}]
        graph = Graph(edges)
        graph.add_edge({2, 2})
        self.assertEqual(3, len(graph.edges))
        graph.add_edge({3, 2})
        self.assertEqual(4, len(graph.edges))

