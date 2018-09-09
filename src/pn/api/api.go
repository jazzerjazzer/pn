package api

import "sync"

type API struct {
	currentFilepath string
	currentFilename string
	sync.Mutex
}
