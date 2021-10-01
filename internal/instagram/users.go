package instagram

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserInfo struct {
	Biography       string `json:"biography"`
	FullName        string `json:"full_name"`
	ID              string `json:"id"`
	ProfilePicUrl   string `json:"profile_pic_url"`
	ProfilePicUrlHD string `json:"profile_pic_url_hd"`
	Username        string `json:"username"`
}

type GetUserInfoResponse struct {
	GraphQL struct {
		User UserInfo `json:"user"`
	} `json:"graphql"`
}

func (client *Client) GetUserInfo(username string) (*UserInfo, error) {
	urlString := fmt.Sprintf("https://instagram.com/%s?__a=1", username)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Cookie", "sessionid="+client.sessionid)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resBody GetUserInfoResponse

	if err := json.Unmarshal(data, &resBody); err != nil {
		return nil, err
	}

	return &resBody.GraphQL.User, nil
}
