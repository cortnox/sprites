/*
FIRST
we want a class to hold all these player components.
We need a GENERIC type that can hold all of these gaming components
so everything in the game is derived from the same type
everything in the game is an instance of the same type
and what differentiates all of these types are the different components that these so called game objects have.
We are going to call this class/type the Element.

This Generic Type called Element is defined in its own file - which is the file that you are looking at right now..
*/
package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

//1 define vector class to compartmentalize and store x,y coordinates will be used to set element sprite positioning when rendered on screen.
type vector struct { //vectors compartmentalize into coodinates which are used on elements, so every single element will have its own set of coordinates it can store. These coordinate will be used to position the element if it needs to appear in the game.
	x, y float64 //store x,y coordinates// this is so that element positions can be referenced as myElement.position.x and myElement.position.y when setting the elements coordinates.
}

//2 define type element as a struct
type element struct { //element will represent base entity for evertything that exists in the game space//file the Element will hold common variables that every gaming elemrnt requires in order to be created
	position   vector      //vector holds position x,y coordinates
	rotation   float64     //field to take an rotation angle necessary for rotating textures such as enemies - alot of elements will benifit from having a rotation property.
	active     bool        //marker for element activity so that elements can be disabled if necessary - if set to false - the rule is that the element should not be visible based on this condition
	components []component //element needs to hold a set of components to provide it functionality - each component is custom written and initialized and added to its respective elemtent and utilized when the element is created.
	collisions []circle    //55 each element can have zero or more defined collision points or circles - which will outpost all collision activiy between other circles - that also have their collisions points - keeping in mind that the circles will be atached to elements
	tag        string
}

//3 define component interface - which every single element is going to implement - these mthods wills basically allow the components to manage themselves. - one might be thinking  how useful can a component really be if it only has two methods? - well each component has a little bit more than just ywo methods
type component interface { //a component can be just about anything - anything that can play the part of a component needs to atleast provide the methods defined in this given interface//methods needs to be generally useful for the vast majority of components//each component can have its own method by this common name
	onUpdate() error                     //-  this method is called every frame and allows elements to do any physical or physics related changes according to their component defined functionality
	onDraw(renderer *sdl.Renderer) error //takes a renderer argument and returns an error - this method renders the component//puts element texture on screen every frame
	onCollision(other *element) error    // define oncollisions metheod all elemtns must have
} //-each method returns an error //error gives ability to signal that something was unsuccessful
//create methods that will manage components
//4 define methods on element - addComponent - the purpose of this methods is to ensure that this new component we want to add to the element doiesnt share the same type with any other componets already attached to the element.
func (elem *element) addComponent(new component) { //takes a component we want to add as an argument - this will allow us to add a component to an element - giving the element the functionality of that component
	for _, existing := range elem.components { //loop through every component in this element that we are trying to add a component to
		// check that there is no component that has already been added to this element that may resemble the very same component we are feeding this element.
		if reflect.TypeOf(new) == reflect.TypeOf(existing) { //we ne go through each element component and do a type assertion//compare new and existing underlying types attached to the element//every component will be of type component but there will be underlying component types that we want to be able differentiate/distinguish each other by
			panic(fmt.Sprintf("element.go/addComponent - attempted to add %T but %T alreadyt exists:\n", //if there is a match we want to panic because this should be resovled before even compiling//what component each element has NEEDS to be known at compile time.//the caller must know about this before hand if a panic does occur - because the element already has the given component.
				reflect.TypeOf(new), reflect.TypeOf(existing))) //reflect returns a type object which describes the underlying type of this new component
		}
	}
	fmt.Printf("element.go/addComponent - added %T to %T:\n",
		reflect.TypeOf(new), reflect.TypeOf(elem))
	elem.components = append(elem.components, new) //add component to element []component if there is no existing component that matches this new component
} //we want to mak sure that this component doesnt share the same type with any other component already attached to the element//tenant of component based design - there is no need for an element to hold multiple components of the same type - each component type provides its own functionality - same functionality is not needed more than once//one would never need to define two methods that define and do exactly the same thing on an object
//5 define getComponent
func (elem *element) getComponent(withType component) component { //take in a template type as an argument to determine if element already has component with this given type - gets a component back from the element if it finds it it returns the componet that matches by type.
	for _, _type := range elem.components { //loop through every component in this element//the function will be called with an empty initialized type or dummy object that will be named something as &MyComponent -representing _type - note pointer address to component. the dummy will represent the component we are looking based on the dummies type
		if reflect.TypeOf(_type) == reflect.TypeOf(withType) { //use reflect.TypeOf() to compare component (dummy component with the element component)//like in addComponent() we just just want to loop through each iteration of component an compoare its type with the given template type - and see if there any matching component to return.
			fmt.Printf("element.go / getComponent() - fetched %T for %T :\n",
				reflect.TypeOf(_type), reflect.TypeOf(elem)) //reflection is never clear according to one of the go proverbs
			return _type
		}
	} //we want want to use reflect as little as possible through out the program
	panic(fmt.Sprintf("element.go / getComponent() - could not find matching type for %T\n", reflect.TypeOf(withType)))
} //the usage of this method would look like - myElement.getComponent(&myComponent{})//were just passig an empty initialized object into the function to represent what we are looking for //components are always pointers - an empty component of the provided type is sent into function to get the component that matches its type.

//30 Define onUpdate Method for elemtnet
func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}
	return nil
}

//31 Define onUpdate method for element
func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

//32 - Recycling Components out of the architecture - adding basic enemy
/*
so by now we have broken down the initial oop and we've converted the architecture into individual and reusable components.
the caller can see that by now the player element is established by just one big creation function containing all the required component initializations. - necessary to define that element
and it is now just a matter of invoking that player function to create a new player element in the main function and then further invoke established methods for that player element such as drawing moving and shooting
It is important to note that all individual components are now reusable for other elements such as the bullet and the enemy. so we want to break down both the enemy and bullet types into elements that cand take in existing components just like we did with the player.
*/

//38 - Define Global variable that will hold all active elements that need updating.
var elements []*element //in one foul swoop we iterate this slicein the main//this wil be a gloable registry of elements which will be iterated over inside the main fuction so that each elements methods can be accessed for updating throughout caller game play//we need to make sure that we append all elements to the global []element slice upon element creation.
//61
//embue the onCollision() method elements require in order to implement the componernt interface//this need restructure and here is where component architechture is starting to break - we need to consider redoing collision system as as interface on its own rather attaching to elements.
func (elem *element) collision(other *element) error {
	for _, comp := range elem.components {
		err := comp.onCollision(other)
		if err != nil {
			return err
		}
	}
	return nil
}
