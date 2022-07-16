package main

import (
	"os"
	"os/exec"
)

const (
	ErrCMD int = iota - 3
	ErrEnvs
	ErrRunApp
	OK
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, envs Environment) (returnCode int) {
	if len(cmd) < 1 {
		return ErrCMD
	}
	if err := UpdateEnvs(envs); err != nil {
		return ErrEnvs
	}

	app := exec.Command(cmd[0], cmd[1:]...)
	app.Stderr = os.Stderr
	app.Stdout = os.Stdout
	app.Stdin = os.Stdin

	if err := app.Run(); err != nil {
		return ErrRunApp
	}

	return app.ProcessState.ExitCode()
}

// Update envs
func UpdateEnvs(envs Environment) error {
	for env, valEnv := range envs {
		if err := os.Unsetenv(env); err != nil {
			return err
		}
		if !valEnv.NeedRemove {
			if err := os.Setenv(env, valEnv.Value); err != nil {
				return err
			}
		}
	}

	return nil
}
