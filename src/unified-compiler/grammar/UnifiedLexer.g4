lexer grammar UnifiedLexer;

// Brackets, braces and parentheses
LBRACE      : '{';
RBRACE      : '}';
LPAREN      : '(';
RPAREN      : ')';
LBRACK      : '[';
RBRACK      : ']';

// Punctuation
SEMI        : ';';
COLON       : ':';
COMMA       : ',';
DOT         : '.';
ARROW       : '->';
DOUBLE_COLON: '::';

// Arithmetic operators
PLUS        : '+';
MINUS       : '-';
STAR        : '*';
DIV         : '/';
MOD         : '%';

// Bitwise operators
BIT_AND     : '&';
BIT_OR      : '|';
BIT_XOR     : '^';
BIT_NOT     : '~';
LSHIFT      : '<<';
RSHIFT      : '>>';

// Logical operators
AND         : '&&';
OR          : '||';
NOT         : '!';

// Comparison operators
EQ          : '==';
NE          : '!=';
LT          : '<';
GT          : '>';
LE          : '<=';
GE          : '>=';

// Assignment operators
ASSIGN      : '=';
PLUS_ASSIGN : '+=';
MINUS_ASSIGN: '-=';
STAR_ASSIGN : '*=';
DIV_ASSIGN  : '/=';
MOD_ASSIGN  : '%=';
LSHIFT_ASSIGN: '<<=';
RSHIFT_ASSIGN: '>>=';
AND_ASSIGN  : '&=';
XOR_ASSIGN  : '^=';
OR_ASSIGN   : '|=';

// Other operators
QUESTION    : '?';
QUESTION_QUESTION : '??';  // Null coalescing
NULL_COND   : '?.';
RANGE       : '..';
RANGE_INCL  : '..=';
INC         : '++';
DEC         : '--';
UNDERSCORE  : '_';

// Lexer rules
MODULE      : 'module' | 'mod';
IMPORT      : 'import';
AS          : 'as';
FN          : 'fn';
STRUCT      : 'struct';
ENUM        : 'enum';
INTERFACE   : 'interface';
IMPL        : 'impl';
ACTOR       : 'actor';
TYPE        : 'type';
REFINE      : 'refine';
CONST       : 'const';
LET         : 'let';
VAR         : 'var';
MUT         : 'mut';
PUB         : 'pub';
SELF        : 'self' | 'Self';
IF          : 'if';
ELSE        : 'else';
LOOP        : 'loop';
WHILE       : 'while';
FOR         : 'for';
IN          : 'in';
BREAK       : 'break';
CONTINUE    : 'continue';
RETURN      : 'return';
SWITCH      : 'switch';
CASE        : 'case';
WHERE       : 'where';
MOVE        : 'move';
AWAIT       : 'await';
ASYNC       : 'async';
TRY         : 'try';
REGION      : 'region';
NEW         : 'new';

// Literals (must come before Identifier to take precedence)
BoolLiteral : 'true' | 'false';
NullLiteral : 'null' | 'nil';

Identifier  : [a-zA-Z_][a-zA-Z0-9_]*;

IntLiteral  : [0-9]+ | '0x' [0-9a-fA-F]+ | '0b' [01]+;
FloatLiteral: [0-9]+ '.' [0-9]+ ([eE] [+-]? [0-9]+)?  // Requires digits after decimal
            | [0-9]+ [eE] [+-]? [0-9]+                 // Scientific notation without decimal
            ;
StringLiteral: '"' (EscapeSequence | ~["\\\r\n])* '"';
CharLiteral : '\'' (EscapeSequence | ~['\\\r\n]) '\'';

//StringLiteral: '"' ( EscapeSequence | ~["\\\r\n] | Interpolation )* '"';
EscapeSequence
    : '\\' [btnfr"'\\]
 //   | '\\' [0-7]{1,3}
 //   | '\\' 'u' [0-9a-fA-F]{4}
    ;

// Skips and comments
WS          : [ \t\r\n\f]+ -> skip;
COMMENT     : '//' ~[\r\n]* -> skip;
MULTILINE_COMMENT: '/*' .*? '*/' -> skip;



//Interpolation
//    : '${' -> pushMode(INTERPOLATION_MODE)
//    ;

//mode INTERPOLATION_MODE;
//    INTERPOLATION_CLOSE: '}' -> popMode;
    // Rules for tokens allowed inside interpolation
//    INTERPOLATION_TEXT: ~[{}]+;
//    NESTED_OPEN: '{' -> pushMode(INTERPOLATION_MODE);
