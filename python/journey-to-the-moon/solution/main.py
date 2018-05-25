#!/usr/bin/python3


def main():
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


class Graph:

    def __init__(self, edges=None):
        self._vertices = set()
        self._vertex_dict = {}
        if edges is not None:
            for edge in edges:
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


if __name__ == "__main__":
    main()
