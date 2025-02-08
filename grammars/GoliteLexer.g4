lexer grammar GoliteLexer; 

TYPE: 'type';

VAR: 'var';

INT: 'int';

BOOL: 'bool';

FUNC: 'func';

STRUCT: 'struct';

DELETE: 'delete';

PRINTF: 'printf';

SCAN: 'scan';

IF: 'if';

ELSE: 'else';

FOR: 'for';

RETURN: 'return';

NEW: 'new';

TRUE: 'true';

FALSE: 'false';

NIL: 'nil';

EQUALS: '==';

NEQUALS: '!=';

GTE: '>=';

LTE: '<=';

OR: '||';

AND: '&&';

GT: '>';

LT: '<';

ASSIGN: '=';

PLUS: '+';

MINUS: '-';

ASTERIX: '*';

FSLASH: '/';

LPAREN: '(';

RPAREN: ')';

LBRACE: '{';

RBRACE: '}';

SEMICOLON: ';';

COMMA: ',';

DOT: '.';

NOT: '!';

STRING: '"' .*? '"';

NUM: '0' | [1-9][0-9]*;

ID: [a-zA-Z][a-zA-Z0-9]*;

WS : [ \t\r\n]+ -> skip;