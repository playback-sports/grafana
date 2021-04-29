package commands

import (
	"fmt"
	"github.com/grafana/grafana"
	"github.com/grafana/grafana/pkg/cmd/grafana-cli/utils"
	"github.com/grafana/grafana/pkg/schema"
	"github.com/grafana/grafana/pkg/schema/load"
)

var p = load.BaseLoadPaths{
	BaseCueFS:       grafana.CoreSchema,
	DistPluginCueFS: grafana.PluginSchema,
}

func (cmd Command) validateScuemataBasics(c utils.CommandLine) error {
	if err := validate(p, load.BaseDashboardFamily); err != nil {
		return err
	}

	if err := validate(p, load.DistDashboardFamily); err != nil {
		return err
	}

	return nil
}

func validate(p load.BaseLoadPaths, loader func(p load.BaseLoadPaths) (schema.VersionedCueSchema, error)) error {
	dash, err := loader(p)
	if err != nil {
		return fmt.Errorf("error while loading dashboard scuemata")
	}

	cueValue := dash.CUE()
	if !cueValue.Exists() {
		return fmt.Errorf("cue value for schema does not exist")
	}

	if err := cueValue.Validate(); err != nil {
		return fmt.Errorf("all schema should be valid with respect to basic CUE rules, %w", err)
	}

	return nil
}
