package domain

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
	getProvinces(name *string) (p []Province, err error)
	addProvince(p Province) (err error)
	updateProvince(id int64, p Province) (err error)
	deleteProvince(id int64) (err error)

	getCities(name *string) (c []City, err error)
	addCity(c City) (err error)
	updateCity(id int, c City) (err error)
	deleteCity(id int) (err error)
}

type basicInfoPostgres interface {
}
