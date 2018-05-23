#!/usr/bin/python3


class Edge:

    def __init__(self, vertexFrom, vertexTo):
        self._vertexSet = {vertexFrom, vertexTo}

    def vertices(self):
        return self._vertexSet
