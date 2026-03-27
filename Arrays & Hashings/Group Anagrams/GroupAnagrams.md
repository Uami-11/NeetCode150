# Group Anagrams

## The solution

Here I have 3 solutions sort of. In my previous solutions of this problem I would make a map that would take in an array that would hold the frequency of each letter. So anagrams would be in the same key in the map. I tried to implement this, but I found out that slices can not be keys in a map, which makes sense as I am typing this. Slices can be resized, and maybe that does some crazy things. Anyway the second approach to this is to make a map of strings. And the string key would be the word but sorted. Since sorted anagrams would be the same word. 

I tried to do this, and I got the logic down pretty quickly. But I had a hard time trying to sort a string. All roads led to sort.Slice or slices.Sort and I was getting confused. You can see evil.go for that. But in my main implementation, I think I got a clean way to sort strings.

Anyway, I also have frequency.go that does the thing I mentioned first. But how? Well, I completely forgot about arrays, which are fixed size and I think that makes them fine. So making the key an array of 26 integers works fine.

Frequency.go is better performing, but I'm keeping my sorted string as the main solution in here and its the one I submitted

## Time & Space Complexity

Time Complexity:
- For each of the n strings in the input list:
  - Converting the string to a rune slice: O(k), where k is the length of the string.
  - Sorting the rune slice using slices.Sort: O(k log k).
  - Appending the original string to the map: O(1) on average.
- Overall, the total time for processing all strings is approximately O(n * k log k).

The final loop to assemble the grouped slices runs in O(n * k) in the worst case, but since it just iterates over the map entries, it's dominated by the previous sorting step.

Total time complexity: O(n * k log k).

Space Complexity:
- The map stores all input strings, so O(n * k) space.
- Additional space for the rune slices during sorting: O(k) per string, which is reused.
- The output list of grouped anagrams also consumes O(n * k) space.

Overall space complexity: O(n * k).
