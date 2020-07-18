package apimodel

type SetMissionParam struct {
	Content string `json:"content" example:"this is a content"`
	Picture string `json:"picture" example:"this/is/a/path/of/picture.jpg"`
	Weight  []int  `json:"weight"`
	Score   int    `json:"score"`
	Active  bool   `json:"active"`
}
