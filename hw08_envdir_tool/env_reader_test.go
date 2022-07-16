package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	tests := []struct {
		envDir   string
		countEnv int
		key      string
		val      EnvValue
	}{
		{"testdata/env", 5, "", EnvValue{}},
		{"", 0, "BAR", EnvValue{"bar", false}},
		{"", 0, "EMPTY", EnvValue{"", false}},
		{"", 0, "FOO", EnvValue{"   foo\nwith new line", false}},
		{"", 0, "HELLO", EnvValue{"\"hello\"", false}},
		{"", 0, "UNSET", EnvValue{"", true}},
	}

	var envs Environment
	var err error

	t.Run("test read env", func(t *testing.T) {
		for _, val := range tests[:1] {
			envs, err = ReadDir(val.envDir)
			require.NoError(t, err)

			assert.Equal(t, val.countEnv, len(envs), "count env", len(envs))
		}
	})

	t.Run("continue read env with compare structs", func(t *testing.T) {
		for _, val := range tests[1:] {
			assert.Equal(t, val.val, envs[val.key])
		}
	})
}
