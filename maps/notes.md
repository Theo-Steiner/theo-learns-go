## What I learned/noticed:

### map

- declaration: `map[string]string{"a_key": "a_value"}`
  - `map` keyword
  - `[string]` key type
    -  must be a *comparable* type
  - `string` value type
    -  can be any type
- lookup: `value, ok := dictionary[key]`
  - always returns string (`""` if not found)
  - second return is `boolean` that indicates if map lookup successful
- key/value pairs can be deleted using `delete(map,key)`
- maps can be globally modified even if passed by reference
  - A map value is a pointer to a runtime.hmap structure
  - therefore maps can be nil

  -> never not-initialize a map ~var m map[string]string]~
  instead initialize with an empty collection:
  `var m = map[string]string{}`

### best practice

create error constants for sentinel errors:

> Sentinel errors are user defined errors that indicated very specific events that you, as a developer, anticipate & identify as adequately important to define and specify. As such, you declare them at the package level and, in doing so, imply that your package functions may return these errors (thereby committing you in the future to maintain these errors as others depending on your package will be checking for them).

see: [digital ocean sentinel errors](https://www.digitalocean.com/community/tutorials/how-to-add-extra-information-to-errors-in-go#handling-specific-errors-using-sentinel-errors) 

A sentinel error can be defined as a var, 
```go
var MySentinelError = error.New("MySentinelError")
```
however that renders it mutable.
There is however a way to declare sentinel errors in an immutable manner:
```go
type ErrorType string // declare your ErrorType based on the default string type

// add an `Error` method to your ErrorType, to implement the error interface
func (e ErrorType) Error() string {
  string(e)
}

// you can now declare sentinel errors in an immutable manner
const (
  MySentinelError = "Oh no!"
  MyOtherBadError = "Oh hell no!"
)
```
