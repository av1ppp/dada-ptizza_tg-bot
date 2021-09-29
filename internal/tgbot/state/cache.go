package state

import "sync"

var cache = []*DialogState{}
var mu sync.Mutex

func saveToCache(ds *DialogState) {
	mu.Lock()
	defer mu.Unlock()

	cache = append(cache, ds)
}

func getFromCache(usedID int) (*DialogState, bool) {
	mu.Lock()
	defer mu.Unlock()

	for _, ds := range cache {
		if ds.UserID == usedID {
			return ds, true
		}
	}
	return nil, false
}
