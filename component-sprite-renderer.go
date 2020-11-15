//8 We want to start first with drawing as this will be a reusable component for every element that needs to be displayed in the screen
package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

//in component architecture we want to be creating different related types.
//9 Define sprite renderer component type
type spriteRenderer struct { //a this a component type for rendering and drawing component - we can create the components and then attach it to an element such as a player.
	tex           *sdl.Texture //sr component needs an sdl texture// this is what will be used to render the sprite imgs
	container     *element     //we need to be able to store a reference of the element that the sr component is attached to so that we can acces the elements postion data (x,y)//container will provide a means to access the parent element.//the container will reference the element that contains this component.
	width, height float64      //the width and height fields will store the dimension data for the sprite texture/these dimensions are necessary for acquiring positioning and movement speed calculations

}

//10 Define Sprite render function which will be used to create a sprite render (sr) component.
func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer { //we want to give it the file name as an argument - which will locate the sprite file which needs rendering and it will then return a *spriteRenderer which will populate the tex (texture) in the sprite renderer component - the *sr will be passed all over the place and we wnt to make we are always refering to the same one with its pointer rather than its value.
	tex := textureFromBMP(renderer, filename) //we need to initialize the texture so that the next set of operations can use the tex to return its dimension data (width and height)
	_, _, width, height, err := tex.Query()   //before creating the componetn we want to perform some necessary function calls that will return texture data that can be stored and used. It it necessary that this data is stored so that it can be accessed by other components attached to the element //Query() returns quite a few thing about the texture including it's width and height.
	if err != nil {
		panic(fmt.Errorf("sprite_renderer.go - newSpriteRenderer() - could not Query() *sr - %v", err))
	}
	return &spriteRenderer{ //here we just want to return the sprite renderer using struct literals
		tex:       textureFromBMP(renderer, filename), //tx is populated by calling the textureFromBMP() function which returns tex texture
		container: container,                          //the container is the element that this component is attached to - the component needs to be able to access and change the element that it is part of.
		width:     float64(width),
		height:    float64(height),
	}
}

//11 Define helper function for loadiing texture from bmp file. - (moved from main.go)
func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture { //takes in sdl renderer with filename as argument and returns
	img, err := sdl.LoadBMP(filename) //init img from sdl.LoadBMP() method - filename is the location where the file is stored on disk.
	if err != nil {                   //panic if given filename cannot be found - notify caller that there is a descrepency in the file naming or the file is missing
		panic(fmt.Errorf("loading %v sprite: %v", filename, err))
	}
	defer img.Free()                                   //house keeping//- before exiting function- release img stored in RAM because it is not needed any more//Free image-surface for player sprite
	tex, err := renderer.CreateTextureFromSurface(img) //Declare player texture*//sdl.Texture-from img now stored in RAM
	if err != nil {                                    //panic that the file is there but there is something wrong with the file maybe the wrong format?
		panic(fmt.Errorf("creating player texture from %v: %v", filename, err))
	}
	return tex //return *tex texture - which is used to set the tex texture field iin the spriteRenderer component when it is created
}

//12 define onDraw method functionality for the SpriteRenderer component which will be used to draw elements on to the screen -  all element that need to be displayed will have an spriteRenderer component
func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error { //the onDraw() method for the SpriteRenderer component takes in a sprite-renderer argument and returns an error - if anything goes wrong. //the sr argument will contain a populated tex texture which have methods that can query the texture
	x := (sr.container.position.x - float64(sr.width/2.0)) //we need to put value in dynamically so all elements can be given a particular position and size//convert coordinates to top left of sprite
	y := (sr.container.position.y - float64(sr.height/2.0))
	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: int32(sr.width),
			H: int32(sr.height),
		},
		&sdl.Rect{
			X: int32(x),
			Y: int32(y),
			W: int32(sr.width),
			H: int32(sr.height),
		},
		sr.container.rotation,
		&sdl.Point{
			X: int32(sr.width / 2),
			Y: int32(sr.height / 2),
		},
		sdl.FLIP_NONE)

	return nil
}

//13 Define the required onUpdate() method that the sr componenent needs to be able to implement
func (sr *spriteRenderer) onUpdate() error { //as long as we have on onUpdate method and onDraw() method the componetn interface will be satisfied - interface will says - ' yo baby if you got these methods(updateanddraw) then you are my type!'
	return nil //this function never needs to be used  - but future changes could end up accomodating this function for updating a sprite such as changing the sprite elemts image that get displayed to another image tht can be displayed.
} //an applicable method for the sr  component to function and accomodate the component interface

//
func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
