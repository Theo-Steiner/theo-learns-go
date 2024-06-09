## `Sync`

## Wait group

- allows you to synchronize async activity
- initialization: `var wg sync.WaitGroup` (initial state `0`)
- a wait group has 3 important methods
  - `wg.Add(n)` -> increase the state for things to wait for by `n`
  - `wg.Done()` -> decrease the state for things to wait for by `1`
  - `wg.Wait()` -> block the thread until the state for things to wait for reaches `0`

Example:

```go
wantedCount := 1000
counter := Counter{}
// define a wait group
var wg sync.WaitGroup
// add a `delta` of how many times we expect wg.Done() to be called
wg.Add(wantedCount)

for i := 0; i < wantedCount; i++ {
    go func() {
        counter.Inc()
        // decrements the `delta` of our wait group by 1
        wg.Done()
    }()
}
// block until the wait group `delta` is 0
wg.Wait()
```

## Mutex

- allows you to ensure that an operation is executed only once at a time across threads
- initialization: `var mu sync.Mutex` (initial state is unlocked mutex)
- two main methods:
  - `mu.Lock()` -> locks a mutex
  - `mu.Unlock()` -> locks a mutex

Example:

```go
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Count) Inc () {
    mu.Lock() // <- lock the mutex (blocks if the mutex is already locked until it becomes available, then locks it)
    defer mu.Unlock() // <- unlock the mutex once the function returns
}
```

DO NOT EMBED MUTEX INTO STRUCTS (embedding means making the mutex a part of the struct like so:)

```go
type BadCounter struct {
    // this is bad, because mutex methods become part of the interface BadCounter
    // users can now call badCounter.Lock() from the outside, allowing tight coupling with internals
    sync.Mutex
    badCount int
}
```

DO NOT MAKE COPIES OF MUTEXES (instead pass the containing struct by reference)

## misc

- `go vet` to find bugs
- the author's convention for returning a pointer to a struct

```go
func NewCounter () *Counter {
  return &Counter{}
}
```
