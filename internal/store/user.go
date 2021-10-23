package store

type Sex string

const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

type SocialNetwork string

const (
	SocialNetworkVK    = "vk"
	SocialNetworkInsta = "insta"
)

type User struct {
	ID int64

	FirstName     string
	LastName      string
	Sex           Sex
	SocialNetwork SocialNetwork
	URL           string
	Picture       []byte

	CountPrivatePhotos int
	CountPrivateVideos int
	CountHiddenData    int
	CountHiddenFriends int
}

func (store *Store) SaveUser(user *User) error {
	result, err := store.db.Exec(
		`INSERT
			INTO users(
				first_name,
				last_name,
				sex,
				social_network,
				url,
				picture,
				count_private_photos,
				count_private_videos,
				count_hidden_data,
				count_hidden_friends
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		user.FirstName,
		user.LastName,
		user.Sex,
		user.SocialNetwork,
		user.URL,
		user.Picture,
		user.CountPrivatePhotos,
		user.CountPrivateVideos,
		user.CountHiddenData,
		user.CountHiddenFriends,
	)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = lastID
	return nil
}

func (store *Store) GetUserByID(id int64) (*User, error) {
	var user User

	row := store.db.QueryRow(
		`SELECT
			id,
			first_name,
			last_name,
			sex,
			social_network,
			url,
			picture,
			count_private_photos,
			count_private_videos,
			count_hidden_data,
			count_hidden_friends
		FROM users WHERE id = $1`,
		id,
	)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Sex,
		&user.SocialNetwork,
		&user.URL,
		&user.Picture,
		&user.CountPrivatePhotos,
		&user.CountPrivateVideos,
		&user.CountHiddenData,
		&user.CountHiddenFriends,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store *Store) GetUserByURL(url string) (*User, error) {
	var user User

	row := store.db.QueryRow(
		`SELECT
			id,
			first_name,
			last_name,
			sex,
			social_network,
			url,
			picture,
			count_private_photos,
			count_private_videos,
			count_hidden_data,
			count_hidden_friends
		FROM users WHERE url = $1`,
		url,
	)
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Sex,
		&user.SocialNetwork,
		&user.URL,
		&user.Picture,
		&user.CountPrivatePhotos,
		&user.CountPrivateVideos,
		&user.CountHiddenData,
		&user.CountHiddenFriends,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUserByID - обновление user
func (store *Store) UpdateUserByID(u *User) error {
	_, err := store.db.Exec(
		`UPDATE users
			SET
				first_name = $1,
				last_name = $2,
				sex = $3,
				social_network = $4,
				url = $5,
				picture = $6,
				count_private_photos = $7,
				count_private_videos = $8,
				count_hidden_data = $9,
				count_hidden_friends = $10
			WHERE
				id = $11`,
		&u.FirstName,
		&u.LastName,
		&u.Sex,
		&u.SocialNetwork,
		&u.URL,
		&u.Picture,
		&u.CountPrivatePhotos,
		&u.CountPrivateVideos,
		&u.CountHiddenData,
		&u.CountHiddenFriends,
		&u.ID)

	return err
}

func (store *Store) DeleteUserByID(id int64) error {
	_, err := store.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (store *Store) UpdateOrSaveUser(u *User) error {
	_, err := store.GetUserByURL(u.URL)
	if err == nil {
		return store.UpdateUserByID(u)
	}
	return store.SaveUser(u)
}
