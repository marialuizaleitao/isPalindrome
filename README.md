# Palindrome Checker

## Overview

This project implements a Palindrome Checker in Go, providing two different implementations to check whether a given string is a palindrome. It also includes comprehensive unit tests, test-driven development (TDD) practices, and benchmark tests to compare the performance of the two implementations.

### Technologies Used

- **Go (Golang)**: The primary programming language used for implementation.
- **Testing**: Go's built-in testing framework for unit testing.
- **Benchmarking**: Go's benchmarking feature to measure the performance of functions.

## What is a Palindrome?

A palindrome is a word, phrase, number, or other sequence of characters that reads the same forward and backward. Examples of palindromes include "radar," "level," "12321," and "A man, a plan, a canal, Panama!".

## Palindrome Checker Implementations

### isPalindrome Function

The `isPalindrome` function checks if a given string is a palindrome by reversing the string and comparing it with the original. If the reversed string matches the original, the function returns `true`; otherwise, it returns `false`. This implementation uses the `invertText` helper function to reverse the string.

### isPalindrome2 Function

The `isPalindrome2` function directly compares characters from the start and end of the string, moving towards the center. It iterates through the string, comparing characters at corresponding positions. If any characters do not match, the function returns `false`; otherwise, it returns `true`.

## Test-Driven Development (TDD) Approach

Both `isPalindrome` and `isPalindrome2` functions were developed using a test-driven development (TDD) approach. Test cases were written to define the expected behavior of the functions before implementation. This ensured that the functions were thoroughly tested and behaved correctly under various scenarios.

## Benchmarking for Performance Comparison

Benchmark tests were conducted to measure the performance of the `isPalindrome` and `isPalindrome2` functions. These benchmarks were used to compare the execution time of the two implementations and determine which one performs better in terms of speed.

## Results and Performance Analysis

Benchmark tests revealed that the `isPalindrome2` function, which directly compares characters, performs faster than the `isPalindrome` function, which involves string reversal. This is because the `isPalindrome2` function avoids the overhead of creating a reversed string and comparing it with the original. The direct character comparison approach of `isPalindrome2` results in better performance, especially for large strings.

## Conclusion

In conclusion, the Palindrome Checker project demonstrates the implementation of two different approaches to check for palindromes in a given string. Through the use of test-driven development (TDD) practices and benchmarking, the performance differences between the two implementations were analyzed, highlighting the importance of choosing efficient algorithms for common programming problems like palindrome detection.
