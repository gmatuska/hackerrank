#!/usr/bin/python3
# from dfs_result import DfsResult


class Graph:

    def __init__(self, edges=None):
        if edges is None:
            edges = []
        self._edges = edges

    def add_edge(self, edge):
        self._edges.append(edge)

    @property
    def vertices(self):
        return []

    """return vertices that have an edge with the given vertex"""
    def get_vertex_edges(self, vertex):
        return set().union(*self._edges).symmetric_difference({vertex})

    def dfs(self):
        counts = []
        visited = []
        for v in self.vertices:
            if filter(lambda x: x == v.value, visited): continue
            result = []  # self.dfs_inside(v, DfsResult(visited, 0))
            # visited = result.is_visited
            counts.append(len(result))

    @staticmethod
    def dfs_inside(vertex, result):
        stack = [vertex.value()]
        while len(stack) > 0:
            stackVertex = stack.pop()
            if filter(lambda v: v != stackVertex, result.visited):
                result.append_to_visited(stackVertex)
                result.count += 1
                """get the edges attached to the current vertex"""
                # edges = self.get_vertex_edges(stackVertex)
