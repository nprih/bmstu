package service

import "lesson22/internal/db"

func CheckClientById(clientId int) bool {
	clients := db.SelectAllClients()
	for _, v := range clients {
		if clientId == v.Id {
			return true
		}
	}
	return false
}
