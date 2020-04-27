# Command Pattern Example

* Build Comamnd Interface
```go
type Command interface {
    execute()
}
```

* Creating Commands
 
```go
// The restaurant contains the total dishes and the total cleaned dishes
type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// `NewRestaurant` constructs a new restaurant instance with 10 dishes,
// all of them being clean
func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

// The MakePizzaCommand is a struct which contains
// the number of pizzas to make, as well as the
// restaurant as its attributes
type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	// Reduce the total clean dishes of the restaurant
	// and print a message once done
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}

// The MakeSaladCommand is similar to the MakePizza command
type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	// Reset the cleaned dishes to the total dishes
	// present, and print a message once done
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}
```

* Add methods to `Restaurant`

```go
func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}
```
* Executing Commands

```go
// A Cook comes with their list of commands as attributes
type Cook struct {
	Commands []Command
}

// The executeCommands method executes all the commands
// one by one
func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}
```

* Putting it All together

```go
func main() {
	// initialize a new resaurant
	r := NewResteraunt()

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	// Assign tasks to cooks alternating between the existing
	// cooks.
	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	// Now that all the cooks have their commands, we can call
	// the `executeCommands` method that will have each cook
	// execute their respective commands
	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
```

* output
```shell
cook 0 :
made 2 pizzas
made 3 pizzas
made 4 pizzas
cook 1 :
made 1 salads
dishes cleaned
dishes cleaned
```
