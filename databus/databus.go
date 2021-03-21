package databus

import "sync"

type Databus struct {
	access sync.Mutex
}

func (D *Databus) WaitAccess() {
	D.access.Lock()
	D.access.Unlock()
}

func (D *Databus) GetAccess() {
	D.access.Lock()
}

func (D *Databus) AllowCPU() {
	D.access.Unlock()
	D.access.Lock()
}
