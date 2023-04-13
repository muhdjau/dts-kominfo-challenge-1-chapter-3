package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	h8HelperRand "github.com/novalagung/gubrak/v2"
)

func main() {
	waterStatus := func(nilai int) {
		fmt.Print("status water : ")
		if nilai <= 5 {
			fmt.Println("aman")
		} else if nilai >= 6 && nilai <= 8 {
			fmt.Println("siaga")
		} else {
			fmt.Println("bahaya")
		}
	}

	windStatus := func(nilai int) {
		fmt.Print("status wind : ")
		if nilai <= 6 {
			fmt.Println("aman")
		} else if nilai <= 7 && nilai <= 15 {
			fmt.Println("siaga")
		} else {
			fmt.Println("bahaya")
		}
	}

	for {
		waterNum := h8HelperRand.RandomInt(1, 100)
		windNum := h8HelperRand.RandomInt(1, 100)

		PostData(waterNum, windNum)
		waterStatus(waterNum)
		windStatus(windNum)

		fmt.Println(strings.Repeat("-", 22))
		time.Sleep(15 * time.Second)
	}
}

func PostData(waterNum, windNum int) {
	data := map[string]interface{}{
		"water": waterNum,
		"wind":  windNum,
	}

	reqJson, err := json.Marshal(data)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
