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
    : MODULE identifier '{' item* '}'
    ;

// Import declarations
importDecl
    : IMPORT importPath (AS identifier)? ';'
    | IMPORT importPath '{' importList '}' ';'
    ;

importPath
    : identifier ('.' identifier)*
    ;

importList
    : identifier (AS identifier)? (',' identifier (AS identifier)?)*
    | '*'
    ;

// Function declarations
functionDecl
    : (PUB)? FN identifier genericParams? '(' paramList? ')' ('->' type)? (whereClause)? block
    ;

paramList
    : parameter (',' parameter)*
    ;

parameter
    : identifier ':' type
    | SELF
    | SELF ':' '&' SELF
    | SELF ':' '&' MUT SELF
    ;

genericParams
    : '<' genericParam (',' genericParam)* '>'
    ;

genericParam
    : identifier (':' typeConstraint)?
    ;

typeConstraint
    : type ('+' type)*
    ;

whereClause
    : WHERE whereConstraint (',' whereConstraint)*
    ;

whereConstraint
    : type ':' typeConstraint
    ;

// Struct declarations
structDecl
    : (PUB)? STRUCT identifier genericParams? '{' structMember* '}'
    ;

structMember
    : (PUB)? identifier ':' type
    | functionDecl
    ;

// Enum declarations
enumDecl
    : (PUB)? ENUM identifier genericParams? '{' enumVariant (',' enumVariant)* '}'
    ;

enumVariant
    : identifier ('(' enumVariantParams ')')?
    ;

enumVariantParams
    : enumVariantParam (',' enumVariantParam)*
    ;

enumVariantParam
    : (identifier ':')? type
    ;

// Interface declarations
interfaceDecl
    : (PUB)? INTERFACE identifier genericParams? '{' interfaceMember* '}'
    ;

interfaceMember
    : functionSig
    | TYPE identifier ';'
    ;

functionSig
    : FN identifier genericParams? '(' paramList? ')' ('->' type)? ';'
    | FN identifier genericParams? '(' paramList? ')' ('->' type)? block
    ;

// Implementation declarations
implDecl
    : IMPL genericParams? type FOR type whereClause? '{' implMember* '}'
    | IMPL genericParams? type whereClause? '{' implMember* '}'
    ;

implMember
    : functionDecl
    | TYPE identifier '=' type ';'
    ;

// Actor declarations
actorDecl
    : (PUB)? ACTOR identifier genericParams? '{' actorMember* '}'
    ;

actorMember
    : (PUB)? VAR identifier ':' type ('=' expr)? ';'
    | functionDecl
    ;

// Type aliases
typeAlias
    : (PUB)? TYPE identifier genericParams? '=' type ';'
    | (PUB)? TYPE identifier genericParams? '=' type REFINE '|' identifier '|' expr ';'
    ;

// Constants
constantDecl
    : (PUB)? CONST identifier ':' type '=' expr ';'
    ;

// Types
type
    : identifier ('<' typeList '>')?                  // Simple type or generic type
    | FN '(' typeList? ')' '->' type                  // Function type
    | '(' typeList ')'                                // Tuple type
    | '&' type                                        // Reference type
    | '&' MUT type                                    // Mutable reference type
    | type '|' type                                   // Union type
    | IMPL identifier                                 // Interface type
    | type ('::' identifier)+                         // Qualified type
    ;

typeList
    : type (',' type)*
    ;

// Statements
statement
    : letStatement
    | varStatement
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
    | tryStatement
    ;

letStatement
    : LET MUT? identifier (':' type)? '=' expr ';'
    ;

varStatement
    : VAR identifier ':' type ('=' expr)? ';'
    ;

regionStatement
    : REGION identifier block
    ;

exprStatement
    : expr ';'
    ;

returnStatement
    : RETURN expr? ';'
    ;

ifStatement
    : IF expr block (ELSE IF expr block)* (ELSE block)?
    ;

loopStatement
    : (identifier ':')? LOOP block
    ;

whileStatement
    : (identifier ':')? WHILE expr block
    ;

forStatement
    : (identifier ':')? FOR identifier IN expr block
    ;

switchStatement
    : SWITCH expr '{' (caseClause)* '}'
    ;

caseClause
    : CASE pattern ('->' | ':') (statement | block)
    ;

breakStatement
    : BREAK identifier? ';'
    ;

continueStatement
    : CONTINUE identifier? ';'
    ;

blockStatement
    : block
    ;

tryStatement
    : TRY block
    ;

// Patterns
pattern
    : identifier                                      // Variable pattern
    | '_'                                             // Wildcard pattern
    | literal                                         // Literal pattern
    | pattern '..' pattern                            // Range pattern
    | identifier '(' patternList? ')'                 // Constructor pattern
    | LET identifier (':' type)? (IF expr)?           // Variable binding pattern
    | identifier '{' fieldPatternList '}'             // Struct pattern
    | '(' patternList? ')'                            // Tuple pattern
    ;

patternList
    : pattern (',' pattern)*
    ;

fieldPattern
    : identifier ':' pattern
    | identifier
    | '..'
    ;

fieldPatternList
    : fieldPattern (',' fieldPattern)*
    ;

// Expressions
expr
    : primary
    | expr '.' identifier                            // Field access
    | expr '.' identifier '(' argList? ')'           // Method call
    | expr '[' expr ']'                              // Index access
    | expr '(' argList? ')'                          // Function call
    | expr ('++' | '--')                             // Postfix inc/dec
    | ('++' | '--' | '+' | '-' | '!' | '~') expr     // Prefix operators
    | MOVE expr                                      // Move expression
    | AWAIT expr                                     // Await expression
    | expr ('*' | '/' | '%') expr                    // Multiplicative
    | expr ('+' | '-') expr                          // Additive
    | expr ('<<' | '>>') expr                        // Shift
    | expr ('<' | '<=' | '>' | '>=') expr            // Relational
    | expr ('==' | '!=') expr                        // Equality
    | expr '&' expr                                  // Bitwise AND
    | expr '^' expr                                  // Bitwise XOR
    | expr '|' expr                                  // Bitwise OR
    | expr '&&' expr                                 // Logical AND
    | expr '||' expr                                 // Logical OR
    | expr '?' expr ':' expr                         // Conditional
    | expr ('=' | '+=' | '-=' | '*=' | '/=' | '%=' 
          | '<<=' | '>>=' | '&=' | '^=' | '|=' | '..' 
          | '..=' | '?.') expr                       // Assignment and range
    | lambdaExpr                                     // Lambda expression
    | asyncExpr                                      // Async block
    | IF expr block (ELSE IF expr block)* (ELSE block)? // If expression
    | SWITCH expr '{' (caseExpr)* '}'                // Switch expression
    ;

caseExpr
    : CASE pattern ('->' | ':') expr
    ;

primary
    : identifier                                     // Variable reference
    | literal                                        // Literal
    | '(' expr ')'                                   // Grouping
    | block                                          // Block expression
    | structExpr                                     // Struct instantiation
    | listExpr                                       // List literal
    | mapExpr                                        // Map literal
    | setExpr                                        // Set literal
    | tupleExpr                                      // Tuple expression
    ;

lambdaExpr
    : ('move')? '|' paramList? '|' (('->' type)? block | expr)
    ;

asyncExpr
    : ASYNC block
    ;

structExpr
    : identifier '{' fieldInitList '}'
    ;

fieldInitList
    : fieldInit (',' fieldInit)* ','?
    ;

fieldInit
    : identifier ':' expr
    | identifier
    | '..' expr
    ;

listExpr
    : '[' (expr (',' expr)*)? ','? ']'
    | '[' expr FOR identifier IN expr (IF expr)? ']'  // List comprehension
    ;

mapExpr
    : '{' (keyValue (',' keyValue)*)? ','? '}'
    ;

keyValue
    : expr ':' expr
    ;

setExpr
    : '{' (expr (',' expr)*)? ','? '}'
    ;

tupleExpr
    : '(' (expr (',' expr)*)? ','? ')'
    ;

block
    : '{' statement* (expr)? '}'
    ;

argList
    : expr (',' expr)*
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

