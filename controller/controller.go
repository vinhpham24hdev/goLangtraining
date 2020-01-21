package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	dat "goClean/config"
	"goClean/models"
) 


func GetRates(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dat.DBMS, dat.User+":"+dat.Password+"@/"+ dat.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	
	rateNew := make(map[string]string)
	currency := ""
	rate := ""
	strQuerry := dat.QueryRateNew
	rows, _ := db.Query(strQuerry)
	for rows.Next() {
		rows.Scan(&currency, &rate)
		rateNew[currency] = rate
	}

	var ratesNew models.Rate
	ratesNew.Rate = rateNew
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratesNew)
}

func GetRate(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dat.DBMS, dat.User+":"+dat.Password+"@/"+ dat.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rateByday := make(map[string]string)
	param := mux.Vars(r)
	currency := ""
	rate := ""
	strQuerry := fmt.Sprintf(dat.QueryRateByDay, param["day"])
	rows, _ := db.Query(strQuerry)
	for rows.Next() {
		rows.Scan(&currency, &rate)
		rateByday[currency] = rate
	}

	var ratesByDay models.RateByDate
	ratesByDay.Time = param["day"]
	ratesByDay.Rate = rateByday
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratesByDay)

}

func GetAvgRate(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dat.DBMS, dat.User+":"+dat.Password+"@/"+ dat.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rateAvg := make(map[string]models.AvgRate)
	currency := ""
	max := ""
	min := ""
	avg := ""
	strQuerry := dat.QueryMaxMin
	rows, _ := db.Query(strQuerry)
	for rows.Next() {
		rows.Scan(&currency, &max, &min, &avg)
		rateAvg[currency] = models.AvgRate{max, min, avg}
	}
	var ratesAvg models.Envelope
	ratesAvg.Name = rateAvg

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratesAvg)
}
