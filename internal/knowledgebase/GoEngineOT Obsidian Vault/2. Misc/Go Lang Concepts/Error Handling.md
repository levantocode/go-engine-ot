
In Go, an error is a very simple Interface: it has a single method called `Error()` that returns a string.

```go
type error interface {
    Error() string
}
```

In Go, every interface can be nil, so whenever we evaluate `if error != nil` we are checking if any kind of `Error` was created. This is, in essence, very simple.


## Creating Error Instances

There is 2 ways to create an Error instance:
- `errors.new(string)` for strings without variables values
- `fmt.Errorf(string, ...variables)` to format variables in the string

 Example of `errors.New(string)`:
```go
func divide (a, b int) (int, error) {
    if b == 0 {
	    return errors.New("divide: Cannot divide by 0")
    }
	return a / b, nill
}
```

It is a good practice to preface/subtitle the error with the name of the function where it was created.

Example of `fmt.Errorf(string, ...variables)`:
```go
func subtract (a, b int) (int, error) {
	result := a - b
    if result <  0 {
	    return result, fmt.Errorf("subtract: %d - %d is a negative number", a, b)
    }
	return result, nill
}
```


## Testing Errors

To test an error besides `nil`, instead of checking by the value of its string, it's best to compare the entire entity. For that we need to have an accessible and shared error value.

The function that compares errors is: `errors.is()`.

```go
var ErrUnsupportedMode = errors.New("This mode is unsupported")

func someOperation(someMode int) (string, error) {
	if mode < 0 {
		return "", ErrUnsupported
	}
	return fmt.Sprintf("Running operation with mode %d", someMode), nil
}

func TestSomeOperationUnsupported(t *testing.T) {
	_, err := operation(-1)
	if errors.Is(err, ErrUnsupportedMode) {
		t.Logf("Expected error occurred: %v", err)
	} else {
		t.Errorf("Unexpected Error. Expected unsupported error, got %v", err)
	}
}
```


## Creating Your Own Error Type

To do that we need to implement the `Error` interface.

You can make it receiving an error message, having it a pre-defined error message that receives nothing or that receives some value that gives more context to the error message.

Example 1:
```go
type InvalidInputError struct {
	message string
}

func (e *InvalidInputError) Error() string {
	return "Invalid input: " + e.message
}
```

Then you can create a instance with `&InvalidInputError{message: "your message here"}`.

Example 2:
```go
type UnsupportedModeError struct {
	mode int
}

func (e *UnsupportedModeError) Error() string {
	return fmt.Sprintf("someOperation: Mode %d is unsupported", e.mode)
}
```

Then you can create a instance with `&UnsupportedModeError{mode: yourVariable}`.


## Testing With Your Own Error Type

To assert against a custom error type, instead of using `errors.Is` we need to use `errors.As` passing the address of the custom error type. For example:
```go
...
var invalidInputError *InvalidInputError
ìf errors.As(err, &invalidInputError) {...}
...
```


## Creating MultiErrors (List of Errors)

This is a little advanced, and a bit outside of the scope at the moment.
Will be documented on the future.


## Wrapped Errors

This is a little advanced, and a bit outside of the scope at the moment.
Will be documented on the future.