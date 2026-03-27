# Two Sum

## The solution

So going in to this, I knew we had to make a map to keep track of all the numbers we have seen then we can compare the difference, and see if the current number's difference existed and return those indecies. But I went about it wrong the first time. I first made the whole map, then I went through it again, and would check if the difference would exist in the map and return the indecies. My thought process was that, "well the index of the difference would be later on so I can just return the index of the current value, and then the difference". But here I did not account for two things

### Oversights

#### Same value numbers

If there were two numbers in the array that had the same value, the map would overwrite it, so when I check for the difference and i need two of the same value (which was one of the test cases in NeetCode), I would be shit out of luck.

#### Maps in Go

What I did not know was that the order of maps in go are randomized. So i thought i was going serially through the map but I really wasn't and the return value would sometimes be right and sometimes wrong.

## The solution (cont)

So what I ended up doing, and what I had done in previous languages, was to make the map and check for the difference in the same loop. This fixes both of my problems. If there is a same value number, when the first time it appears shows up it gets stored, and when the second time it comes along, and I need that value, I already have the index on loop and i can just grab the index of the previous instance from the map. And it seems to NeetCode, we just need the last instance of the value, if both is not needed

And the randomization is no longer needed, because I know the difference is already stored in the map from an earlier iteration in the loop so i can just put the index of difference first

## Time & Space Complexity
The time complexity of the twoSum function is O(n), where n is the number of elements in the input slice nums. This is because the function iterates through the list once, performing constant-time operations (map lookups and insertions) during each iteration.

The space complexity is also O(n), since in the worst case, the map numMap can store all n elements from the input list. This additional space is used to keep track of the indices of the numbers encountered so far.
