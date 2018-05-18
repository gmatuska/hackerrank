#!/usr/bin/python3
import os


class Graph:

    def __init__(self):
        self._edgeList = []

    def add_edge(self, edge):
        self._edgeList.append(edge)


class Vertex:

    def __init__(self, value):
        self._value = value


class Edge:

    def __init__(self, vertexFrom, vertexTo):
        self._vertexList = [vertexFrom, vertexTo]


if __name__ == '__main__':
    tokens = input().strip().split(' ')
    n = int(os.sys.argv[0])
    m = int(os.sys.argv[1])
    graph = Graph()
    for _ in range(m):
        a, b = input().strip().split(' ')
        graph.add_edge(Edge(Vertex(int(a)), Vertex(int(b))))
