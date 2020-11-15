package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed        = 5.00
	playerSize         = 64
	playerShotCooldown = time.Millisecond * 100
)

/*
//Declare player struct
type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}
*/
//6. create new player element and attach necessary components to it
func newPlayer(renderer *sdl.Renderer) *element { //takes in an sr argument and returns a player element created. element will always be referenced with pointers because we always want to be refering to the same component anywhere - having a refernce type rather than a value type will reinforce this through type assertion and acceptance
	player := &element{}      //initialize the player//declare an empty player element//want to se the elememts common variables that every gaming elemrnt requires i
	player.position = vector{ //set player position - center botttom of screen.
		x: screenwidth / 2.0,             //center screen
		y: screenheight - playerSize/2.0, //bottom edge
	}
	player.active = true //player is immediately visible and usable on screen

	//7. write individual drawing moving and shooting components and attach them to the player - these must be components that are defined for drawing moving shooting the player
	//14 initialize the spriteRenderer component for attaching to the player element so that it can be drawn.
	sr := newSpriteRenderer(player, renderer, "sprite/player.bmp") //newSpriteRenderer() functionality is located in the sprite_renderer.go file//player is the element argument, renderer is the render argument, "player.bmp" is bmp filename argument
	//15 add the sr component to the player
	player.addComponent(sr) //the addComponent() method functionalty is in the element.go file//now that we got our sr sprite renderer the sprite is ready to be drawn which will be initiated somehow in the main() function in the main.go file
	//16 write the update() method component needed for the player element to move left anf right upon keyboard press and then attach it to the player. - see keyboardMover onUpdate() method inside player_control.go
	//21 initialize the keyboard mover component
	mover := newKeyboardMover(player, playerSpeed) //the newKeyBoardMover function takes the player as element argument along with the playerSpeed constant to initialize the new mover component for the player
	//22 attach the new keyboardMover component to the player element using the addComponent() method
	player.addComponent(mover)
	//23 write the update() component method for the component that enables the player element to shoot bullets - see onUpdate() method for keyboardShooter
	//27 initialize the keyboard shooter component
	shooter := newKeyboardShooter(player, playerShotCooldown) //
	//28 attach the new keyboard shooter component to the player element
	player.addComponent(shooter)

	return player
}

/*
//create new player
func newPlayer_old(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromBMP(renderer, "sprite/player.bmp")
	p.x = screenwidth / 2.0 //Set Player position
	p.y = screenheight - playerSize/2.0
	return p
}

func (p *player) draw(renderer *sdl.Renderer) { //Draw player onto window
	x := (p.x - (playerSize / 2.0)) //convert coordinates to top left of sprite
	y := (p.y - (playerSize / 2.0))
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},               //source coordinates of sprite/texture
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize}) //destination cooridnates of sprite/texture
}

func (p *player) update() { //player keyboard input
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 { //is left key down
		if p.x-(playerSize/2.0) > 0 { //is player within the screean? -to the left
			p.x -= playerSpeed //	//move player left - if within screen
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 { //is right key being presesd
		if p.x+(playerSize/2.0) < screenwidth { //is player within screen to thr right
			p.x += playerSpeed //move player right - if withing screen
		}
	}
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= playerShotCooldown { //allow shooting if enough time has passed
			p.shoot(p.x, p.y-20)
			p.shoot(p.x-32, p.y-20)
			//p.shoot(p.x, p.y)
			p.lastShot = time.Now() //update last shot to current time
		}
	}
}

func (p *player) shoot(x, y float64) {
	if bul, ok := bulletFromPool(); ok {
		bul.active = true //,ok idiom - shoot bullet
		bul.x = x         //set bullet position
		bul.y = y
		//p.lastShot = time.Now()    //singleshot
		bul.angle = 270 * (math.Pi / 180) //convert degrees to radians and feed radian anglr to the bullet update.
	}
}*/

/*player.go -//three basic things about the player they can move, shoot, be drawn -  we wnat to now create seperate components for these player functionalities
FIRST
we want a class to hold all these player components.
We need a GENERIC type that can hold all of these gaming components
so everything in the game is derived from the same type
everything in the game is an instance of the same type
and what differentiates all of these types are the different components that these so called game objects have.
We are going to call this class the Element.

THis Generic Type called Element is defined in its own file








*/
