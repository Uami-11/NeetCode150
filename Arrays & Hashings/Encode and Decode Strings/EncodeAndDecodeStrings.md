# Encode and Decode Strings 

## The solution

So here we have to encode an array of strings into a single string, and then be able to decode it back to the original string array. We can do this through delimeters, where we separate the strings with a special character. Though just doing this poses a problem, since the given strings can have any character, the delimeter can be one of the characters. So doing that could separate strings that were together. We fix this by also giving how long the original string is before the delimeter. Now we need both the length and the delimeter, because numbers can also be in the string characters, so while decoding the length can be seen as much larger than it actually is. So having a delimeter after the number lets us know how long the original string was and where it starts from.

As you can see I have two files, the main one and a program help by ducttape and glue. The original solution was the latter. Where I would check for the length and strings all in a for loop and such

### Duct Tape and Glue

For encoding, I go through every string in the given array, and add it to this other string with the length and the delimeter in front. Now i do this with strings builder. I initially just added the string to another string using the plus operator, but I learnt that was inefficient, since it basically creates a new string every time. So I found out string builders, and basically you can use Write function and specifically the WriteString function to add to the string you're calling it from. I do this later with WriteByte and WriteRune which can append characters to strings. 

Then I manually go through each letter to build up the exact length of the string from the characters, because the length can be multiple digits. Then i use that length to loop through the string for that amount of length, and append to the array from that loop.

And it works, but it looks like a mess and is a mess in general. So I needed to know how to clean this up because there had to be a better way than to loop through the string with magic numbers

Then I found out why slices are called slices

## The solution (cont)

So, while WriteString is useful, it's not even the best way to write to strings with strings builder. What the heck. A string builder type is an io writer, so we can use fmt.Fprintf to edit the string and add to it, with string formatting. So basically encoded is improved with that.

Now decoded is much more abstract that the first one looks like it was written by a caveman who has never known how to use slices. 
First we find the how long the number string is through strings.IndexByte, which finds a rune's index in a given string now im making the string a slice of the encoded string from the current i to the end. And in later loops, the i is not the start, so we have to add th e index so it is universal since the IndexByte returns the index relative to the given string which is only a slice of the original string. 

so now that we know how long the digit actually is, we can store it in a length variable using Atoi in the strconv package. We do it from i to the index we got because slices is not inclusive of the latter index

now that we have the length we have to take out the string from the encoded string. First we have to set a starting point on how to slice it out of the string. We take the start from where the delimeter is + 1, since we want whats after it. So we take a slice that starts at the start (duh), then until the starting index + the length we got. And then we append that word

Now we must update the index so that it goes past this word and onto the number for the next one
we can do this by making i equal start + length. Because start is the index the previous word started at, and length is how long it is. So the new index should land where it ends


## Time & Space Complexity

The Encode method processes each string in the input slice exactly once. For each string, it appends its length and content to a StringBuilder, which is an O(1) amortized operation per append. Therefore, the overall time complexity of Encode is O(N * L), where N is the number of strings and L is the average length of each string, since concatenation and formatting operations are proportional to string length.

The space complexity of Encode is also O(N * L), as it constructs a new string that contains all input strings along with their length prefixes.

The Decode method processes the encoded string sequentially. For each segment, it finds the position of the '*' character (which is O(1) on average), extracts the length, and then slices the substring of that length. Each character in the encoded string is processed exactly once, so the time complexity is O(M), where M is the total length of the encoded string, which is proportional to the sum of the lengths of all original strings plus their length prefixes.

The space complexity of Decode is O(N * L), as it reconstructs all original strings, storing them in a slice.
