#!/bin/bash
# Create GitHub Issues for Unified Compiler Phases
# This script creates issues for Phases 2-15 using the GitHub CLI

set -e  # Exit on error

# Check if gh CLI is installed
if ! command -v gh &> /dev/null; then
    echo "Error: GitHub CLI (gh) is not installed."
    echo "Install it from: https://cli.github.com/"
    echo ""
    echo "Alternatively, create issues manually using the templates in docs/issues/"
    exit 1
fi

# Check if authenticated
if ! gh auth status &> /dev/null; then
    echo "Error: Not authenticated with GitHub CLI."
    echo "Run: gh auth login"
    exit 1
fi

echo "Creating GitHub issues for Unified compiler phases..."
echo ""

# Array of phase definitions
# Format: "phase_number|title|labels|duration"
declare -a phases=(
    "2|Phase 2: Control Flow and Parser Grammar Updates|phase-2,enhancement,parser,vm,control-flow|1-2 weeks"
    "3|Phase 3: Variables, Mutability, and Assignment|phase-3,enhancement,parser,vm,semantic-analysis|2 weeks"
    "4|Phase 4: Functions and Advanced Expressions|phase-4,enhancement,parser,vm,functions|2-3 weeks"
    "5|Phase 5: Structs and Basic Types|phase-5,enhancement,parser,vm,types|3-4 weeks"
    "6|Phase 6: Enums and Pattern Matching|phase-6,enhancement,parser,vm,pattern-matching|3-4 weeks"
    "7|Phase 7: Error Handling with ? Operator|phase-7,enhancement,parser,vm,error-handling|1-2 weeks"
    "8|Phase 8: Arrays and Slices|phase-8,enhancement,parser,vm,arrays|2-3 weeks"
    "9|Phase 9: String Operations|phase-9,enhancement,vm,stdlib,strings|2 weeks"
    "10|Phase 10: Generics Basics|phase-10,enhancement,parser,vm,generics|3-4 weeks"
    "11|Phase 11: Modules and Imports|phase-11,enhancement,parser,vm,modules|2-3 weeks"
    "12|Phase 12: Basic Ownership|phase-12,enhancement,semantic-analysis,ownership|3-4 weeks"
    "13|Phase 13: Standard Library Foundation|phase-13,enhancement,stdlib|2-3 weeks"
    "14|Phase 14: Concurrency Basics|phase-14,enhancement,parser,vm,concurrency|4-6 weeks"
    "15|Phase 15: Tooling and Polish|phase-15,enhancement,tooling|3-4 weeks"
)

# Counter for created issues
created=0
failed=0

# Create each issue
for phase_def in "${phases[@]}"; do
    IFS='|' read -r phase_num title labels duration <<< "$phase_def"
    
    # Construct filename
    filename="docs/issues/phase-$(printf "%02d" $phase_num)-*.md"
    filepath=$(ls $filename 2>/dev/null | head -1)
    
    if [ ! -f "$filepath" ]; then
        echo "⚠️  Warning: Template file not found for Phase $phase_num"
        echo "    Expected: $filename"
        ((failed++))
        continue
    fi
    
    echo "Creating issue for Phase $phase_num..."
    echo "  Title: $title"
    echo "  File: $filepath"
    #echo "  Labels: $labels"
    
    # Create the issue
    if gh issue create \
        --title "$title" \
        --body-file "$filepath" \
        2>/dev/null; then
        echo "✅ Created issue for Phase $phase_num"
        ((created++))
    else
        echo "❌ Failed to create issue for Phase $phase_num"
        ((failed++))
    fi
    
    echo ""
done

# Summary
echo "========================================="
echo "Issue Creation Summary"
echo "========================================="
echo "✅ Successfully created: $created issues"
if [ $failed -gt 0 ]; then
    echo "❌ Failed: $failed issues"
fi
echo ""

if [ $created -gt 0 ]; then
    echo "View all issues at:"
    echo "https://github.com/kristofer/Unified/issues"
    echo ""
    echo "Next steps:"
    echo "1. Review the created issues"
    echo "2. Create milestones for each phase"
    echo "3. Assign issues as needed"
    echo "4. Start working on Phase 2!"
fi

exit 0
