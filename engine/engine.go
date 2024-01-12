package engine

import (
	"time"

	"github.com/GuardianOfWynn/eco-simulator/territory"
)

type Engine struct {
	Map    *territory.GuildMap
	ticker *time.Ticker
}

func (e *Engine) Start() {
	e.ticker = time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-e.ticker.C:
				for _, terr := range e.Map.Claims {
					terr.Tick()
				}
			}
		}
	}()
}

func (e *Engine) Stop() {
	e.ticker.Stop()
}
