package application

import (
	"os"
)

func listFiles() (FileIndex, error) {
	reader, err := os.Open(".")

	if err != nil {
		return nil, err
	}

	files, err := reader.Readdir(0)

	if err != nil {
		return nil, err
	}

	var fileIndex FileIndex

	for _, file := range files {
		if !file.IsDir() {
			fileInfo := FileInfo{
				Name:         file.Name(),
				CreationTime: getFileCreationTime(file),
			}

			fileIndex = append(fileIndex, fileInfo)
		}
	}

	return fileIndex, nil
}

func createFolder(dirPath string) error {
	return os.MkdirAll("./"+dirPath, os.ModePerm)
}

func moveFile(sourcePath string, destinationPath string) error {
	return os.Rename("./"+sourcePath, "./"+destinationPath)
}

func changeDirectory(dirPath string) error {
	return os.Chdir(dirPath)
}

func getCurrentDirectory() (string, error) {
	return os.Getwd()
}
