package application

type CreationTime struct {
	Day   string
	Month string
	Year  string
}

type FileInfo struct {
	Name         string
	Keyword      string
	CreationTime CreationTime
}

type FileIndex []FileInfo
