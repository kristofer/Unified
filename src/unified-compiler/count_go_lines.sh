#!/bin/bash

# Check if a directory is provided
if [ $# -eq 0 ]; then
    # Use current directory if none provided
    DIR="."
else
    DIR="$1"
fi

# Check if the provided path is a directory
if [ ! -d "$DIR" ]; then
    echo "Error: '$DIR' is not a directory"
    exit 1
fi

echo "Counting lines in Go files under directory: $DIR"

# Find all .go files and count their lines
TOTAL_FILES=0
TOTAL_LINES=0

while IFS= read -r file; do
    if [ -f "$file" ]; then
        LINES=$(wc -l < "$file")
        echo "$file: $LINES lines"
        TOTAL_LINES=$((TOTAL_LINES + LINES))
        TOTAL_FILES=$((TOTAL_FILES + 1))
    fi
done < <(find "$DIR" -type f -name "*.go")

echo "----------------------"
echo "Total Go files: $TOTAL_FILES"
echo "Total lines: $TOTAL_LINES"