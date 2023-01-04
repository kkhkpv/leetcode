def twoSum(numbers: list[int], target: int) -> list[int]:
    left, right = 0, len(numbers) - 1
    while left < right:
        temp_sum = numbers[left] + numbers[right]

        if temp_sum == target:
            return [left, right]
        elif temp_sum > target:
            right -= 1
        else:
            left += 1


l = [2, 7, 11, 15]
target = 9

print(twoSum(l, target))
