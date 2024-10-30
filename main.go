package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Welcome to File System Utility Command Line Tool")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nEnter command (ls, mkdir, mv, exit): ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			continue
		}
		command = trimNewline(command)

		switch command {
		case "ls":
			listFiles()
		case "mkdir":
			makeDirectory(reader)
		case "mv":
			moveFile(reader)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command. Please try again")
		}
	}
}

func listFiles() {
	// retrieve a list of file names that match a pattern —, "*" means  all files in the current directory.
	files, err := filepath.Glob("*")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println("-", file)
	}
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
	err := os.Rename(srcFileName, filepath.Join(destDir, srcFileName))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File moved successfully.")
}

// input = string, output = string, returns a sliced string
func trimNewline(s string) string {
	return strings.TrimSpace(s)
}
