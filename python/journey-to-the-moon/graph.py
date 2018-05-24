#!/usr/bin/python3
# from dfs_result import DfsResult
from dfs_result import DfsResult


class Graph:

    def __init__(self, edges=None):
        if edges is None:
            self._edges = []
        self._edges = edges

    def add_edge(self, edge):
        if edge not in self._edges:
            self._edges.append(edge)

    @property
    def edges(self):
        return self._edges

    @property
    def vertices(self):
        return set().union(*self._edges)

    """return vertices that have an edge with the given vertex"""
    def get_vertex_edge_vertices(self, vertex):
        return self.vertices.symmetric_difference({vertex})

    def dfs(self):
        visited = set()
        for v in list(self.vertices):
            if set(visited).intersection({v}) == {v}:
                continue
            result = self.dfs_inside(v, DfsResult(visited, 0))
            visited = result.visited
        return result.count

    def dfs_inside(self, vertex, result):
        stack = [vertex]
        while len(stack) > 0:
            stackVertex = stack.pop()
            if result.visited.intersection({stackVertex}) == set():
                result.append_to_visited(stackVertex)
                result.count += 1
                """get the vertices that have an edge attached to the current vertex"""
                for vertex in self.get_vertex_edge_vertices(stackVertex):
                    if result.visited.intersection({vertex}) == set():
                        stack.append(vertex)
        return result
