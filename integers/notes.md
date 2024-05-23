## What I learned/noticed:

### Syntax

- for integers there is type `int`
- multiple arguments can be typed in one go `func a(a, b int)`

### Testing

```go
// adder_test.go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

```

- a `*_test.go` file can hold examples that
  - An example for a function has to be named `Example{function_name}`
  - fail a test if the Output is other than expected
    - expected output is defined via a code comment and then matched with stdout:
    `// Output: 6`
  - Examples are also automatically added to `godoc`

### Documentation

godoc can be installed via: `go install golang.org/x/tools/cmd/godoc@latest`
To view the docs in the browser run: `godoc -http=:8000` to have godoc run on localhost:8000

both examples and code comments are automatically displayed in godoc
