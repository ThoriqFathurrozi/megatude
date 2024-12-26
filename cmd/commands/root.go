package commands

import (
	"github.com/ThoriqFathurrozi/megatude/configs"
	"github.com/ThoriqFathurrozi/megatude/internal/core"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	cfgFile  string
	megatude *core.Megatude

	rootCmd = &cobra.Command{
		Use:   "megatude",
		Short: "Simple Rest API for BMKG Earthquake",
	}
)

func Execute() error {
	return rootCmd.Execute()

}

func init() {
	cobra.OnInitialize(initApp)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "../config.yaml", "config file (default is ./configs/config.yml)")

	serveCmd := serveCmd()
	dbCmd := newDBCmd()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(dbCmd)

}

func initApp() {
	sugar, _ := zap.NewDevelopment()
	defer sugar.Sync()
	if cfgFile == "" {
		sugar.Fatal("Config file not set")
	}

	configs.LoadConfig(cfgFile)
	cfg := configs.GetConfig()

	app := core.NewEcho()

	cron := core.NewCron()

	db, err := core.NewDB()
	if err != nil {
		sugar.Fatal("unable to initialize db: ", zap.Error(err))
	}

	megatude = &core.Megatude{
		Config: cfg,
		App:    app,
		DB:     db,
		Corn:   cron,
	}

	core.Init(megatude)

}
