## What I learned/noticed:

### Concurrency

#### Goroutines

- functions can be run concurrently using the `go` keyword

```go
for _, url := range urls {
  go func(u string) {
    fetchData(u)
  }(url)
}
```

- go routines are commonly written with inline anonymous functions (which inherit the lexical scope)
- however, there are a few gotcha's with concurrent execution
  - when in a for loop, loop variables from the outer scope can be overwritten before the inner scope accesses them.
    - this makes it necessary to pass changing variables in as arguments to the anonymous goroutine function
  - writing to outputs can result in race conditions
    - to avoid them we can use `channels`
    - to test for race conditions run `go test -race`

#### Channels

- channels allow you to coordinate the communication between multiple processes
- channels are created like so: `myChannel := make(chan string)` -> a channel that takes in strings
  - to send to a channel `myChannel <- "hello"` (this sends a single `"hello"` to the channel, almost like pushing to an array)
  - to receive from a channel `result <- myChannel` ("result" receives a single result from a channel, almost like popping from an array)

```go
type result struct {
  url string
  data string
}
dataByUrl := make(map[string]string)
dataResult := make(chan result)
for _, url := range urls {
  go func(u string) {
    data := fetchData(u)
    dataResult <- result{u, data}
  }(url)
}
for _, _ := range urls {
  result := <- dataResult
  dataByUrl[result.url] = result.data
}
```
