package commands

import (
	"github.com/ThoriqFathurrozi/megatude/internal/db/migrations"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newDBCmd() *cobra.Command {
	var (
		shouldPurge   bool
		shouldMigrate bool
	)

	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	cmd := &cobra.Command{
		Use:   "db",
		Short: "Database management",
		Run: func(cmd *cobra.Command, args []string) {

			migrator := migrations.NewMigrator(megatude.DB)

			if shouldPurge {
				sugar.Info("Purging database")
				if err := migrator.Purge(); err != nil {
					sugar.Fatal("Failed to purge database", zap.Error(err))
				}
			}

			if shouldMigrate {
				sugar.Info("Migrating database")
				if err := migrator.Migrate(megatude.DB); err != nil {
					sugar.Fatal("Failed to migrate database", zap.Error(err))
				}
			}
		},
	}

	cmd.PersistentFlags().BoolVar(&shouldPurge, "purge", false, "Purge database")
	cmd.PersistentFlags().BoolVar(&shouldMigrate, "migrate", false, "Migrate database")

	return cmd
}
