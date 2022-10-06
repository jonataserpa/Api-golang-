package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type User struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	Avatar_ul string `json:"avatar_ul"`
	Url       string `json:"url"`
}

type Repos struct {
	ID          int    `json:"id"`
	Node_id     string `json:"login"`
	Name        string `json:"avatar_ul"`
	Full_name   string `json:"full_name"`
	Html_url    string `json:"html_url"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

/*
* Load .env file
 */
func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

/*
* Execution main
 */
func main() {

	rota := mux.NewRouter()

	rota.HandleFunc("/user", getUsers).Methods("GET")
	rota.HandleFunc("/repos/{username}", getRepositorys).Methods("GET")
	rota.HandleFunc("/repos/{username}/{type}/{per_page}/{page}/{sort}/{direction}", getRepositorys).Methods("GET")
	rota.HandleFunc("/user/{username}", getUserDinamic).Methods("GET")

	port := goDotEnvVariable("PORT")
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":8089", rota))
}

/*
* Get user dinamic
* @param username
 */
func getUserDinamic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	apiUserDinamic := goDotEnvVariable("API_URL_GITHUB_USER_DINAMIC") + username

	response, err := http.Get(apiUserDinamic)

	if err != nil {
		fmt.Println("Error getting user dinamic: ", err)
		return
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	var data User
	json.Unmarshal(responseData, &data)
	fmt.Printf("Results: %v\n", data)
	json.NewEncoder(w).Encode(data)
}

/*
* Get user default
 */
func getUsers(w http.ResponseWriter, r *http.Request) {

	api := goDotEnvVariable("API_URL_GITHUB_USER_DEFAULT")

	response, err := http.Get(api)

	if err != nil {
		fmt.Println("Error getting user")
		return
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	var data User
	json.Unmarshal(responseData, &data)
	fmt.Printf("Results: %v\n", data)

	json.NewEncoder(w).Encode(data)
}

/*
* Get repos dinamic
* @param username
 */
func getRepositorys(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]
	typeQuery := vars["type"]
	per_page := vars["per_page"]
	page := vars["page"]
	sort := vars["sort"]
	direction := vars["direction"]

	apiUserDinamic := goDotEnvVariable("API_URL_GITHUB_USER_DINAMIC") + username + "/repos?type=" + typeQuery + "&per_page=" + per_page + "&page=" + page + "&sort=" + sort + "&direction=" + direction
	response, err := http.Get(apiUserDinamic)

	if err != nil {
		fmt.Println("Error getting repository user dinamic: ", err)
		return
	}

	responseData, _ := ioutil.ReadAll(response.Body)
	var x []Repos
	json.Unmarshal([]byte(responseData), &x)
	json.NewEncoder(w).Encode(x)
}
