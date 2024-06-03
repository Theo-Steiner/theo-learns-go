## Things I noticed

### Reflection

> the ability of a program to examine its own structure, particularly through types

- form of metaprogramming
- allows you to determine types at runtime
- less performant than pre-typed

```go
func (x interface{}) {
val := reflect.ValueOf(x)
switch val.Kind() {
case reflect.String:
  str := val.String() // convert to string
case reflect.Pointer:
  el := val.Elem() // convert to element that pointer points to
case reflect.Slice, reflect.Array:
    for i := 0; i < val.Len(); i++ {
        val.Index(i) // get element at index
    }
case reflect.Map:
    for _, key := range val.MapKeys() {
        val.MapIndex(key) // get element for key
    }
// struct case   -> recurse over every element in slice
case reflect.Struct:
    for i := 0; i < val.NumField(); i++ {
        val.Field(i) // get struct property
    }
}
```

### interface{}

- `interface{}` is `any` (any is an alias)
- therefore using interface{} means you lose typesafety
