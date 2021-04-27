package databus

import (
	"go6502/video"
	"sync"
)

type Bus struct {
	Access sync.Mutex
	level  bool // True: CPU / False: VIC
	vic    *video.Video
}

func (B *Bus) Init(vic *video.Video) {
	B.level = false
	B.vic = vic
}

func (B *Bus) Release() {
// KEEPBUS:
	// B.vic.Run()
	// if !B.vic.BA {
	// 	goto KEEPBUS
	// }
}
