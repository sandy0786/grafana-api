package errors

type ApiError struct {
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

var ErrorMap = map[string]string{
	"InvalidQueryParam":           "Error : Invalid query param: ",
	"InvalidQueryParamOccurences": "Error : Query param is allowed only once: ",
	"InvalidColumn":               "Error : Invalid column: ",
	"InvalidTimestamp":            "Error : Invalid timestamp: ",
	"InvalidFilter":               "Error : Invalid filter: ",
	"MissingMandatoryQP":          "Error : Mandatory QP missing: ",
}
