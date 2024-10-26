from typing import List


def ThreeSums(numbers: List[int]) -> List[List[int]]:
    numbers.sort()
    res = []

    for index, number in enumerate(numbers):
        if index > 0 and number == numbers[index-1]:
            continue

        left, right = index+1, len(numbers)-1

        while left < right:
            three_sum = number + numbers[left] + numbers[right]

            if three_sum > 0:
                right -= 1
            elif three_sum < 0:
                left += 1
            else:
                res.append([number, numbers[left], numbers[right]])
                left += 1
                while numbers[left] == numbers[left-1] and left < right:
                    left += 1
    return res
