Go学习线索.

### io
    os.Create
    os.Open
    bytes.buffer
    bufio.NewReader
    fmt.Scanf
### args
    flag.String
    flag.NewFlagSet
### time
    timer/ticker
### json

### error
    errors.New
### composite types:
    slice
    map
    struct
### reflect
    reflect.Type
    reflect.Value
### functions
    anonymous functions
### method
    method is function with a inplicit receiver
### interface
    sort.Interface
### go routine
    自动为main函数创建一个goroutine。
    所有goroutine在main函数结束时会一同结束。
    使用go关键之创建goroutine时，被调用函数的返回值会被忽略。
    go runtime实现了一个小型的任务调度器,类似于操作系统调度线程。
### network
    http.Get
    http.Post
    net.Listen
    Listener.Accept
### unit test
    go test -v --run TestCaseA
    go test -v --bench=. BenchmarkA

