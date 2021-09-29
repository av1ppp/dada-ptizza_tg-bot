package state

type DialogState struct {
	UserID int
	State  State
}

func (ds *DialogState) IsSelectUser() bool {
	return ds.State == SELECT_USER_INSTAGRAM ||
		ds.State == SELECT_USER_VKONTAKTE ||
		ds.State == SELECT_USER_TELEGRAM ||
		ds.State == SELECT_USER_WHATSAPP ||
		ds.State == SELECT_USER_VIBER
}

func Get(userID int) *DialogState {
	ds, find := getFromCache(userID)
	if !find {
		ds = &DialogState{
			UserID: userID,
			State:  EMPTY,
		}
		saveToCache(ds)
	}
	return ds
}

func Save(userID int, state State) {
	ds, find := getFromCache(userID)
	if find {
		ds.State = state
	} else {
		saveToCache(&DialogState{
			UserID: userID,
			State:  state,
		})
	}
}
