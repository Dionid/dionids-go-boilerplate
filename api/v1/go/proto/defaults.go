package proto

import emptypb "google.golang.org/protobuf/types/known/emptypb"

type Request interface {
	GetId() string
}

func NewDefaultCallResponse(request Request) *DefaultCallResponse {
	return &DefaultCallResponse{
		Id: request.GetId(),
		Result: &DefaultCallResponse_Result{
			Result: &DefaultCallResponse_Result_Success{
				Success: &emptypb.Empty{},
			},
		},
	}
}
