## What I learned/noticed:

### Syntax

- variable declaration is done using `:=` (colon, equal)
- function signature:
  1. `func` keyword
  2. function name (Pascal case seems to be the convention)
  3. paranthesis with arguments & argument type (separated by a space)
- tests are run using `go test` inside a module
- inside string format `%s` inserts the string while `%q` inserts the string inside quotations
- string concatenation can be done with `+`
- `const` available for constants

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
