package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	envs, _ := ReadDir("testdata/env")

	tests := []struct {
		cmd  []string
		envs Environment
		wait int
	}{
		{[]string{"env"}, envs, 0},
		{[]string{"echo", envs["FOO"].Value}, envs, 0},
		{[]string{"no commmand", envs["FOO"].Value}, envs, -1},
	}
	t.Run("test for start cmd with env", func(t *testing.T) {
		for _, val := range tests {
			require.Equal(t, val.wait, RunCmd(val.cmd, val.envs), val)
		}
	})
}
