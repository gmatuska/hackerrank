#!/usr/bin/python3

from graph import Graph


if __name__ == '__main__':
    graph = Graph()
    tokens = input().strip().split(' ')
    n = int(tokens[0])
    m = int(tokens[1])
    for _ in range(m):
        a, b = input().strip().split(' ')
        edge = {int(a), int(b)}
        graph.add_edge(edge)
    singletons = n - len(graph.vertices)
    results = graph.dfs()
    combinations = 0
    if len(results) > 1:
        for p in range(len(results)):
            for q in range(p + 1, len(results)):
                combinations += (results[p] * results[q])
    blah = 0
    for result in results:
        blah += result * singletons
    combinations += blah
    """connect disjoint sets to themselves and aggregate combinations"""
    for p in range(singletons - 1, 0, -1):
        combinations += p
    print(combinations)
