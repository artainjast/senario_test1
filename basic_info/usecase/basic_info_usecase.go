package usecase

import (
	"context"
	"senario_test1/domain"
)

type basicInfoUseCase struct {
	BasicInfoRepository domain.BasicInfoPostgres
}

func (b basicInfoUseCase) GetProvinces(ctx context.Context, name *string) (p []domain.Province, err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) addProvince(ctx context.Context, p domain.Province) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) updateProvince(ctx context.Context, id int64, p domain.Province) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) deleteProvince(ctx context.Context, id int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) getCities(ctx context.Context, id int64, name *string) (c []domain.City, err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) addCity(ctx context.Context, c domain.City) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) updateCity(ctx context.Context, id int, c domain.City) (err error) {
	//TODO implement me
	panic("implement me")
}

func (b basicInfoUseCase) deleteCity(ctx context.Context, id int) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewbasicInfoUseCase(repository domain.BasicInfoPostgres) domain.BasicInfoUseCase {
	return &basicInfoUseCase{
		BasicInfoRepository: repository,
	}
}
