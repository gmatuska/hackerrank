#!/usr/bin/python3
import math

from graph import Graph


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


if __name__ == '__main__':
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
