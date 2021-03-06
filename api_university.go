package cloudAss1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// HandlerUniversity detects wether a call is a post or get request, though this
// progam only has get requests this is only for potential future use
func HandlerUniversity(startTime int64) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			//This is incase the program is evert to have a post option
			//In other word this function is very redundant and everything
			//in handleUniInfoget couldve been put in the below case
		case http.MethodGet:
			handleUniGet(w, r, startTime)
		}
	}
}

// handleUniInfoGet utility function, package level, to handle GET request to university route
func handleUniGet(w http.ResponseWriter, r *http.Request, startTime int64) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	// error handling

	//This is extremely redundant and only for magical cases where someone manages to alter the URL after entering the function
	if len(parts) < 4 || parts[1]+"/"+parts[2] != UNI_BASE_PATH {
		fmt.Println("ERROR IN: HandleUniGet")
		http.Error(w, "Malformed URL", http.StatusBadRequest)
		return
	}
	//Checks for what

	switch parts[3] {
	case UNI_INFO_PATH:
		if 4 < len(parts) && parts[4] != "" {
			toScreen, _ := handleUniInfoGet(w, parts[4])
			w.Write(toScreen)
		} else {
			http.Error(w, "Page Empty, not enough mandatory URL components, example:\n"+HEROKU_PATH+"/"+UNI_BASE_PATH+"/"+UNI_INFO_PATH+"/{:partial_or_complete_university_name}/", http.StatusNotFound)
		}
	case UNI_NEIGH_PATH:
		if 5 < len(parts) && parts[4] != "" && parts[5] != "" {
			handleUniNeighbourGet(w, r)
		} else {
			http.Error(w, "Page Empty, not enough mandatory URL components, example:\n"+HEROKU_PATH+"/"+UNI_BASE_PATH+"/"+UNI_NEIGH_PATH+"/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}", http.StatusNotFound)
		}
	case UNI_DIAG_PATH:
		handleUniDiagGet(w, r, startTime)

	case "":
		http.Error(w, "Page Empty, example:\n"+HEROKU_PATH+"/"+UNI_BASE_PATH+"/"+UNI_NEIGH_PATH+"/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}\n"+HEROKU_PATH+"/"+UNI_BASE_PATH+"/"+UNI_INFO_PATH+"/{:partial_or_complete_university_name}/", http.StatusNoContent)
	}
}

// handleUniInfoGet utility function, package level, to handle GET request to university route
func handleUniInfoGet(w http.ResponseWriter, urlData string) ([]byte, []University) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	//I have a cleaning function, but since this is already split its a bit overkill
	urlData = strings.ReplaceAll(urlData, " ", "%20")
	data, err := http.Get("http://universities.hipolabs.com/search?name=" + urlData)
	if err != nil {
		fmt.Println("ERROR IN: handleUniInfoGet", urlData)
		log.Fatal(err)
	}
	defer data.Body.Close()
	//Prepears data for writing to screen
	parsedData, err2 := ioutil.ReadAll(data.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	//Creates an empty slot for the Uni data to go and fills it
	var Uni []University
	json.Unmarshal(parsedData, &Uni)
	if len(Uni) != 0 {
		countDat := getCountryDB(w, true, Uni[0].Country)
		for i := 0; i < len(Uni); i++ {
			Uni[i].Coutryholder = countDat[0]
		}
	}

	screenDat, err3 := json.Marshal(&Uni)
	if err3 != nil {
		log.Fatal(err3)
	}

	return screenDat, Uni
}

/*
handleUniNeighbourGet collects information about a country and its universities following a pattern from URL:
neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}}
then it gets data about universities in the surroudning countries
*/
func handleUniNeighbourGet(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	urlData := removeUrlExcess(r)                     //Cleans URL before split
	urlData = strings.ReplaceAll(urlData, " ", "%20") //Removes spaces in URL

	Parts := strings.Split(urlData, "/") //Splits URL Into workable strings

	countDb := getCountryDB(w, true, Parts[0])

	//Creates a list of all country codes with original country
	countCodeQuery := "," + countDb[0].Isocode
	for _, bord := range countDb[0].Bordering {
		countCodeQuery += "," + bord
	}
	//Overwrites countDB with all new data
	countDb = getCountryDB(w, false, countCodeQuery)

	//Gets Number from text
	limit := 0
	var err1 error
	limitString := r.URL.Query().Get("limit")
	if limitString != "" {
		limit, err1 = strconv.Atoi(limitString)
		if err1 != nil {
			log.Fatal(err1)
		} else if limit == 0 {
			limit = 9999 //Sets a default value if no limit is supplied, India has most universities in the world with over 5000
		}
	} else {
		limit = 9999
	}
	//Define empty University list
	var uniNeighDb []University
	for _, countUni := range countDb {
		//Gets information about all unis in correct country
		_, tempUni := handleUniInfoGet(w, Parts[1]+"&country="+countUni.CountryName.CommonName)
		for o := 0; o < limit && o < len(tempUni); o++ {
			//Appends university to the emptry list
			uniNeighDb = append(uniNeighDb, tempUni[o])
		}
	}
	//Prepares data for writing to webpage
	screenDat, err2 := json.Marshal(&uniNeighDb)
	if err2 != nil {
		log.Fatal(err2)
	}
	w.Write(screenDat)
}

// handleUniInfoGet utility function, package level, to handle GET request to university route
func handleUniDiagGet(w http.ResponseWriter, r *http.Request, startTime int64) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	diag := PrePopulateDiagnostics()
	diag.Uptime = time.Now().Unix() - startTime

	//Checks for status of the given website
	uniResp, err := http.Get("http://universities.hipolabs.com/")
	if err != nil {
		log.Fatal(err)
	}
	diag.UniversitiesApi = uniResp.StatusCode
	//Checks for status of the given website
	resp, err2 := http.Get("https://restcountries.com/")
	if err2 != nil {
		log.Fatal(err2)
	}
	diag.CountrisApi = resp.StatusCode

	diagData, err3 := json.Marshal(&diag)
	if err3 != nil {
		log.Fatal(err3)
	}

	w.Write(diagData)
}

func getCountryDB(w http.ResponseWriter, nameAlpha bool, countCodeName string) []Country {
	http.Header.Add(w.Header(), "content-type", "application/json")
	var countDat *http.Response //Reserve name

	if nameAlpha { //Checks if this is the original loop of the program
		//gets iniital data from first country
		tempDat, err := http.Get("https://restcountries.com/v3.1/name/" + countCodeName + "?fullText=true")
		if err != nil {
			log.Fatal(err)
		}
		countDat = tempDat
	} else {
		//Gets data from all bordering nations in second loop
		tempDat, err := http.Get("https://restcountries.com/v3.1/alpha?codes=cca3" + countCodeName)
		if err != nil {
			log.Fatal(err)
		}
		countDat = tempDat
	}

	pars, err2 := ioutil.ReadAll(countDat.Body)
	if err2 != nil {
		log.Fatal(err2)
	}
	//Enters json data into new Country list
	var countryDB []Country
	json.Unmarshal(pars, &countryDB)

	return countryDB
}

//Cleans URLs by removing the unecessary prefixes,
func removeUrlExcess(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")
	updatedURL := ""
	for _, part := range parts {
		switch part {
		//Simply checks if one segment of the sliced string is a
		//known constant or not
		case UNI_BASE_PATH_A:
		case UNI_BASE_PATH_B:
		case UNI_INFO_PATH:
		case UNI_NEIGH_PATH:
		case UNI_DIAG_PATH:
		case "":
		default:
			updatedURL += part
			if 0 < len(updatedURL) {
				//Ensures that there isnt two slashes at the end
				if updatedURL[len(updatedURL)-1] != '/' {
					updatedURL += "/"
				}
			}
		}
	}
	return updatedURL
}
