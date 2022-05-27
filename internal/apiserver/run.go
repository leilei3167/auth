package apiserver

import "auth/internal/apiserver/config"

// Run 启动APIserver服务,常驻不应该退出
func Run(cfg *config.Config) error {
	server, err := createAPIServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun().Run()
}
