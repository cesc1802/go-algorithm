## Write an algorithm to determine if a number num is a happy number.
### We use the following process to check if a given number is a happy number:
    - starting with the given number num, replace the number with the sum of the squares of its digits.
    - repeat the process until:
        - the number equal to 1, which will depict that the given number num is a happy number
        - it enters a cycle,, which will depict that the given number num is not a happy number

### Return TRUE if num is a happy number, and FALSE if not.