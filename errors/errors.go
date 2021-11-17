package errors

// Common error
// swagger:response CommonError
type ApiError struct {
	// out:string
	// Detailed error message
	Message string `json:"message"`
}

// Validation error
// swagger:response ValidationError
type ValidationError struct {
	// out:string
	// Detailed validation error message
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
