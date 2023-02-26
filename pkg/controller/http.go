package controller 


type HttpHandler interface {
	Serve() error
}