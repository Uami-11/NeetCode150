# Products of Array Except Self

Here we get an array of integers and we have to return an array where each index is the product of every other number except for the number at that index.

## The solution

So the way to go about this is to have prefix and postfix arrays. By multiplying them and by first starting with just 1, we skip the first value for the prefix so that the first value does not get multiplied by itself, and same thing with the postfix

## Time & Space Complexity

The function productExceptSelf has a time complexity of O(n), where n is the length of the input array nums. This is because it performs three separate passes over the array: one to initialize the result array, one to compute prefix products, and one to compute postfix products, each of which iterates through the array once.

The space complexity is O(n) as well. The result array is of size n, and additional variables (prefix and postfix) use constant space. No extra data structures proportional to n are used, so the overall space complexity is dominated by the result array.

Though in NeetCode, the resulting array does not get counted in the memory, so it is technically O(1)
