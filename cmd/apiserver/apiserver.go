package main

import (
	"auth/internal/apiserver"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	//设置随机seed 和P的数量等于核心数
	rand.Seed(time.Now().UnixNano())
	if len(os.Getenv("GOMAXPROCESS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	apiserver.NewApp("iam-apiserver")

}
