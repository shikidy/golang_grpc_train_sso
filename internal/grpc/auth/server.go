package auth

import (
	"context"

	ssoy1 "github.com/shikidy/golang_grpc_train_proto/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssoy1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssoy1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssoy1.LoginRequest) (*ssoy1.LoginResponse, error) {
	return &ssoy1.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssoy1.RegisterRequest) (*ssoy1.RegisterResponse, error) {
	panic("implement me")
}
func (s *serverAPI) IsAdmin(ctx context.Context, req *ssoy1.IsAdminRequest) (*ssoy1.IsAdminResponse, error) {
	panic("implement me")
}
