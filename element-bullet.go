//42 - break the bullet down into a component architecture - this is the remainder of the oop functionality
/*
The bullet will be made as an element just like the basic enemy and the player elements
We need to build a component to
draw the bullet
move the bullet when the player shoots based on speed and angle with which the bullet is given
..eventually we will work on physics and object collision with the bulltets but we want to get the early architecture prequisite first so that later more components can be built and applied.
and then we need to attach each component to the bullet most of the functional code below will be reused -we are mainly defusing the bullet struct type usage so that its fuctionality accomodates an element instead.
And in doing that because we need a bullet element we need to make sure that the bullet contains methods to implement the component interface such onDraw() an onUpdate()
*/

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 4.2
)

/*
type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	active bool
	angle  float64 //inform us how to move bullett
}
*/
//47 Define new bullet element function which is used to invoke and create a bullet element and add the necessary components to the bullet for drawing and moving
func NewBullet(renderer *sdl.Renderer) *element {
	b := &element{}

	//48 initialize the sr component for drawing the bullet
	sr := newSpriteRenderer(b, renderer, "sprite/bullet.bmp")
	sr.width /= 2
	sr.height /= 2
	//49 add the draw component to the bullet
	b.addComponent(sr)
	//50 - initialize the bullet mover component functionality for the bullet
	mover := newBulletMover(b, bulletSpeed)
	//51 - add the bullet mover component to the bullet
	b.addComponent(mover)

	//b.active = false
	//55 - estable the bullet collsion point
	col := circle{
		center: b.position,
		//the radius will be set to the size of the bullet
		radius: 8,
	}
	//56 - append the collision to the bulets []collision //we want bullets to kill enemies//so enemies are going to need a collision circle as well- so that we can detect when the bullet hits the enemy
	b.collisions = append(b.collisions, col)
	b.tag = "player-bullet"

	return b
}

/*
func NewBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprite/bullet.bmp")
	return bul
}


func (bul *bullet) draw(renderer *sdl.Renderer) {
	if !bul.active {
		return
	}
	x := (bul.x - (bulletSize / 2.0)) //convert coordinates to top left of sprite
	y := (bul.y - (bulletSize / 2.0))
	renderer.Copy(bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},               //source coordinates of sprite/texture
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize}) //destination cooridnates of sprite/texture
}

//update method for thew bullet to visibly mive
func (b *bullet) update() { //this get called every frame
	b.x += (bulletSpeed * math.Cos(b.angle) * (math.Pi * math.Pow(1.618, 2))) //update bullet x and y positions
	b.y += (bulletSpeed * math.Sin(b.angle) * (math.Pi * math.Pow(0.619, 2)))
	if b.x > screenwidth || b.x < 0 || b.y > screenheight || b.y < 0 { //check if bullet is offscreen
		b.active = false //deactivate bullet an set as reusable
	}
}
*/
var bulletPool []*element //GLOABAL VARIABLE: declare slice to store prepared bullets

func initBulletPool(renderer *sdl.Renderer) { //prepare bullet and load into pool
	for i := 0; i < 30; i++ {
		bul := NewBullet(renderer)
		bulletPool = append(bulletPool, bul)
		elements = append(elements, bul)
	}
}

func bulletFromPool() (*element, bool) { //convenience function that returns an unused bullet from pool
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true //success - next blluet loaded
		}

	}
	return nil, false //no bullets available - bullet pool is exhausted
}
