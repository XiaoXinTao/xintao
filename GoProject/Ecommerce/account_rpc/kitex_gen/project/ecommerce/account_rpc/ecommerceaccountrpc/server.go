// Code generated by Kitex v0.2.0. DO NOT EDIT.
package ecommerceaccountrpc

import (
	"github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler account_rpc.EcommerceAccountRpc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
