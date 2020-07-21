package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/hashids"
	"time"
)

type PublicMission struct {
	UID     string `json:"missionUID"`
	Content string `json:"content"`
	Picture string `json:"picturePath"`
	Weight  string `json:"weight"`
	Score   int    `json:"score"`
	Date    string `json:"date"`
	Active  bool   `json:"active"`
}

func (m *Mission) Public() PublicMission {
	p := PublicMission{
		UID:     m.UID,
		Content: m.Content,
		Picture: m.Picture,
		Weight:  m.Weight,
		Score:   m.Score,
		Date:    m.Date.String()[0:10],
		Active:  m.Active,
	}
	return p
}

func CreateMission(userID uint, content string, picture string, weight string, score int) (*Mission, error) {
	mission := Mission{
		Content: content,
		Picture: picture,
		Weight:  weight,
		Score:   score,
		Date:    time.Now(),
		Active:  true,
	}

	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Create(&mission).Error; err != nil {
		return nil, fmt.Errorf("mission creation failed: %w", err)
	}
	uid, err := hashids.EncodeMissionUID(mission.ID)
	if err != nil {
		return nil, fmt.Errorf("mission uid generation failed: %w", err)
	}
	if err := tx.Model(&mission).Update("uid", uid).Error; err != nil {
		return nil, fmt.Errorf("mission uid update failed: %w", err)
	}
	tx.Commit()
	return &mission, nil
}

func UpdateMission(id uint, content string, picture string, weight string, score int, active bool) (*Mission, error) {
	mission := Mission{}
	updateMission := map[string]interface{}{
		"content": content,
		"picture": picture,
		"weight":  weight,
		"score":   score,
		"active":  active,
	}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&mission, id).Error; err != nil {
		return nil, fmt.Errorf("mission query failed: %w", err)
	}
	if err := tx.Model(&mission).Updates(updateMission).Error; err != nil {
		return nil, fmt.Errorf("mission update failed: %w", err)
	}
	tx.Commit()
	return &mission, nil
}

func DeleteMission(missionID uint) error {
	mission := Mission{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&mission, missionID).Error; err != nil {
		return fmt.Errorf("mission query failed: %w", err)
	}
	if err := tx.Delete(&mission).Error; err != nil {
		return fmt.Errorf("mission deletion failed: %w", err)
	}
	tx.Commit()
	return nil
}

func GetMissionsByDate(dateFrom string, dateTo string) ([]Mission, error) {
	missions := []Mission{}

	query := db.LilHelperDB
	if err := query.Where("created_at BETWEEN ? and ?", dateFrom, dateTo).Find(&missions).Error; err != nil {
		return nil, fmt.Errorf("query missions failed: %w", err)
	}
	return missions, nil
}
