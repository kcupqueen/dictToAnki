package anki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
)

type ConnectClient struct {
	Endpoint string `json:"endpoint,omitempty"`
}

func NewConnectClient(endpoint string) *ConnectClient {
	return &ConnectClient{
		Endpoint: endpoint,
	}
}
func (c *ConnectClient) GetDecksNames() ([]string, error) {
	data := map[string]interface{}{
		"action":  "deckNames",
		"version": 6,
	}
	body, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(resBody))
	var resp struct {
		Result []string    `json:"result"`
		Error  interface{} `json:"error"`
	}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()
	if resp.Error != nil {
		return nil, fmt.Errorf("%v", resp.Error)
	}
	return resp.Result, nil
}

func (c *ConnectClient) AddNote(deckName, front, end string, tags []string) (int64, error) {
	note := &Note{
		DeckName:  deckName,
		ModelName: "Basic",
		Fields: struct {
			Front string `json:"Front"`
			Back  string `json:"Back"`
		}{
			Front: front,
			Back:  end,
		},
		Tags: tags,
	}
	noteReq := &SimpleNoteRequest{
		Action:  "addNote",
		Version: 6,
		Params: struct {
			Note *Note `json:"note"`
		}{
			Note: note,
		},
	}

	body, err := json.Marshal(&noteReq)
	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest(http.MethodPost, c.Endpoint, bytes.NewReader(body))
	if err != nil {
		return -1, err
	}

	res, err := http.DefaultClient.Do(req)
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	resp := &SimpleNoteResponse{}
	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		return -1, err
	}
	fmt.Println(resp)
	_ = res.Body.Close()

	return resp.Result, nil
}
