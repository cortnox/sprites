package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	basicEnemySpeed = 0.4
	basicEnemySize  = 105
)

////COMPONENT BASED FUNCTIONALTIY
//34 Define new basic enemy fine and reuse component for drawinfg and diplaying the basic enemy
//func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
func newBasicEnemy(renderer *sdl.Renderer, x, y float64) *element {
	basicEnemy := &element{}
	//basicEnemy.position = position
	basicEnemy.position = vector{x: x, y: y}

	basicEnemy.rotation = 180
	//35 intiialze spriter renderer component for basic enemy element
	sr := newSpriteRenderer(basicEnemy, renderer, "sprite/eplayer.bmp")
	//36 add sr component to the basic enemy
	basicEnemy.addComponent(sr)

	//57 define the collision circle for the basic enemy
	col := circle{
		center: basicEnemy.position,
		radius: 38,
	}
	//58 append the collision circle to the basic enemy []collision
	basicEnemy.collisions = append(basicEnemy.collisions, col)
	basicEnemy.active = true
	//init btv
	vtb := newBulletVulnerability(basicEnemy)
	basicEnemy.addComponent(vtb)
	return basicEnemy
}

/////////////////////////////////////////////////
/*
type basic_enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x, y float64) (be basic_enemy) { //Create image-surface
	be.tex = textureFromBMP(renderer, "sprite/basic_enemy.bmp")
	be.x = x
	be.y = y

	return be

}

//Add a draw method
func (be *basic_enemy) draw(renderer *sdl.Renderer) {
	//convert coordinates to top left of sprite
	x := (be.x - (basicEnemySize / 2.0))
	y := (be.y - (basicEnemySize / 2.0))
	renderer.CopyEx(be.tex, //funtion that also take agle float to flip enemy sprite
		&sdl.Rect{X: 0, Y: 0, W: basicEnemySize, H: basicEnemySize},               //source coordinates of sprite/texture
		&sdl.Rect{X: int32(x), Y: int32(y), W: basicEnemySize, H: basicEnemySize}, //destination coordinates /positions and size
		180, //angle in degrees to flip spreite
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE) //destination cooridnates of sprite/texture
}
*/
