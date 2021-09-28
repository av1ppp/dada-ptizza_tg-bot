package state

type DialogState struct {
	UserID uint64
	State  State
}

func Get(userID uint64) *DialogState {
	ds, find := getFromCache(userID)
	if !find {
		ds = &DialogState{
			UserID: userID,
			State:  SELECT_SOCIAL_NETWORK,
		}
		saveToCache(ds)
	}
	return ds
}
