package apimodel

type SetMissionParams struct {
	Title      string `json:"title" example:"this is a title"`
	Content    string `json:"content" example:"this is a content"`
	Picture    string //path
	Weight     []int  `json:"weight"`
	Score      int    `json:"score"`
	Active     bool   `json:"active"`
	Activeat   string `json:"active_at" example:"2020-02-02"`
	Inactiveat string `json:"inactive_at" example:"2020-02-02"`
}

type SetAutoTimeParams struct {
	Hour   int `json:"hour`
	Minute int `json:"minute"`
}
