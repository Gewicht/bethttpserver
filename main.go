package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

/*********************************
Definicija JSON elemenata liga

********************************/
type LigeList struct {
	Lige []Lige `json:"lige"`
}

/*type LigeList struct {
	Lige Lige `json:"lige"`
}*/

type Lige struct {
	Naziv   string    `json:"naziv"`
	Razrade []Razrade `json:"razrade"`
}
type Razrade struct {
	Tipovi []Tipovi `json:"tipovi"`
	Ponude []int    `json:"ponude"`
}

type Tipovi struct {
	Naziv string `json:"naziv"`
}

func main() {

	/*********************************

	  Check if DB connection can be established
	  Make connection to DB
	  Truncate tables
	  ********************************/

	//////////////////////////////////////////////////////////////////////////////////////////
	/*********************************




	  Execute request to : https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json
	  Get JSON
	  Parse JSON
	  import to mysql
	  *********************************/

	getLeaguesHTTP("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json")

	//////////////////////////////////////////////////////////////////////////////////////////
	/*********************************
	  Execute request to : https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json
	  Get JSON
	  Parse JSON
	  import to mysql
	  *********************************/

	//////////////////////////////////////////////////////////////////////////////////////////
	/*********************************
	  Start HTTP server
	  *********************************/
	http.HandleFunc("/", BetServerListener)
	http.ListenAndServe(":1234", nil)
}

func BetServerListener(w http.ResponseWriter, req *http.Request) {

	// Use URL to decide what API call is triggered and call function

	apiCall := req.URL.Path[1:]
	//	log.Println("New request arrived on path: " + apiCall)
	if apiCall == "GetLeagues" {
		log.Println("New request arrived on path: " + apiCall + " -> Server was asked to provide Leagues")
		w.Write([]byte("List of Leagues"))
	} else if apiCall == "GetOffer" {
		log.Println("New request arrived on path: " + apiCall + " -> Server was asked to provide Offer")
		w.Write([]byte("Offer"))
	} else if apiCall == "AddOffer" {
		log.Println("New request arrived on path: " + apiCall + " -> Server was asked add new offer to list")
		w.Write([]byte("Adding Offer"))
	} else if apiCall == "NewTicket" {
		log.Println("New request arrived on path: " + apiCall + " -> New ticket received")
		w.Write([]byte("Verifying ticket"))
	} else {
		log.Println("New request arrived on path: " + apiCall + " -> API CALL unknown")
		w.WriteHeader(404)
		w.Write([]byte("Unknown call"))

	}
}

func getLeaguesHTTP(u string) {

	log.Println("Sending request to: " + u)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	for name, value := range resp.Header {
		log.Println("", name, " : ", value)
	}
	if resp.Header.Get("Content-Type") == "application/json" {

		log.Println(string(body))

		//////procitaj JSON

		leagueJSON := LigeList{}
		err := json.Unmarshal(body, &leagueJSON)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(leagueJSON)

		log.Println("Ucitan JSON")
		for i := 0; i < len(leagueJSON.Lige); i++ {

			log.Println(leagueJSON.Lige[i].Naziv)
			for j := 0; j < len(leagueJSON.Lige[i].Razrade[0].Tipovi); j++ {
				log.Println(leagueJSON.Lige[i].Razrade[0].Tipovi[j])
			}

			for j := 0; j < len(leagueJSON.Lige[i].Razrade[0].Ponude); j++ {
				log.Println(leagueJSON.Lige[i].Razrade[0].Ponude[j])
			}
		}

	}
}
