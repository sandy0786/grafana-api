package entities

// HTTP status code 200 and message
// swagger:response APIStatus
type Status struct {
	// HTTP status code 200 - Status OK
	// out: string
	Status string `json:"status"`
	// Message whether working or not - Datasource is working
	// out: string
	Message string `json:"message"`
}

// List of aggregate functions
// swagger:response AggregationList
type AggregationList struct {
	// list of aggregate functions
	// out: []string
	Aggregate []string `json:"aggregate"`
}

// List of filters
// swagger:response FiltersList
type FilterList struct {
	// list of aggregate functions
	// out: []string
	Filters []string `json:"filters"`
}

// List of distinct column values
// swagger:response SuggestList
type SuggestList struct {
	// list of column values
	// out: []string
	Suggest []string `json:"suggest"`
}

// List of rows
// swagger:response DataList
type DataList struct {
	// list of column values
	// out:body
	Body struct {
		// time in epoch milliseconds
		Time int64 `json:"time"`
		// metric
		Metric string `json:"metric"`
		// hostname
		Hostname string `json:"hostname"`
		// value
		Value float64 `json:"value"`
	}
}

// List of rows
// swagger:response DataList2
type DataList2 struct {
	// list of column values
	// out:body
	Body struct {
		// time in epoch milliseconds
		Time int64 `json:"time"`
		// metric
		Metric string `json:"metric"`
		// hostname
		Hostname string `json:"hostname"`
		// value
		Value float64 `json:"value"`
	}
}
