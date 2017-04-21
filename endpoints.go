package testrepo

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type SayHelloRequest struct {
	Name string
}

type SayHelloResponse struct {
	Message string `json: "message"`
	Err     error  `json:"err,omitempty"`
}

//endp wrapper
type Endpoints struct {
	SayHelloEndpoint endpoint.Endpoint
}

func MakeSayHelloEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		var (
			name string
		)
		name = req.Name

		return SayHelloResponse{Message: "hello " + name}, nil

	}
}
