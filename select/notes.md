### Things I noticed

## syntax

### defer

- prefixing a function call with `defer` will call that function at the end of the containing functions
  - good for cleanup like closing a server / file
  - `defer` is written close to the initial invocation, but executed at the end
  - being close is better because that gives it context

### select

- syntax to wait on multiple channels and "select" the first channel that receives a message

```go
// select takes whichever value comes first
select {
// ping() is implemented so that it returns a channel,
// to which it will send a message as soon as it get's a response
case <-ping(url):
    return a, nil
// time.After returns a channel that will receive a message after the timeout
case <-time.After(timeout):
    return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
}
```

## standard mock library for testing http stuff

- go provides a std library one can mock http servers/ requests with: `httptest`

```go
httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    time.Sleep(delay)
    w.WriteHeader(http.StatusOK)
}))
```
