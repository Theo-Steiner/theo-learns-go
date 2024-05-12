## What I learned/noticed:

- variable declaration is done using `:=` (colon, equal)
- function signature:
  1. `func` keyword
  2. function name (Pascal case seems to be the convention)
  3. paranthesis with arguments & argument type (separated by a space)
- modules are created with: `go mod init {module_name}`
- module main functions are run using `go run {file_name.go}`
- tests are run using `go test` inside a module
- inside string format `%s` inserts the string while `%q` inserts the string inside quotations
- ~go doc can be run locally `godoc -http :8000`~ on older versions of go

## Questions

- what is the asterisk in front of the argument in `(t *testing.T)`
