package main

import (
	"sync"
)

/*
import(
	"fmt"
	"reflect"
)
*/

//"strconv"
//"example.com/go-demo-1/mascot"
//"example.com/go-demo-1/variable"

var i int = 45
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	// Variables
	/*
			i := 42 // var i int = 42
			var j string
			j = strconv.Itoa(i)
			// Print string
			fmt.Println(mascot.BestMascot())
			// Print variable
			fmt.Printf("The number is : %v and the type is: %T\n", variable.PrintVar(i), i)
			fmt.Printf("The number is : %v and the type is: %T", j, j)


		// Primitives
		/*
			n := 1 == 1
			m := 1 == 2
			fmt.Printf("n: %v , Type %T\n", n, n)
			fmt.Printf("m: %v , Type %T\n", m, m)
			// tipuri diferite
			var a int = 10
			var b int = 3
			fmt.Println(a + int(b))
			// bit operators
			a := 10             // 1010
			b := 3              // 0011
			fmt.Println(a & b)  // and
			fmt.Println(a | b)  // or
			fmt.Println(a ^ b)  // xor
			fmt.Println(a &^ b) //and not
			//bit shifting
			a := 8              // 2^3
			fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
			fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0
			// Runes aka int32
			r := 'a'
			fmt.Printf("%v %T\n", r, r)
	*/
	// Constants
	/*
		const myConstant int = 42
		fmt.Println(myConstant)

		// IOTA aka enumerated constants
		const (
			a = iota
			b //= iota
			c //= iota
		)

		const (
			a2 = iota
		)
		fmt.Printf("%v\n", a)
		fmt.Printf("%v\n", b)
		fmt.Printf("%v\n", c)
		fmt.Printf("%v\n", a2)

		// Check if a value has been asignad to a constant yet
		const (
			_ = iota + 5 // errorSpecialist = iota
			catSpecialist
			dogSpecialist
			snakeSpecialist
		)
		var specialistType int
		fmt.Printf("%v\n", specialistType == 0)
		fmt.Printf("%v\n", catSpecialist)

	*/

	// Arrays and Slices
	/*
		// Array
		grades := [3]int{97, 85, 93}
		fmt.Printf("Grades: %v", grades)

		var students [5]string
		fmt.Printf("Students:%v\n", students)
		students[0] = "Lisa"
		students[1] = "Ahmed"
		students[2] = "Arnold"
		fmt.Printf("Student #1: %v\n", students[2])
		fmt.Printf("Number of Students: %v", len(students))

		a := [...]int{1, 2, 3}
		b := a // b := &a in loc sa creeze o copie a arrayului care poate fi modificata fara sa influenteze arrayul
		//original , apeleaza arrayul original si il modifica
		b[1] = 5
		fmt.Println(a)
		fmt.Println(b)

		//Slices
		a := []int{1, 2, 3}
		fmt.Println(a)
		fmt.Printf("Length a: %v\n", len(a))
		fmt.Printf("Capacity a: %v", cap(a))

		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b := a[:]   //slice of all elements
		c := a[3:]  //slice from 4th element to end
		d := a[:6]  //slice from 6 elements
		e := a[3:6] //slice the 4th,5th and 6th elements
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
		fmt.Println(e)

		a := make([]int, 3, 100)
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
		a = append(a, 1)
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))

		a := []int{1, 2, 3, 4, 5}
		fmt.Println(a)
		b := append(a[:2], a[3:]...) // elimina element din interior slice cu concatenarea a 2 sliecuri, se adauga ...
		fmt.Println(b)
		fmt.Println(a)
	*/

	// Maps and structs
	// Maps
	/*


		statePopulations := make(map[string]int)
		statePopulations = map[string]int{
			"California": 39250017,
			"Texas":      27862596,
			"Florida":    20612439,
			"New York":   12902503,
			"Illinois":   12801539,
			"Ohio":       11614373,
		}
		statePopulations["Georgia"] = 10310371
		delete(statePopulations, "Georgia")
		fmt.Println(statePopulations)

		pop, ok := statePopulations["Oho"] // sau _, ok :=....
		fmt.Println(pop, ok)               // sau fmt.Println( ok)

		fmt.Println(len(statePopulations))
	*/
	// Struct
	/*

		type Doctor struct {
			number     int
			actorName  string
			companions []string
		}
		aDoctor := Doctor{
			number:    3,
			actorName: "Jon Pertwee",
			companions: []string{
				"Liz Shaw",
				"Jo Grant",
			},
		}
		fmt.Println(aDoctor)
		fmt.Println(aDoctor.actorName)

		aDoctor := struct{ name string }{name: "John Pertwee"}
		anotherDoctor := aDoctor
		anotherDoctor.name = "Tom Baker"
		fmt.Println(aDoctor)
		fmt.Println(anotherDoctor) // comparat cu slice nu afecteaza primul map ca in cazul sliceurilor
	*/
	// Embedding
	/*
		type Animal struct {
			Name   string
			Origin string
		}
		type Bird struct {
			Animal
			SpeedKPH float32
			CanFly   bool
		}

			b := Bird{}
			b.Name = "Emu"
			b.Origin = "Australia"
			b.SpeedKPH = 48
			b.CanFly = false
			fmt.Println(b)
		// sau
		b := Bird{
			Animal:   Animal{Name: "Emu", Origin: "Australia"},
			SpeedKPH: 48,
			CanFly:   false,
		}
		fmt.Println(b.Name)
		// Tags
		type Animal struct {
			Name   string `required max:"100"`
			Origin string
		}

		t := reflect.TypeOf(Animal{})
		field, _ := t.FieldByName("Name")
		fmt.Println(field.Tag)
	*/
	// If and switch statements
	/*

		if true {
			fmt.Println("Test is true")
		}

		statePopulations := map[string]int{
			"California": 39250017,
			"Texas":      27862596,
			"Florida":    20612439,
			"New York":   12902503,
			"Illinois":   12801539,
			"Ohio":       11614373,
		}
		if pop, ok := statePopulations["Florida"]; ok {
			fmt.Println(pop)
		}


		number := 50
		guess := 12
		if guess < 1 || guess > 100 {
			fmt.Println("The guess must be between 1 and 100")
		} else {
			//if guess >= 1 && guess <= 100 {
			if guess < number {
				fmt.Println("Too low")
			}
			if guess > number {
				fmt.Println("Too high")
			}
			if number == guess {
				fmt.Println("You got it")
			}
			// sau
			fmt.Println(number <= guess, number >= guess, number != guess)
			//}
		}

		switch 1 {
		case 1, 5, 10:
			fmt.Println("one, five or ten")
		case 2, 4, 6:
			fmt.Println("two, four or six")
		default:
			fmt.Println("another number")
		}
		var i interface{} = [3]int{}
		switch i.(type) {
		case int:
			fmt.Println("i is an int")
		case float64:
			fmt.Println("i is a float64")
		case string:
			fmt.Println("i is string")
		case [3]int:
			fmt.Println("i is [3]int")
		default:
			fmt.Println("i is another type")
		}
	*/
	// Looping
	/*

		// for loop
		for i := 0; i < 5; i = i + 2 {
			fmt.Println(i)
		}

		for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
			fmt.Println(i, j)
		}

		Loop:
			for i := 1; i <= 3; i++ {
				for j := 1; j <= 3; j++ {
					fmt.Println(i * j)
					if i*j >= 3 {
						break Loop
					}

				}
			}

		//for range loop
		s := []int{1, 2, 3} //
		for k, v := range s {
			fmt.Println(k, v) // key and value
		}

		x := "Hello Go!"
		for k, v := range x {
			fmt.Println(k, string(v)) // key and value
		}
		// only keys
		for k := range...
		// only values
		for _, v := range...
	*/
	// Control flow
	/*
		// defer
			defer fmt.Println("Start")
			defer fmt.Println("Middle")
			defer fmt.Println("End")

			res, err := http.Get("http://www.google.com/robots.txt")
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			robots, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)

			a := "start"
			defer fmt.Println(a)
			a = "end"

		// pannic
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Go!"))
		})
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("start")
		defer fmt.Println("this was deffered")
		panic("something bad happened")
		fmt.Println("end")
		fmt.Println("start")
		panicker()
		fmt.Println("end")
	*/
	// Pointers

	/*
		var a int = 42
		var b *int = &a
		fmt.Println(a, *b)
		a = 27
		fmt.Println(a, *b)
		*b = 14
		fmt.Println(a, *b)

		var ms *myStruct
		ms = new(myStruct)
		ms.foo = 42
		fmt.Println(ms.foo)
	*/
	// Functions
	/*
			for i := 0; i < 5; i++ {
			sayMessage("Hello Go!", i)

			greeting := "Hello"
			name := "Stacey"            // pointer
			sayGreeting(greeting, name) //  sayGreeting(&greeting, &name)
			fmt.Println(name)

			s := sum(1, 2, 3, 4, 5)
			fmt.Println("The sum is", s) // *s

			d, err := divide(5.0, 3.0)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(d)
		// Method
			g := greeter{
				greeting: "Hello",
				name:     "Go",
			}
			g.greet()
			fmt.Println("The new name is:", g.name)

	*/
	//Interfaces
	/*

		var w Writer = ConsoleWriter{}
		w.Write([]byte("Hello Go!"))

		myInt := IntCounter(0)
		var inc Incrementer = &myInt
		for i := 0; i < 10; i++ {
			fmt.Println(inc.Increment())
		}
	*/
	// GO ROUTINEs
	/*
		var wg = sync.WaitGroup{}
		var msg = "Hello"
		wg.Add(1)
		go func(msg string) {
			fmt.Println(msg)
			wg.Done()
		}(msg)
		wg.Wait()
	*/
	/*
		runtime.GOMAXPROCS(100)

		for i := 0; i < 10; i++ {
			wg.Add(2)
			m.RLock()
			go sayHello()
			m.Lock()
			go increment()
		}
		wg.Wait()
	*/
	// CHANNELS
	/*
		ch := make(chan int)
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 42
			wg.Done()
		}(ch)
		wg.Wait()
	*/
	/*
		ch := make(chan int, 50)
		wg.Add(2)
		go func(ch <-chan int) {
			for i := range ch {
				fmt.Println(i)
			}
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 42
			ch <- 27
			close(ch)
			wg.Done()
		}(ch)
		wg.Wait()
	*/

}

/*
func sayHello() {
	fmt.Printf("hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
*/
/*
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
*/

/*
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
*/

/*
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
*/
/*
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
*/

/*
func sum(values ...int) int { // ...int) *int {....}
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}
*/

/*
func sayGreeting(greeting, name string) { //func sayGreeting(greeting, name *string)
	fmt.Println(greeting, name) //fmt.Println(*greeting, *name)
	name = "Ted"
	fmt.Println(name)
}
*/
/*
func sayMessage(msg string, index int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", index)
}
*/

/*
type myStruct struct {
	foo int
}
*/

/*
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			panic(err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done pancking")
}
*/
