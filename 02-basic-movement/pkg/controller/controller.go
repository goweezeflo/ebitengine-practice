package controller

import "github.com/hajimehoshi/ebiten/v2"

type Controller struct {
	Id ebiten.GamepadID
}

func CreateController(id ebiten.GamepadID) Controller {
	return Controller{
		Id: id,
	}
}
