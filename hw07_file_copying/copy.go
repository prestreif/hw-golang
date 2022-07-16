package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	// This error doesn't need because
	// if offset equal zero, that app works good
	// else app will be able to return error ErrOffsetExceedsFileSize
	ErrUnsupportedFile = errors.New("unsupported file")
	// This error will be able to return if offset is set and offset more then size file or size file isn't possible counting
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Open In file for read
	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer fileFrom.Close()

	// Get info about size file
	infoFileFrom, err := fileFrom.Stat()
	if err != nil {
		return err
	}
	sizeFileFrom := infoFileFrom.Size()

	// If size file less then offset
	if offset > 0 && sizeFileFrom < offset {
		return ErrOffsetExceedsFileSize
	}

	// Create out file for write
	fileOut, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	// Set offset file
	if _, err := fileFrom.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	// Get reader from In file
	reader := io.Reader(fileFrom)

	// if limit not equal zero that set limit for reader
	if limit != 0 {
		reader = io.LimitReader(reader, limit)
	}

	// create progress bar
	bar := pb.Full.Start64(sizeFileFrom)
	defer bar.Finish()

	// wrapp reader into progress bar
	barReader := bar.NewProxyReader(reader)

	// use base function for copy
	io.Copy(fileOut, barReader)

	// return nil if everythin ok
	return nil
}
