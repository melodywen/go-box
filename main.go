package main

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/http"
	"github.com/sirupsen/logrus"
)

func main() {
	var httpKernel http.KernelInterface
	var ok bool
    k :=	App.Make(&httpKernel)

	if httpKernel,ok =  k.(http.KernelInterface);!ok{
		logrus.Panicln("获取 http kernel 失败")
	}

	httpKernel.Bootstrap()




	fmt.Println(httpKernel)

}
