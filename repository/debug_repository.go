package repository

import (
	"debugpedia-api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IDebugRepository interface {
	GetAllDebugs(debugs *[]model.Debug, userId uint) error
	GetDebugById(debug *model.Debug, userId uint, debugId uint) error
	CreateDebug(debug *model.Debug) error
	UpdateDebug(debug *model.Debug, userId uint, debugId uint) error
	DeleteDebug(userId uint, debugId uint) error
}

type debugRepository struct {
	db *gorm.DB
}

func NewDebugRepository(db *gorm.DB) IDebugRepository {
	return &debugRepository{db}
}

func (dr *debugRepository) GetAllDebugs(debugs *[]model.Debug, userId uint) error {
	// 一番新しくできたタスクを返す。
	if err := dr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(debugs).Error; err != nil {
		return err
	}
	return nil
}

func (dr *debugRepository) GetDebugById(debug *model.Debug, userId uint, debugId uint) error {
	if err := dr.db.Joins("User").Where("user_id=?", userId).First(debug, debugId).Error; err != nil {
		return err
	}
	return nil
}

func (dr *debugRepository) CreateDebug(debug *model.Debug) error {
	if err := dr.db.Create(debug).Error; err != nil {
		return err
	}
	return nil
}

func (dr *debugRepository) UpdateDebug(debug *model.Debug, userId uint, DebugId uint) error {
	// NOTE:clausesにより更新後のデバッグポインタを変更前のデバッグのポイントに変えておいてくれる
	result := dr.db.Model(debug).Clauses(clause.Returning{}).Where("id=? AND user_id=?", DebugId, userId).Save(&debug)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (dr *debugRepository) DeleteDebug(userId uint, debugId uint) error {
	result := dr.db.Where("id=? AND user_id=?", debugId, userId).Delete(&model.Debug{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
