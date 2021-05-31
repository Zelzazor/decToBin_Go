# DecToBin for Go language

## Description
This program allows the user to enter a natural number (unsigned integer) in decimal (base-10) and convert it to binary (base-2).

## Requirements 

Req-1: The function should accept an integer and return a string representing the integer in base-2.  
Req-2: The program must allow the user to input their own number.  

## Acceptance Criteria

Cri-1-1: The function must accept number zero (0) as a base case.  
Cri-1-2: The function must accept as input every unsigned integer (within the language chosen range) and succesfully convert it to base-2.  
Cri-1-3: The function must throw an exception if tried to input a number lower than zero (0).  
Cri-2-1: The program must throw an exception if tried to input a value that isn't a number.  
Cri-2-2: The program must throw an exception if tried to input a floating-point number.  

## Scenarios

Sce-1-1-1: Input: 0 - Output: 0  
Sce-1-2-1: Input: 1 - Output: 1  
Sce-1-2-2: Input: 2 - Output: 10  
Sce-1-2-3: Input: 8 - Output: 1000  
Sce-1-2-4: Input: 1024 - Output: 10000000000  
Sce-1-2-5: Input: 6278956 - Output: 010111111100111100101100  
Sce-1-2-6: Input: 2147483647 - Output: 1111111111111111111111111111111  
Sce-1-3-1: Input: -1 - Output: "Exception: Value out of range."  
Sce-1-3-2: Input: -64 - Output: "Exception: Value out of range."  
Sce-2-1-1: Input: "Hello, world" - Output: "Exception: Value is not a number."  
Sce-2-2-1: Input: 1.54 - Output: "Exception: Value is not an integer."