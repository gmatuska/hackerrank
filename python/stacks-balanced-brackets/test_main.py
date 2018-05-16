from unittest import TestCase


class TestMain(TestCase):

    def test_main(self):
        i = '{{{}}}'
        """self.assertEqual(['{', '{', '{', '}', '}', '}'], list(i))"""
        """reversedStack = sorted(stack, reverse=True)"""
        self.assertTrue(self.is_balanced(i))

    def is_balanced(self,brackets):
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
                    return False
                _ = stack.pop()
        return len(stack) == 0
