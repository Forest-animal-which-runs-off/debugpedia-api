package usecase

import (
	"debugpedia-api/model"
	"debugpedia-api/repository"
)

type IDebugUsecase interface {
	GetAllDebugs(userId uint) ([]model.DebugResponse, error)
	GetDebugById(userId uint, debugId uint) (model.DebugResponse, error)
	CreateDebug(debug model.Debug) (model.DebugResponse, error)
	UpdateDebug(debug model.Debug, userId uint, debugId uint) (model.DebugResponse, error)
	DeleteDebug(userId uint, debugId uint) error
}

type debugUsecase struct {
	dr repository.IDebugRepository
}

func NewDebugUsecase(dr repository.IDebugRepository) IDebugUsecase {
	return &debugUsecase{dr}
}

func (du *debugUsecase) GetAllDebugs(userId uint) ([]model.DebugResponse, error) {
	debugs := []model.Debug{}
	if err := du.dr.GetAllDebugs(&debugs, userId); err != nil {
		return nil, err
	}
	resDebugs := []model.DebugResponse{}
	for _, v := range debugs {
		d := model.DebugResponse{
			ID:        v.ID,
			Title:     v.Title,
			Body:      v.Body,
			Links:     v.Links,
			Techs:     v.Techs,
			Cause:     v.Cause,
			Resolve:   v.Resolve,
			User:      v.User,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resDebugs = append(resDebugs, d)
	}
	return resDebugs, nil
}
func (du *debugUsecase) GetDebugById(userId uint, debugId uint) (model.DebugResponse, error) {
	debug := model.Debug{}
	if err := du.dr.GetDebugById(&debug, userId, debugId); err != nil {
		return model.DebugResponse{}, err
	}
	resDebug := model.DebugResponse{
		ID:        debug.ID,
		Title:     debug.Title,
		Body:      debug.Body,
		Links:     debug.Links,
		Techs:     debug.Techs,
		Cause:     debug.Cause,
		Resolve:   debug.Resolve,
		User:      debug.User,
		CreatedAt: debug.CreatedAt,
		UpdatedAt: debug.UpdatedAt,
	}
	return resDebug, nil
}

func (du *debugUsecase) CreateDebug(debug model.Debug) (model.DebugResponse, error) {
	if err := du.dr.CreateDebug(&debug); err != nil {
		return model.DebugResponse{}, err
	}
	resDebug := model.DebugResponse{
		ID:        debug.ID,
		Title:     debug.Title,
		Body:      debug.Body,
		Links:     debug.Links,
		Techs:     debug.Techs,
		Cause:     debug.Cause,
		Resolve:   debug.Resolve,
		User:      debug.User,
		CreatedAt: debug.CreatedAt,
		UpdatedAt: debug.UpdatedAt,
	}
	return resDebug, nil
}

func (du *debugUsecase) UpdateDebug(debug model.Debug, userId uint, debugId uint) (model.DebugResponse, error) {
	if err := du.dr.UpdateDebug(&debug, userId, debugId); err != nil {
		return model.DebugResponse{}, err
	}
	resDebug := model.DebugResponse{
		ID:        debug.ID,
		Title:     debug.Title,
		Body:      debug.Body,
		Links:     debug.Links,
		Techs:     debug.Techs,
		Cause:     debug.Cause,
		Resolve:   debug.Resolve,
		User:      debug.User,
		CreatedAt: debug.CreatedAt,
		UpdatedAt: debug.UpdatedAt,
	}
	return resDebug, nil
}

func (du *debugUsecase) DeleteDebug(userId uint, debugId uint) error {
	if err := du.dr.DeleteDebug(userId, debugId); err != nil {
		return err
	}
	return nil
}
