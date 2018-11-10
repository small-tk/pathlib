package pathlib

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// base Path
type Path struct {
	Path string
}

// New Returns a new path
func New(path string) *Path {
	p := new(Path)
	p.Path = path
	return p
}

// Absolute Returns an absolute representation of path
func (p *Path) Absolute() (string, error) {
	return filepath.Abs(p.Path)
}

// Cwd Return a new path pointing to the current working directory
func (p *Path) Cwd() (*Path, error) {
	path, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "get Cwd failed")
	}
	newP := New(path)
	return newP, nil
}

// Exists Return current path parent exists
func (p *Path) Exists() bool {
	_, err := os.Stat(p.Path)
	return err == nil || os.IsExist(err)
}

// Parent Return a new path for current path parent
func (p *Path) Parent() (*Path, error) {
	path, err := p.Absolute()
	if err != nil {
		return nil, errors.Wrap(err, "get Parent failed")
	}
	newP := New(path)
	return newP, nil
}
