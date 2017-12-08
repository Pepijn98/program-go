package pgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
)

const (
	baseURI   = "http://api.program-o.com/v2/chatbot"
	userAgent = "program-go/v0.0.1 - (github.com/KurozeroPB/program-go)"
)

func get(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", "application/vnd.api+json")
	request.Header.Set("Content-Type", "application/vnd.api+json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Expected status %d; Got %d\nResponse: %#v", 200, response.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}

// Response holds all the returned data from program-o
type Response struct {
	ConvoID string `json:"convo_id"`
	UserSay string `json:"usersay"`
	BotSay  string `json:"botsay"`
}

// Say sends the query to program-o
func Say(botID int, query string, convoID string) (*Response, error) {
	var (
		err     error
		res     []byte
		parJSON *gabs.Container
		resJSON []byte
	)

	const realBotIDs = "6, 10, 12, 15"
	if strings.Contains(realBotIDs, strconv.Itoa(botID)) != true {
		err = fmt.Errorf("Invalid bot id, either use 6, 10, 12 or 15.\nTo see what type of bots these are visit: https://program-o.com/v2/api.php")
		return nil, err
	}

	if len(query) > 255 {
		err = fmt.Errorf("The query cannot be longer then 255 characters")
		return nil, err
	}
	if len(convoID) > 128 {
		err = fmt.Errorf("The convo id cannot be longer then 128 characters")
		return nil, err
	}
	newQuery := url.QueryEscape(query)
	queryURL := fmt.Sprintf("%s?bot_id=%d&say=%s&convo_id=%s", baseURI, botID, newQuery, convoID)
	res, err = get(queryURL)
	if err != nil {
		return nil, err
	}
	parJSON, err = gabs.ParseJSON(res)
	if err != nil {
		return nil, err
	}
	data := parJSON.Data().(interface{})
	resJSON, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp := new(Response)
	err = json.Unmarshal(resJSON, &resp)
	if err != nil {
		return nil, err
	}

	return resp, err
}
