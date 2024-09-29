# Isigo

Isigo is a simple compiler in Go that converts your grammar into an executable binary.

The Isigo Language grammar is based on suggestions from the compilers course taught by Prof. Dr. Francisco Isidro (Federal University of ABC).

The initial suggestion was adapted by Kayky de Brito and Igor Mozetic to accommodate the presented requirements.

Isigo also implements an Interpreter (Runtime) that allows the execution of .isi scripts and a REPL that allows the execution of command lines in Isigo, useful for language development.

> 🇧🇷 LEIA A DOCUMENTAÇÃO EM PORTUGUÊS:
> [Documentação do Isigo [PT-BR]](https://www.notion.so/Documenta-o-do-Isigo-PT-BR-af7044b07f404e03b023ed27c3ae01a5?pvs=21)

## Language Features

- Compiler for executable binaries, based on Golang
- Interpreter/runtime for executing Isigo scripts
- Variables of types `inteiro`, `decimal`, `texto`
- Conditional structures `se` and `senao`
- Control structures `enquanto`/ `faca...enquanto`
- Arithmetic operations and structured expressions
- Value assignments to symbols
- Input and output operations `leia` and `escreva`
- Error checks during compilation time

# Using the compiler

The Isigo compiler works by receiving a source code file in Isigo, performing its lexical, syntactic, and semantic analysis, converting it to an AST.

This generated AST can be evaluated in the Isigo runtime or transformed into intermediate Go code. This intermediate code is then compiled into an executable binary by the Golang compiler.

### Hello world

Let's test Isigo in practice.

Make sure you have Go `>1.22.0` installed and have cloned this repository to your device.

In the root directory of Isigo, create a `hello_world.isi` file that you want to compile. We'll do the classic Hello world example:

```jsx
programa
    escreva("Ola, mundo!").
fimprog.
```

And we'll compile it to a binary using:

```jsx
go run . -c=build -f=hello_world.isi
```

You'll see a new hidden folder in the root directory: `./.isi_output`. This new directory will have the following structure:

```jsx
.isi_output
|- std
|- go.mod
|- isigoprogram
|- main.go
```

- `std` is a directory for a Go module with Isigo's standard functions, such as `leia` and `escreva`
- `go.mod` is the declaration of the generated program for the Go compiler.
- `main.go` is the file with the compiled code
- `isigoprogram` is an executable binary

Now you can run the `isigoprogram` program with `./isi_output/isigoprogram` for example.

When running the program in our terminal, we should see something like:

```
$ ./isi_output/isigoprogram
Ola, mundo!
```

> ⚠️ You may need to change the binary mode to executable depending on your operating system's security rules:
> 
> ```chmod +x ./isi_output/isigoprogram```

### Running Hello world in the Runtime

You can also skip the compilation step and run an Isigo program through the interpreter:

```
$ go run . -c=run -f=hello_world.isi
Ola, mundo!
->  <nil>
[ Programa encerrado. ]
```

## REPL

You can test Isigo instructions using the REPL:

```
$ go run. -c=repl
ISIGO REPL
isigo> escreva(1 + 5.4).
6.4
🟢 <nil>
isigo> 
```

> ⚠️ The REPL only accepts one command line at a time at the moment!

# Grammar

The final grammar of the Isigo Language is described by:

Prog → `programa` Block `fimprog.`

Block → VariableContext

VariableContext →  DeclareContext | ExecutionContext

DeclareContext → Declare VariableContext

ExecutionContext → (DeclareOrCommand)+

Declare → `declare` Id Type (, Id Type)* `.`

DeclareOrCommand → VariableContext | Command

Command → IfCommand | WhileCommand | DoWhileCommand | ReadCommand | WriteCommand | AssignmentCommand

IfCommand →   `se`  `(`  RelationalExpr `)` `{`  Block  `}`  (ElseCommand)?

ElseCommand → `senao`  `{`  Block `}`

WhileCommand → `enquanto` `(` RelationalExpr `)` `{` Block `}`

DoWhile → `faca` `{` Block `}` `enquanto` `(` RelationalExpr `)`

ReadCommand → `leia` `(` Id `)` `.`

WriteCommand → `enquanto` `(` Expr `)` `.`

AssignmentCommand → Id `:=` Expr `.`

Expr → Term ExprAux | Term

ExprAux → `+` Term (ExprAux)? | `-` Term (ExprAux)?

Term → Factor TermAux | Factor

TermAux → `*` Factor (TermAux)? | `/` Factor (TermAux)?

Factor → Num | Id | `(` Expr `)`

RelationalExpr → Expr Op_rel Expr

Op_rel → `<` |`>` |`<=` |`>=` |`!=` |`==`

Text → `"` (0..9 | a..z | A..Z | ' ' )+ `"`

Num → Integer | Decimal

Integer → (0…9)+

Decimal → Integer `.` Decimal

Id → (a..z | A..Z)(a..z | A..Z | 0..9)*

# Features

## Program

Every compiled Isigo program must be wrapped in `programa` and `fimprog.`

```
programa
		...
fimprog.
```

When compiled, the `programa` declaration is replaced by the import of Isigo's standard library.

## Variable Declaration

Variable declarations can happen at any time in the program.

A variable must always be declared with its respective type.

A variable cannot be used before being declared and initialized with a value.

To declare a variable:

```
declare nome1 tipo, nome2 tipo.
```

You can declare as many variables as you want in a single `declare`.

Type must be one of the existing types in Isigo: `inteiro` `decimal` or `texto`.

Example:

```
declara a inteiro, b decimal, c texto.
```

## End of Sentence

After a Declare, Read, Write, or Assignment command, it is necessary to insert the end of sentence character: `.`

```
declare a inteiro.
a := a + 1.
leia(a).
escreva(a).
```

## Assignment

Value assignment can be done using the `:=` operator and an expression on the right side to be assigned to the symbol on the left side.

Assignments can only be made to variables of the same type as the result of the expression or convertible without loss of precision.

```
programa
    declare a decimal.
    a := 40.
    a := a + 3.5.
    escreva(a).
fimprog.
```

## Expression

Expressions represent mathematical operations, respecting the order of operations:

```
programa
    escreva(8 / 2 * (2 + 2)).
fimprog.
```

## Reading

The `leia` command receives an identifier of a previously declared variable to assign the value that the user types in the terminal.

`leia` works for `inteiro`, `decimal`, and `texto` and automatically detects the type to be read from the standard input.

```
programa
    declare senha texto.
    leia(senha).
fimprog.
```

## Writing

The `escreva` command receives an expression, evaluates its value, and writes its result to the standard output.

```
programa
    declare senha texto.
    leia(senha).
    escreva(senha).
    
    escreva(50 + 3.4).
fimprog.
```

## If/Else

The if/else se/senao structure is written as follows:

```
programa
    declare senha texto, tentativa texto.
    senha := "segredo".
    leia(tentativa).
    
    se (senha == tentativa) {
        escreva("Parabens!").
    } senao {
        escreva("Tente novamente!").
    }
fimprog.
```

## While

The while (enquanto) structure is written as follows:

```
programa
    declare i inteiro.
    i := 0.
    enquanto (i < 4) {
        escreva(i).
        i := i + 1.
    }
fimprog.
```

### Do/While

Alternatively, the `faca` `enquanto` form can be used, which executes the block first and then checks the condition:

```
programa
    declare i inteiro.
    i := 0.
    faca {
        escreva(i).
        i := i + 1.
    } enquanto (i < 4)
fimprog.
```
