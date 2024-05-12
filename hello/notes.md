## What I learned/noticed:

### Syntax

### vars
- variable declaration is done using `:=` (colon, equal)

#### functions

- function signature:
  1. `func` keyword
  2. function name (Pascal case seems to be the convention)
  3. paranthesis with arguments & argument type (separated by a space)

- named returns
  -  can be implicitly returned (just `return`)
  - automatically created with the `zero` value initially
  -> zero value: int = 0, string = ""
  - documented more nicely by godoc

#### consts

- `const` available for constants
- multiple const declarations can be grouped in a block
```go
const(
  a = 'a'
  b ='b'
)
```

#### misc

- switch block
  -  seems not to be indented

- tests are run using `go test` inside a module
- inside string format `%s` inserts the string while `%q` inserts the string inside quotations
- string concatenation can be done with `+`

### Testing

- modules are created with: `go mod init {module_name}`
- module main functions are run using `go run {file_name.go}`
- subtests can be written using t.Run

- t.Helper() -> tells the test suite that the called method is a helper method
  - helps show the correct line number

- *testing.T -> T in Test
- *testing.B -> B in Benchmark
- testing.TB -> Interface that is satisfied by testing.T & testing.B

### Other

- ~go doc can be run locally `godoc -http :8000`~ on older versions of go


## Questions

- what is the asterisk in front of the argument in `(t *testing.T)`

## TDD

- write test
- make compiler pass
- run test & check if error message is meaningful
- write code & make test pass
- refactor
