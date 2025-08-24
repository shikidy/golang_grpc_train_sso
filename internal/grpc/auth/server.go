package auth

import (
	"context"

	ssoy1 "github.com/shikidy/golang_grpc_train_proto/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (string, error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (int64, error)
	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)
}
type serverAPI struct {
	ssoy1.UnimplementedAuthServer
	auth Auth
}

const (
	emptyValue = 0
)

func Register(gRPC *grpc.Server, auth Auth) {

	ssoy1.RegisterAuthServer(gRPC, &serverAPI{
		auth: auth,
	})
}

func (s *serverAPI) Login(ctx context.Context, req *ssoy1.LoginRequest) (*ssoy1.LoginResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssoy1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssoy1.RegisterRequest) (*ssoy1.RegisterResponse, error) {
	if req.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	userId, err := s.auth.RegisterNewUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssoy1.RegisterResponse{
		UserId: userId,
	}, nil
}
func (s *serverAPI) IsAdmin(ctx context.Context, req *ssoy1.IsAdminRequest) (*ssoy1.IsAdminResponse, error) {
	if req.GetUserId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "userid is required")
	}
	isAdmin, err := s.auth.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssoy1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
