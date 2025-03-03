package filehandler

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileHandler struct {
	InputFile  string
	OutputFile string
}

func NewFileHandler(inputFile string, outputFile string) *FileHandler {
	return &FileHandler{
		InputFile:  inputFile,
		OutputFile: outputFile,
	}
}

func (f *FileHandler) ReadLines() ([]string, error) {
	file, err := os.Open(f.InputFile)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("file not found")
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("error reading file")
	}

	file.Close()

	return lines, nil
}

func (f *FileHandler) WriteResult(data any) error {
	file, err := os.Create(f.OutputFile)

	if err != nil {
		return errors.New("error creating file")
	}

	err = json.NewEncoder(file).Encode(data)

	if err != nil {
		return errors.New("error writing to file")
	}

	file.Close()
	return nil
}
