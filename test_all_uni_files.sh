#!/bin/bash

# Script to test all .uni files with the WASM compiler
# This script will run all .uni files and categorize them as working or failing
#
# macOS Note: This script uses the 'timeout' command which is not available on macOS by default.
# To use on macOS, install coreutils: brew install coreutils
# Then either:
#   - Use 'gtimeout' instead of 'timeout' (change line 45 below)
#   - Create an alias: alias timeout=gtimeout
#   - Or run without timeout by commenting out the timeout wrapper

COMPILER="./src/unified-compiler/bin/unified"
OUTPUT_FILE="test_results.txt"
WORKING_FILE="test_working.txt"
FAILING_FILE="test_failing.txt"

# Detect timeout command (Linux has 'timeout', macOS with coreutils has 'gtimeout')
if command -v timeout &> /dev/null; then
    TIMEOUT_CMD="timeout"
elif command -v gtimeout &> /dev/null; then
    TIMEOUT_CMD="gtimeout"
else
    echo "Warning: 'timeout' command not found. Tests will run without timeout protection."
    echo "On macOS, install with: brew install coreutils"
    TIMEOUT_CMD=""
fi

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
TOTAL=0
WORKING=0
FAILING=0

# Clear output files
> "$OUTPUT_FILE"
> "$WORKING_FILE"
> "$FAILING_FILE"

echo "======================================" | tee -a "$OUTPUT_FILE"
echo "Testing All .uni Files with WASM Backend" | tee -a "$OUTPUT_FILE"
echo "======================================" | tee -a "$OUTPUT_FILE"
echo "" | tee -a "$OUTPUT_FILE"

# Find all .uni files
UNI_FILES=$(find . -name "*.uni" -type f | sort)

for file in $UNI_FILES; do
    TOTAL=$((TOTAL + 1))
    echo -n "Testing: $file ... "
    
    # Run the compiler with a 5 second timeout and capture output
    if [ -n "$TIMEOUT_CMD" ]; then
        OUTPUT=$($TIMEOUT_CMD 5 $COMPILER --input "$file" 2>&1)
        EXIT_CODE=$?
    else
        # No timeout available - run without it
        OUTPUT=$($COMPILER --input "$file" 2>&1)
        EXIT_CODE=$?
    fi
    
    # Check for timeout (exit code 124 from timeout command)
    if [ -n "$TIMEOUT_CMD" ] && [ $EXIT_CODE -eq 124 ]; then
        echo -e "${YELLOW}TIMEOUT${NC}"
        echo "⏱️  $file (TIMEOUT)" >> "$FAILING_FILE"
        echo "⏱️  $file (TIMEOUT)" >> "$OUTPUT_FILE"
        echo "Test timed out after 5 seconds" >> "$OUTPUT_FILE"
        echo "" >> "$OUTPUT_FILE"
        FAILING=$((FAILING + 1))
    # Check if compilation was successful (look for "Program completed successfully")
    elif echo "$OUTPUT" | grep -q "Program completed successfully"; then
        echo -e "${GREEN}PASS${NC}"
        echo "✅ $file" >> "$WORKING_FILE"
        echo "✅ $file" >> "$OUTPUT_FILE"
        echo "$OUTPUT" >> "$OUTPUT_FILE"
        echo "" >> "$OUTPUT_FILE"
        WORKING=$((WORKING + 1))
    else
        echo -e "${RED}FAIL${NC}"
        echo "❌ $file" >> "$FAILING_FILE"
        echo "❌ $file" >> "$OUTPUT_FILE"
        echo "$OUTPUT" >> "$OUTPUT_FILE"
        echo "" >> "$OUTPUT_FILE"
        FAILING=$((FAILING + 1))
    fi
done

echo "" | tee -a "$OUTPUT_FILE"
echo "======================================" | tee -a "$OUTPUT_FILE"
echo "Summary" | tee -a "$OUTPUT_FILE"
echo "======================================" | tee -a "$OUTPUT_FILE"
echo "Total tests: $TOTAL" | tee -a "$OUTPUT_FILE"
echo -e "${GREEN}Working: $WORKING${NC}" | tee -a "$OUTPUT_FILE"
echo -e "${RED}Failing: $FAILING${NC}" | tee -a "$OUTPUT_FILE"
echo "Success rate: $(awk "BEGIN {printf \"%.1f\", ($WORKING/$TOTAL)*100}")%" | tee -a "$OUTPUT_FILE"
echo "" | tee -a "$OUTPUT_FILE"
echo "Results saved to:" | tee -a "$OUTPUT_FILE"
echo "  - Full results: $OUTPUT_FILE" | tee -a "$OUTPUT_FILE"
echo "  - Working tests: $WORKING_FILE" | tee -a "$OUTPUT_FILE"
echo "  - Failing tests: $FAILING_FILE" | tee -a "$OUTPUT_FILE"
