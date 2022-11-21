# Learnings about structs, types and Interfaces

### struct
- These are just key value pairs
- We usually use them to define a type, since different types have different properties, structs help us to add these properties in the type.
```
type Rectangle struct {
	width float64
	height float64
}
```
- above is an example of a type Rectangle.
```
type Circle struct {
	radius float64
}
```
- above is an example of type Circle.
- As you can see the basic difference between these types is the properties present in struct.

### types
- We use types to **encapsulate** logic together, best example is, area, perimeter etc of different shapes.
- There are different formulae for different shapes in order to calculate these properties.
- So, instead of making a bunch of functions, and then deciding what to call when, passing the arguments. Why not make something that is self containing ? want to calculate area of a triangle ? use triangle.area() and similarly for rectangle, use rectangle.area()
- That is where types come in, you can attach methods to them.

### methods
- These are similar to normal functions, only differnce is they have a reciever parameter.
```
type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64 {
	 return r.width * r.height
}
```
- Above we have a type, of rectangle. below it there is a method which can be only called by an instance of type Rectange, using rect.Area(). Because it a method of type Rectangle.
- There is a reciever ```(r Rectangle)```, it tells that the function is attached to type Rectangle, and it also gives gives access to the method, to instance which called the function in the first place. eg: rect in out above case. 

### Interfaces:
- Interfaces help in making functions that can be invoked by different types.
- The thing is, eventually we will need something, through which we can calculate stuff like area, perimeter etc. All the user cares about is passing in what is needed and getting the result.
- So, it doesn't matter if it is a rectangle, triangle etc.
- Interfaces help us decide what is to be done, so an interface for calculating area will look like this, 
```
type Shape interface {
	Area () float64
}
```
- All the above interface cares about is a function, that calculates the area and returns a float value.
- Now, in order to implement this, we will need **type**(s).
- Each type, that has a method which does the same thing as needed i.e. has a method named **Area** and returns a float64, can implement our interface.
```
type Circle struct {
	radius float64
}

func (c Circle) Area() float64{
	return math.pi * r.radius * r.radius
}
```
- The above struct implements the Shape interface, since there is a method name Area, that returns a float64 value.
- So, in order to calculate that Area we just have to call the function using any type, that can implement it.

# Table driven tests:
- It is all about running similar test cases using a loop, instead of specifically writing one for each.
- we have done it in the structs_test.go example.









