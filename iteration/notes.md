## What I learned/noticed:

### Syntax

- for loop: `for i := 0; i < 5; i++ {}`
- `+=` assignment
- `var` keyword (as opposed to const)

### Testing

#### Benchmarking

```go
// repeat_test.go
func BenchmarkRepeat(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Repeat("a")
  }
}
```

to run: `go test ./iteration -bench="."`
this will tell you how many ns per ops your function takes
