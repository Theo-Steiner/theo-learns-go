## What I learned/noticed:

### Syntax

#### Float

- float64 -> 64 bit float
- format strings:
  - `%.2f` for floats with 2 digit precision
  - `%g` for floats with high precision


#### Structs

- a named collection of fields
- can be printed with `fmt.Printf("%#v", myStruct)`

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
Interfaces allow you to defined functions that can be used by different types (parametric polymorphism)


### Testing

#### Table driven tests

```go
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.want {
			t.Errorf("%#v, got %.2f want %.2f", tt.shape, got, tt.want)
		}
	}
}
```
- areaTests slice is declared with an "anonymous struct""
- tt is short for table test
