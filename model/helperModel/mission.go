package helpermodel

import (
	"encoding/json"
	"fmt"
	"lil-helper-backend/config"
	"lil-helper-backend/db"
	"lil-helper-backend/goroutine"
	"lil-helper-backend/hashids"
	"lil-helper-backend/pkg/e"
	"lil-helper-backend/pkg/utils"
	"lil-helper-backend/scheduler"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jmcvetta/randutil"
)

type PublicMission struct {
	UID        string `json:"missionUID"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Picture    string `json:"picturePath"`
	Weight     string `json:"weight"`
	Score      int    `json:"score"`
	Activeat   string `json:"activeat"`
	Inactiveat string `json:"inactiveat"`
	Active     bool   `json:"active"`
}

type MissionsStat struct {
	Active   []string `json:"active"`
	Inactive []string `json:"inactive"`
}

func (m *Mission) Public() PublicMission {
	p := PublicMission{
		UID:        m.UID,
		Title:      m.Title,
		Content:    m.Content,
		Picture:    m.Picture,
		Weight:     m.Weight,
		Score:      m.Score,
		Activeat:   m.Activeat.String()[0:10],
		Inactiveat: m.Inactiveat.String()[0:10],
		Active:     m.Active,
	}
	return p
}

func CreateMission(userID uint, title string, content string, weightstr string, score int, activeat string, inactiveat string) (*Mission, error) {
	mission := Mission{
		Title:      title,
		Content:    content,
		Weight:     weightstr,
		Score:      score,
		Active:     false,
		Activeat:   utils.ParseTime(activeat),
		Inactiveat: utils.ParseTime(inactiveat),
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
	// SetTotalMissionWeight(weight, 1)
	return &mission, nil
}

func AddMissionPath(id uint, picture string) (*Mission, error) {
	mission := Mission{}
	updateMission := map[string]interface{}{
		"picture": picture,
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

func UpdateMission(id uint, title string, content string, weight string, score int, activeat string, inactiveat string) (*Mission, error) {
	mission := Mission{}
	updateMission := map[string]interface{}{
		"title":      title,
		"content":    content,
		"weight":     weight,
		"score":      score,
		"activeat":   utils.ParseTime(activeat),
		"inactiveat": utils.ParseTime(inactiveat),
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

	utils.ParseTimeLocation(&mission.Activeat)
	utils.ParseTimeLocation(&mission.Inactiveat)
	return &mission, nil
}

func DeleteMission(id uint) error {
	VTool := config.VTool
	config := config.Config.Mission
	mission := Mission{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.First(&mission, id).Error; err != nil {
		return fmt.Errorf("mission query failed: %w", err)
	}
	RemoveFile(mission.Picture)
	if err := tx.Delete(&mission).Error; err != nil {
		return fmt.Errorf("mission deletion failed: %w", err)
	}
	tx.Commit()
	_, ok := config.Weight[mission.UID]
	if ok {
		delete(config.Weight, mission.UID)
	}
	DeleteAssignmentByMission(mission.ID)
	VTool.Set("mission.totalweight", config.Weight)
	VTool.WriteConfig()
	return nil
}

func GetMissionsByDate(dateFrom string, dateTo string, titleKeyword string, contentKeyword string) ([]Mission, error) {
	missions := []Mission{}

	query := db.LilHelperDB
	query = query.Where("title LIKE ? AND content LIKE ?", titleKeyword, contentKeyword)
	if err := query.Where("(? BETWEEN activeat and inactiveat) or (? BETWEEN activeat and inactiveat) or (activeat BETWEEN ? and ?)", dateFrom, dateTo, dateFrom, dateTo).Find(&missions).Error; err != nil {
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
	// addvar := 1
	// if !active {
	// 	addvar = -1
	// }
	// SetTotalMissionWeight(weight, addvar)
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
		if len(weight) <= int(level) {
			continue
		}
		choice := randutil.Choice{
			Weight: weight[level],
			Item:   m,
		}
		choices = append(choices, choice)
	}
	return choices, nil
}

// func SetTotalMissionWeight(weight []int, addvar int) error {
// 	VTool := config.VTool
// 	config := config.Config.Mission
// 	for i := 0; i <= config.Maxlevel; i++ {
// 		config.Totalweight[i] = config.Totalweight[i] + weight[i]*addvar
// 	}
// 	VTool.Set("mission.totalweight", config.Totalweight)
// 	VTool.WriteConfig()
// 	return nil
// }

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

func ReorganizeMission() (*MissionsStat, error) {
	VTool := config.VTool
	config := config.Config.Mission
	activeMissions := []Mission{}
	inactiveMissions := []Mission{}
	currentDate := time.Now().String()[0:10]
	fmt.Println(currentDate)

	var active, inactive []string

	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	activeQuery := tx.Where("activeat = ?", currentDate)
	if err := activeQuery.Error; err != nil {
		return nil, fmt.Errorf("Missions query failed: %w", err)
	}
	if err := activeQuery.Model(&activeMissions).Update("active", true).Error; err != nil {
		return nil, fmt.Errorf("Missions activation failed: %w", err)
	}
	activeQuery.Find(&activeMissions)
	for _, m := range activeMissions {
		var weight []int
		if err := json.Unmarshal([]byte(m.Weight), &weight); err != nil {
			return nil, fmt.Errorf("json unmarshal weight failed: %w", err)
		}
		config.Weight[m.UID] = weight
		active = append(active, m.UID)
	}

	inactiveQuery := tx.Where("inactiveat = ?", currentDate)
	if err := inactiveQuery.Error; err != nil {
		return nil, fmt.Errorf("Missions query failed: %w", err)
	}
	if err := inactiveQuery.Model(&inactiveMissions).Update("active", false).Error; err != nil {
		return nil, fmt.Errorf("Missions inactivation failed: %w", err)
	}
	inactiveQuery.Find(&inactiveMissions)
	for _, m := range inactiveMissions {
		_, ok := config.Weight[m.UID]
		if ok {
			delete(config.Weight, m.UID)
		}
		inactive = append(inactive, m.UID)
		DeleteAssignmentByMission(m.ID)
	}

	VTool.Set("mission.totalweight", config.Weight)
	VTool.WriteConfig()
	tx.Commit()

	stat := MissionsStat{
		Active:   active,
		Inactive: inactive,
	}

	return &stat, nil
}

func AutoReorganizeMission() error {
	updateat := config.Config.Mission.Updateat
	fmt.Println("do auto at", updateat)
	scheduler.Cron.Every(1).Day().At(updateat).Do(func() {
		fmt.Println("auto-reorganizing mission")
		ReorganizeMission()
	})
	<-scheduler.Cron.Start()
	goroutine.Wg.Done()
	return nil
}
