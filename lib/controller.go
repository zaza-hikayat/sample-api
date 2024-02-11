package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

type controller struct {
	app *App
}

var CollectionPerson = []Person{}

func NewController(app *App) *controller {
	return &controller{app}
}

// soal 1: print data
func (c *controller) PrintData(w http.ResponseWriter, r *http.Request) {

	// receive query param total of sequence
	sequenceStr := r.URL.Query().Get("total")

	// validate input only number
	totalSeq, err := strconv.Atoi(sequenceStr)
	if err != nil {
		ResponseJson(w, http.StatusBadRequest, map[string]interface{}{"message": "query param `total` only number"})
		return
	}

	// looping data sequences
	var seqAggregate []string
	for i := 1; i <= totalSeq; i++ {
		seq := ""
		if i%2 == 0 {
			seq = GetReverseSequence(1, totalSeq)
		} else {
			seq = GetSequence(1, totalSeq)
		}
		fmt.Println(seq)
		seqAggregate = append(seqAggregate, seq)
	}

	ResponseJson(w, http.StatusOK, map[string]interface{}{"data": seqAggregate})
}

// get data start number up to total number
func GetSequence(startNumber, endNumber int) (res string) {
	for i := startNumber; i <= endNumber; i++ {
		res += strconv.Itoa(i)
	}
	return res
}

// get data reverse from endNumber to startNumber
func GetReverseSequence(startNumber, endNumber int) (res string) {
	for i := endNumber; i >= startNumber; i-- {
		res += strconv.Itoa(i)
	}
	return res
}

// helper return response json
func ResponseJson(w http.ResponseWriter, httpStatusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp, _ := json.Marshal(data)
	w.Write(resp)
}

// soal 2
func (c *controller) InsertData(w http.ResponseWriter, r *http.Request) {
	var req Person
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ResponseJson(w, http.StatusBadRequest, map[string]interface{}{"message": "invalid format request"})
		return
	}
	// insert data to data store
	CollectionPerson = append(CollectionPerson, req)

	// show data return json
	ResponseJson(w, http.StatusOK, ParticipantResponse{
		Data:    CollectionPerson,
		Success: true,
	})
}

// soal 4
func (c *controller) ContactChip(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "contact_chips.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// soal 8
func (c *controller) PrintFizzbuzz(w http.ResponseWriter, r *http.Request) {
	startNum := 1
	endNum := 100
	var results []string

	for i := startNum; startNum <= endNum; i++ {
		s := ""
		switch true {
		case i%3 == 0 && i%5 == 0:
			s = "FizzBuzz"
		case i%3 == 0:
			s = "Fizz"
		case i%5 == 0:
			s = "Buzz"
		}

		if s != "" {
			fmt.Println(s)
			results = append(results, s)
		}
	}

	// show data return json
	ResponseJson(w, http.StatusOK, map[string]interface{}{"data": results})
}

func (c *controller) DataMahasiswa(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}
	query := `SELECT * FROM "Mahasiswa"`
	c.app.DB.Raw(query).Scan(&results)
	// show data return json
	ResponseJson(w, http.StatusOK, GeneralResponse{
		Data:    results,
		Success: true,
	})
}

func (c *controller) DataDosen(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}
	query := `SELECT * FROM "Dosen"`
	c.app.DB.Raw(query).Scan(&results)
	// show data return json
	ResponseJson(w, http.StatusOK, GeneralResponse{
		Data:    results,
		Success: true,
	})
}

func (c *controller) DataMataKuliah(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}
	query := `SELECT * FROM "MataKuliah"`
	c.app.DB.Raw(query).Scan(&results)
	// show data return json
	ResponseJson(w, http.StatusOK, GeneralResponse{
		Data:    results,
		Success: true,
	})
}

func (c *controller) DataNilai(w http.ResponseWriter, r *http.Request) {
	var results []map[string]interface{}
	query := `SELECT * FROM "Nilai"`
	c.app.DB.Raw(query).Scan(&results)
	// show data return json
	ResponseJson(w, http.StatusOK, GeneralResponse{
		Data:    results,
		Success: true,
	})
}
