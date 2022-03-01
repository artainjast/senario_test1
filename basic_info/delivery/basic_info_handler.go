package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"senario_test1/domain"
)

type BasicInfo struct {
	basicInfoUseCase domain.BasicInfoUseCase
}

func Basicinfohandler(apiRouter *mux.Router, basicInfo domain.BasicInfoUseCase) {
	handler := &BasicInfo{
		basicInfoUseCase: basicInfo,
	}

	apiRouter.HandleFunc("/province", handler.getProvinces).Methods(http.MethodGet)
	apiRouter.HandleFunc("/province", handler.addProvince).Methods(http.MethodPost)
	apiRouter.HandleFunc("/province/:id", handler.updateProvince).Methods(http.MethodPut)
	apiRouter.HandleFunc("/province/:id", handler.deleteProvince).Methods(http.MethodDelete)
}

func (b *BasicInfo) getProvinces(w http.ResponseWriter, r *http.Request) {
	var resByte []byte
	var name *string
	err, a := b.basicInfoUseCase.GetProvinces(r.Context(), name)
	if err != nil {
		panic(err)
	}
	resByte, _ = json.Marshal(a)
	w.WriteHeader(200)
	_, _ = w.Write(resByte)
	return
}

func (b *BasicInfo) addProvince(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) updateProvince(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) deleteProvince(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) getCities(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) addCity(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) updateCity(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b *BasicInfo) deleteCity(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
