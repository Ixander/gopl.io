package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	person()
	str := `{
		"header": {
		"code": 0,
			"message": ""
	},
		"data": [{
	"type": "user",
	"id": 100,
	"attributes": {
	"email": "bob@yandex.ru",
	"article_ids": [10, 11, 12]
	}
	}]
}`
	res, _ := ReadResponse(str)

	fmt.Println(res)
}

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	c := Response{}
	b := []byte(rawResp)
	err := json.Unmarshal(b, &c)
	if err != nil {
		log.Fatalln("unable unmarshal to json")
	}
	return c, nil
}

func person() {
	type Person struct {
		Name        string    `json:"Ім'я"`
		Email       string    `json:"Пошта"`
		DateOfBirth time.Time `json:"-"`
	}
	per := Person{
		Name:        "Aлекс",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}

	jper, err := json.Marshal(per)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}

	fmt.Printf("Man %v\n", string(jper))

}
