package main //main.go

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenwidth          = 800
	screenheight         = 800
	targetTicksPerSecond = 60.00 //nr of physics engine ticks
)

//
//init delta value Global
var delta float64 //we need to delta value to be calculated when a frame update is timed so that the delta value can accerelated hard running this code else where.

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil { //initialize sdl
		fmt.Println("Initializing SDL:", err)
		return //discontinue out
	}

	window, err := sdl.CreateWindow( //create window
		"Game Space",            //window title
		sdl.WINDOWPOS_UNDEFINED, //window position
		sdl.WINDOWPOS_UNDEFINED, //window position
		screenwidth,             //window width
		screenheight,            //window height
		sdl.WINDOW_OPENGL)       //open gl accelerated window -flag
	if err != nil {
		fmt.Println("Initializing Window:", err)
		return
	}
	defer window.Destroy()                                                    //defer means it will be left to execute last before other functions have executed
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED) //givewindowtoputrendererin-tellrendereri wanttobeaccelerated
	if err != nil {
		fmt.Println("Initializing Renderer:", err)
		return //return to main
	}
	defer renderer.Destroy() //destroy renderer when we are done with it0
	//29 - invoke new Player element with the newPLayer() function to create a player
	plyr := newPlayer(renderer) //creat new player struct - instance
	//40 append the new player to the Global []element
	elements = append(elements, plyr)
	var basicEnemies []*element

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			x := ((float64(i))/4)*screenwidth + basicEnemySize
			y := float64(j)*basicEnemySize + basicEnemySize
			//36 invoke basic enemy creation
			//basicEnemy := newBasicEnemy(renderer, vector{x, y})
			basicEnemy := newBasicEnemy(renderer, x, y)

			basicEnemies = append(basicEnemies, basicEnemy)
			//41 append basic enemy elements the the Global []element
			elements = append(elements, basicEnemy)
			//time.Sleep(time.Millisecond * 140)
		}
	}

	initBulletPool(renderer) //init bullet pool - to store bullets

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() { //ALT+F4 - closing capability
			switch event.(type) {
			case *sdl.QuitEvent:
				return //re
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255) //set the back ground color
		renderer.Clear()                          //fill window with draw color

		//37 Globalisation
		/*

		   in order to nullfy alot of vebose and repeated lines of updating the same type of element it just makes reasonable sense
		   to optimize this somehow by creating a global variable that stores all element inside a slice - of []elements
		   and by just basically adding every game element and component to this global []element
		   we then just need to iterate over all of these global elements  -the active ones only and we can then call each elements component methods that needs updating.
		   Eventully we could look at adding concurrency to this iterative loop at a later stage to experiment putting bullets inside channels with players and enemies.

		*/
		/*
			//31 invoke the draw method which which initiate all components that are ready to be dranw and displayed
			err = plyr.draw(renderer)
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, benemy := range basicEnemies {
				err = benemy.draw(renderer) //draw single basic enemy
				if err != nil {
					fmt.Println(err)
					return
				}
				err = benemy.update()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			//33 invoke the update method which manages all components that need updating every frame
			err = plyr.update()
			if err != nil {
				fmt.Println(err)
				return
			}
		*/
		//39 - iterate and range through the active Global elemtents in the []element slice
		for _, e := range elements {
			if e.active { //check that the element is marked as active before proceeding
				//33 invoke the update method which manages all components that need updating every frame
				err = e.update() //call the update() method on elemenet - this will take the components attached and look for updating methods that need executing so that is refreshes the state of the element to the caller
				if err != nil {
					fmt.Println(err)
					return
				}
				//31 invoke the draw method which initiates all components that are ready to be drawn and displayed
				err = e.draw(renderer)
				if err != nil {
					fmt.Println(err)
					return

				}
			}
		}
		/*
			for _, bul := range bulletPool { //init bullet pool - for drawing out bullets from
				bul.draw(renderer)
				bul.update()
			}

		*/

		err = checkCollisions()
		if err != nil {
			fmt.Println(err)
			return
		}
		renderer.Present()

		//fmt.Println(time.Since(frameStartTime))//learn how how a frame takes to process in ms
		delta = time.Since(frameStartTime).Seconds() * float64(targetTicksPerSecond) //set bench mark value on delta based on time it take to load frame - game speeds are affected by this delta on different hardware architechtures
	}
}
