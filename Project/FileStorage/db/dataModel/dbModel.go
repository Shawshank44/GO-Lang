package datamodel

import (
	"os"
	"sync"

	"github.com/gofrs/flock"
)

type DB struct {
	Path    string
	Delim   rune
	Mu      sync.RWMutex
	NextID  int64
	Perm    os.FileMode
	LoadMax bool

	Filelock *flock.Flock // 3rd party dependency
}
