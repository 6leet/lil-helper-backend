package helpermodel

import (
	"encoding/json"
	"fmt"
	"lil-helper-backend/config"
	"lil-helper-backend/db"
	"lil-helper-backend/hashids"
	"lil-helper-backend/pkg/e"

	"github.com/jinzhu/gorm"
	"github.com/jmcvetta/randutil"
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
		Date:    m.CreatedAt.String()[0:10],
		Active:  m.Active,
	}
	return p
}

func CreateMission(userID uint, content string, picture string, weightstr string, score int) (*Mission, error) {
	mission := Mission{
		Content: content,
		Picture: picture,
		Weight:  weightstr,
		Score:   score,
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
	var weight []int
	if err := json.Unmarshal([]byte(mission.Weight), &weight); err != nil {
		return nil, fmt.Errorf("json unmarshal weight failed: %w", err)
	}
	SetTotalMissionWeight(weight, 1)
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

func DeleteMission(id uint) error {
	mission := Mission{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&mission, id).Error; err != nil {
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

func ActivateMission(id uint, active bool) (*Mission, error) {

	mission := Mission{}
	updateMission := Mission{
		Active: active,
	}

	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&mission, id).Error; err != nil {
		return nil, fmt.Errorf("mission query failed: %w", err)
	}
	if err := tx.Model(&mission).Update(updateMission).Error; err != nil {
		return nil, fmt.Errorf("mission activation failed: %w", err)
	}
	tx.Commit()
	var weight []int
	if err := json.Unmarshal([]byte(mission.Weight), &weight); err != nil {
		return nil, fmt.Errorf("json unmarshal weight failed: %w", err)
	}
	addvar := 1
	if !active {
		addvar = -1
	}
	SetTotalMissionWeight(weight, addvar)
	return &mission, nil
}

func GetMissionsWeight(level uint) ([]randutil.Choice, error) {
	choices := []randutil.Choice{}
	missions := []Mission{}

	query := db.LilHelperDB
	if err := query.Where("active = ?", true).Find(&missions).Error; err != nil {
		return nil, fmt.Errorf("query mission failed: %w", err)
	}
	for _, m := range missions {
		var weight []int
		if err := json.Unmarshal([]byte(m.Weight), &weight); err != nil {
			return nil, fmt.Errorf("json unmarshal weight failed: %w", err)
		}
		choice := randutil.Choice{
			Weight: weight[level],
			Item:   m,
		}
		choices = append(choices, choice)
	}
	return choices, nil
}

func SetTotalMissionWeight(weight []int, addvar int) error {
	VTool := config.VTool
	config := config.Config.Mission
	for i := 0; i <= config.Maxlevel; i++ {
		config.Totalweight[i] = config.Totalweight[i] + weight[i]*addvar
	}
	VTool.Set("mission.totalweight", config.Totalweight)
	VTool.WriteConfig()
	return nil
}

func GetMission(id uint) (*Mission, error) {
	mission := Mission{}

	query := db.LilHelperDB.Where("id = ?", id)
	if err := query.First(&mission).Error; err == gorm.ErrRecordNotFound {
		return nil, e.ErrMissionNotExist
	} else if err != nil {
		return nil, fmt.Errorf("Mission query failed: %w", err)
	} else {
		return &mission, nil
	}
}
