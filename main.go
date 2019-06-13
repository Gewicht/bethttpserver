package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	//	_ "github.com/denisenkom/go-mysqldb"
)

// GLOBAL VARIABLE DECLARATION

var LoadedLeagueListGlobal LigeList
var OffersGlobal []Ponude

type LigeList struct {
	Lige []Lige `json:"lige"`
}
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

type LeagueResp struct {
	NazivLige   string `json:"nazivlige"`
	PonudeULigi []int  `json:"ponudeuligi"`
}

type Ponude struct {
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
	log.Println("\n****************************************************\n******************* Starting APP *******************\n****************************************************")
	/*
		log.Println("Trying to bind to DB")
		//	db, err := sqlx.Connect("mysql", "betserver:zmrU9QkIFd4qGszl@(localhost:3306)/betserver")
		db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/mysql")

		if err != nil {
			log.Println("Error occur. Action stopped\n",err)
		}
		db.MustExec("insert into login_log values ('login');")
	*/

	getLeaguesFromServerHTTP("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/lige.json")
	getOffersFromServerHTTP("https://minus5-dev-test.s3.eu-central-1.amazonaws.com/ponude.json")
	//log.Println("Pronadena ", OffersGlobal[getOfferByID(8679702)])

	http.HandleFunc("/", BetServerListener)
	http.ListenAndServe(":1234", nil)
}

func BetServerListener(w http.ResponseWriter, req *http.Request) { // Use URL to decide what API call is triggered and call function

	apiCall := req.URL.Path[1:]
	//	log.Println("New request arrived on path: " + apiCall)
	if apiCall == "GetOffersPerLeagues" { // davanje popisa liga na GET method zahtjev
		log.Println("New request arrived on path: " + apiCall + " -> Server was asked to provide Leagues and offers for leagues")
		lr := getLeaguesResp()
		log.Println(lr)
		resptext, err := json.Marshal(lr)
		if err != nil {
			log.Println("Error occur. Action stopped\n", err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(resptext)
	} else if apiCall == "GetOfferById" { //davanje ponude na GET method zahtjev
		keys, ok := req.URL.Query()["offer_id"]
		if !ok || len(keys[0]) < 1 { //ako ne postoji key ili je vrijednost kraca od 1 znak
			log.Println("New request arrived on path: ", apiCall, " -> Server was asked to provide Offer but parameter offer_id is missing")
			w.Write([]byte("offer_id value is missing"))
		} else { //ako je key matchan i ima pozitivnu vrijednost
			var oid int
			oid, _ = strconv.Atoi(keys[0])
			oid = getOfferByID(oid)
			if oid > 0 { // ako je key vrijednost pozitivan intiger
				log.Println("New request arrived on path: ", apiCall, " -> Server was asked to provide Offer with ID:", keys[0], "/n", OffersGlobal[oid])
				resptext, err := json.Marshal(OffersGlobal[oid])
				if err != nil {
					log.Println("Error occur. Action stopped\n", err)
				}
				w.Header().Add("Content-Type", "application/json")
				w.Write(resptext)
			} else {
				log.Println("New request arrived on path: ", apiCall, " -> Server was asked to provide Offer with ID:", keys[0], " but it is invalid")
				w.Write([]byte("offer_id value is invalid"))
			}
		}
	} else if apiCall == "AddOffer" { // dodavanje nove ponude na POST metodu s json bodyem
		//log.Panicln(req.Method)
		if req.Method != "POST" {
			log.Println("New request arrived on path: " + apiCall + " that is not POST method.")
			w.Write([]byte("Only POST method allowed"))
		} else {
			log.Println("New POST request arrived on path: " + apiCall + " -> Server was asked add new offer to list")
			/*
				//Ispis Headera
				for name, value := range req.Header {
					log.Println("", name, " : ", value)
				}
			*/

			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Println("Error occur. Action stopped\n", err)
			} else if req.Header.Get("Content-Type") == "application/json" {
				//log.Println("POST body is json:\n", string(body))
				var P []Ponude
				err := json.Unmarshal(body, &P)
				if err != nil {
					var Po Ponude
					err2 := json.Unmarshal(body, &Po)
					if err2 != nil {
						log.Println(err2)
					}
					P = append(P, Po)
				}
				//log.Println("POST body is json:\n" + string(body))
				//log.Println(P)

				//log.Println(provjeraPopisaPonuda())
				cnt := 0
				for i := 0; i < len(P); i++ {
					if !checkExistingOffer(P[i].ID) && P[i].ID > 0 {
						OffersGlobal = append(OffersGlobal, P[i])
						cnt++
					}
				}
				for i := 0; i < len(P); i++ {
					log.Println(P[i])
				}

				log.Println(provjeraPopisaPonuda())
				log.Println("POST body contains", len(P), "possible offers. Added", cnt, "offers to existing list")

				w.Write([]byte("Offers added!"))
			} else {
				log.Println("POST body is not json\n", string(body))
				w.Write([]byte("No json file in POST body"))
			}
		}
	} else if apiCall == "NewTicket" {
		log.Println("New request arrived on path: " + apiCall + " -> New ticket received")
		w.Write([]byte("Verifying ticket"))
	} else {
		log.Println("New request arrived on path: " + apiCall + " -> API CALL unknown")
		w.WriteHeader(404)
		w.Write([]byte("Unknown call: " + apiCall))

	}
}
func getLeaguesFromServerHTTP(u string) { // Funkcija kojom dohvacam lige.json

	log.Println("Sending request to: " + u)
	resp, err := http.Get(u)
	if err != nil {
		log.Println("Error occur. Action stopped\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	/*
		//Ispis Headera
		for name, value := range resp.Header {
			log.Println("", name, " : ", value)
		}
	*/
	if resp.Header.Get("Content-Type") == "application/json" {

		//Ispis Bodya
		//log.Println(string(body))

		//////procitaj JSON

		err := json.Unmarshal(body, &LoadedLeagueListGlobal)
		if err != nil {
			log.Println("Error occur. Action stopped\n", err)
		}
		//log.Println(leagueJSON)
		log.Println("lige.json received and parsed. Leagues loaded:", len(LoadedLeagueListGlobal.Lige))
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
}
func getLeaguesResp() []LeagueResp { //funkcija koja vraca ime lige i sve ponude u ligi  --> API CALL 1

	lr := []LeagueResp{}
	for i := 0; i < len(LoadedLeagueListGlobal.Lige); i++ {
		L := LeagueResp{LoadedLeagueListGlobal.Lige[i].Naziv, LoadedLeagueListGlobal.Lige[i].Razrade[0].Ponude}
		//log.Println("Nasao ligu: ", L.NazivLige)
		//log.Println(L.PonudeULigi)
		lr = append(lr, L)
	}
	//log.Println("Found ", len(lr), " leagues to deliver")
	return lr
}
func getOffersFromServerHTTP(u string) { // Funkcija kojom dohvacam ponude.json

	log.Println("Sending request to: " + u)
	resp, err := http.Get(u)
	if err != nil {
		log.Println("Error occur. Action stopped\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if resp.Header.Get("Content-Type") == "application/json" {
		err := json.Unmarshal(body, &OffersGlobal)
		if err != nil {
			log.Println("Error occur. Action stopped\n", err)
		}
		log.Println("ponude.json received and parsed. Overall", len(OffersGlobal), "offer loaded")
	}
}
func getOfferByID(id int) int {

	for i := 0; i < len(OffersGlobal); i++ {
		if OffersGlobal[i].ID == id {
			return i
		}
	}
	return -1
}
func checkExistingOffer(newId int) bool {
	for i := 0; i < len(OffersGlobal); i++ {
		if OffersGlobal[i].ID == newId {
			return true
		}
	}
	return false
}

func provjeraPopisaPonuda() string {
	str := ""
	for i := 0; i < len(OffersGlobal); i++ {
		if i == 0 {
			str = strconv.Itoa(OffersGlobal[i].ID)
		} else {
			str += ", " + strconv.Itoa(OffersGlobal[i].ID)
		}
	}
	return str
}
