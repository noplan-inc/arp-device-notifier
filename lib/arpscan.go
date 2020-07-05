package lib

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/chrismarget/arp"
)

func Run(o *PostOptions) {
	log.Println("Start!")
	var client http.Client

	for {
		table := arp.Table()
		if table == nil {
			log.Fatal("Failed to get arp table.")
		}
		if o.Verbose {
			for k, v := range table {
				log.Printf("%v @ %v\n", k, v)
			}
		}

		params, err := json.Marshal(table)
		if err != nil {
			log.Println("a")
		}

		req, err := http.NewRequest("POST", o.Endpoint, bytes.NewReader(params))
		if err != nil {
			log.Println("Failed to create request.")
			continue
		}
		req.Header.Set("Content-Type", "application/json")

		if o.Authorization != "" {
			req.Header.Set("Authorization", o.Authorization)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to request json. \n\nBody: %v\n\nError: %v", string(params), err)
			continue
		}

		if resp != nil && o.Verbose {
			log.Printf("Status Code: %d", resp.StatusCode)
		}

		time.Sleep(time.Duration(o.PostInterval) * time.Second)
	}
}
