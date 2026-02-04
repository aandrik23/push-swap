#!/bin/bash

# Generate 100 unique random numbers between -5000 and 5000
generate_random_numbers() {
    local count=$1
    local numbers=()
    
    while [ ${#numbers[@]} -lt $count ]; do
        # Generate random number between -5000 and 5000
        num=$((RANDOM % 10001 - 5000))
        
        # Check if number already exists
        if [[ ! " ${numbers[@]} " =~ " ${num} " ]]; then
            numbers+=($num)
        fi
    done
    
    echo "${numbers[@]}"
}

# Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo "========================================"
echo "100 ELEMENT TEST"
echo "========================================"
echo

# Generate 100 random numbers
echo -e "${BLUE}Generating 100 unique random numbers...${NC}"
ARG=$(generate_random_numbers 100)
echo -e "${GREEN}Generated!${NC}"
echo
echo "Sample (first 10): $(echo $ARG | cut -d' ' -f1-10)..."
echo

# Test push-swap with 100 elements
echo -e "${BLUE}Running push-swap with 100 elements...${NC}"
OPERATIONS=$(./push-swap "$ARG")
COUNT=$(echo "$OPERATIONS" | wc -l)

echo "Operations count: $COUNT"

if [ $COUNT -lt 700 ]; then
    echo -e "${GREEN}✓ PASS${NC}: Less than 700 operations ($COUNT)"
else
    echo -e "${RED}✗ FAIL${NC}: More than 700 operations ($COUNT)"
fi
echo

# Validate with checker
echo -e "${BLUE}Validating with checker...${NC}"
VALIDATION=$(echo "$OPERATIONS" | ./checker "$ARG")

if [ "$VALIDATION" = "OK" ]; then
    echo -e "${GREEN}✓ PASS${NC}: Checker validation OK"
else
    echo -e "${RED}✗ FAIL${NC}: Checker validation failed"
    echo "  Got: $VALIDATION"
fi
echo

# Summary
echo "========================================"
echo "SUMMARY - 100 Element Test"
echo "========================================"
echo "Numbers: 100"
echo "Operations: $COUNT"
echo "Target: < 700"
echo "Status: $([ $COUNT -lt 700 ] && echo -e "${GREEN}PASS${NC}" || echo -e "${RED}FAIL${NC}")"
echo "Validation: $VALIDATION"
echo "========================================"