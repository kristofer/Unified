parser grammar UnifiedParser;

options { tokenVocab=UnifiedLexer; }

program
    : item*
    ;

item
    : moduleDecl
    | functionDecl
    | structDecl
    | enumDecl
    | interfaceDecl
    | implDecl
    | actorDecl
    | typeAlias
    | importDecl
    | constantDecl
    ;

// Module declaration
moduleDecl
    : MODULE identifier LBRACE item* RBRACE
    ;

// Import declarations
importDecl
    : IMPORT importPath (AS identifier)? SEMI?
    | IMPORT importPath LBRACE importList RBRACE SEMI?
    ;

importPath
    : identifier (DOT identifier)*
    ;

importList
    : identifier (AS identifier)? (COMMA identifier (AS identifier)?)*
    | '*'
    ;

// Function declarations
functionDecl
    : (PUB)? FN identifier genericParams? LPAREN paramList? RPAREN (ARROW type)? (whereClause)? block
    ;

paramList
    : parameter (COMMA parameter)*
    ;

parameter
    : identifier COLON type
    | SELF
    | SELF COLON BIT_AND SELF
    | SELF COLON BIT_AND MUT SELF
    ;

genericParams
    : LT genericParam (COMMA genericParam)* GT
    ;

genericParam
    : identifier (COLON typeConstraint)?
    ;

typeConstraint
    : type (PLUS type)*
    ;

whereClause
    : WHERE whereConstraint (COMMA whereConstraint)*
    ;

whereConstraint
    : type COLON typeConstraint
    ;

// Struct declarations
structDecl
    : (PUB)? STRUCT identifier genericParams? LBRACE structMember* RBRACE
    ;

structMember
    : (PUB)? identifier COLON type
    | functionDecl
    ;

// Enum declarations
enumDecl
    : (PUB)? ENUM identifier genericParams? LBRACE enumVariant (COMMA enumVariant)* RBRACE
    ;

enumVariant
    : identifier (LPAREN enumVariantParams RPAREN)?
    ;

enumVariantParams
    : enumVariantParam (COMMA enumVariantParam)*
    ;

enumVariantParam
    : (identifier COLON)? type
    ;

// Interface declarations
interfaceDecl
    : (PUB)? INTERFACE identifier genericParams? LBRACE interfaceMember* RBRACE
    ;

interfaceMember
    : functionSig
    | TYPE identifier SEMI?
    ;

functionSig
    : FN identifier genericParams? LPAREN paramList? RPAREN (ARROW type)? SEMI?
    | FN identifier genericParams? LPAREN paramList? RPAREN (ARROW type)? block
    ;

// Implementation declarations
implDecl
    : IMPL genericParams? type FOR type whereClause? LBRACE implMember* RBRACE
    | IMPL genericParams? type whereClause? LBRACE implMember* RBRACE
    ;

implMember
    : functionDecl
    | TYPE identifier ASSIGN type SEMI?
    ;

// Actor declarations
actorDecl
    : (PUB)? ACTOR identifier genericParams? LBRACE actorMember* RBRACE
    ;

actorMember
    : (PUB)? VAR identifier COLON type (ASSIGN expr)? SEMI?
    | functionDecl
    ;

// Type aliases
typeAlias
    : (PUB)? TYPE identifier genericParams? ASSIGN type SEMI?
    | (PUB)? TYPE identifier genericParams? ASSIGN type REFINE BIT_OR identifier BIT_OR expr SEMI?
    ;

// Constants
constantDecl
    : (PUB)? CONST identifier COLON type ASSIGN expr SEMI?
    ;

// Types
type 
    : identifier (LT typeList GT)?                   // Simple type or generic type
    | FN LPAREN typeList? RPAREN ARROW type          // Function type
    | LPAREN typeList RPAREN                         // Tuple type
    | BIT_AND type                                   // Reference type
    | BIT_AND MUT type                               // Mutable reference type
    | type BIT_OR type                               // Union type
    | IMPL identifier                                // Interface type
    | type (DOUBLE_COLON identifier)+                // Qualified type
    | LBRACK type SEMI IntLiteral RBRACK             // Array type: [Int; 10]
    | LBRACK type RBRACK                             // Slice type: [Int]
    ;

typeList
    : type (COMMA type)*
    ;

// Statements
statement
    : letStatement
    | varStatement
    | assignmentStatement
    | regionStatement
    | exprStatement
    | returnStatement
    | ifStatement
    | loopStatement
    | whileStatement
    | forStatement
    | switchStatement
    | breakStatement
    | continueStatement
    | blockStatement
    ;

letStatement
    : LET MUT? identifier (COLON type)? ASSIGN expr SEMI?
    ;

varStatement
    : VAR identifier COLON type (ASSIGN expr)? SEMI?
    ;

assignmentStatement
    : identifier assignmentOp expr SEMI?
    ;

assignmentOp
    : ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | STAR_ASSIGN 
    | DIV_ASSIGN | MOD_ASSIGN
    ;

regionStatement
    : REGION identifier block
    ;

exprStatement
    : expr SEMI?
    ;

returnStatement
    : RETURN expr? SEMI?
    ;

ifStatement
    : IF expr block (ELSE IF expr block)* (ELSE block)?
    ;

loopStatement
    : (identifier COLON)? LOOP block
    ;

whileStatement
    : (identifier COLON)? WHILE expr block
    ;

forStatement
    : (identifier COLON)? FOR identifier IN expr block
    ;

switchStatement
    : SWITCH expr LBRACE (caseClause)* RBRACE
    ;

caseClause
    : CASE pattern (ARROW | COLON) (statement | block)
    ;

breakStatement
    : BREAK identifier? SEMI?
    ;

continueStatement
    : CONTINUE identifier? SEMI?
    ;

blockStatement
    : block
    ;

// Patterns
pattern
    : identifier                                      // Variable pattern
    | UNDERSCORE                                      // Wildcard pattern
    | literal                                         // Literal pattern
    | pattern RANGE pattern                           // Range pattern
    | identifier LPAREN patternList? RPAREN           // Constructor pattern
    | LET identifier (COLON type)? (IF expr)?         // Variable binding pattern
    | identifier LBRACE fieldPatternList RBRACE       // Struct pattern
    | LPAREN patternList? RPAREN                      // Tuple pattern
    ;

patternList
    : pattern (COMMA pattern)*
    ;

fieldPattern
    : identifier COLON pattern
    | identifier
    | RANGE
    ;

fieldPatternList
    : fieldPattern (COMMA fieldPattern)*
    ;

// Expressions
expr
    : primary
    | expr DOT  (identifier | IntLiteral)            // Field/index access
    | expr DOT identifier LPAREN argList? RPAREN     // Method call
    | expr LBRACK expr RBRACK                        // Index access
    | expr LPAREN argList? RPAREN                    // Function call
    | expr (INC | DEC)                               // Postfix inc/dec
    | (INC | DEC | PLUS | MINUS | NOT | BIT_NOT) expr  // Prefix operators
    | MOVE expr                                      // Move expression
    | AWAIT expr                                     // Await expression
    | expr (STAR | DIV | MOD) expr                   // Multiplicative
    | expr (PLUS | MINUS) expr                       // Additive
    | expr (LSHIFT | RSHIFT) expr                    // Shift
    | expr (LT | LE | GT | GE) expr                  // Relational
    | expr (EQ | NE) expr                            // Equality
    | expr BIT_AND expr                              // Bitwise AND
    | expr BIT_XOR expr                              // Bitwise XOR
    | expr BIT_OR expr                               // Bitwise OR
    | expr AND expr                                  // Logical AND
    | expr OR expr                                   // Logical OR
    | expr QUESTION                                   // Error propagation
    | expr QUESTION_QUESTION expr                     // Null coalescing
    | expr QUESTION expr COLON expr                   // Conditional
    | expr (ASSIGN | PLUS_ASSIGN | MINUS_ASSIGN | STAR_ASSIGN | DIV_ASSIGN | MOD_ASSIGN 
          | LSHIFT_ASSIGN | RSHIFT_ASSIGN | AND_ASSIGN | XOR_ASSIGN | OR_ASSIGN | RANGE 
          | RANGE_INCL | NULL_COND) expr             // Assignment and range
    | lambdaExpr                                     // Lambda expression
    | asyncExpr                                      // Async block
    | IF expr block (ELSE IF expr block)* (ELSE block)? // If expression
    | SWITCH expr LBRACE (caseExpr)* RBRACE          // Switch expression
    // In the expr rule, add null coalescing with appropriate precedence
    ;


caseExpr
    : CASE pattern (ARROW | COLON) expr
    ;

primary
    : constructorExpr                                // New expression (must be before structExpr)
    | structExpr                                     // Struct instantiation (must be before identifier)
    | listExpr                                       // List literal
    | mapExpr                                        // Map literal
    | setExpr                                        // Set literal
    | tupleExpr                                      // Tuple expression
    | identifier                                     // Variable reference
    | SELF                                           // Self reference
    | literal                                        // Literal
    | LPAREN expr RPAREN                             // Grouping
    | block                                          // Block expression
    ;

lambdaExpr
    : (MOVE)? BIT_OR paramList? BIT_OR ((ARROW type)? block | expr)
    ;

asyncExpr
    : ASYNC block
    ;

constructorExpr
    : NEW identifier (LT typeList GT)? LPAREN argList? RPAREN
    ;

structExpr
    : identifier LBRACE fieldInitList RBRACE
    ;

fieldInitList
    : fieldInit (COMMA fieldInit)* COMMA?
    ;

fieldInit
    : identifier COLON expr
    | identifier
    | RANGE expr
    ;

listExpr
    : LBRACK (expr (COMMA expr)*)? COMMA? RBRACK
    | LBRACK expr FOR identifier IN expr (IF expr)? RBRACK  // List comprehension
    ;

mapExpr
    : LBRACE (keyValue (COMMA keyValue)*)? COMMA? RBRACE
    ;

keyValue
    : expr COLON expr
    ;

setExpr
    : LBRACE (expr (COMMA expr)*)? COMMA? RBRACE
    ;

// tupleExpr
//     : LPAREN (expr (COMMA expr)*)? COMMA? RPAREN
//     ;

tupleExpr
    : LPAREN (namedTupleField (COMMA namedTupleField)*)? COMMA? RPAREN
    ;

namedTupleField
    : identifier COLON expr    // named field
    | expr                     // unnamed field
    ;

block
    : LBRACE statement* (expr)? RBRACE
    ;

argList
    : expr (COMMA expr)*
    ;

// Literals
literal
    : IntLiteral
    | FloatLiteral
    | StringLiteral
    | CharLiteral
    | BoolLiteral
    | NullLiteral
    ;

// Basic tokens
identifier
    : Identifier
    ;
