package cmd

import (
	"github.com/spf13/cobra"

	"chat/configs"
	"chat/pkg/log"
)

var (
	key    string
	secret string
	level  string
)

func init() {
	log.NewDefault()

	flags := rootCmd.Flags()
	flags.StringVarP(&key, "key", "k", "", "dingtalk applet key")
	flags.StringVarP(&secret, "secret", "s", "", "dingtalk applet secret")
	flags.StringVarP(&level, "level", "l", "info", "log lever")

	err := rootCmd.MarkFlagRequired("key")
	if err != nil {
		log.Errorf("rootCmd.MarkFlagRequired err:%v", err)
		return
	}
	err = rootCmd.MarkFlagRequired("secret")
	if err != nil {
		log.Errorf("rootCmd.MarkFlagRequired err:%v", err)
		return
	}
}

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "chat server",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Info("chat server start...")
		return start()
	},
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		log.Errorf("rootCmd.Execute err:%v", err)
	}
}

func start() error {
	log.New(level)
	app := &configs.Applet{Key: key, Secret: secret}
	engine, err := initApplication(level, app)
	if err != nil {
		return err
	}
	return engine.Run(":8100")
}
