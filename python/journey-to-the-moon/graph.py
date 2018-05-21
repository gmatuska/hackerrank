#!/usr/bin/python3


class Graph:

    def __init__(self):
        self._edgeList = []

    def add_edge(self, edge):
        self._edgeList.append(edge)

    @property
    def vertices(self):
        return []

    def dfs(self):
        counts = []
        visited = []
        for v in self.vertices:
            if filter(lambda x: x == v.value, visited): continue
            result = []  # self.dfs_inside(v, DfsResult(visited, 0))
            visited = False  # result.IsVisited()
            counts.append(len(result))

    @staticmethod
    def dfs_inside(vertex, result):
        stack = [vertex.value()]
        while len(stack) > 0:
            stackVertex = stack.pop()
            if filter(lambda v: v != stackVertex, result.visited()):
                result.visited().append(stackVertex)
                # len(result)
