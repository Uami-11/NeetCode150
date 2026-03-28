# Top K Frequent Elements

## The solution

For this problem, we need to basically sort all values but not based on their number, but by their frequency. The way this can be done is through a bucket sort. But not in the sense where the indexes are the lowest to the highest values in the given array, but where the indexes are the frequency of the array. So the 'Top K Elements', would be at the end of the array, and what would need to be done is basically to loop backwards until K elements have been added. 

So, to actually find the frequency of each number before we can add it, I made a map to hold each number and their frequencies. Pretty simple. Then I needed to make a slice that was as long as the input array. Now it actually needed to be the length of the array + 1, since the frequency could be the length of it, which would be out of bounds. 

Then we go through the map, and we would put the numbers in the index of their frequencies. Now there can be multiple numbers with the same frequency, so the slice is a slice that holds a slice of integers

## Time & Space Complexity
Time Complexity:
- Counting frequencies: O(n), where n is the length of nums.
- Initializing the frequency list: O(n), since it creates a list of size up to n+1.
- Populating the frequency list: O(n), as it iterates over the count map.
- Collecting the top k elements: In the worst case, it may iterate through the entire frequency list, which is O(n).

Overall, the dominant factors are the counting and the iteration over the frequency list, leading to an overall time complexity of O(n).

Space Complexity:
- The count map stores at most n entries, so O(n).
- The frequency list is of size O(n+1), with each element being a list of numbers sharing the same frequency.
- The result list stores up to k elements, so O(k).

Combined, the space complexity is dominated by the count map and the frequency list, resulting in O(n).
