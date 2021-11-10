package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	constants "github.com/sandy0786/grafana-api/constants"
	data "github.com/sandy0786/grafana-api/data"
)

func GetData(queryParams url.Values) []map[string]interface{} {
	log.Println("inside getData ")
	var response []map[string]interface{}
	columns := queryParams["columns"][0]
	filter := queryParams["filter"][0]
	var start int64
	start, err := strconv.ParseInt(queryParams["start"][0], 10, 0)
	if err != nil {
		log.Println("error >> ", err)
	}

	var end int64
	end, err = strconv.ParseInt(queryParams["end"][0], 10, 0)

	var filterKey string
	var filterVal string

	// extract filters
	if len(filter) > 0 {
		re, _ := regexp.Compile(constants.FILTER_REGEX1)
		match := re.FindStringSubmatch(filter)

		re, _ = regexp.Compile(constants.FILTER_REGEX2)
		match = re.FindStringSubmatch(filter)
		filterKey = match[1]
		filterVal = match[2]
	}

	var dataMap []map[string]interface{}
	inrec, _ := json.Marshal(data.Data)
	json.Unmarshal(inrec, &dataMap)

	for _, dMap := range dataMap {
		if dMap[constants.TIME].(float64) > float64(start) && dMap[constants.TIME].(float64) < float64(end) {
			resp := make(map[string]interface{})
			splits := strings.Split(columns, ",")
			for _, col := range splits {
				if len(filterKey) > 0 {
					if filterKey == col {
						if strings.Contains(fmt.Sprintf("%v", dMap[col]), filterVal) {
							resp[col] = dMap[col]
						}
					} else {
						resp[col] = dMap[col]
					}
				} else {
					resp[col] = dMap[col]
				}

			}
			if resp[filterKey] != nil {
				response = append(response, resp)
			}
		} else {
			// log.Println("time else >>")
		}

	}
	return response
}
