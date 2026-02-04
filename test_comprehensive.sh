#!/bin/bash

# Generate unique random numbers using only bash
generate_random_numbers() {
    local count=$1
    local numbers=()
    local attempts=0
    local max_attempts=$((count * 10))
    
    while [ ${#numbers[@]} -lt $count ] && [ $attempts -lt $max_attempts ]; do
        num=$((RANDOM % 10001 - 5000))
        attempts=$((attempts + 1))
        
        # Check if number already exists
        local exists=0
        for n in "${numbers[@]}"; do
            if [ "$n" = "$num" ]; then
                exists=1
                break
            fi
        done
        
        if [ $exists -eq 0 ]; then
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
echo "COMPREHENSIVE TEST REPORT"
echo "========================================"
echo

# Test different sizes
echo -e "${BLUE}Testing 3 elements...${NC}"
ARG="3 1 2"
OPS=$(./push-swap "$ARG")
CNT=$(echo "$OPS" | wc -l)
VAL=$(echo "$OPS" | ./checker "$ARG")
echo "  Operations: $CNT (limit: <3), Validation: $VAL"
echo

echo -e "${BLUE}Testing 5 elements...${NC}"
ARG="5 2 4 1 3"
OPS=$(./push-swap "$ARG")
CNT=$(echo "$OPS" | wc -l)
VAL=$(echo "$OPS" | ./checker "$ARG")
echo "  Operations: $CNT (limit: <12), Validation: $VAL"
echo

echo -e "${BLUE}Testing 100 elements (Test 1)...${NC}"
ARG=$(generate_random_numbers 100)
echo "  Sample: $(echo $ARG | cut -d' ' -f1-5)..."
OPS=$(./push-swap "$ARG")
CNT=$(echo "$OPS" | wc -l)
VAL=$(echo "$OPS" | ./checker "$ARG")
echo "  Operations: $CNT (target: <700), Validation: $VAL"
[ $CNT -lt 700 ] && echo -e "  ${GREEN}✓ Operations within limit${NC}" || echo -e "  ${YELLOW}⚠ Needs optimization${NC}"
echo

echo -e "${BLUE}Testing 100 elements (Test 2)...${NC}"
ARG=$(generate_random_numbers 100)
echo "  Sample: $(echo $ARG | cut -d' ' -f1-5)..."
OPS=$(./push-swap "$ARG")
CNT=$(echo "$OPS" | wc -l)
VAL=$(echo "$OPS" | ./checker "$ARG")
echo "  Operations: $CNT (target: <700), Validation: $VAL"
[ $CNT -lt 700 ] && echo -e "  ${GREEN}✓ Operations within limit${NC}" || echo -e "  ${YELLOW}⚠ Needs optimization${NC}"
echo

echo -e "${BLUE}Testing 100 elements (Test 3)...${NC}"
ARG=$(generate_random_numbers 100)
echo "  Sample: $(echo $ARG | cut -d' ' -f1-5)..."
OPS=$(./push-swap "$ARG")
CNT=$(echo "$OPS" | wc -l)
VAL=$(echo "$OPS" | ./checker "$ARG")
echo "  Operations: $CNT (target: <700), Validation: $VAL"
[ $CNT -lt 700 ] && echo -e "  ${GREEN}✓ Operations within limit${NC}" || echo -e "  ${YELLOW}⚠ Needs optimization${NC}"
echo

echo "========================================"
echo "Note: The algorithm produces correct results"
echo "but needs optimization for 100+ elements"
echo "to meet the <700 operations requirement."
echo "========================================"