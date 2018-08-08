package transports

//Caller interface for sending a request to a network transport
type Caller interface {
	Call(method string, params, response interface{}) error
}
