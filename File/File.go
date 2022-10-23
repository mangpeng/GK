package File

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadFile(fileName string, onSuccess func(file *os.File), onFail func(err error)) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		onFail(err)
	} else {
		onSuccess(file)
	}
}

// AppendText If the file doesn't exist, create it, or append to the file
func AppendText(fileName string, texts ...string) error {

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	var data string
	for _, t := range texts {
		data += t + " "
	}
	data += "\n"

	if _, err := f.Write([]byte(data)); err != nil {
		return err
	}

	return nil
}

//func main() {
//	if len(os.Args) < 3 {
//		fmt.Println("arguments are less than 3")
//	}
//
//	command := os.Args[0]
//	word := os.Args[1]
//	fmt.Printf("%s %s\n", command, word)
//	files := os.Args[2:]
//
//	PrintAllFiles(files)
//}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("error")
			return
		}

		for _, name := range fileList {
			fmt.Println(name)
		}
	}
}
