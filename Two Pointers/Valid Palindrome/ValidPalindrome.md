# Valid Palindrome

This is my first time doing something on my own, without trying to understand the content beforehand. And my first step beyond Arrays & Hashings. There is something great about that feeling. 


## The solution

I think I figured out what two pointers mean with this question. But I saw that the question was to see if a string was palindrome was not and I thought "How hard could that be?", the answer was shocking! Just a little harder than I thought.

So a normal way to check if a string is palindrome or not, not caring to ignore special characters is to loop through half the array and check each character with its other half counterpart

A similar concept can be used here, but instead of looping through it like normal, we have to ignore characters that are not alphanumeric. This can be done by adding an index for the latter half of the string as well, not just the first. So we have the first half looping until the half of the string, and the latter half following suit. Then we increment both pointers if the character they are pointing at are the same. But if they arent we have to check if one of the characters is a special character. If it is then we move the pointer thats pointing to it to the next step so that it points to the next character. And in the next loop we can compare that new character and the old normal character (if it wasnt normal then that pointer also shifts). But if all the characters are normal then that just means the string is not a palindrome

## The revamp

My previous solution did pass all of NeetCode's test, but in LeetCode it was a different story. I realized I did not cover many edge cases and in the end when all my code worked, it was a mess of if statements.

So then I decided to learn what real two pointers is. Where we have a left and right, which increments and decrements through the string respectively. And though that thought came to mind I did not implement it so what does that matter.

Then if the left pointer is pointing to a special character then we make it loop forward until it isnt

The same with the right and we make it loop backwards.

Then we check if the characters is not equal and since we know they are not special characters then we can just return false. Then if it isnt then we just move the pointers along

It's so much simpler it's embarassing. I need to check for all my solution's LeetCode submissions from now

## Time & Space Complexity

Time Complexity:
- The function uses two pointers, left and right, which traverse the string from both ends towards the center.
- Each character is visited at most once by either pointer, and the inner while loops skip non-alphanumeric characters efficiently.
- Therefore, the overall time complexity is O(n), where n is the length of the string, since each character is processed at most once.

Space Complexity:
- The function uses a constant amount of extra space for variables left, right, and temporary variables within the loops.
- No additional data structures proportional to the input size are used.
- Hence, the space complexity is O(1).
