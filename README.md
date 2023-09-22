# goscangen

Generate source code for a tokenizer from a given configuration, designed with simplicity, performance and readability in mind

## Usage

Create a file postfixed with `.scangen`:

```
$ cat ini.scangen
ignore: ;.*$
braket_left: \[
braket_right: \]
ident: [\a _ \d]*$
equal: =
```

Pass the file to `scangen` via the `-i` parameter, provide an output path via `-o`

```sh
$ scangen -i ini.scangen -o IniLexer.go
```

Use the Lexer :)

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
ident: [\a _ \d]*$
equal: =
```

Example ini file:

```ini
; comment
[Section]
1key123=value_12
```
