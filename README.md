# GoLD - GoLite compiler

## Current status
Fully working semantic analysis. Working on LLVM IR.

## Build compiler
```bash
go build -o gold golite/golite.go
```

## Print out lexer tokens
```bash
./gold -l benchmarks/<file_name>
```
An example for dead code elimination is provided in benchmarks/deadcode.golite

## Print out lexer tokens
```bash
./gold -l benchmarks/<file_name>
```

## Print out AST tree from parser
```bash
./gold -ast benchmarks/<file_name>
```

## Generate LLVM IR (stack-based) 
```bash
./gold -llvm-ir=stack benchmarks/<file_name>
```
File saved as <file_name>_stack.ll

## Generate LLVM IR (register-based)
```bash
./gold -llvm-ir=reg benchmarks/<file_name>
```
File saved as <file_name>_reg.ll

## Generate ARM machine code file
```bash
./gold -S benchmarks/<file_name>
```
File saved as <file_name>_arm.s
