# The Longest Consecutive Sequence

I had to use a hint for this question, I used hint 2, which suggested to keep track of the starting numbers of any sequence. So from there I figured out what to do.

I kept a hashmap of each number so i could see if any number existed in the array.

Then from that i checked for any number that didnt have a number less than it, so a candidate for starter. Then put that in a seperate hashmap that holds an array, which would be the sequence

Then i check the length of each sequence and return the length of the longest one.

I think maybe i dont have to keep track of any individual sequence because thats not what im returning, maybe that could save space
