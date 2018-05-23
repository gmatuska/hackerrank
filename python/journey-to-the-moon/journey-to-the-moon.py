#!/usr/bin/python3
import math

from graph import Graph

if __name__ == '__main__':
    tokens = input().strip().split(' ')
    n = int(tokens[0])
    m = int(tokens[1])
    graph = Graph()
    for _ in range(m):
        a, b = input().strip().split(' ')
        graph.add_edge({int(a), int(b)})
    singletons = graph.vertices
    results = graph.dfs()
    combinations = 0
    for p in range(len(results)):
        for q in range(p+1, len(results)):
            combinations += (results[p] * results[q])
    combinations += math.fsum(list(filter(lambda t: t * singletons, results)))
    """connect disjoint sets to themselves and aggregate combinations"""
    for p in range(len(singletons)-1, 0, -1):
        combinations += p
    print(combinations)
