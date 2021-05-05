package schema

// TODO tests for this stuff! Everything in this package is totally generic,
// nothing is specific to Grafana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"cuelang.org/go/cue"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/txtar"
)

type Case struct {
	Name       string
	CUE        string
	JsonSchema string
}

func TestGenerate(t *testing.T) {
	var CasesDir = filepath.Join("load", "testdata", "artifacts", "dashboards", "cue2jsonschema")
	cases, err := loadCases(CasesDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		t.Run(c.Name+" convert cue to jsonschema", func(t *testing.T) {
			var r cue.Runtime
			scmInstance, err := r.Compile(c.Name+".cue", c.CUE)
			if err != nil {
				t.Fatal(err)
			}
			reschema, err := ConvertCueToJsonSchema(scmInstance)
			if err != nil {
				t.Fatal(err)
			}

			if err != nil {
				t.Fatal(err)
			}
			if s := cmp.Diff(reschema, c.JsonSchema); s != "" {
				t.Fatal(s)
			}
		})
	}
}

func loadCases(dir string) ([]Case, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var cases []Case

	for _, fi := range files {
		file := filepath.Join(dir, fi.Name())
		a, err := txtar.ParseFile(file)
		if err != nil {
			return nil, err
		}

		if len(a.Files) != 2 {
			return nil, fmt.Errorf("Malformed test case '%s': Must contain exactly two files (CUE, JsonSchema), but has %d", file, len(a.Files))
		}

		jsonSchemaBuffer := new(bytes.Buffer)
		jsonSchema := a.Files[1].Data
		if err := json.Compact(jsonSchemaBuffer, jsonSchema); err != nil {
			return nil, err
		}

		cases = append(cases, Case{
			Name:       fi.Name(),
			CUE:        string(a.Files[0].Data),
			JsonSchema: jsonSchemaBuffer.String(),
		})
	}
	return cases, nil
}
