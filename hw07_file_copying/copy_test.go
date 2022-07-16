package main

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testFiles struct {
	fromPath, toPath, cmpFile string
	offset, limit             int64
	expect                    interface{}
}

func TestCopy(t *testing.T) {
	tests := []testFiles{
		{"", "", "", 0, 0, "open : no such file or directory"},
		{"testdata/input.txt", "", "", 0, 0, "open : no such file or directory"},
		{"testdata/input.txt", "testdata/res_out_offset0_limit0.txt", "testdata/out_offset0_limit0.txt", 0, 0, nil},
		{"testdata/input.txt", "testdata/res_out_offset0_limit0.txt", "testdata/out_offset0_limit0.txt", 7000, 0, ErrOffsetExceedsFileSize.Error()},
		{"/dev/urandom", "testdata/res_out_urandom_offset0_limit1000.txt", "testdata/out_offset0_limit1000.txt", 1, 1000, ErrOffsetExceedsFileSize.Error()},
		{"testdata/input.txt", "testdata/res_out_offset0_limit1000.txt", "testdata/out_offset0_limit1000.txt", 0, 1000, nil},
		{"testdata/input.txt", "testdata/res_out_offset0_limit10000.txt", "testdata/out_offset0_limit10000.txt", 0, 10000, nil},
		{"testdata/input.txt", "testdata/res_out_offset100_limit1000.txt", "testdata/out_offset100_limit1000.txt", 100, 1000, nil},
		{"testdata/input.txt", "testdata/res_out_offset6000_limit1000.txt", "testdata/out_offset6000_limit1000.txt", 6000, 1000, nil},
	}
	t.Run("Test open files", func(t *testing.T) {
		for _, val := range tests {
			if err := Copy(val.fromPath, val.toPath, val.offset, val.limit); err != nil {
				val_err, _ := val.expect.(string)
				assert.EqualError(t, err, val_err)
			} else {
				assert.Equal(t, val.expect, err)
			}
		}
	})

	t.Run("Test compare size created files", func(t *testing.T) {
		var err error
		var cmpFile, toFile *os.File
		var cmpFileInfo, toFileInfo fs.FileInfo

		for _, val := range tests[5:] {
			require.FileExists(t, val.cmpFile)
			require.FileExists(t, val.toPath)

			cmpFile, err = os.Open(val.cmpFile)
			require.NoError(t, err)
			cmpFileInfo, err = cmpFile.Stat()
			require.NoError(t, err)

			toFile, err = os.Open(val.toPath)
			require.NoError(t, err)
			toFileInfo, err = toFile.Stat()
			require.NoError(t, err)

			require.Equal(t, cmpFileInfo.Size(), toFileInfo.Size())

			textToFile, err := io.ReadAll(toFile)
			require.NoError(t, err)

			textCmpFile, err := io.ReadAll(cmpFile)
			require.NoError(t, err)

			require.Equal(t, 0, bytes.Compare(textCmpFile, textToFile), val)
		}
	})

	t.Run("Clear all created files after tests", func(t *testing.T) {
		for _, val := range tests {
			os.Remove(val.toPath)
		}
	})
}
