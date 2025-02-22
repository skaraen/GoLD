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

## Print out AST
```bash
./gold -ast benchmarks/<file_name>
```
