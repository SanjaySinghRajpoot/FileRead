package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all Users from the Database
func GetData(ctx *gin.Context) {
	n := ctx.Query("n")
	m := ctx.Query("m")

	fmt.Println(n)
	fmt.Println(m)

	filePath := fmt.Sprintf("tmp/data/%s.txt", n)

	file, err := os.Open(filePath)
	if err != nil {
		ctx.String(500, "Failed to open file")
		return
	}
	defer file.Close()

	if m != "" {

		lineNumber, err := strconv.Atoi(m)
		if err != nil || lineNumber <= 0 {
			ctx.String(400, "Invalid line number")
			return
		}

		// Create a scanner to read the file line by line
		scanner := bufio.NewScanner(file)
		line := ""
		currentLineNumber := 1
		for scanner.Scan() {
			if currentLineNumber == lineNumber {
				line = scanner.Text()
				break
			}
			currentLineNumber++
		}
		if err := scanner.Err(); err != nil {
			ctx.String(500, "Failed to read file")
			return
		}

		// If the line number is greater than the number of lines in the file, return not found
		if line == "" {
			ctx.String(404, "Line not found")
			return
		}

		// Send the line as a response
		ctx.String(200, line)
	}

	// Wrap the file reader with a buffered reader for more efficient reading
	reader := bufio.NewReader(file)

	// Copy the file to the response writer using bufio.Writer
	// This will stream the file in chunks instead of loading it into memory
	_, err = reader.WriteTo(ctx.Writer)
	if err != nil {
		ctx.String(500, err.Error())
		return
	}

	// Send the file content as a response
	ctx.String(200, "")
	return

}
