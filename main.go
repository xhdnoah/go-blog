package main

import (
	"fmt"
	"net/http"

	"github.com/xhdnoah/go-blog/pkg/setting"
	"github.com/xhdnoah/go-blog/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 2 的 20 次幂 1048576
	}
	s.ListenAndServe()
}
