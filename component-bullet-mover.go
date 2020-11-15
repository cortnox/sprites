//43 implement bullet mover component

package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

//44 Define bulet mover function for creating new bullet mover components
func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}

}

//45 Define onDraw() method for the bullet mover comonent - there no current usage for this bullet mover method however it is needed in order the the bullet mover to implement the component interface
func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

//46 Define onUpdate() method for the bullet mover component - RE-IMPLEMENT THE MOVEMENT OF THE BULLET EVERY FRAME
func (mover *bulletMover) onUpdate() error {
	mover.container.position.x += (bulletSpeed * math.Cos(mover.container.rotation) * (math.Pi * math.Pow(1.618, 1))) //update bullet x and y positions
	mover.container.position.y += (bulletSpeed * math.Sin(mover.container.rotation) * (math.Pi * math.Pow(0.619, 2)))
	if mover.container.position.x > screenwidth || mover.container.position.x < 0 || mover.container.position.y > screenheight || mover.container.position.y < 0 { //check if bullet is offscreen
		mover.container.active = false //deactivate bullet an set as reusable
	}

	//
	mover.container.collisions[0].center = mover.container.position
	return nil
}

//make bullets disapear upon impact / collision with the enemy
func (mover *bulletMover) onCollision(other *element) error {
	mover.container.active = false
	return nil
}
