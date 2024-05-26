## What I learned/noticed:

### Syntax

#### Casing

if a symbol (variable, type, function...) starts with a lowercase letter, it becomes private
```go
package wallet

type Wallet struct { // Wallet is public
  amount int // Wallet.amount is private to the wallet package
  Owner  string // Wallet.Owner is public
}
```

#### Pointers

- go has pass by copy (as opposed to javascript pass by reference)
- pointer can be manually retrieved by prefixing with `&`
  - pointers can be formatted in strings using `fmt.Printf("%p", &foo),`
- example with method receiver:
  - value:   `func (w Wallet) Balance() {...}`
  - pointer: `func (w *Wallet) Balance() {...}`
- go has `automatic dereferencing`
  - meaning: a pointer is automatically turned back into it's value when used
  - manual dereferencing: `func (w *Wallet) Deposit (amount int) { *w.balance += amount  }` // also valid
  - auto dereferencing: `func (w *Wallet) Deposit (amount int) { w.balance += amount  }` 
- by convention: if one method uses a struct pointer, all methods should use them (to keep receiver type consistent)

#### Type Declaration

- allows you to define types based on other types
- allows you to declare methods for types

```go
type Bitcoin int // declare `Bitcoin type based on int type`

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}

networth := Bitcoin(99999)
fmt.Printf("I own %s", networth)
```

#### Errors

- returned as value
- created using `err := errors.New("error message")`
- error message retrieved via `err.Error()`
- error return value does not need to be typed

```go
import "errors"

func ThrowError() error {
  return errors.New("oh no!")
}

err := ThrowError() // error{"oh no!"}
err.Error() // "oh no!"
```

#### nil

- accessing nil throws a runtime panic -> check for nil

#### Testing

- `t.Fatal` fail a test immediately without making other assertions
  - necessary to prevent nil pointer panic
