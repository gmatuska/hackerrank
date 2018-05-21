#!/usr/bin/python3
import math


class Graph:

    def __init__(self):
        self._edgeList = []

    def add_edge(self, edge):
        self._edgeList.append(edge)

    @staticmethod
    def vertices():
        return []

    def dfs(self):
        counts = []
        visited = []
        for v in self.vertices():
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


class DfsResult:

    def __init__(self, visited, count):
        self._visited = visited
        self._count = count

    def visited(self):
        return self._visited

    def count(self):
        return self._count


class Vertex:

    def __init__(self, value):
        self._value = value

    def value(self):
        return self._value


class Edge:

    def __init__(self, vertexFrom, vertexTo):
        self._vertexList = [vertexFrom, vertexTo]


if __name__ == '__main__':
    tokens = input().strip().split(' ')
    n = int(tokens[0])
    m = int(tokens[1])
    graph = Graph()
    for _ in range(m):
        a, b = input().strip().split(' ')
        graph.add_edge(Edge(Vertex(int(a)), Vertex(int(b))))
    singletons = graph.vertices()
    results = []  # graph.dfs()
    combinations = 0
    for p in range(len(results)):
        for q in range(p+1, len(results)):
            combinations += (results[p] * results[q])
    combinations += math.fsum(list(filter(lambda t: t * singletons, results)))
    """connect disjoint sets to themselves and aggregate combinations"""
    for p in range(len(singletons)-1, 0, -1):
        combinations += p
    print(combinations)
