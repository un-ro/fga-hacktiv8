package main

import (
	"bytes"
	"encoding/json"
	"environment-service/models"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	URL      string = "https://jsonplaceholder.typicode.com/posts"
	DURATION        = time.Minute * 5
)

func main() {
	for {
		// Menampilkan waktu saat ini
		fmt.Println("Waktu:", time.Now().Format("2006-01-02 15:04:05"))

		// Generate data
		water := rand.Float64() * 100
		wind := rand.Float64() * 100
		waterStatus := getStatus(water, 5, 8)
		windStatus := getStatus(wind, 6, 15)

		// Buat struct
		data := models.Data{
			Water: water,
			Wind:  wind,
		}

		// Encode data
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error encoding data", err.Error())
			continue
		}

		// Buat request
		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request", err.Error())
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		// Kirim request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request", err.Error())
			continue
		}

		// Baca response
		body := make([]byte, resp.ContentLength)
		_, err = resp.Body.Read(body)
		if err != nil {
			fmt.Println("Error reading response", err.Error())
			continue
		}
		fmt.Println("Response:", string(body))

		// Print ke console
		fmt.Printf("Water: %.2f m, Status: %s \nWind: %.2f m/s, Status: %s \n", water, waterStatus, wind, windStatus)

		// Tutup koneksi
		client.CloseIdleConnections()
		fmt.Println("=====================================")

		// Tunggu 15 detik
		time.Sleep(DURATION)
	}
}

func getStatus(value, lowerBound, upperBound float64) string {
	if value < lowerBound {
		return "Aman"
	} else if value >= lowerBound && value <= upperBound {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}
