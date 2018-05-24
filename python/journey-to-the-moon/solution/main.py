#!/usr/bin/python3
import math


def main():
    graphs = []
    results = []
    tokens = input().strip().split(' ')
    n = int(tokens[0])
    m = int(tokens[1])
    for _ in range(m):
        a, b = input().strip().split(' ')
        graphs = add_edge(graphs, int(a), int(b))
    vertices_total = 0
    for graph in graphs:
        vertices_total += len(graph.vertices)
    singletons = n - vertices_total
    [results.append(graph.dfs()) for graph in graphs]
    combinations = 0
    for p in range(len(results)):
        for q in range(p + 1, len(results)):
            combinations += (results[p] * results[q])
    combinations += math.trunc(math.fsum(list(filter(lambda t: t * singletons, results))))
    """connect disjoint sets to themselves and aggregate combinations"""
    for p in range(singletons - 1, 0, -1):
        combinations += p
    print(combinations)


def add_edge(graphList, x, y):
    """don't add edges that are singletons"""
    if x == y:
        return graphList
    index = 0
    added = False
    while index < len(graphList):
        if graphList[index].vertices.intersection({x, y}) == {x, y} or \
                        graphList[index].vertices.intersection({x, y}) == {x} or \
                        graphList[index].vertices.intersection({x, y}) == {y}:
            graphList[index].add_edge({x, y})
            added = True
            break
        index += 1
    if not added:
        graphList.append(Graph([{x, y}]))
    return graphList


class DfsResult:
    def __init__(self, visited, count):
        self._visited = visited
        self._count = count

    @property
    def visited(self):
        return self._visited

    def append_to_visited(self, value):
        self._visited = self._visited.union({value})

    @property
    def count(self):
        return self._count

    @count.setter
    def count(self, value):
        self._count = value

    def is_visited(self, vertex):
        return True if self._visited.intersection({vertex}) == {vertex} else False


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


if __name__ == '__main__':
    main()
