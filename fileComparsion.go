package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	files := []*os.File{}

	fmt.Print("Type file's names to compare their MD5 hash sums: ")
	scanner.Scan()
	filenames := scanner.Text()

	filename := strings.Fields(filenames)

	for _, name := range filename {
		file, err := os.Open(name)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		files = append(files, file)
	}

	checkSums := []string{}
	for _, file := range files {
		file.Seek(0, 0)
		sum, err := getMD5SumString(file)
		if err != nil {
			panic(err)
		}

		checkSums = append(checkSums, sum)
	}

	compareCheckSum(checkSums[0], checkSums[1])
}

func getMD5SumString(f *os.File) (string, error) {
	file1Sum := md5.New()
	_, err := io.Copy(file1Sum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", file1Sum.Sum(nil)), nil
}

func compareCheckSum(sum1, sum2 string) {
	match := "match"
	if sum1 != sum2 {
		match = " doesn't match"
	}
	fmt.Printf("MD5: %s and MD5: %s %s\n", sum1, sum2, match)
}
