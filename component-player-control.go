//In this file we will define new components to give the caller their ability to control the player element - such as moving and shooting

package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

//17 We must define a keyboard mover component type which which will be attached to the player element so that it is able to be controlled with a keyboard
type keyboardMover struct { //this component could be added to a bullet or a player  so that movement control can be applied as soon as or during the time when the key is hit.
	container *element        //holds pointer reference to the element that this keyboardMover component is attached to// the component needs to be able to access the elements properties - position  so that position of the element can be changed and updated
	speed     float64         //we need to define a speed that that container element can move according to relative element position
	sr        *spriteRenderer //we need to hold a reference to a sprite renderer component so that the texture properties needed(width and height of sprite) can accessed and used to dynamically populate the required element size which is formulated in moving the element

}

//18 define the onUpdate() method necessary for this component to move the element every frame
func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 { //is left key down
		if mover.container.position.x-(mover.sr.width/2.0) > 0 { //is player within the screean? -to the left
			mover.container.position.x -= mover.speed * delta //	//move player left - if within screen
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 { //is right key being presesd
		if mover.container.position.x+(mover.sr.width/2.0) < screenwidth { ////is player within screen to thr right
			mover.container.position.x += mover.speed * delta //move player right - if within screen
		}
	}
	return nil
}

//19 Define a new KeyboardMover Function which is used to initialize and create a new keynoard mover component which will be attached to the player element so that it can be moved by the keyboard
func newKeyboardMover(container *element, speed float64) *keyboardMover { //a container argument is needed to identify/reference the element that will respond to keyboard control initiated by the caller.
	return &keyboardMover{ //this function will simply return the component in the form of a struct literal and the arguments passed throught the function will be used to construct and populate  the contents/properties held by the component -
		container: container,                                                   //the container could be a parent element such as a player or a bullet
		speed:     speed,                                                       //the speed will determine how fast the element it is attached to can move
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer), //the .(*spriteRenderer) TYPE CASTING indicates the the caller knows what type this component is which is what we are looking for//here we are using the getComponent() method inside element.go from the container to fetch the spriteRenderer for the given element
	}
}

//20 Define the onDraw() method necessary for the element to implement keyboardMover conponent interface - because the method is a required definition of the component interface
func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error { //note the same method definition in the interface defined in element.go
	return nil
}

//22 Define keyboard shooter component type which when attached to a player element, will allow the player to shoot bullet upon space key press
type keyboardShooter struct {
	container *element      //like all other component this container will reference the element that this component is contained by
	cooldown  time.Duration //cooldown time to prevent spamming bullets//the cooldown is a configurable duration of time to allow a certain number of bullets to be fired per second when space key is hit
	lastShot  time.Time     //the last shot is a captured time stamp which stores the current time, the cooldown duration time is evaluated against this time value to determine if more bullets can be fired at a certain point in time.
}

//23 Define key board shooter function to initialize and create a new keyboard shooter component so that it can be attached to a player element
func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
		lastShot:  time.Now(),
	}
}

//24 Define the necessary onUpdate() method for the keyboardShooter component - which when attached to the the player element enables the player to shoot bullets upon key press
func (mover *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= playerShotCooldown { //allow shooting if enough time has passed
			mover.shoot(mover.container.position.x, mover.container.position.y-20)
			mover.shoot(mover.container.position.x-32, mover.container.position.y-20)
			mover.lastShot = time.Now() //update last shot to current time
		}
	}
	return nil
}

//25 Define the shoot() method for the keyBoardShooter component
func (mover *keyboardShooter) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true  //,ok idiom - shoot bullet
		bul.position.x = x //set bullet position
		bul.position.y = y
		bul.rotation = 270 * (math.Pi / 180) //convert degrees to radians and feed radian anglr to the bullet update.
	}
}

//26 Define the necessary onDraw() method for the keyboardShooter component - which will allow the player element to implement the keyboard shooter component
func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

//
func (mover *keyboardShooter) onCollision(other *element) error {
	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}
