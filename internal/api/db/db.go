package db

import (
	"log"
	"sync"
)

type Clients struct {
	Users map[string]string
	mu    sync.Mutex
}

func NewClients() *Clients {
	bd := make(map[string]string)
	return &Clients{Users: bd}
}

func (cs *Clients) addUser(login string, pass string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	cs.Users[login] = pass
}

func (cs *Clients) ActiveSessions(login string, pass string) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if _, ok := cs.Users[login]; !ok {
		cs.addUser(login, pass)
	} else {
		return false
	}
	return true
}
func (cs *Clients) CloseAuth(l string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	delete(cs.Users, l)
}
func (cs *Clients) ViewActiveSessions() {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	i := 1
	for k := range cs.Users {
		log.Printf("user#%v %v\n", i, k)
		i++
	}
}
