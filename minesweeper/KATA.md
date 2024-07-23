# The Minesweeper Kata

You can find the details of this Kata at [Minesweeper Kata](https://codingdojo.org/kata/Minesweeper/).
The task is to calculate the number of adjacent mines for each field on a Minesweeper board.

## Suggested Test Cases

Those are meant for acceptance tests, not steps in TDD!

### Input

    4 4
    *...
    ....
    .*..
    ....
    3 5
    **...
    .....
    .*...
    0 0

### Output

    Field #1:
    *100
    2210
    1*10
    1110
    
    Field #2:
    **100
    33200
    1*100

__Remember__:

- Do TDD.
- Do one requirement at a time, don’t read ahead.
- Test first, there is no try.

## Task

Provide the following function

    func FindBombs(reader io.Reader, writer io.Writer)

Read the input from the reader and output the results to the writer.

## Requirements

1. The method takes a single integer parameter
2. and returns the prime factors of this integer in numerical order.
