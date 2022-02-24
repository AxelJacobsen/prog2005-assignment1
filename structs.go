package cloudAss1

// UniversityStorage is how we can access University data.
/* type UniversityStorage interface {
	Add(u University) error
	Count() int
	Get(key string) (University, bool)
	GetAll() []University
}
*/
/*
University is the main data structure for the API.
It follows this format:
{
	"name": 	"Norwegian University of Science and Technology",
    "country": 	"Norway",
    "web_pages":
    Countryholder Country
}
*/
type University struct {
	Name         string   `json:"name"`
	Country      string   `json:"country"`
	Webpages     []string `json:"web_pages"`
	Coutryholder Country
}

/*
Country is the secondary data structure, an example of the structure is:
{
    "name": 	"Norway",
    "cca3": 	"NOR",
    "languages":{"nno": "Norwegian Nynorsk",
                "nob": "Norwegian Bokmål",
                "smi": "Sami"},
	"borders"
    "maps": "https://www.openstreetmap.org/relation/2978650"
}
*/
type Country struct {
	CountryName CountCommonName   `json:"name"`
	Isocode     string            `json:"cca3"`
	Languages   map[string]string `json:"languages"`
	Bordering   []string          `json:"borders"`
	Map         MapOpenStreetMap  `json:"maps"`
}

/*
Stores diagnostics data with the structure:
{
    "universitiesapi": 	"<http status code for universities API>",
    "countriesapi": 	"<http status code for restcountries API>",
    "version":			"v1",
	"uptime":			"<time in seconds from the last service restart>"
}
*/
type Diagnostics struct {
	UniversitiesApi int    `json:"universitiesapi"`
	CountrisApi     int    `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          int64  `json:"uptime"`
}

//these are simply to acces sub folder of data, since the json contains
//Official and common name, for most nations the universities APi uses common name
type CountCommonName struct {
	CommonName string `json:"common"`
}

//Similar to the above
type MapOpenStreetMap struct {
	OpenStreetMap string `json:"openStreetMaps"`
}

/*
Initializes standard data in diagnostics struct
*/
func PrePopulateDiagnostics() *Diagnostics {
	db := Diagnostics{Version: CURRENT_VER}
	return &db
}
