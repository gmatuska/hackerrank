#!usr/bin/python3


def staircase(num):
    for i in range(0, num):
        message = ''
        for j in range(0, num-i-1):
            message = message + ' '

        for j in range(0, i+1):
            message = message + '#'

        print("{!s}".format(message))

if __name__ == '__main__':
    n = int(input())
    staircase(n)
