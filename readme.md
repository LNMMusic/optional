# Optional

The `optional` package is a Golang package that provides a type called `Option` for representing optional values. It allows you to handle `scenarios` where a `value` `may or may not be present`. This include compatibility with dynamic types such as maps and JSON where the `existence of a key` may not be guaranteed. In this case the `Option` type can be used to represent the existence of a key and its value.

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
opt := optional.Some[int](value)
```

#### None

The `None` function creates an optional value with a nil inner value.

```go
opt := optional.None[int]()
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

To retrieve the inner value of a `Some` optional value, you can use the `Unwrap` method. It returns the inner value. If the optional value is `None`, the function panics. You should handle this appropriately in your code.

```go
if opt.IsSome() {
	value := opt.Unwrap()
	// Do something with the value
} else {
	// The optional value is None
}
```

### Example

Here's an example that demonstrates the usage of the `optional` package:

```go
package main

import (
	"fmt"

	"github.com/LNMMusic/optional"
)

func main() {
	opt := optional.Some[int](42)
	if opt.IsSome() {
		value := opt.Unwrap()
		fmt.Printf("The optional value is Some and its value is %d\n", value)
	} else {
		fmt.Println("The optional value is None")
	}
}
```

## Error Handling

If you attempt to call the `Unwrap` method on a `None` optional value, it will panic. You should handle this error appropriately in your code, by first checking if the optional value is `Some` or `None`.

```go
if opt.IsSome() {
	value := opt.Unwrap()
	// Do something with the value
} else {
	// The optional value is None
}
```

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

#### *Rules*
- If json is `null`, the optional value will be `None`. This means that optional DOES NOT DISTINGUISH between the `absence of the key` and the `presence of the key with a null value`. All `null` values are treated as `non-existent-values`. To avoid this, later will be available a new type called `Nullable` that will be able to extend the values on go with the `null` value. This can be combined with the `optional` type to create a `NullableOptional` type that will be able to distinguish between the `absence of the key` and the `presence of the key with a null value`.

```go
var optionalNullable := optional.Some[Nullable[int]](Nullable[int]{Value: nil})
```

#
---

# Nullable

## About Nullable and Optionals

In modern programming and data handling, understanding and differentiating the presence of a value from its content is crucial. This repository contains two distinct Go packages—`nullable` and `optionals`—designed to address and manipulate these aspects with clarity and efficiency in structured and semi-structured data environments.

### Why Nullable?

In structured data systems like SQL databases or data represented in CSV files, fields are consistently present; however, their values can explicitly be `null`. The `nullable` package provides a robust way to handle such scenarios in Go programming. It introduces a `Null[T]` generic type that encapsulates the presence and actual content of a value through a dual-state mechanism:

- **Null State**: Represents the absence of a value (akin to `NULL` in databases), clearly differentiating from...
- **Not Null State**: A valid, user-defined value of type `T`.

Using `nullable`, developers can explicitly handle and differentiate between "no data" and "zero data," enhancing the robustness and clarity of data handling operations.

### Importance of Optionals

For semi-structured data like JSON, where fields may not be consistently present, the theory behind and implementation of `optionals` is vital. Unlike `nullable`, which deals with the nuance of value presence within statically existing fields, `optionals` tackles the dynamism of fields themselves — they can either exist or not. This is particularly relevant in programming environments dealing with dynamic types and memory management (like slices and maps in Go), where the structure is managed in heap memory and can change at runtime.

### Conclusion

Together, these packages offer a structured approach to data handling that aligns with the intrinsic properties of different data schemas, from rigidly structured to flexible, schema-less formats. They promote safer, more predictable code by explicitly managing the nuances of data presence and value nullability, making them indispensable for developers working across various data-intensive applications.

#
--- 

## License

This package is licensed

 under the MIT License. See the [LICENSE](LICENSE) file for more information.