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
	Data struct {
		Kode, Nama, KodeMaster string
	}
}

var wilayah2 []Wilayah

func main() {
	// Ambil data provinsi.
	provinsi := addWilayah(getJSON(URLroot + PropPath))
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
	return fmt.Sprintf("%s", jsonString)
}

func addWilayah(jsonString string) []Wilayah {
	var umsWilayah2 []Wilayah

	err := json.Unmarshal([]byte(jsonString), &umsWilayah2)
	if err != nil {
		fmt.Println("Error unmarshal:", err)
	}

	return umsWilayah2
}
