#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Test counter
PASS=0
FAIL=0

# Helper function to print test results
print_test() {
    if [ $1 -eq 0 ]; then
        echo -e "${GREEN}✓ PASS${NC}: $2"
        ((PASS++))
    else
        echo -e "${RED}✗ FAIL${NC}: $2"
        ((FAIL++))
    fi
}

echo "========================================"
echo "PUSH-SWAP TEST SUITE"
echo "========================================"
echo

# Test 1: ./push-swap with no arguments
echo -e "${BLUE}Test 1:${NC} ./push-swap (no arguments)"
OUTPUT=$(./push-swap 2>&1)
if [ -z "$OUTPUT" ]; then
    print_test 0 "No output with no arguments"
else
    print_test 1 "Should display nothing"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 2: ./push-swap "2 1 3 6 5 8"
echo -e "${BLUE}Test 2:${NC} ./push-swap \"2 1 3 6 5 8\""
OUTPUT=$(./push-swap "2 1 3 6 5 8")
COUNT=$(echo "$OUTPUT" | wc -l)
VALID=$(echo "$OUTPUT" | ./checker "2 1 3 6 5 8")
echo "Operations count: $COUNT"
echo "Validation: $VALID"
if [ "$VALID" = "OK" ] && [ $COUNT -lt 9 ]; then
    print_test 0 "Valid solution with less than 9 instructions"
else
    print_test 1 "Should be valid and <9 instructions"
fi
echo "$OUTPUT"
echo

# Test 3: ./push-swap "0 1 2 3 4 5"
echo -e "${BLUE}Test 3:${NC} ./push-swap \"0 1 2 3 4 5\" (already sorted)"
OUTPUT=$(./push-swap "0 1 2 3 4 5" 2>&1)
if [ -z "$OUTPUT" ]; then
    print_test 0 "No output for already sorted array"
else
    print_test 1 "Should display nothing"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 4: ./push-swap "0 one 2 3"
echo -e "${BLUE}Test 4:${NC} ./push-swap \"0 one 2 3\" (invalid input)"
OUTPUT=$(./push-swap "0 one 2 3" 2>&1)
if [ "$OUTPUT" = "Error" ]; then
    print_test 0 "Error displayed for invalid input"
else
    print_test 1 "Should display 'Error'"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 5: ./push-swap "1 2 2 3"
echo -e "${BLUE}Test 5:${NC} ./push-swap \"1 2 2 3\" (duplicates)"
OUTPUT=$(./push-swap "1 2 2 3" 2>&1)
if [ "$OUTPUT" = "Error" ]; then
    print_test 0 "Error displayed for duplicates"
else
    print_test 1 "Should display 'Error'"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 6: 5 random numbers
echo -e "${BLUE}Test 6:${NC} ./push-swap with 5 random numbers"
NUMBERS="3 5 1 4 2"
OUTPUT=$(./push-swap "$NUMBERS")
COUNT=$(echo "$OUTPUT" | wc -l)
VALID=$(echo "$OUTPUT" | ./checker "$NUMBERS")
echo "Numbers: $NUMBERS"
echo "Operations count: $COUNT"
echo "Validation: $VALID"
if [ "$VALID" = "OK" ] && [ $COUNT -lt 12 ]; then
    print_test 0 "Valid solution with less than 12 instructions"
else
    print_test 1 "Should be valid and <12 instructions"
fi
echo

# Test 7: 5 different random numbers
echo -e "${BLUE}Test 7:${NC} ./push-swap with 5 different random numbers"
NUMBERS="9 2 7 4 6"
OUTPUT=$(./push-swap "$NUMBERS")
COUNT=$(echo "$OUTPUT" | wc -l)
VALID=$(echo "$OUTPUT" | ./checker "$NUMBERS")
echo "Numbers: $NUMBERS"
echo "Operations count: $COUNT"
echo "Validation: $VALID"
if [ "$VALID" = "OK" ] && [ $COUNT -lt 12 ]; then
    print_test 0 "Valid solution with less than 12 instructions"
else
    print_test 1 "Should be valid and <12 instructions"
fi
echo

# Test 8: ./checker with no arguments
echo -e "${BLUE}Test 8:${NC} ./checker (no arguments)"
OUTPUT=$(./checker 2>&1)
if [ -z "$OUTPUT" ]; then
    print_test 0 "No output with no arguments"
else
    print_test 1 "Should display nothing"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 9: ./checker "0 one 2 3"
echo -e "${BLUE}Test 9:${NC} ./checker \"0 one 2 3\" (invalid input)"
OUTPUT=$(./checker "0 one 2 3" 2>&1)
if [ "$OUTPUT" = "Error" ]; then
    print_test 0 "Error displayed for invalid input"
else
    print_test 1 "Should display 'Error'"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 10: Invalid operations
echo -e "${BLUE}Test 10:${NC} echo -e \"sa\\npb\\nrrr\\n\" | ./checker \"0 9 1 8 2 7 3 6 4 5\""
OUTPUT=$(echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5")
if [ "$OUTPUT" = "KO" ]; then
    print_test 0 "KO displayed for incorrect sorting"
else
    print_test 1 "Should display 'KO'"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 11: Valid operations
echo -e "${BLUE}Test 11:${NC} echo -e \"pb\\nra\\npb\\nra\\nsa\\nra\\npa\\npa\\n\" | ./checker \"0 9 1 8 2\""
OUTPUT=$(echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2")
if [ "$OUTPUT" = "OK" ]; then
    print_test 0 "OK displayed for correct sorting"
else
    print_test 1 "Should display 'OK'"
    echo "  Got: '$OUTPUT'"
fi
echo

# Test 12: Integration test
echo -e "${BLUE}Test 12:${NC} ARG=\"4 67 3 87 23\"; ./push-swap \"\$ARG\" | ./checker \"\$ARG\""
ARG="4 67 3 87 23"
OUTPUT=$(./push-swap "$ARG" | ./checker "$ARG")
if [ "$OUTPUT" = "OK" ]; then
    print_test 0 "Integration test passed"
else
    print_test 1 "Should display 'OK'"
    echo "  Got: '$OUTPUT'"
fi
echo

echo "========================================"
echo "SUMMARY"
echo "========================================"
echo -e "Passed: ${GREEN}$PASS${NC}"
echo -e "Failed: ${RED}$FAIL${NC}"
echo "========================================"