package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math"
	"time"
)

var (
	currentEpochTimeSeconds       int64
	previousEpochTimeSeconds      int64
	currentEpochTimeMilliseconds  int64
	previousEpochTimeMilliseconds int64
)

type Game struct {
	tickInSeconds      uint64
	tickInMilliseconds uint64
	sinTick            float64
	cosTick            float64
	tanTick            float64
}

func init() {
	currentEpochTimeSeconds = time.Now().Unix()
	previousEpochTimeSeconds = currentEpochTimeSeconds
	currentEpochTimeMilliseconds = time.Now().UnixMilli()
	previousEpochTimeMilliseconds = currentEpochTimeMilliseconds
}

func (g *Game) Update() error {
	currentEpochTimeSeconds = time.Now().Unix()
	if currentEpochTimeSeconds != previousEpochTimeSeconds {
		previousEpochTimeSeconds = currentEpochTimeSeconds
		g.tickInSeconds++
	}
	currentEpochTimeMilliseconds = time.Now().UnixMilli()
	if currentEpochTimeMilliseconds != previousEpochTimeMilliseconds {
		previousEpochTimeMilliseconds = currentEpochTimeMilliseconds
		g.tickInMilliseconds++
	}
	g.sinTick = math.Sin(float64(g.tickInMilliseconds) / 100)
	g.cosTick = math.Cos(float64(g.tickInMilliseconds) / 100)
	g.tanTick = math.Tan(float64(g.tickInMilliseconds) / 100)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Epoch: %d", time.Now().Unix()), 10, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick (seconds): %d", g.tickInSeconds), 10, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tick (milliseconds): %d", g.tickInMilliseconds), 10, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sin Tick: %.2f", g.sinTick), int(110+math.Abs(g.tanTick*50)), int(70+math.Abs(g.sinTick*50)))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Cos Tick: %.2f", g.cosTick), int(70+g.sinTick*50), int(90+math.Abs(g.cosTick*50)))
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tan Tick: %.2f", g.tanTick), int(150+g.cosTick*50), int(110+math.Abs(g.tanTick*50)))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ebitengine - 09 - Time")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
