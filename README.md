# goscangen

Generate source code for a tokenizer from a given configuration, designed with simplicity, performance and readability in mind

## Usage

### 1. Create a file postfixed with `.scangen`:

```
$ cat ini.scangen
ignore: ;.*$
braket_left: \[
braket_right: \]
ident: [\a _ \d]*
equal: =
```

### 2. Pass the file to `scangen` via the `-i` parameter, provide an output path via `-o` and a package with `-p`

```sh
$ scangen -i examples/ini.scangen -o examples/lexer.go -p main
```

### 3. Call the Lexer:

```go
// examples/main.go
package main

var test = `
; comment
[Section]
1key123=value_12
`

func main() {
    lexer := Lexer{Builder: &strings.Builder{}}
    err := lexer.NewInput(test)
    if err != nil {
        panic(err)
    }
    tok, err := lexer.Lex()
    if err != nil {
        panic(err)
    }
    Debug(tok)
}
```

Running the above results in the following debug output:

```
$ cd examples/
$ go run .
| line | pos |                 type |                                                raw |
|    - |   - |                    - |                                                  - |
| 0002 | 001 |          braket_left |                                                    |
| 0002 | 002 |                ident |                                            Section |
| 0002 | 009 |         braket_right |                                                    |
| 0003 | 001 |                ident |                                            1key123 |
| 0003 | 008 |                equal |                                                    |
| 0003 | 009 |                ident |                                           value_12 |
```

## Syntax:

### Special characters:

| Symbol | Description                                      | Example                                                         |
| ------ | ------------------------------------------------ | --------------------------------------------------------------- |
| `$`    | Match EOL, depending on OS either `\n` or `\r\n` | `test$` → match all text starting with test and ending with EOL |
| `.`    | Match any character                              | `.` → matches anything                                          |
| `*`    | Match the symbol before multiple times           | `//.*$ `→ match anything between `//` and the `EOL`             |
| `[]`   | Match any of the instructions between            | `[abc]`, matches either a, b or c                               |
| `\d`   | Match a digit character                          | `\d` → matches a character if either 0-9, \_, e or .            |
| `\a`   | Match a letter character                         | `\a` → matches a character between a-z or A-Z                   |

### Example for the ini format:

Scangen definition:

```
ignore: ;.*$
braket_left: \[
braket_right: \]
ident: [\a _ \d]*
equal: =
```

Example ini file:

```ini
; comment
[Section]
1key123=value_12
```
