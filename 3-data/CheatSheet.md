# Cheat Sheet

## Structs, Pointers

```go
type Person struct { // public
	Name string      // public
	age  int         // private
}

type user struct {   // private
	Person //embedding

	Id int
}

func update(p *Person) { //private
	p.Name = "Igor"
}

func (p Person) Speak() { //public
	fmt.Println("Hello!")
}

func (p *Person) HappyBirthday() {
	p.Age += 1
}

func main() {
	user := User{
		Person: Person{"Igor", 33},
		Id:     42,
	}
	user.Speak()
	user.HappyBirthday()

	fmt.Printf("%s is %d years old", user.Name, user.Age)
}

```


## Interfaces
```go
type Speaker interface {
	Speak()
}


//empty interface
func printStuff(something interface{} ) {
	fmt.Printf(something)
}
printStuff(5)
printStuff(Person{})

```

