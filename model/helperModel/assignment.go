package helpermodel

import (
	"fmt"
	"lil-helper-backend/db"
	"lil-helper-backend/pkg/e"

	"github.com/jinzhu/gorm"
)

func CreateAssignment(userID uint, missionID uint) error {
	assignment := Assignment{
		UserID:    userID,
		MissionID: missionID,
	}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()

	if err := tx.Create(&assignment).Error; err != nil {
		return fmt.Errorf("assignment creation failed: %w", err)
	}
	tx.Commit()
	return nil
}

func GetAssignment(userID uint) (*Assignment, error) {
	assignment := Assignment{}
	query := db.LilHelperDB
	if err := query.Where("user_id = ?", userID).First(&assignment).Error; err == gorm.ErrRecordNotFound {
		return nil, e.ErrAssignmentNotExist
	} else if err != nil {
		return nil, fmt.Errorf("queyr assignment failed: %w", err)
	}
	return &assignment, nil
}

func DeleteAssignmentByMission(missionID uint) error {
	assignments := Assignment{}
	tx := db.LilHelperDB.Begin()
	defer tx.RollbackUnlessCommitted()
	if err := tx.Where("mission_id = ?", missionID).Find(&assignments).Error; err != nil {
		return fmt.Errorf("query assignment failed: %w", err)
	}
	if err := tx.Delete(&assignments).Error; err != nil {
		return fmt.Errorf("assignment deletion failed: %w", err)
	}
	tx.Commit()
	return nil
}
