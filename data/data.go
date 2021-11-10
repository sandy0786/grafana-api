package data

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

type DataBase struct {
	Time     int64   `json:"time"`
	Metric   string  `json:"metric"`
	Value    float64 `json:"value"`
	Hostname string  `json:"hostname"`
}

var Data []DataBase

func SetData() {
	log.Println("inside SetData")

	// loop through data and set sample data
	totalRange := 360 // minutes
	dateTime := getDates(int64(totalRange))
	for i := 0; i < totalRange; i++ {
		d := DataBase{
			Time:     dateTime[i],
			Metric:   "metric" + strconv.Itoa(rand.Intn(10)),
			Value:    rand.Float64(),
			Hostname: "host" + strconv.Itoa(rand.Intn(10)),
		}
		Data = append(Data, d)
	}
}

func getDates(totalRange int64) (dateTime []int64) {
	log.Println("inside getDates")
	now := time.Now().UTC()
	startTime := now.Add(time.Duration(-6) * time.Hour)

	// process time
	start := startTime.Unix() * 1000

	var i int64
	for i = 1; i <= totalRange; i++ {
		dateTime = append(dateTime, start+(i*1000*60))
	}

	return
}
