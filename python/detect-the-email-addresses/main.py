import re


def main():
    n = int(input())
    lines = ' '.join(input() for _ in range(n))
    matches = re.findall('(\w+@\w+.\w+)', lines)
    print(*sorted(list(set(matches))), sep=';')

main()
