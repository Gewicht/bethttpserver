package main

import (
"net/http"
"log"
)

func main(){

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
	http.ListenAndServe(":1234",nil)
}

func BetServerListener(w http.ResponseWriter, req *http.Request){


// Use URL to decide what API call is triggered and call function

	apiCall :=req.URL.Path[1:]
//	log.Println("New request arrived on path: " + apiCall)
	if apiCall == "GetLeagues"{
		log.Println("New request arrived on path: " + apiCall +" -> Server was asked to provide Leagues")
		w.Write([]byte("List of Leagues"))
	}else if apiCall == "GetOffer" {
		log.Println("New request arrived on path: " + apiCall +" -> Server was asked to provide Offer")
		w.Write([]byte("Offer"))
	}else if apiCall == "AddOffer"{
		log.Println("New request arrived on path: " + apiCall +" -> Server was asked add new offer to list")
		w.Write([]byte("Adding Offer"))
	}else if apiCall == "NewTicket"{
		log.Println("New request arrived on path: " + apiCall +" -> New ticket received")
		w.Write([]byte("Verifying ticket"))
	}else{
		log.Println("New request arrived on path: " + apiCall +" -> API CALL unknown")
		w.WriteHeader(404)
		w.Write([]byte("Unknown call"))

	}
}