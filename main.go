package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Welcome to File System Utility Command Line Tool")
	//reader := bufio.NewReader(os.Stdin)
	//makeDirectory(reader)
}

func makeDirectory(reader *bufio.Reader) {
	fmt.Print("Enter directory name: ")
	dirName, _ := reader.ReadString('\n')
	dirName = trimNewline(dirName)

	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func moveFile(reader *bufio.Reader) {
	fmt.Print("Enter source file name: ")
	srcFileName, _ := reader.ReadString('\n')
	srcFileName = trimNewline(srcFileName)
	fmt.Print("Enter destination directory: ")
	destDir, _ := reader.ReadString('\n')
	destDir = trimNewline(destDir)

	// Error handling
	// moves file using rename function
	err := os.Rename(srcFileName, filepath.Join(destDir, srcFileName))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File moved successfully.")
}

// input = string, output = string, returns a sliced string
func trimNewline(s string) string {
	return s[:len(s)-1]
}
