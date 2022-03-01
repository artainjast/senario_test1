package domain

type City struct {
	ID         int64  `json:"id"`
	NAME       string `json:"name"`
	PROVINCEID int64  `json:"provinceid"`
	CREATED_AT int64  `json:"created_at"`
	UPDATED_AT int64  `json:"updated_at"`
	DELETED_AT int64  `json:"deleted_at"`
}

type Province struct {
	ID         int64  `json:"id"`
	NAME       string `json:"name"`
	CREATED_AT int64  `json:"created_at"`
	UPDATED_AT int64  `json:"updated_at"`
	DELETED_AT int64  `json:"deleted_at"`
}

type BasicDataUseCase interface {
	addProvince(p Province) (err error)
	getProvinces(name *string) (p []Province, err error)
	updateProvince(id int64, p Province) (err error)
	deleteProvince(id int64) (err error)
}
