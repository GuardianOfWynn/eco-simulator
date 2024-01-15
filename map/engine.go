package territory

import (
	"fmt"
	"time"
)

var EngineInstance *Engine = nil

type Engine struct {
	Map    *GuildMap
	ticker *time.Ticker
}

func (e *Engine) Start() {
	for range time.Tick(time.Second * 1) {
		fmt.Println("ticou")
		for _, terr := range e.Map.Territories {
			terr.Tick()
		}
	}

}
