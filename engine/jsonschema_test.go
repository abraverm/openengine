package engine

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/qri-io/jsonschema"
)

// TestSet is a json-based set of tests
// JSON-Schema comes with a lovely JSON-based test suite:
// https://github.com/json-schema-org/JSON-Schema-Test-Suite
type TestSet struct {
	Description string             `json:"description"`
	Schema      *jsonschema.Schema `json:"schema"`
	Tests       []TestCase         `json:"tests"`
}

type TestCase struct {
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	Valid       bool        `json:"valid"`
}

func runJSONTests(t *testing.T, testFilepaths []string) {
	tests := 0
	passed := 0
	ctx := context.Background()
	for _, path := range testFilepaths {
		t.Run(path, func(t *testing.T) {
			base := filepath.Base(path)
			testSets := []*TestSet{}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				t.Errorf("error loading test file: %s", err.Error())
				return
			}

			if err := json.Unmarshal(data, &testSets); err != nil {
				t.Errorf("error unmarshaling test set %s from JSON: %s", base, err.Error())
				return
			}

			for _, ts := range testSets {
				sc := ts.Schema
				for i, c := range ts.Tests {
					tests++
					validationState := sc.Validate(ctx, c.Data)
					if validationState.IsValid() != c.Valid {
						t.Errorf("%s: %s test case %d: %s. error: %s", base, ts.Description, i, c.Description, *validationState.Errs)
					} else {
						passed++
					}
				}
			}
		})
	}
	t.Logf("%d/%d tests passed", passed, tests)
}

func TestOeKeywords(t *testing.T) {
	jsonschema.LoadDraft2019_09()
	jsonschema.RegisterKeyword("oeProperties", NewOeProperties)
	jsonschema.RegisterKeyword("oeRequired", NewOeRequired)
	runJSONTests(t, []string{
		"testdata/oeRequired.json",
		"testdata/oeProperties.json",
	})
}
