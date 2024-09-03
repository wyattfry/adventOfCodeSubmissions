# Changelog

## 2024-09-03

Refactored with ChatGPT

### App Code

1. Merged Similar Functions: Created a generic calculateSum function to compute the sum, accepting different functions to find the first and last digits. This avoids duplicating similar logic.

1. Used Pre-Compiled Regex Patterns: Compiled the regex patterns globally to avoid re-compilation on every function call. This makes the code more efficient.

1. Simplified stringToInt Function: This function is now concise and directly returns results using a map lookup.

1. Removed Redundant or Debug Code: Removed commented-out or debug lines to make the code cleaner and more readable.

1. Refactored Recursive Logic in getLastMatchDigit: Changed from a recursive approach to an iterative one for finding the last matching digit, which is simpler and avoids potential stack overflows.


### Tests

1. Reused a Helper Function for Running Tests: The runTestCases helper function is used to run test cases and print detailed error messages, keeping the test code DRY (Don't Repeat Yourself).

1. Combined Similar Test Cases: Merged similar test cases into a single function to reduce redundancy and improve maintainability. For example, TestGetLastDigitPart2, TestGetLastDigitPart2_2, etc., were combined into TestGetLastMatchDigit.

1. Made Test Names Descriptive: Improved the naming convention of tests for clarity.

1. Used t.Errorf Instead of t.Fatalf: t.Errorf allows the test to continue running other test cases even if one fails. It provides better overall feedback on multiple test cases.