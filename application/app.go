package application

import (
	"path/filepath"
	"strings"
)

func OrganizeFilesByYear(targetDirectoryPath string) error {
	initialDirectory, err := getCurrentDirectory()

	if err != nil {
		return err
	}

	err = changeDirectory(targetDirectoryPath)

	if err != nil {
		return err
	}

	fileIndex, err := listFiles()

	if err != nil {
		return err
	}

	for _, fileInfo := range fileIndex {
		err := createFolder(filepath.Join(
			fileInfo.CreationTime.Year,
		))

		if err != nil {
			return err
		}

		err = moveFile(
			fileInfo.Name,
			filepath.Join(
				fileInfo.CreationTime.Year,
				fileInfo.Name,
			),
		)

		if err != nil {
			return err
		}
	}

	err = changeDirectory(initialDirectory)

	if err != nil {
		return err
	}

	return nil
}

func OrganizeFilesByMonth(targetDirectoryPath string) error {
	initialDirectory, err := getCurrentDirectory()

	if err != nil {
		return err
	}

	err = changeDirectory(targetDirectoryPath)

	if err != nil {
		return err
	}

	fileIndex, err := listFiles()

	if err != nil {
		return err
	}

	for _, fileInfo := range fileIndex {
		err := createFolder(filepath.Join(
			fileInfo.CreationTime.Year,
			fileInfo.CreationTime.Month,
		))

		if err != nil {
			return err
		}

		err = moveFile(
			fileInfo.Name,
			filepath.Join(
				fileInfo.CreationTime.Year,
				fileInfo.CreationTime.Month,
				fileInfo.Name,
			),
		)

		if err != nil {
			return err
		}
	}

	err = changeDirectory(initialDirectory)

	if err != nil {
		return err
	}

	return nil
}

func OrganizeFilesByDay(targetDirectoryPath string) error {
	initialDirectory, err := getCurrentDirectory()

	if err != nil {
		return err
	}

	err = changeDirectory(targetDirectoryPath)

	if err != nil {
		return err
	}

	fileIndex, err := listFiles()

	if err != nil {
		return err
	}

	for _, fileInfo := range fileIndex {
		err := createFolder(filepath.Join(
			fileInfo.CreationTime.Year,
			fileInfo.CreationTime.Month,
			fileInfo.CreationTime.Day,
		))

		if err != nil {
			return err
		}

		err = moveFile(
			fileInfo.Name,
			filepath.Join(
				fileInfo.CreationTime.Year,
				fileInfo.CreationTime.Month,
				fileInfo.CreationTime.Day,
				fileInfo.Name,
			),
		)

		if err != nil {
			return err
		}
	}

	err = changeDirectory(initialDirectory)

	if err != nil {
		return err
	}

	return nil
}

func OrganizeFilesByName(targetDirectoryPath string) error {
	initialDirectory, err := getCurrentDirectory()

	if err != nil {
		return err
	}

	err = changeDirectory(targetDirectoryPath)

	if err != nil {
		return err
	}

	fileIndex, err := listFiles()

	if err != nil {
		return err
	}

	for index, fileInfo := range fileIndex {
		title := strings.ToUpper(strings.Split(fileInfo.Name, "_")[0])

		splittedTitle := strings.Split(title, " ")

		keyword := strings.ToUpper(splittedTitle[0])

		fileIndex[index].Keyword = keyword
	}

	for _, fileInfo := range fileIndex {
		err := createFolder(fileInfo.Keyword)

		if err != nil {
			return err
		}

		oldPath := fileInfo.Name
		newPath := filepath.Join(fileInfo.Keyword, fileInfo.Name)

		err = moveFile(oldPath, newPath)

		if err != nil {
			return err
		}
	}

	err = changeDirectory(initialDirectory)

	if err != nil {
		return err
	}

	return nil
}
