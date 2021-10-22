package tables

import "os"

func openFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}
