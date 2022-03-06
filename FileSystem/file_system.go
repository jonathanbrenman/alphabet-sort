package filesystem

import (
	"bufio"
	"os"
)

type FileSystem interface {
	ReadFile() ([]string, error)
	WriteFile(data []string) error
}

type fs struct {
	Input  string
	Output string
}

func NewFS(input, output string) FileSystem {
	return &fs{input, output}
}

func (f *fs) ReadFile() ([]string, error) {
	file, err := os.Open(f.Input)
	if err != nil {
		return []string{}, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	list := []string{}

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	return list, nil
}

func (f *fs) WriteFile(data []string) error {
	file, err := os.OpenFile(f.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	datawriter := bufio.NewWriter(file)

	defer file.Close()
	defer datawriter.Flush()

	for _, data := range data {
		_, _ = datawriter.WriteString(data + "\n")
	}

	return nil
}
