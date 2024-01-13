# The String Calculator Kata

This Kata was created by Roy Osherove: [TDD Kata - String Caluclator](http://osherove.com/kata)

Step by step you are supposed to implement a function that takes a string containing numbers and calculates their sum.

__Remember__:

- Do TDD.
- Do one requirement at a time, do not read ahead.
- Test first, there is no try.

## Task

Provide the following method

    int Add(string numbers)

## Requirements

1. The method returns `0` given an empty string.

       Add("") == 0

2. The method returns the number given the string only contains a single number.

       Add("6") == 6

3. Given the string contains two comma-separated numbers, then the method returns their sum.

       Add("1,2") == 3

4. The method should deal with an arbitrary amount of numbers as well.

       Add("1,2,3,4,5") == 15

5. The method should allow new lines instead of commas between numbers.

       Add("1\n2,3") == 6

6. Allow different delimiters. Given the input string starts with a comment line like `//[delim]\n[numbers...]`, then `delim` delimiter should work too: 

       Add("//;\n1;2") == 3`

7. If the input string contains a negative number, then throw an exception or return an error like `negatives not allowed` including the passed in negative number.

       Add("-1,2,3") produces "negatives not allowed: -1"

8. If there are multiple negative numbers, then show them all within the exception or error.

       Add("1,-2,-3") produces "negatives not allowed: -2, -3"

9. Add a function to the calculator that returns the count how many times the `Add()` function was called. Implement it Test-First, the TDD way.

       int GetCalledCount()

10. Numbers bigger than 1000 should be ignored.

        Add("1,1001,2") == 3

11. Allow delimiters of any length, like 

        "//[***]\n1***2***3" == 6

12. Allow multiple delimiters, like

        "//[*][&]\n1*2&3" == 6

13. Ensure that delimiters of variable length are still supported, like 

        "//[##][$]\n1$2##3" == 6