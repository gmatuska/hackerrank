#!/usr/bin/python3

from graph import Graph


def add_edge(graphList, x, y):
    index = 0
    added = False
    while index < len(graphList):
        if graphList[index].intersection({x, y}) == {x, y} or \
                        graphList[index].intersection({x, y}) == {x} or \
                        graphList[index].intersection({x, y}) == {y}:
            graphList[index].add_edge({x, y})
            added = True
            break
        index += 1
    if not added:
        graphList.append(Graph({x, y}))
    return graphList


if __name__ == '__main__':
    tokens = input().strip().split(' ')
    n = int(tokens[0])
    m = int(tokens[1])
    graphs = []
    for _ in range(m):
        a, b = input().strip().split(' ')
        graphs = add_edge(graphs, {int(a), int(b)})
    # singletons = graph.vertices
    # results = graph.dfs()
    combinations = 0
    # for p in range(len(results)):
    #     for q in range(p + 1, len(results)):
    #         combinations += (results[p] * results[q])
    # combinations += math.sum(list(filter(lambda t: t * singletons, results)))
    # """connect disjoint sets to themselves and aggregate combinations"""
    # for p in range(len(singletons) - 1, 0, -1):
    #     combinations += p
    print(combinations)
