
- Early Returns
- Avoid Raw Strings
- Use Sentinel Errors for Reusable Errors
- Integrate Structured Logging
- Centralize Error Messages
- Categorize Errors
- Create Unifying Error Handling Functions
- Test Error Scenarios


## Early Returns

Return Errors as soon as possible.
Generally, in the next line that it was found.


## Avoid Raw Strings

Define Custom Error Types instead of using just strings.
Encapsulate relevant details in the error type.

Custom Error Types are identifiable via type assertions or `errors.As`.

This allows you to provide additional context, have better Pattern Match and avoid runtime string parsing.


## Use Sentinel Errors for Reusable Errors

Create them with `var ErrNotFound = errors.New("...")`.
Compare them with `errors.Is`.


## Integrate Structured Logging

Use Structured Logging with metadata to include context in each error for easier debugging and monitoring.

Use logging libraries like `zap` or `logrus` to structure logs.


## Centralize Error Messages

Define Error Messages in one place to ensure Consistency & Reusability.


## Categorize Errors

Differentiate between **Recoverable & Unrecoverable** Errors.
Provide Error Categories for easy handling:
- **User Error** (e.g., invalid input)
- **System Error** (e.g., I/O failures)
- **Network Error** (e.g., connection failures, database timeouts)

If necessary, define interfaces to group them.


## Create Unifying Error Handling Functions

Centralize error handling with an unified function that differentiates error types.

Example:
```go
func handleErrorExample(err error) {
    switch {
    case errors.As(err, &ValidationError{}):
        // log...
    case errors.As(err, &NetworkError{}):
        // log...
    default:
        // log...
    }
}
```


## Test Errors

Test to ensure proper error handling.

