package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Result struct {
	TrainDetailInfo []TrainDetailInfo
}
type TrainDetailInfo struct {
	TrainStation string
	ArriveTime   string
	StartTime    string
	KM           string
}

func main() {
	content, err := ioutil.ReadFile("../static/file/test.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result Result
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
