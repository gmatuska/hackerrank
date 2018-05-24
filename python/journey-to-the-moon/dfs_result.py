#!/usr/bin/python3


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
