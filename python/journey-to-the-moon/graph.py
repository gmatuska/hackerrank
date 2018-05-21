#!/usr/bin/python3
# from dfs_result import DfsResult


class Graph:

    def __init__(self):
        self._edgeList = []

    def add_edge(self, edge):
        self._edgeList.append(edge)

    @staticmethod
    def vertex_edges(vertex):
        """go through edge list and return edges containing vertex"""
        return vertex

    @property
    def vertices(self):
        return []

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
                result.visited.append(stackVertex)
                result.count += 1
                """get the edges attached to the current vertex"""
                # edges = filter(lambda v: v.value == stackVertex, self.vertices)
