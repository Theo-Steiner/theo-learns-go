## What I learned/noticed:

### Syntax

#### Float

- float64 -> 64 bit float
- format strings:
  - `%.2f` for floats with 2 digit precision
  - `%g` for floats with high precision


#### Structs

- a named collection of fields

```go
type Rectangle struct {
  Width  float64
  Height float64
}

rectangle := Rectangle{12.0,6.0}
```

#### Methods

- a method is a special type of function that has a "receiver"
- "receiver" variable that a method can be called on
- receiver argument usually first letter of receiver type
  - Rectangle -> `func (r ReceiverType) Method() ReturnType {}`

```go
func (r Rectangle) Area() (area float64) {
  area = r.Height * r.Width
  return
}

rectangle := Rectangle{10.0, 10.0}
rectangleArea := rectangle.Area() // 100
```

#### Interfaces

- An interface that defines what methods something needs to become an instance of that type
- Go has *implicit interface resolution*:
  -> If the type you pass in matches what the interface is asking for, it will compile

```go
type Shape interface {
  Area() float64
}
```

Interfaces should declare only what is needed, in order to *decouple* implementation details.


### Testing

#### Table driven tests
