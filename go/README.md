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
    Marshal/UnMarshal
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
    匿名函数
    可变参函数
    closure：函数+捕获到的自由变量，可以看作是包含一个函数及自由变量的对象体。
### method
    method is function with a inplicit receiver
### interface
    sort.Interface
    状态模式
### go routine
    自动为main函数创建一个goroutine。
    所有goroutine在main函数结束时会一同结束。
    使用go关键之创建goroutine时，被调用函数的返回值会被忽略。
    go runtime实现了一个小型的任务调度器,类似于操作系统调度线程。
### network
    http.Get
    http.Post
    net.Listen
    net.Dial
    Listener.Accept
    tcp粘包问题
### unit test
    go test -v --run TestCaseA
    go test -v --bench=. BenchmarkA

