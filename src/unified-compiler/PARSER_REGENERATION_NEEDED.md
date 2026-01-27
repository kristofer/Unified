# Parser Regeneration Required

## Status

The grammar files have been updated to support the `new` keyword:
- ✅ `UnifiedLexer.g4` - Added `NEW : 'new';` keyword
- ✅ `UnifiedParser.g4` - Added `newExpr` rule for `new Type<T>(args)` syntax
- ⚠️  Parser code needs to be regenerated from grammar

## Why Parser Regeneration is Needed

The ANTLR-generated parser files in `internal/parser/` are auto-generated from the `.g4` grammar files. When we modify the grammar (which we did to add the `new` keyword), we must regenerate the parser to reflect those changes.

## How to Regenerate the Parser

### Option 1: Using ANTLR4 (Recommended)

1. Install ANTLR4:
   ```bash
   # On macOS with Homebrew
   brew install antlr
   
   # On Ubuntu/Debian
   sudo apt-get install antlr4
   
   # Or download the jar directly
   cd /usr/local/lib
   sudo curl -O https://www.antlr.org/download/antlr-4.13.2-complete.jar
   sudo chmod +x /usr/local/bin/antlr4
   ```

2. Run the generation script:
   ```bash
   cd src/unified-compiler
   make parser
   ```

### Option 2: Using Python antlr4-tools

1. Install antlr4-tools:
   ```bash
   pip install antlr4-tools
   ```

2. Generate the parser:
   ```bash
   cd src/unified-compiler/grammar
   antlr4 -Dlanguage=Go -package parser -visitor -listener -o ../internal/parser UnifiedParser.g4 UnifiedLexer.g4
   ```

### Option 3: Using Docker

If you have Docker installed:
```bash
cd src/unified-compiler
docker run --rm -v $(pwd):/work -w /work/grammar \
  openjdk:11 bash -c \
  "curl -O https://www.antlr.org/download/antlr-4.13.2-complete.jar && \
   java -jar antlr-4.13.2-complete.jar -Dlanguage=Go -package parser \
   -visitor -listener -o ../internal/parser UnifiedParser.g4 UnifiedLexer.g4"
```

## What Happens After Regeneration

Once the parser is regenerated, you'll need to:

1. **Update the AST Visitor** (`internal/ast/visitor.go`):
   - Add `VisitNewExpr()` method to handle the new expression type
   - Update `VisitPrimary()` to check for `newExpr` context

2. **Update Bytecode Generator** (`internal/bytecode/generator.go`):
   - Add handling for `*ast.NewExpr` in the expression generator
   - Generate appropriate bytecode for constructor calls

3. **Build and Test**:
   ```bash
   make build
   make test
   ./bin/unified --input test/new_keyword_basic.uni
   ```

## Current Limitations

Until the parser is regenerated:
- The compiler cannot parse the `new` keyword syntax
- Test files using `new Type()` will fail to compile
- All documentation has been updated to use the new syntax

## Files Modified

Grammar files updated:
- `grammar/UnifiedLexer.g4` - Line 105: Added `NEW : 'new';`
- `grammar/UnifiedParser.g4` - Lines 318-319, 340-342: Added `newExpr` rule

AST updated:
- `internal/ast/ast.go` - Lines 500-517: Added `NewExpr` struct

Test file created:
- `test/new_keyword_basic.uni` - Test cases for new keyword

Documentation updated (all occurrences of old syntax replaced):
- `spec/UnifiedSpecifiation.md`
- `docs/issues/*.md` (7 files)
- `lib/List.uni` and `lib/Tree.uni`
- `ClassicExercises.md`
- `PHASES_2_9_STATUS.md`
