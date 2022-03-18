package main

import (
	"context"
	"github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc"
)

// EcommerceAccountRpcImpl implements the last service interface defined in the IDL.
type EcommerceAccountRpcImpl struct{}

// VerifyUser implements the EcommerceAccountRpcImpl interface.
func (s *EcommerceAccountRpcImpl) VerifyUser(ctx context.Context, req *account_rpc.VerifyUserRequest) (resp *account_rpc.VerifyUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the EcommerceAccountRpcImpl interface.
func (s *EcommerceAccountRpcImpl) Login(ctx context.Context, req *account_rpc.LoginRequest) (resp *account_rpc.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Logout implements the EcommerceAccountRpcImpl interface.
func (s *EcommerceAccountRpcImpl) Logout(ctx context.Context, req *account_rpc.LogoutRequest) (resp *account_rpc.LogoutResponse, err error) {
	// TODO: Your code here...
	return
}
