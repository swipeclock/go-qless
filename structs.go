package qless

//easyjson:json
type jobData struct {
	Jid          string
	Klass        string
	State        string
	Queue        string
	Worker       string
	Tracked      bool
	Priority     int
	Expires      int64
	Retries      int
	Remaining    int
	Data         interface{}
	Tags         StringSlice
	History      []History
	Failure      interface{}
	Dependents   StringSlice
	Dependencies interface{}
}

