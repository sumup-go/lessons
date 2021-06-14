# Challenges

## Merge two channels

Write a function with the following signature:

```go
func Merge2Channels(f func(int) int, in1 <- chan int, in2 <- chan int, out chan <-int, n int)
```

that `n` times does the following:

* read one number from each channel `in1` and `in2`; let's call them `x1` and `x2`.
* calculate `f(x1) + f(x2)`.
* send the result to channel `out`.

The function must non-blocking, i.e. returning the control back immediately. Function `f` could be time-consuming. 

[Source](https://srgbl.ru/merge2channels/) of the task. 