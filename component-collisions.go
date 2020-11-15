//52
//the collisions pkg will hold a lot of collision information
package main

import "math"

//53
//we need to define and describe a type within a certain region that a collision can happen
type circle struct { //detecting collisions between two circles is very easy//we are going to stick with a basic circle collision system not to get too tied up with the extreaneous and tedious mathethmatical calculations involved with varoius polymetric collision detection systems
	center vector  //a circle is a form defined as having a center point - we therefore want this center point store in a vector type, the vector will stores xy coordinates in relation to it postion determined on screen
	radius float64 //radius is defined as the furtherest point out from the centre that the circle is at in all directions//all circles have a radius which is defined as the distance between any outer edge and the center point - which is equivilant to the length accross the circle -  if doubled
}

//54
//a function for detcting if two circles collided
func collides(c1, c2 circle) bool { //the function will return an evealuted condition of the two circles distances between each to determine if a collision was detected.
	//standard euclidian disance algorithm//we will use a simple distance algortithm to detect circle cillision
	dist := math.Sqrt(
		math.Pow(c2.center.x-c1.center.x, 2) +
			math.Pow(c2.center.y-c1.center.y, 2)) // we need to give the function a means to compare/evaluate and indicate distances between each circle radius
	return dist <= c1.radius+c2.radius //if their distance is less than the combined radius between the two circles dist distance above - then we can concur that a collision occured between these two circle and the function will return true
} //in the scenario that there is a collision between two circles//one circle is inside another circle- meaning both of their radiuses intercept//the slightestan furthest collision is determined by the two circle points with the radius of each circle away from each other touching each other right at the very edge

/*
by now we shold be in a situation where the basic enemy and the bullet are setup to have collisions so that we can now detect whether they collide or not//so this is a good spot here and now all we need to write code that compare elements to each other detects if they have collided.
*/

//59 Define function to compare elements and detect collisions
func checkCollisions() error { //function will iterate every global []element -within registry of gloabal elements//and chck them aghainst each other and see if they collide
	for i := 0; i < len(elements); i++ {
		for j := i + 1; j < len(elements)-1; j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if collides(c1, c2) && elements[i].active && elements[j].active {
						err := elements[i].collision(elements[j])
						if err != nil {
							return err
						}
						//elements[i].active = false
						err = elements[j].collision(elements[i])
						if err != nil {
							return err
						}
						//elements[j].active = false
					}
				}
			}
		}
	}
	return nil
}
