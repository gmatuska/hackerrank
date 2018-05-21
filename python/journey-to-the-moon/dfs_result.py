#!/usr/bin/python3


class DfsResult:
    def __init__(self, visited, count):
        self._visited = visited
        self._count = count

    @property
    def visited(self):
        return self._visited

    @visited.setter
    def visited(self, value):
        self._visited = value

    @property
    def count(self):
        return self._count

    @count.setter
    def count(self, value):
        self._count = value

    def is_visited(self, vertex):
        result = list(filter(lambda v: v == vertex, self._visited))
        if len(result) > 0:
            return True
        else:
            return False
