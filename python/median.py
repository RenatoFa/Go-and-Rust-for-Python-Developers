numbers = [2, 1, 3]

nums = sorted(numbers)

i = len(nums) // 2

if len(nums) % 2 == 1:
    median = nums[i]
else:
    median = (nums[i-1] + nums[i]) / 2

print(median)
