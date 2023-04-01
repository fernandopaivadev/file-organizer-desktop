//go:build linux

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

	fileStat := file.Sys().(*syscall.Stat_t)
	cTim := fileStat.Ctim
	creationTime := time.Unix(int64(cTim.Sec), int64(cTim.Nsec))

	return CreationTime{
		Day:   fmt.Sprint(creationTime.Day()),
		Month: months[creationTime.Month()-1],
		Year:  fmt.Sprint(creationTime.Year()),
	}
}
