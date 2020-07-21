package apimodel

type AuditScreenshotParams struct {
	Approve bool `json:"approve"`
}

type SetScreenshotParams struct {
	MissionUID string `json:"missionUID" example:"screenshot_missionUID"`
	Picture    string `json:"picture" example:"this/is/a/path/of/picture.jpg"`
}
