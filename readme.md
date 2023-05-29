# Optional

The `optional` package is a Golang package that provides a type called `Option` for representing optional values. It allows you to handle scenarios where a value may or may not be present.

## Installation

To use the `optional` package in your Golang project, you can use the following command to install it:

```
go get github.com/LNMMusic/optional
```

## Usage

Import the `optional` package in your Go code:

```go
import "github.com/LNMMusic/optional"
```

### Creating an Optional Value

To create an optional value, you can use the `Some` and `None` functions.

#### Some

The `Some` function creates an optional value with a non-nil inner value.

```go
value := 42
opt := optional.Some(value)
```

#### None

The `None` function creates an optional value with a nil inner value.

```go
opt := optional.None()
```

### Checking if an Optional Value is Some

You can check if an optional value is a `Some` value by using the `IsSome` method.

```go
if opt.IsSome() {
    // The optional value is Some
} else {
    // The optional value is None
}
```

### Unwrapping an Optional Value

To retrieve the inner value of a `Some` optional value, you can use the `Unwrap` method. It returns the inner value and an error if the optional value is `None`.

```go
value, err := opt.Unwrap()
if err != nil {
    // Handle the error
} else {
    // Use the inner value
}
```

### Example

Here's an example that demonstrates the usage of the `optional` package:

```go
package main

import (
	"fmt"

	"github.com/your-username/optional"
)

func main() {
	opt := optional.Some(42)
	if opt.IsSome() {
		value, err := opt.Unwrap()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Value:", value)
		}
	} else {
		fmt.Println("The optional value is None")
	}
}
```

## Error Handling

If you attempt to call the `Unwrap` method on a `None` optional value, it will return the `ErrUnwrapNone` error. You should handle this error appropriately in your code.

## JSON Marshalling and Unmarshalling

The `Option` type in the `optional` package supports JSON marshalling and unmarshalling. Here are some important points to consider:

- When unmarshalling JSON into an optional value, make sure to work with a pointer to the optional value. This allows the JSON unmarshalling process to modify the inner value correctly.
- When marshalling an optional value to JSON, work with a non-pointer value. This ensures that Go's JSON marshalling rules are followed.

### Unmarshalling JSON

The `UnmarshalJSON` method allows you to unmarshal JSON data into an optional value.

```go
data := []byte(`{"Value":42}`)
var opt optional.Option[int]
err := json.Unmarshal(data, &opt)
if err != nil {
    // Handle the error
}
```

### Marshalling JSON

The `MarshalJSON` method allows you to marshal an optional value to JSON.

```go
opt := optional.Some(42)
data, err := json.Marshal(opt)
if err != nil {
    // Handle the error
} else {
    fmt.Println(string(data))
}
```

The output of the above code is `42` due to encodes straight forward the inner value of the optional

Please note that marshalling a `None` optional value will result in `null` in the JSON output.

## License

This package is licensed

 under the MIT License. See the [LICENSE](LICENSE) file for more information.