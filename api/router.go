package main

import "github.com/gorilla/mux"

// SetWeaponSubRouter routes weapon stuff
func SetWeaponSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateWeaponEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetWeaponEndpoint).Methods("GET")
	subRouter.HandleFunc("type/{weapontype}", GetWeaponTypeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneWeaponEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteWeaponEndpoint).Methods("DELETE")
}

// SetArmorSubRouter routes weapon stuff
func SetArmorSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateArmorEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetArmorEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneArmorEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteArmorEndpoint).Methods("DELETE")
}

// SetFairySubRouter routes weapon stuff
func SetFairySubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateFairyEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetFairyEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneFairyEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteFairyEndpoint).Methods("DELETE")
}

// SetFashionaccessoireSubRouter routes weapon stuff
func SetFashionaccessoireSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateFashionaccessoireEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetFashionaccessoireEndpoint).Methods("GET")
	subRouter.HandleFunc("type/{fashionaccessoiretype}", GetFashionaccessoireTypeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneFashionaccessoireEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteFashionaccessoireEndpoint).Methods("DELETE")
}

// SetResistanceSubRouter routes weapon stuff
func SetResistanceSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateResistanceEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetResistanceEndpoint).Methods("GET")
	subRouter.HandleFunc("type/{fashionaccessoiretype}", GetResistanceTypeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneResistanceEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteResistanceEndpoint).Methods("DELETE")
}

// SetAccessoireSubRouter routes weapon stuff
func SetAccessoireSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateAccessoireEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetAccessoireEndpoint).Methods("GET")
	subRouter.HandleFunc("type/{fashionaccessoiretype}", GetAccessoireTypeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneAccessoireEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteAccessoireEndpoint).Methods("DELETE")
}

// SetCostumeSubRouter routes weapon stuff
func SetCostumeSubRouter(subRouter *mux.Router) {
	subRouter.HandleFunc("", CreateCostumeEndpoint).Methods("POST")
	subRouter.HandleFunc("", GetCostumeEndpoint).Methods("GET")
	subRouter.HandleFunc("type/{fashionaccessoiretype}", GetCostumeTypeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", GetOneCostumeEndpoint).Methods("GET")
	subRouter.HandleFunc("{id}", DeleteCostumeEndpoint).Methods("DELETE")
}
