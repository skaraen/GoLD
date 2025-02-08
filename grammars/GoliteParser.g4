parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer;
}

program         : types declarations functions EOF                                                  ;       
types           : typedecl*                                                                         ;
typedecl        : TYPE ID STRUCT LBRACE fields RBRACE SEMICOLON                                     ;
fields          : decl SEMICOLON (decl SEMICOLON)*                                                  ;
decl            : ID type                                                                           ;
type            : INT | BOOL | ASTERIX ID                                                           ;
declarations    : declaration*                                                                      ;
declaration     : VAR ids type SEMICOLON                                                            ;
ids             : ID (COMMA ID)*                                                                    ;
functions       : (function)*                                                                       ;
function        : FUNC ID parameters returntype? LBRACE declarations statements RBRACE              ;
parameters      : LPAREN (decl (COMMA decl)*)? RPAREN                                               ;
returntype      : type                                                                              ;
statements      : statement*                                                                        ;
statement       : assignment | print | read | delete | conditional | loop | return | invocation     ;
block           : LBRACE statements RBRACE                                                          ;
delete          : DELETE expression SEMICOLON                                                       ;
read            : SCAN lvalue SEMICOLON                                                             ;
assignment      : lvalue ASSIGN expression SEMICOLON                                                ;
print           : PRINTF LPAREN STRING (COMMA expression)* RPAREN SEMICOLON                         ;
conditional     : IF LPAREN expression RPAREN block (ELSE block)?                                   ;
loop            : FOR LPAREN expression RPAREN block                                                ;
return          : RETURN expression SEMICOLON                                                       ;
invocation      : ID arguments SEMICOLON                                                            ;                                                                  
arguments       : LPAREN (expression (COMMA expression)*)? RPAREN                                   ;
lvalue          : ID (DOT ID)*                                                                      ;
expression      : boolterm (OR boolterm)*                                                           ;
boolterm        : equalterm (AND equalterm)*                                                        ;
equalterm       : relationterm ((EQUALS | NEQUALS) relationterm)*                                   ;
relationterm    : simpleterm ((GT | LT | GTE | LTE) simpleterm)*                                    ;
simpleterm      : term ((PLUS | MINUS) term)*                                                       ;
term            : unaryterm ((ASTERIX | FSLASH) unaryterm)*                                         ;
unaryterm       : NOT selectorterm | MINUS selectorterm | selectorterm                              ;
selectorterm    : factor (DOT ID)*                                                                  ;
factor          : LPAREN expression RPAREN | ID (arguments)? | NUM | NEW ID | TRUE | FALSE | NIL    ;