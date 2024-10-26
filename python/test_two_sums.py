import unittest

from two_sums import TwoSums


class TestTwoSums(unittest.TestCase):
    def test_two_sums(self):

        self.assertEqual(TwoSums([2, 7, 11, 15], 9), [1, 2])
