package main

import (
	account_rpc "github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc/ecommerceaccountrpc"
	"log"
)

func main() {
	svr := account_rpc.NewServer(new(EcommerceAccountRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
