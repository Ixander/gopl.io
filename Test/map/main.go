package main

import "fmt"

const (
	EventsStatusNew       = "new"
	EventsStatusCreated   = "created"
	EventsStatusVerifying = "verifying"
	EventsStatusClosed    = "closed"

	EventPaymentSigned     = 10
	EventPaymentAuthorized = 13
	EventPaymentPaying     = 21
	EventPaymentCheck      = 5
	EventPaymentConveyor   = 51
	EventPaymentClosed     = 30
	EventPaymentReady      = 7
	EventPaymentPart       = 24
	EventPaymentPayed      = 22
)

func main() {

	testEvents := map[int]string{
		EventPaymentSigned:     EventsStatusVerifying,
		EventPaymentAuthorized: EventsStatusVerifying,
		EventPaymentPaying:     EventsStatusVerifying,
		EventPaymentCheck:      EventsStatusCreated,
		EventPaymentConveyor:   EventsStatusCreated,
		EventPaymentReady:      EventsStatusCreated,
		EventPaymentClosed:     EventsStatusClosed,
		EventPaymentPart:       EventsStatusClosed,
		EventPaymentPayed:      EventsStatusClosed,
		//1:                                   "new",
	}

	key, isexistvalue := testEvents[5]

	fmt.Println(isexistvalue, key)
	fmt.Println(testEvents)

	_, ok := testEvents[5]
	fmt.Println(ok)

	/*if _, ok := testEvents[5]; ok{
		fmt.Println(ok)
	}*/

	ages := map[string]int{
		"Максим": 20,
		"Олег":   25,
		"Саня":   28,
	}

	// видалити елемент з мапки
	delete(ages, "Олег")

	fmt.Println(ages)
	_, exists := ages["Антон"]

	if !exists {
		fmt.Println("Антона нет в списке")
		return
	}

	fmt.Printf("Антону %d лет\n", ages["Антон"])

}
