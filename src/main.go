package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	cpu Cpu
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	printDebug(g, screen)

}

func printDebug(g *Game, screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("AF: %.4x %.16b", g.cpu.AF, g.cpu.AF), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("BC: %.4x %.16b", g.cpu.BC, g.cpu.BC), 0, 15)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("DE: %.4x %.16b", g.cpu.DE, g.cpu.DE), 0, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("HL: %.4x %.16b", g.cpu.HL, g.cpu.HL), 0, 45)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("PC: %.4x %.16b", g.cpu.PC, g.cpu.PC), 0, 70)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("SP: %.4x %.16b", g.cpu.SP, g.cpu.SP), 0, 85)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("gbemu")

	cpu := Cpu{
		AF: 0x1c1f,
		BC: 0x0001,
		DE: 0xcfcf,
		HL: 0x1234,
		PC: 0xff01,
		SP: 0xcc12,
	}

	if err := ebiten.RunGame(&Game{cpu: cpu}); err != nil {
		log.Fatal(err)
	}
}
