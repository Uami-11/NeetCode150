# Valid Palindrome

This is my first time doing something on my own, without trying to understand the content beforehand. And my first step beyond Arrays & Hashings. There is something great about that feeling. 


## The solution

I think I figured out what two pointers mean with this question. But I saw that the question was to see if a string was palindrome was not and I thought "How hard could that be?", the answer was shocking! Just a little harder than I thought.

So a normal way to check if a string is palindrome or not, not caring to ignore special characters is to loop through half the array and check each character with its other half counterpart

A similar concept can be used here, but instead of looping through it like normal, we have to ignore characters that are not alphanumeric. This can be done by adding an index for the latter half of the string as well, not just the first. So we have the first half looping until the half of the string, and the latter half following suit. Then we increment both pointers if the character they are pointing at are the same. But if they arent we have to check if one of the characters is a special character. If it is then we move the pointer thats pointing to it to the next step so that it points to the next character. And in the next loop we can compare that new character and the old normal character (if it wasnt normal then that pointer also shifts). But if all the characters are normal then that just means the string is not a palindrome

## Time & Space Complexity

Time Complexity:
- The function iterates through approximately half of the string (len(str)/2).
- Each iteration involves constant-time operations: character comparisons, checks for digits or letters, and index updates.
- Therefore, the overall time complexity is O(n), where n is the length of the input string.

Space Complexity:
- The function creates a new string 'str' which is a uppercase version of the input string, requiring O(n) space.
- Other variables (pointers, indices) use constant space.
- Hence, total space complexity is O(n).
