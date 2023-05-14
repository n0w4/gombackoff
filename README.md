# GOMBACKOFF

an other very simple tool, this time solve exponential and linear backoff.

## LinearBackoff example

- the **LinearBackoff** method, factor is the time that we want to wait
- the **Times** method receive how many attempts we can do

```go
retryPolicy := gombackoff.NewRetryPolicy().LinearBackoff(2).Times(2)
...
if ok := retryPolicy.Wait(); ok {
    continue
}
```

## Exponential example

- the **ExponentialBackoff** method, factor is multiplicated to current attempt.
- the **Times** method receive how many attempts we can do.

```go
retryPolicy := gombackoff.NewRetryPolicy().ExponentialBackoff(2).Times(2)
...
if ok := retryPolicy.Wait(); ok {
    continue
}
```