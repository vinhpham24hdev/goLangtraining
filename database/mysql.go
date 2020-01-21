package database

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	dat "goClean/config"
	str "database"
)

source := "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"
strQuerry := "INSERT INTO goApi.Envelope VALUES "

func main() {

	response, err := http.Get(source)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	bytes, _ := ioutil.ReadAll(response.Body)
	string_body := string(bytes)
	// fmt.Println(string_body)
	var result str.Envelope
	xml.Unmarshal([]byte(string_body), &result)

	// connectdatabase
	db, err := sql.Open(dat.DBMS, dat.User+":"+dat.Password+"@/"+ dat.DbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// -------- LƯU DỮ LIỆU VÀO DATABASE -------------

	
	for i := 0; i < len(result.Cube.Cube); i++ {
		for j := 0; j < len(result.Cube.Cube[i].Cube); j++ {
			strQuerry += fmt.Sprintf("('%s','%s',%s),", result.Cube.Cube[i].Time,
				result.Cube.Cube[i].Cube[j].Currency, result.Cube.Cube[i].Cube[j].Rate)
			if err != nil {
				panic(err.Error())
			}
		}

	}
	strQuerry = strQuerry[0 : len(strQuerry)-1]
	db.Query(strQuerry)

	//--------------------------------------------------

}
