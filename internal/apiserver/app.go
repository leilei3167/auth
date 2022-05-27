package apiserver

import (
	"auth/internal/apiserver/config"
	"auth/internal/apiserver/options"
	"auth/pkg/app"
	"github.com/marmotedu/log"
)

const commandDesc = `The IAM API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.

Find more iam-apiserver information at:
    https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iam-apiserver.md`

// NewApp 用默认配置创建App,App是通用的一个结构,使用选项模式,根据不同的选项构建不同的app
func NewApp(basename string) *app.App {
	opts := options.NewOptions() //构建客户端选项

	application := app.NewApp("I AM APISERVER", basename,
		app.WithOptions(opts), //此处接收的是CliOption这个接口
		app.WithDescription(commandDesc),
		app.WithDeaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)
	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}

}
