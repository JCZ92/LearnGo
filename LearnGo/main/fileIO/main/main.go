package main
import (
	"LearnGo/LearnGo/main/fileIO"
)


func main() {
	// myFile.ReadFile("../../../../README.md")
	// myFile.ReadFileWithBuffer("../../../../README.md")
	// myFile.WriteFile("../../../../README.md", "hello")
	myFile.CopyFile("../../../../README.md", "../../../../README2.md")
}