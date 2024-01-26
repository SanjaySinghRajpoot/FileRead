package controller

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ErrLineNotFound represents an error when the requested line is not found in the file.
var ErrLineNotFound = fmt.Errorf("line not found")

// readLine reads a specific line from the given file.
func readLine(file *os.File, lineNumber int) (string, error) {
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		if i == lineNumber {
			return scanner.Text(), nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", ErrLineNotFound
}

// @Summary 	Get File data for a given param
// @Param		n path string true "file name"
// @Description Returns the File Content or the desired line
// @Produce     application/json
// @Success     200 {string} string  "ok"
// @Router      /data?n={file_name}&m={line_number} [get]
func GetData(ctx *gin.Context) {
	n := ctx.Query("n")
	m := ctx.Query("m")

	// Validate inputs
	if n == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parameter 'n' is required"})
		return
	}

	filePath := fmt.Sprintf("tmp/data/%s.txt", n)

	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file - " + err.Error(),
		})
		return
	}
	defer file.Close()

	if m != "" {
		lineNumber, err := strconv.Atoi(m)
		if err != nil || lineNumber <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid line number", "details": err.Error()})
			return
		}

		// Read specific line
		line, err := readLine(file, lineNumber)
		if err != nil {
			if err == ErrLineNotFound {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Line not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file", "details": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, line)
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename="+n+".txt")
	ctx.Header("Content-Type", "text/plain")
	ctx.File(filePath)
}
