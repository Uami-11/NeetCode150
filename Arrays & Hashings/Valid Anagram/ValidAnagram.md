# Valid Anagram

## First time solution

When doing this problem before i created 2 hash maps, and then incremented the chars as keys and int as values and at the end I would compare the two maps.

But that solution takes too much space I think

## This solution

Here I take one single map, Freq, and increment the char values as i go through the first string, and decrement the char values as I go to the second. Then after that I check if any character is not 0, and if there are none then its a valid anagram

### Time and Space Complexity

#### Time Complexity:
- The function iterates through string s once to build the frequency map, which takes O(n) time, where n is the length of s.
- It then iterates through string t once to decrement counts, also O(n).
- Finally, it iterates through the frequency map to check if all counts are zero, which in the worst case is O(n) (if all characters are unique).
- Overall, the total time complexity is O(n) + O(n) + O(n) = O(n).

#### Space Complexity:
- The function uses a map to store character counts, which in the worst case can hold all unique characters in s.
- The size of the map is proportional to the number of unique characters, which is at most the size of the input strings, so O(n).
- Therefore, the space complexity is O(n).
