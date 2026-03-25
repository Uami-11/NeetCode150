# Contains Duplicates

This program needs to see if there are any duplicate integers in the array given in parameter.

## First solution that might come to mind

One way to do it is nest over two loops over the array, checking each and every number to see if any match. And this will get you an answer, but the time complexity will be O(n<sup>2</sup>)

We want a more efficient one, one that can do it in one fell swoop

## My previous solutions

When doing this before using python, a way to do it was with a set. A list of elements that could not have any duplicates whatsoever. So just make a set out of this nums list and compare the lengths. If there are any differences then there were some duplicate values in the nums list. 

I thought that was pretty nice, and when redoing this for go, I thought to do the same. There aren't any sets in go, but making one through a map with empty structs as values worked fine.

But there is a slightly better way to do it I think

## My Solution

Here, I still make the "set", but instead of comparing the length of each list at the end I have a condition to exit early if the key that im trying to insert already exists. This lookup is done through the map which lookup time is O(1), so its fine

That way in the very best case scenario, the loop doesn't keep going on. 

### Time and Space Complexities
Final time complexity is O(n)
Because we are looping through the whole nums at the worst case

Space complexity: O(n), in the worst case when all elements are unique, the map will store all n elements.
And since looping through nums at the worst case is adding all of it to the set, the space complexity is so
