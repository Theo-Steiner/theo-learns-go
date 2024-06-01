## What I noticed

### multiline strings

```go
multiline := `backtick for
multiline string
yay!`
```

### stdlib misc

- sleep for 5 seconds: `time.Sleep(5 * time.Second)`

## Testing / Mocking

Issues with mocks usually stem from testing implementation details instead of behavior.
always test *expected behavior*, not implementation details

- rule of thumb: if a test uses more than 3 mocks the design might be overly convoluted
- spies create a tight coupling between test code and implementation. Only spy on things you actually care about

### Spying

injecting mocks to spy if certain details are implemented correctly
