// Concurrency Task Worker.
// Studi kasus: data museum di Indonesia.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	URLroot  string = "http://jendela.data.kemdikbud.go.id/api/index.php/"
	PropPath string = "CWilayah/wilayahGET"
	KotaPath string = "CWilayah/wilayahGET?mst_kode_wilayah="
)

type Wilayah struct {
	Data []struct {
		Kode       string `json:"kode_wilayah"`
		Nama       string `json:"nama"`
		KodeMaster string `json:"mst_kode_wilayah"`
	} `json:"data"`
}

var wilayah2 []Wilayah

func main() {
	// Ambil data provinsi.
	provinsi := getWilayah(getJSON(URLroot + PropPath))
	fmt.Println(provinsi)
}

func getJSON(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	jsonString, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", jsonString)
	return fmt.Sprintf("%s", jsonString)
}

func getWilayah(jsonString string) Wilayah {
	var umsWilayah Wilayah

	err := json.Unmarshal([]byte(jsonString), &umsWilayah)
	if err != nil {
		fmt.Println("Error unmarshal:", err)
	}

	return umsWilayah
}
