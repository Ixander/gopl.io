package main

import "fmt"

func main() {

	//var voice func(string) string
	//voice = Say

	Print("dog", Say)

	f := func(s string) string { return s }
	fmt.Println(f("test"))

	Print("dog", Do(true))
	Print("animal", Do(true))
}
func Print(who string, how func(string) string) {
	fmt.Println(how(who))

}
func Do(say bool) func(string) string {
	if say {
		return Say
	}
	return func(s string) string { return s }
}
func Say(animal string) (v string) {
	switch animal {
	default:
		v = "heh"
	case "dog":
		v = "gav"
	case "cat":
		v = "myau"
	case "cow":
		v = "mu"
	}
	return
}
