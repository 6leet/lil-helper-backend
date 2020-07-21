package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/hashids"
)

type PublicScreenshot struct {
	UID        string `json:"screenshotUID"`
	UserUID    string `json:"userUID"`
	MissionUID string `json:"missionUID"`
	Picture    string `json:"picturePath"`
	Audit      bool   `json:"audit"`
	Approve    bool   `json:"approve"`
	Date       string `json:"date"`
}

func (s *Screenshot) Public() PublicScreenshot {
	p := PublicScreenshot{
		UID:     s.UID,
		Picture: s.Picture,
		Audit:   s.Audit,
		Approve: s.Approve,
		Date:    s.CreatedAt.String()[0:10],
	}
	p.UserUID, _ = hashids.EncodeUserUID(s.UserID)
	p.MissionUID, _ = hashids.EncodeMissionUID(s.MissionID)
	return p
}

func CreateScreenshot(userID uint, missionID uint, picture string) (*Screenshot, error) {
	screenshot := Screenshot{
		UserID:    userID,
		MissionID: missionID,
		Picture:   picture,
		Audit:     false,
		Approve:   false,
	}
	mission := Mission{}

	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Find(&mission, missionID).Error; err != nil {
		return nil, fmt.Errorf("query mission failed: %w", err)
	}
	if !mission.Active {
		return nil, fmt.Errorf("mission not active")
	}

	if err := tx.Create(&screenshot).Error; err != nil {
		return nil, fmt.Errorf("screenshot creation failed: %w", err)
	}
	uid, err := hashids.EncodeScreenshotUID(screenshot.ID)
	if err != nil {
		return nil, fmt.Errorf("screenshot uid generation failed: %w", err)
	}
	if err := tx.Model(&screenshot).Update("uid", uid).Error; err != nil {
		return nil, fmt.Errorf("mission uid update failed: %w", err)
	}
	tx.Commit()
	return &screenshot, nil
}

func GetScreenshotsByDate(userID uint, dateFrom string, dateTo string, audit bool, careAudit bool) ([]Screenshot, error) {
	screenshots := []Screenshot{}

	query := db.LilHelperDB
	if userID != 0 {
		query = query.Where("user_id = ?", userID)
	}
	query = query.Where("created_at BETWEEN ? and ?", dateFrom, dateTo)
	if careAudit {
		query = query.Where("audit = ?", audit)
	}
	if err := query.Find(&screenshots).Error; err != nil {
		return nil, fmt.Errorf("query screenshots failed: %w", err)
	}
	return screenshots, nil
}

func SetScreeshotApprove(id uint, approve bool) (*Screenshot, error) {
	screenshot := Screenshot{}
	updateScreenshot := map[string]interface{}{
		"audit":   true,
		"approve": approve,
	}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&screenshot, id).Error; err != nil {
		return nil, fmt.Errorf("screenshot query failed: %w", err)
	}

	userID := screenshot.UserID
	preapprove := screenshot.Approve
	missionID := screenshot.MissionID
	addvar := 0
	if preapprove && !approve {
		fmt.Println("case 1")
		addvar = -1
	} else if !preapprove && approve {
		fmt.Println("case 2")
		addvar = 1
	}
	if err := tx.Model(&screenshot).Updates(updateScreenshot).Error; err != nil {
		return nil, fmt.Errorf("screenshot update failed: %w", err)
	}
	tx.Commit()
	_, err := SetUserScore(userID, missionID, addvar)
	if err != nil {
		return nil, err
	}
	return &screenshot, nil
}

func DeleteScreenshot(screenshotID uint) error {
	screenshot := Screenshot{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&screenshot, screenshotID).Error; err != nil {
		return fmt.Errorf("screenshot query failed: %w", err)
	}
	if err := tx.Delete(&screenshot).Error; err != nil {
		return fmt.Errorf("screenshot deletion failed: %w", err)
	}
	tx.Commit()
	return nil
}
