package response

type Response interface {
	Handle(d interface{}) ResponseStruct
}

type ResponseStruct struct {
	Data interface{}
}

func (r ResponseStruct) Handle(d interface{}) ResponseStruct {
	r.Data = d
	return r
}

func Get(d interface{}) Response {
	var r ResponseStruct
	return r.Handle(d)
}
