package myFile

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func ReadFile(fileName string) {

	// note: function ReadFile already contains file open/close operations
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func ReadFileWithBuffer(fileName string) {
	// with buffer = 4096 bytes
	file, err := os.Open(fileName) // file is the pointer to the file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()// close the file when the io operations end

	// define a input stream
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(str)
	}
}

func WriteFile(fileName string, strToWrite string) {
	// note that we will need to call OpenFile, not ReadFile or Open
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	// 0666 here is only for creating new file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// define a output stream with buffer
	writer := bufio.NewWriter(file)
	// write to the buffer
	writer.WriteString(strToWrite)
	// flush the buffer
	writer.Flush()
}	

func CopyFile(srcFileName string, dstFileName string) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dstFile.Close()

	// define a output stream with buffer
	writer := bufio.NewWriter(dstFile)
	// write to the buffer
	io.Copy(writer, srcFile)
	// flush the buffer
	writer.Flush()
}