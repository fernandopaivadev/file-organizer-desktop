//go:build windows

package application

import (
	"fmt"
	"io/fs"
	"syscall"
	"time"
)

func getFileCreationTime(file fs.FileInfo) CreationTime {
	var months = [12]string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	fileStat := file.Sys().(*syscall.Win32FileAttributeData)
	creationTime := time.Unix(0, fileStat.CreationTime.Nanoseconds())

	return CreationTime{
		Day:   fmt.Sprint(creationTime.Day()),
		Month: months[creationTime.Month()-1],
		Year:  fmt.Sprint(creationTime.Year()),
	}
}
