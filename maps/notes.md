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

