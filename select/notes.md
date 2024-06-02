### Things I noticed

## syntax

### defer

- prefixing a function call with `defer` will call that function at the end of the containing functions
  - good for cleanup like closing a server / file
  - `defer` is written close to the initial invocation, but executed at the end
  - being close is better because that gives it context

## mock library for testing http stuff

- go provides a std library one can mock http servers/ requests with: `httptest`

```go
httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
}))
```
