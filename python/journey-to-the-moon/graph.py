#!/usr/bin/python3
from dfs_result import DfsResult


class Graph:

    def __init__(self, edges=None):
        self._vertices = set()
        self._vertex_dict = {}
        if edges is not None:
            for edge in list(edges):
                self.add_edge(edge)

    def add_edge(self, edge):
        """add to dictionary"""
        edge_list = list(edge)
        if edge_list[0] == edge_list[1]:
            return
        key = edge_list[0]
        value = edge_list[1]
        try:
            self._vertex_dict[key] = self._vertex_dict[key].union({value})
        except KeyError:
            self._vertex_dict[key] = {value}
        try:
            self._vertex_dict[value] = self._vertex_dict[value].union({key})
        except KeyError:
            self._vertex_dict[value] = {key}

        self._vertices = self._vertices.union(edge)

    @property
    def vertices(self):
        return self._vertices

    """return vertices that have an edge with the given vertex"""
    def get_vertex_edge_vertices(self, vertex):
        return self._vertex_dict[vertex]

    def dfs(self):
        counts = []
        visited = set()
        for v in list(self.vertices):
            if set(visited).intersection({v}) == {v}:
                continue
            result = self.dfs_inside(v, DfsResult(visited, 0))
            visited = visited.union(result.visited)
            counts.append(result.count)
        return counts

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
