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
	var filter string
	if queryParams["filter"] != nil {
		filter = queryParams["filter"][0]
	}

	var start int64
	start, err := strconv.ParseInt(queryParams["start"][0], 10, 0)
	if err != nil {
		log.Println("error >> ", err)
	}

	var end int64
	end, err = strconv.ParseInt(queryParams["end"][0], 10, 0)

	var filterKey string
	var filterVal string

	var dataMap []map[string]interface{}
	inrec, _ := json.Marshal(data.Data)
	json.Unmarshal(inrec, &dataMap)

	// extract filters
	if len(filter) > 0 {
		re, _ := regexp.Compile(constants.FILTER_REGEX1)
		match := re.FindStringSubmatch(filter)

		re, _ = regexp.Compile(constants.FILTER_REGEX2)
		match = re.FindStringSubmatch(filter)
		filterKey = match[1]
		filterVal = match[2]
	}

	for _, dMap := range dataMap {
		if dMap[constants.TIME].(float64) > float64(start) && float64(end) < dMap[constants.TIME].(float64) {
			resp := make(map[string]interface{})
			splits := strings.Split(columns, ",")
			if len(filterKey) > 0 {
				if strings.Contains(fmt.Sprintf("%v", dMap[filterKey]), filterVal) {
					for _, col := range splits {
						resp[col] = dMap[col]
					}
					response = append(response, resp)
				}
			} else {
				for _, col := range splits {
					resp[col] = dMap[col]
				}
				response = append(response, resp)
			}
		}
	}
	return response
}
