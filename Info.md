# Push-Swap Project Structure

```
push-swap/
├── cmd/
│   ├── push-swap/
│   │   └── main.go          # Main entry point for push-swap program
│   └── checker/
│       └── main.go          # Main entry point for checker program
├── internal/
│   ├── stack/
│   │   ├── stack.go         # Stack data structure and operations
│   │   └── stack_test.go    # Unit tests for stack operations
│   ├── parser/
│   │   ├── parser.go        # Input parsing and validation
│   │   └── parser_test.go   # Unit tests for parser
│   ├── solver/
│   │   ├── solver.go        # Sorting algorithm implementation
│   │   ├── small_sort.go    # Optimized sorting for small stacks
│   │   └── solver_test.go   # Unit tests for solver
│   └── operations/
│       ├── operations.go    # Operation definitions and execution
│       └── operations_test.go # Unit tests for operations
├── go.mod                   # Go module file
├── Makefile                 # Build automation
└── README.md               # Project documentation
```

## Implementation Plan

### 1. Stack Implementation
- Create a stack structure using Go slices
- Implement all 11 operations (sa, sb, ss, pa, pb, ra, rb, rr, rra, rrb, rrr)
- Add utility methods for stack manipulation

### 2. Input Parser
- Parse command-line arguments
- Validate integers and check for duplicates
- Handle edge cases (empty input, invalid formats)

### 3. Sorting Algorithm
- Implement different strategies based on stack size:
  - 2-3 elements: Direct optimal solutions
  - 4-5 elements: Small optimization algorithms
  - Larger stacks: More complex algorithms (possibly radix sort or chunk-based approach)

### 4. Operation Optimization
- Track and minimize operation count
- Implement operation combination logic (e.g., using rr instead of ra + rb)

### 5. Checker Implementation
- Read operations from stdin
- Execute operations on the input stack
- Verify final state (stack A sorted, stack B empty)


# Sample 100 random unique numbers for testing
# You can use these to reproduce test results

# Test Set 1 (use this for consistent testing)
42 -87 91 -23 56 -78 12 -45 67 -34 89 -12 23 -67 45 -89 78 -56 34 -91 65 -42 87 -65 21 -88 76 -21 54 -76 98 -54 32 -98 10 -32 88 -10 43 -43 99 -99 11 -11 55 -55 77 -77 33 -33 66 -66 22 -22 44 -44 9 -9 8 -8 7 -7 6 -6 5 -5 4 -4 3 -3 2 -2 1 -1 100 -100 51 -51 62 -62 73 -73 84 -84 95 -95 16 -16 27 -27 38 -38 49 -49 60 -60 71 -71 82 -82 93 -93 14 -14 25 -25 36 -36

# Test Set 2 (another set for variety)
1500 -2300 800 -1200 3400 -900 2100 -3100 500 -4200 1800 -600 2900 -1900 700 -2700 3300 -1100 400 -3900 1300 -4100 2600 -800 3800 -1600 1100 -2800 4000 -1400 200 -3600 1700 -2200 3200 -1000 600 -3400 2400 -1800 900 -2600 3700 -1500 300 -3300 2000 -4300 1400 -700 3000 -2100 1900 -3700 2500 -1300 3900 -500 100 -4000 1600 -2900 3500 -1700 800 -2400 2200 -3800 4100 -400 1200 -3000 2700 -1100 3100 -600 400 -4400 2300 -2000 3600 -900 1000 -3200 2800 -1400 4200 -300 1500 -3500 3300 -1600 700 -2500 4300 -200 1100 -4100 2900 -1200

# How to use these:
# Copy a line of numbers and paste into your test:
# 
# ./push-swap "42 -87 91 -23 ... (all numbers)" | wc -l
# ./push-swap "42 -87 91 -23 ... (all numbers)" | ./checker "42 -87 91 -23 ... (all numbers)"
#
# Or create a variable:
# ARG="42 -87 91 -23 ... (all numbers)"
# ./push-swap "$ARG" | wc -l
# ./push-swap "$ARG" | ./checker "$ARG"