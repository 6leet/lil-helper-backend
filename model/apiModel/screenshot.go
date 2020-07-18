package apimodel

type AuditScreenshotParams struct {
	UserID    string `json:"userID" example:"screenshot_userID"`
	MissionID string `json:"missionID example:"screenshot_missionID"`
	Approve   bool   `json:"approve"`
}
