## Shift Zeros to the End

Given an array of integers, modify the array in place to move all zeros to the end while maintaingin the relative order of non-zero elements.

Example:
Input: nums = [0, 1, 0, 3, 2]

Output: [1, 3, 2, 0, 0]

### Intuition

This problem has three main requirements:

1. Move all zeros to the end of the array.
2. Maintain the relative order of the non-zero elements.
3. Perform the modification in place.

A naive approach to this problem is to build the output using a separate array (temp). We can add all non-zero elements from the left of nums to this temporary arrayn and leave the rest of it as zeros.

Then, we just set the input array equal to temp.

By identifying and moving the non-zero elements from the left side of the array first, we ensure their order is preserved when we add them to the output.
