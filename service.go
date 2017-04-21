package testrepo

type Service interface {
	SayHello(str string) string 
}

// Implement service with empty struct
type SayHelloService struct {
}

func (SayHelloService) SayHello(str string) string {
	return str
}
