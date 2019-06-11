package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	//	_ "github.com/denisenkom/go-mysqldb"
	"github.com/jmoiron/sqlx"
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

/*********************************
Definicija JSON elemenata ponuda

********************************/
type Ponude []struct {
	Broj     string    `json:"broj"`
	ID       int       `json:"id"`
	Naziv    string    `json:"naziv"`
	Vrijeme  time.Time `json:"vrijeme"`
	Tecajevi []struct {
		Tecaj float64 `json:"tecaj"`
		Naziv string  `json:"naziv"`
	} `json:"tecajevi"`
	TvKanal       string `json:"tv_kanal,omitempty"`
	ImaStatistiku bool   `json:"ima_statistiku,omitempty"`
}

func main() {

	/*********************************

	  Check if DB connection can be established
	  Make connection to DB
	  Truncate tables
	  ********************************/

	log.Println("Trying to bind to DB")
	//	db, err := sqlx.Connect("mysql", "betserver:zmrU9QkIFd4qGszl@(localhost:3306)/betserver")
	db, err := sqlx.Connect("mysql", "root@tcp(localhost:3306)/mysql")

	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec("insert into login_log values ('login');")
	//////////////////////////////////////////////////////////////////////////////////////////
	/*********************************
	  Execute request to : https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json
	  Get JSON
	  Parse JSON
	  import to mysql
	  *********************************/

	ll := getLeaguesHTTP("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json")
	log.Println(ll)

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

func getLeaguesHTTP(u string) LigeList {

	log.Println("Sending request to: " + u)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	/*
		//Ispis Headera
		for name, value := range resp.Header {
			log.Println("", name, " : ", value)
		}
	*/
	leagueJSON := LigeList{}
	if resp.Header.Get("Content-Type") == "application/json" {

		//Ispis Bodya
		//log.Println(string(body))

		//////procitaj JSON

		err := json.Unmarshal(body, &leagueJSON)
		if err != nil {
			log.Fatal(err)
		}
		//log.Println(leagueJSON)

		log.Println("Ucitan JSON")

		/*
			for i := 0; i < len(leagueJSON.Lige); i++ {

				log.Println(leagueJSON.Lige[i].Naziv)
				for j := 0; j < len(leagueJSON.Lige[i].Razrade[0].Tipovi); j++ {
					log.Println(leagueJSON.Lige[i].Razrade[0].Tipovi[j])
				}

				for j := 0; j < len(leagueJSON.Lige[i].Razrade[0].Ponude); j++ {
					log.Println(leagueJSON.Lige[i].Razrade[0].Ponude[j])
				}
			}
		*/

	}
	return leagueJSON
}
