package qless

import "github.com/mailru/easyjson"

//easyjson:json
type jobData struct {
	JID          string              `json:"jid"`
	Class        string              `json:"klass"`
	State        string              `json:"state"`
	Queue        string              `json:"queue"`
	Worker       string              `json:"worker"`
	Tracked      bool                `json:"tracked"`
	Priority     int                 `json:"priority"`
	Expires      int64               `json:"expires"`
	Retries      int                 `json:"retries"`
	Remaining    int                 `json:"remaining"`
	Data         easyjson.RawMessage `json:"data"`
	Tags         StringSlice         `json:"tags"`
	History      []History           `json:"history"`
	Failure      interface{}         `json:"failure"`
	Dependents   StringSlice         `json:"dependents"`
	Dependencies StringSlice         `json:"dependencies"`
}

//easyjson:json
type History struct {
	When   int64  `json:"when"`
	Queue  string `json:"q"`
	What   string `json:"what"`
	Worker string `json:"worker"`
}

//easyjson:json
type QueueInfo struct {
	Name      string `json:"name"`
	Paused    bool   `json:"paused"`
	Waiting   int    `json:"waiting"`
	Running   int    `json:"running"`
	Stalled   int    `json:"stalled"`
	Scheduled int    `json:"scheduled"`
	Recurring int    `json:"recurring"`
	Depends   int    `json:"depends"`
}
