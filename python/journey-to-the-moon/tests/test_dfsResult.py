from unittest import TestCase

from dfs_result import DfsResult


class TestDfsResult(TestCase):
    def test_is_visited(self):
        result = DfsResult([], 0)
        result.append_to_visited(1)
        result.append_to_visited(2)
        self.assertEqual(2, len(result.visited))
        self.assertTrue(result.is_visited(1))
        self.assertTrue(result.is_visited(2))
        self.assertFalse(result.is_visited(3))
