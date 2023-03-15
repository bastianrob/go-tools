# Go Tools

Collection of utility struct and functions

> Require go v1.19

## Arrays

### 1. ForEach

Loop an array

```go
arr := []int{10, 20, 30}
gotools.ForEach(arr, func(row, i int) {
    fmt.Printf("(%d, %d)", i, row)
})

// prints (0,10) (1,20) (2,30)
```

### 2. Map

Loop an array and map each entry into a new value, return a new array with the same size

```go
arr := []int{10, 20, 30}
mapped := gotools.Map(arr, func(row, i int) int {
    return row + i + 10
})

// mapped = [20, 31, 42]
```

### 3. Reduce

Reduce an array into a new value, based on its initial value

```go
arr := []int{10, 20, 30}
reduced := gotools.Reduce(
    arr,
    // sum all rows
    func(acc, row, i int) int {
        return acc + row
    },
    0, // initial value
)

// reduced = 60
```

### 4. Find

Find value and index in an array based on given predicate

```go
arr := []int{10, 20, 30}
found, index := gotools.Find(
    arr,
    func(row int) bool {
        return row == 20
    },
)

// found = 20, index = 1
```

Returns `zero value` and `index = -1` if not found

```go
arr := []int{10, 20, 30}
found, index := gotools.Find(
    arr,
    func(row int) bool {
        return row == 1000
    },
)

// found = 0, index = -1
```


### 5. Filter

Find all values in an array based on given predicate.
Returns empty array if nothing is found


```go
arr := []int{10, 20, 30}
found := gotools.Filter(
    arr,
    func(row int) bool {
        return row > 15
    },
)

// found = [20, 30]
```

### 6. Any

Check whether array contains given predicate.

```go
arr := []int{10, 20, 30}
exists := gotools.Any(
    arr,
    func(row int) bool {
        return row > 15
    },
)

// exists = true
```

### 7. All

Check whether all entries in the array matches given predicate

```go
arr := []int{10, 20, 30}
all := gotools.All(
    arr,
    func(row int) bool {
        return row > 5
    },
)

// all = true
```

### 8. Chunk

Split an array into smaller chunks based on given `chunkSize`

```go
arr := []int{10, 20, 30, 40, 50}
chunks := gotools.Chunk(arr, 3)

// chunks = [[10 20 30], [40 50]]
```

---

## Concurrencies

### 1. Runnable

Runnable job that utilize channel and go routines to spawn multiple worker


```go
runnable := gotools.NewRunnable(
    10, // number of workers
    func(id int, this *Runnable[int], message int) {
        time.Sleep(1000 * time.Millisecond)
        fmt.Printf("id: %d, message: %d\n", id, message)
    },
)
runnable.Run() // Spawn the worker

// Publish messages
go runnable.Publish(1)
go runnable.Publish(2)
go runnable.Publish(3)
go runnable.Publish(4)
go runnable.Publish(5)
go runnable.Publish(6)
go runnable.Publish(7)
go runnable.Publish(8)
go runnable.Publish(9)
go runnable.Publish(10)

// Wait for all message to finish processing by worker
runnable.Wait()
```

## Utilities

### 1. If

Lazy ternary operator :P

```go
res := gotools.If(1 == 1, "one", "not one")
fmt.Println(res)

// prints one
```
