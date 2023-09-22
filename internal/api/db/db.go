package db

import "sync"

type Users struct {
	Map map[string]string
	mu  sync.RWMutex
}

func NewDataBase() *Users {
	u := new(Users)
	u.Map = make(map[string]string)
	return u
}

func (u *Users) AddUser(user, password string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if _, ok := u.Map[user]; !ok {
		u.Map[user] = password
	} else {
		return false
	}
	return true
}

func (u *Users) CloseUserConnection(user string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	delete(u.Map, user)
}
