## What I learned/noticed:

### Syntax

#### variadic functions

- functions that can take a variable number of arguments

```go
func Sum(numbersToSum ...int) (sum int) {
  for _, number := range numbersToSum {
    sum += number
  }
  return
}
```

  -> variadic argument becomes slice of the defined type

#### Arrays

- always fixed capacity
- capacity is integral part of array type
- two ways to define arrays:
  -  `arr := [3]int{1, 2, 3}` explicit array capacity
  -  `arr := [...]int{1, 2, 3}` implicit array capacity

iterate over array items using range:
```go
for index, el := range arr {
  fmt.Printf("element at array position %d was %d", index, el)
}
```

#### Slices

- variable capacity
- defined like: `slice := []int{1, 2, 3}` explicit array capacity
  - almost like arrays, but without capacity
- you can append to them: `slice = append(slice, item)`
  - appending requires assignment (because **sometimes** a new slice is created)
- create a slice based on a runtime known capacity:
  - `slice = make([]int, capacity)`
- slices can be 'sliced' into partial slices:
  - syntax: like index access, but with a colon, that denotes the lower and upper index bounds of the slice `slice[lower:upper]`

```go
    mySlice := []int{1,2,3} // {1,2,3}
    subSlice := mySlice[1:] // {2,3}
    subSlice = mySlice[:2]  // {1,2}
    subSlice = mySlice[1:2] // {2}
```

> "The make built-in function allocates and initializes an object of type slice, map, or chan (only)."

### misc

- inside string format `%v` uses the `default` format
