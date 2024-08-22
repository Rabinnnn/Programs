package piscine

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . -c N file1.txt [file2.txt ...]")
		os.Exit(1)
	}

	count := os.Args[2]
	files := os.Args[3:]

	for _, file := range files {
		fmt.Printf("==> %s <==\n", file)
		if err := tail(file, count); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(files) > 1 {
			fmt.Println()
		}
	}
}

func tail(filename, count string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	offset := stat.Size()
	size, err := parseCount(count, offset)
	if err != nil {
		return err
	}

	if size > offset {
		size = offset
	}

	_, err = file.Seek(offset-size, io.SeekStart)
	if err != nil {
		return err
	}

	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

	fmt.Print(string(buffer))
	return nil
}

func parseCount(count string, offset int64) (int64, error) {
	var size int64

	if count[len(count)-1] == '%' {
		percentage := count[:len(count)-1]
		var percent int
		_, err := fmt.Sscanf(percentage, "%d", &percent)
		if err != nil {
			return 0, err
		}

		size = offset * int64(percent) / 100
	} else {
		_, err := fmt.Sscanf(count, "%d", &size)
		if err != nil {
			return 0, err
		}
	}

	return size, nil
}
