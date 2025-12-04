package file

import (
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Repository struct {
	FilePath string
}

func NewRepository(filePath string) *Repository {
	return &Repository{FilePath: filePath}
}

func (r *Repository) ListFolder(path string) ([]fs.DirEntry, error) {
	return os.ReadDir(filepath.Join(r.FilePath, path))
}

func (r *Repository) CreateFolder(path string) error {
	return os.Mkdir(filepath.Join(r.FilePath, path), 0755)
}

func (r *Repository) Delete(path string) error {
	return os.RemoveAll(filepath.Join(r.FilePath, path))
}

func (r *Repository) UploadFile(path string, fileHeader *multipart.FileHeader) error {

	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(r.FilePath, path, fileHeader.Filename))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}

	return nil
}
