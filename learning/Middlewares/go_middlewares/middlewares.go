package go_middlewares

type Request struct {
	index       int
	middlewares []HandleFunc
}

type HandleFunc func(*Request)

//实例化Request
func NewRequest() *Request {
	return &Request{
		index:       0,
		middlewares: make([]HandleFunc, 0),
	}
}

//注册中间件
func (this *Request) Register(middlewares ...HandleFunc) {
	for _, middleware := range middlewares {
		this.middlewares = append(this.middlewares, middleware)
	}
}

//执行请求
func (this *Request) Next() {
	index := this.index
	if index >= len(this.middlewares) {
		return
	}

	this.index++
	this.middlewares[index](this)
}