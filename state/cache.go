package state

var cache = []*DialogState{}

func saveToCache(ds *DialogState) {
	cache = append(cache, ds)
}

func getFromCache(usedID uint64) (*DialogState, bool) {
	for _, ds := range cache {
		if ds.UserID == usedID {
			return ds, true
		}
	}
	return nil, false
}
