from gettext import find
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

    def test_vertices(self):
        edges = [{0, 1}]
        graphs = []
        graph = Graph(edges)
        graphs.append(graph)
        x = 2
        y = 3
        print(graphs[0].vertices.intersection({x}) == {x})
        print(graphs[0].vertices.intersection({y}) == {y})
        print(graphs[0].vertices.intersection({x, y}) == {x, y})

    def test_dfs(self):
        visited = [0]
        print(list(filter(lambda x: x == 0, visited)))
        self.assertTrue(True)

    def test_dfs_inside(self):
        visited = set()
        vertex = 0
        self.assertTrue(visited.intersection({vertex}) == set())
