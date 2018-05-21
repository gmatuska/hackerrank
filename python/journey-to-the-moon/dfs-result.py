#!/usr/bin/python3


class DfsResult:

    def __init__(self, visited, count):
        self._visited = visited
        self._count = count

    @property
    def visited(self):
        return self._visited

    @property
    def count(self):
        return self._count
