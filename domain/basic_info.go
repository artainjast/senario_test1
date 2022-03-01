package domain

import "context"

type City struct {
	ID         int64    `json:"id"`
	NAME       string   `json:"name"`
	PROVINCE   Province `json:"province"`
	CREATED_AT int64    `json:"created_at"`
	UPDATED_AT int64    `json:"updated_at"`
	DELETED_AT int64    `json:"deleted_at"`
}

type Province struct {
	ID         int64  `json:"id"`
	NAME       string `json:"name"`
	CREATED_AT int64  `json:"created_at"`
	UPDATED_AT int64  `json:"updated_at"`
	DELETED_AT int64  `json:"deleted_at"`
}

type BasicInfoUseCase interface {
	getProvinces(ctx context.Context, name *string) (p []Province, err error)
	addProvince(ctx context.Context, p Province) (err error)
	updateProvince(ctx context.Context, id int64, p Province) (err error)
	deleteProvince(ctx context.Context, id int64) (err error)

	getCities(ctx context.Context, id int64, name *string) (c []City, err error)
	addCity(ctx context.Context, c City) (err error)
	updateCity(ctx context.Context, id int, c City) (err error)
	deleteCity(ctx context.Context, id int) (err error)
}

type basicInfoPostgres interface {
	getProvinces(ctx context.Context, name *string) (p []Province, err error)
	addProvince(ctx context.Context, p Province) (err error)
	updateProvince(ctx context.Context, id int64, p Province) (err error)
	deleteProvince(ctx context.Context, id int64) (err error)

	getCities(ctx context.Context, id int64, name *string) (c []City, err error)
	addCity(ctx context.Context, c City) (err error)
	updateCity(ctx context.Context, id int, c City) (err error)
	deleteCity(ctx context.Context, id int) (err error)
}
