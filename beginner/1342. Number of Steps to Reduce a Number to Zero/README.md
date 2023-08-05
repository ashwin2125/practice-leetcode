# 1342. Number of Steps to Reduce a Number to Zero

## Intuition

- The goal is to reduce the given integer to zero in the minimum number of steps.
- Each step involves either dividing an even number by 2 or subtracting 1 from an odd number.

## Approach 

- We utilize the properties of bitwise operations (binary representation) to optimize the counting process.
- The binary representation of an even number always ends with a 0.
- The binary representation of an odd number always ends with a 1 and has trailing 0s after subtracting 1.
- Right-shift operations are faster than division, leveraging CPU optimizations for bitwise operations.
- Right-shifting an even number by one position is equivalent to division by 2.
- By subtracting 1 from an odd number, we ensure that its trailing zeroes are not affected in the counting process.

## Complexity
- Time complexity: $$O(log n)$$

- Space complexity: $$O(1)$$
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

## Code
```
func numberOfSteps(num int) int {
    stepsRequired := 0
    for num != 0 {
        // using bitwise operators
        if num & 1 == 0 {
            // rightshift to divide the num/2.
            num >>= 1
        } else {          
            num -= 1
        }
        stepsRequired++
    }
    return stepsRequired
}
```