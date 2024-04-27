grammar VanLang;

// Define tokens
VAR         : 'var' ;
FUNC        : 'func' ;
BEGIN       : 'begin' ;
END         : 'end' ;
IF          : 'if' ;
THEN        : 'then' ;
ELSEIF      : 'elseif' ;
ELSE        : 'else' ;
RETURN      : 'return' ;
PRINTLN     : 'println' ;
READLN      : 'readln' ;
REPEAT      : 'repeat' ;
UNTIL       : 'until' ;
WHILE       : 'while' ;
DO          : 'do' ;
NOT         : 'not' ;
AND         : 'and' ;
OR          : 'or' ;
XOR         : 'xor' ;
IMP         : 'imp' ;
EQV         : 'eqv' ;
DIV         : 'div' ;
MOD         : 'mod' ;
BOOL        : 'bool' ;
INTTYPE     : 'int' ;
STRTYPE     : 'string' ;
TUPLETYPE   : 'tuple' ;
ARRAYTYPE   : 'array' ;
FLOAT       : 'float' ;
UNSIGNED    : 'unsigned' ;

// Define operators
PLUS        : '+' ;
MINUS       : '-' ;
MULTIPLY    : '*' ;
DIVIDE      : '/' ;
EQUAL       : '=' ;
NOTEQUAL    : '<>' ;
LESS        : '<' ;
GREATER     : '>' ;
LESSEQUAL   : '<=' ;
GREATEREQUAL: '>=' ;
AND_OP: AND;
OR_OP: OR;
XOR_OP: XOR;
EQV_OP: EQV;
IMP_OP: IMP;
NOT_OP: NOT;

// Define tokens
ID          : [a-zA-Z]+ ;        // Variable names and function names
INT         : [0-9]+ ;            // Integer literals
STRING      : '"' .*? '"' ;       // String literals
LPAREN      : '(' ;
RPAREN      : ')' ;
LBRACK      : '[' ;
RBRACK      : ']' ;
COMMA       : ',' ;
SEMICOLON   : ';' ;
ARROW       : '->' ;

// Define rules
program     : statement* ;

statement   : variableDeclaration
            | functionDeclaration
            | logicStatement
            | loopStatement
            | conditionStatement
            | expressionStatement
            ;

variableDeclaration : VAR ID LBRACK datatype RBRACK (COMMA ID LBRACK datatype RBRACK)* SEMICOLON ;

functionDeclaration : FUNC ID LPAREN argumentList? RPAREN ARROW datatype SEMICOLON ;

argumentList : ID LBRACK datatype RBRACK (COMMA ID LBRACK datatype RBRACK)* ;

datatype    : INTTYPE
            | STRTYPE
            | ARRAYTYPE LPAREN datatype (COMMA datatype)* RPAREN
            | TUPLETYPE LPAREN datatype (COMMA datatype)* RPAREN
            | BOOL
            | FLOAT
            | UNSIGNED
            ;

logicStatement : ID EQUAL expression SEMICOLON ;

loopStatement : REPEAT statement* UNTIL expression SEMICOLON
              | WHILE expression DO BEGIN statement* END SEMICOLON
              ;

conditionStatement : IF expression THEN statement (ELSEIF expression THEN statement)* (ELSE statement)? SEMICOLON ;

expressionStatement : PRINTLN LPAREN STRING (COMMA expression)* RPAREN SEMICOLON
                     | ID EQUAL READLN LPAREN STRING RPAREN SEMICOLON
                     | RETURN expression SEMICOLON
                     ;

expression  : term ((AND_OP|OR_OP|XOR_OP|IMP_OP|EQV_OP) term)* SEMICOLON ;

term        : factor ((MULTIPLY|DIVIDE|DIV|MOD) factor | (EQUAL|NOTEQUAL|LESS|GREATER|LESSEQUAL|GREATEREQUAL) factor)* ;

factor      : atom ((MULTIPLY|DIVIDE|DIV|MOD) atom | (EQUAL|NOTEQUAL|LESS|GREATER|LESSEQUAL|GREATEREQUAL) atom)* ;

atom        : ID
            | INT
            | LPAREN expression RPAREN
            ;

// Ignored characters
WS          : [ \t\r\n]+ -> skip ;
