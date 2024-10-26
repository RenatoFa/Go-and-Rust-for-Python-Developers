import unittest

from three_sums import ThreeSums


class TestThreeSums(unittest.TestCase):
    def test_three_sums(self):

        self.assertEqual(
            ThreeSums([-1, 0, 1, 2, -1, -4]), [[-1, -1, 2], [-1, 0, 1]])
