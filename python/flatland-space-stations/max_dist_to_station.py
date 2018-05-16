#!/usr/bin/python3
import math


def max_dist_to_station(numCities: int, numStations: int) -> int:
    stations = sorted([int(c_temp) for c_temp in input().strip().split(' ')])
    """start from 0 and go to next station"""
    max_dist = stations[0] - 0
    distRight = numCities - 1 - stations[numStations - 1]
    if distRight > max_dist:
        max_dist = distRight
    index = 0
    for station in stations:
        if index < m - 1:
            dist = math.floor((stations[index+1] - station) / 2)
            if dist > max_dist:
                max_dist = dist
            index = index + 1
    return max_dist


if __name__ == "__main__":
    n, m = input().strip().split(' ')
    n, m = [int(n), int(m)]
    print(max_dist_to_station(n, m))
