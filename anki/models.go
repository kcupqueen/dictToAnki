package anki

type SimpleNoteRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  struct {
		Note *Note `json:"note"`
	} `json:"params"`
}

type Note struct {
	DeckName  string `json:"deckName"`
	ModelName string `json:"modelName"`
	Fields    struct {
		Front string `json:"Front"`
		Back  string `json:"Back"`
	} `json:"fields"`
	Tags []string `json:"tags"`
}

type SimpleNoteResponse struct {
	Result int64       `json:"result"`
	Error  interface{} `json:"error"`
}
