# Intuition
- Use a hash map to store numbers and their indices.
- Check for the complement of each element in the hash map.

# Approach
- Create an empty hash map to store numbers and indices.
- Loop through the array, calculate the complement for each element.
- Check if the complement exists in the hash map.
- Return the indices of the current element and its complement if found.
- Store each element and its index in the hash map for future lookups

# Complexity

## Time complexity: $$O(n)$$

- The solution iterates through the array once, performing constant time operations on the hash map.

## Space complexity: $$O(n)$$
- The hash map can store at most n elements, where n is the length of the input array.

# Code
```
func twoSum(nums []int, target int) []int {

    numsMap := make(map[int]int)

    for i, num := range nums {
        rem := target - num
        if index, found := numsMap[rem]; found {
            return []int{index, i}
        }
        numsMap[num] = i
    }

    return []int{}
}
```