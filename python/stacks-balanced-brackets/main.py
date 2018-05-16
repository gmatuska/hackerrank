from pprint import pprint


def main():
    tests = input()
    """[pprint(is_balanced(input())) for _ in tests]"""
    for _ in tests:
        pprint(is_balanced(input()))


def is_balanced(brackets):
    stack = []
    for i in list(brackets):
        if i == '[':
            stack.append(']')
        elif i == '{':
            stack.append('}')
        elif i == '(':
            stack.append(')')
        else:
            if len(stack) == 0 or i != stack[len(stack)-1]:
                return "NO"
            _ = stack.pop()
    return "YES" if len(stack) == 0 else "NO"


main()
