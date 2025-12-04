package file

import (
	"io/fs"
	"mime/multipart"
)

type Filer interface {
	ListFolder(path string) ([]fs.DirEntry, error)
	CreateFolder(path string) error
	Delete(path string) error
	UploadFile(path string, fileHeader *multipart.FileHeader) error
}
