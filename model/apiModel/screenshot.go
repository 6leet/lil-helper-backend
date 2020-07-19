package apimodel

type AuditScreenshotParams struct {
	UserID    string `json:"userID" example:"screenshot_userID"`
	MissionID string `json:"missionID" example:"screenshot_missionID"`
	Approve   bool   `json:"approve"`
}

type SetScreenshotParams struct {
	MissionID string `json:"missionID" example:"screenshot_missionID"`
	Picture   string `json:"missionID" example:"this/is/a/path/of/picture.jpg"`
}
