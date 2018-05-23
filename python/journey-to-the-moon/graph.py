#!/usr/bin/python3
# from dfs_result import DfsResult
from dfs_result import DfsResult


class Graph:

    def __init__(self, edges=None):
        if edges is None:
            edges = []
        self._edges = edges

    def add_edge(self, edge):
        self._edges.append(edge)

    @property
    def vertices(self):
        return list(set().union(*self._edges))

    """return vertices that have an edge with the given vertex"""
    def get_vertex_edge_vertices(self, vertex):
        return set().union(*self._edges).symmetric_difference({vertex})

    def dfs(self):
        counts = []
        visited = []
        for v in self.vertices:
            if filter(lambda x: x == v, visited): continue
            result = self.dfs_inside(v, DfsResult(visited, 0))
            visited = result.visited
            counts.append(result.count)
        return counts

    def dfs_inside(self, vertex, result):
        stack = [vertex.value]
        while len(stack) > 0:
            stackVertex = stack.pop()
            if filter(lambda v: v != stackVertex, result.visited):
                result.append_to_visited(stackVertex)
                result.count += 1
                """get the vertices that have an edge attached to the current vertex"""
                for vertex in self.get_vertex_edge_vertices(stackVertex):
                    if filter(lambda v: v != vertex, result.visited):
                        stack.append(vertex)
        return result
