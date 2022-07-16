package main

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

var (
	ErrEnvNameHasEqual = errors.New("Error: Env name has equals")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	var filesInDir []fs.DirEntry
	var err error

	// Read all files from directory or retur err if somthing wrong
	if filesInDir, err = os.ReadDir(dir); err != nil {
		return nil, err
	}

	// Make map with envs
	envs := make(Environment, len(filesInDir))

	var envName string

	// Check and fill envs
	for _, fEnv := range filesInDir {
		if fEnv.IsDir() {
			continue
		}

		envName = fEnv.Name()
		if strings.Contains(envName, "=") {
			return nil, ErrEnvNameHasEqual
		}

		fInfo, err := fEnv.Info()
		if err != nil {
			return nil, err
		}

		env, err := getValFromFile(path.Join(dir, envName))
		if err != nil {
			return nil, err
		}

		envs[envName] = EnvValue{
			Value:      env,
			NeedRemove: fInfo.Size() == 0,
		}
	}

	return envs, nil
}

func getValFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	env := strings.Builder{}
	buff := make([]byte, 1)
	for {
		count, err := file.Read(buff)
		// If end string or \n then stop read file
		if err == io.EOF || buff[0] == uint8('\n') || count == 0 {
			break
		}
		if err != nil {
			return "", err
		}

		// If symbol which one ' ', '\t' that continue
		// if buff[0] == ' ' || buff[0] == '\t' {
		// 	continue
		// }
		// If symbol equal \x00 that need to insert \n
		if buff[0] == '\x00' {
			buff[0] = '\n'
		}

		env.WriteByte(buff[0])
	}

	return strings.Trim(env.String(), " \t\n"), nil
}
