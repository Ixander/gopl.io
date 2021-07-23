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

func main()  {

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

	fmt.Println(testEvents)

	_, ok := testEvents[5]
	fmt.Println(ok)



	/*if _, ok := testEvents[5]; ok{
		fmt.Println(ok)
	}*/


}
