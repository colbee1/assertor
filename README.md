# assertor

Assertor is a tiny golang helper to quickly validate parameters.

## How to use it

Create a new assertor

```go
v := assertor.New()
```

Make assertions on input parameters

```go
v.Assert(limit > 0, "invalid limit: %d", limit)
v.Assert(stop > 0 && stop > start, "inconsistant stop value: %v", stop)
v.Assert(ctx != nil, "context is missing")
```

Validate

```go
err := v.Validate()
```

Full example

```go
func Example(ctx context.Context, start int, stop int, limit int) error {
    v := assertor.New()
    v.Assert(limit > 0, "invalid limit: %d", limit)
    v.Assert(stop > 0 && stop > start, "inconsistant stop value: %d", stop)
    v.Assert(ctx != nil, "context is missing")
    if err := v.Validate(); err != nil {
        return err
    }

    /* some code */

    return nil
}
```
