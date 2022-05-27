package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

// App 是客户端的主要结构!,推荐使用NewApp()来创建,注意全是非公开字段
type App struct {
	basename    string
	name        string
	description string
	options     CliOptions //客户端的配置,不同客户端的关键
	runFunc     RunFunc
	silence     bool
	noVersion   bool
	noConfig    bool
	commands    []*Command
	args        cobra.PositionalArgs //cobra命令的执行函数
	cmd         *cobra.Command       //命令
}

// NewApp 使用选项模式创建一个App
func NewApp(name string, basename string, opts ...Option) *App {
	a := &App{
		name:     name,
		basename: basename,
	}
	for _, o := range opts {
		o(a)
	}
	a.buildCommand()
	return a
}

// Option 参数为App的地址,这样才能在其中进行修改App的字段
type Option func(*App)

func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

func WithDescription(desc string) Option {
	return func(a *App) {
		a.description = desc
	}
}

func WithDeaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		}
	}
}

// RunFunc defines the application's startup callback function.
type RunFunc func(basename string) error

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}
