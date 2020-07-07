package cmd

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Gqlgen run",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfigFromDefaultLocations()
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
			os.Exit(2)
		}

		p := modelgen.Plugin{
			MutateHook: mutateHook,
		}

		err = api.Generate(cfg,
			api.NoPlugins(),
			api.AddPlugin(&p),
		)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(3)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, model := range b.Models {
		for _, field := range model.Fields {
			field.Tag = `json:"` + field.Name + `,omitempty" bson:"` + field.Name + `,omitempty"`
		}
	}
	return b
}
