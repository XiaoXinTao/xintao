// Code generated by Kitex v0.2.0. DO NOT EDIT.

package ecommerceaccount

import (
	"context"
	"github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/xintao/project/ecommerce"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return ecommerceAccountServiceInfo
}

var ecommerceAccountServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "EcommerceAccount"
	handlerType := (*ecommerce.EcommerceAccount)(nil)
	methods := map[string]kitex.MethodInfo{
		"Login":  kitex.NewMethodInfo(loginHandler, newEcommerceAccountLoginArgs, newEcommerceAccountLoginResult, false),
		"Logout": kitex.NewMethodInfo(logoutHandler, newEcommerceAccountLogoutArgs, newEcommerceAccountLogoutResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "ecommerce",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.2.0",
		Extra:           extra,
	}
	return svcInfo
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ecommerce.EcommerceAccountLoginArgs)
	realResult := result.(*ecommerce.EcommerceAccountLoginResult)
	success, err := handler.(ecommerce.EcommerceAccount).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEcommerceAccountLoginArgs() interface{} {
	return ecommerce.NewEcommerceAccountLoginArgs()
}

func newEcommerceAccountLoginResult() interface{} {
	return ecommerce.NewEcommerceAccountLoginResult()
}

func logoutHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*ecommerce.EcommerceAccountLogoutArgs)
	realResult := result.(*ecommerce.EcommerceAccountLogoutResult)
	success, err := handler.(ecommerce.EcommerceAccount).Logout(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newEcommerceAccountLogoutArgs() interface{} {
	return ecommerce.NewEcommerceAccountLogoutArgs()
}

func newEcommerceAccountLogoutResult() interface{} {
	return ecommerce.NewEcommerceAccountLogoutResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Login(ctx context.Context, req *ecommerce.LoginRequest) (r *ecommerce.LoginResponse, err error) {
	var _args ecommerce.EcommerceAccountLoginArgs
	_args.Req = req
	var _result ecommerce.EcommerceAccountLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Logout(ctx context.Context, req *ecommerce.LogoutRequest) (r *ecommerce.LogoutResponse, err error) {
	var _args ecommerce.EcommerceAccountLogoutArgs
	_args.Req = req
	var _result ecommerce.EcommerceAccountLogoutResult
	if err = p.c.Call(ctx, "Logout", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
