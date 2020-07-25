package apimodel

type SetMissionParams struct {
	Content    string `json:"content" example:"this is a content"`
	Picture    string `json:"picture" example:"this/is/a/path/of/picture.jpg"`
	Weight     []int  `json:"weight"`
	Score      int    `json:"score"`
	Active     bool   `json:"active"`
	Activeat   string `json:"active_at" example:"2020-02-02"`
	Inactiveat string `json:"inactive_at" example:"2020-02-02"`
}
