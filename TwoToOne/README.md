Function: longest

Description:


Given two strings s1 and s2, which contain only lowercase letters from 'a' to 'z', the function longest returns a new string that consists of distinct letters taken from either s1 or s2. The resulting string is sorted in alphabetical order.

Parameters:

s1: A string containing only lowercase letters.

s2: A string containing only lowercase letters.

Returns:

A string composed of distinct letters from s1 and s2, sorted alphabetically.

Examples:

Example 1:
```go
a := "xyaabbbccccdefww"
b := "xxxxyyyyabklmopq"
longest(a, b) // Output: "abcdefklmopqwxy"
