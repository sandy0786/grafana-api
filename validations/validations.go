package validations

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	constants "github.com/sandy0786/grafana-api/constants"
	errors "github.com/sandy0786/grafana-api/errors"
	utils "github.com/sandy0786/grafana-api/utils"
)

func ValidateRequest(queryParams url.Values) (bool, errors.ApiError) {
	log.Println("inside ValidateRequest")

	// validate mandatory queryarams
	for _, qp := range constants.MANDATORY_QUERYPARAMS {
		_, found := queryParams[qp]
		if !found {
			return false, errors.ApiError{
				Message: errors.ErrorMap["MissingMandatoryQP"] + strings.Join(constants.MANDATORY_QUERYPARAMS, ","),
			}
		}
	}

	// validate queryarams
	for key, _ := range queryParams {
		if !utils.FindItemInList(constants.ALLOWED_QUERYPARAMS, key) {
			return false, errors.ApiError{
				Message: errors.ErrorMap["InvalidQueryParam"] + key,
			}
		}
	}

	// validate occurences
	for key, val := range queryParams {
		if len(val) > 1 {
			return false, errors.ApiError{
				Message: errors.ErrorMap["InvalidQueryParamOccurences"] + key,
			}
		}
	}

	// validate data
	for key, val := range queryParams {
		// validate columns
		switch key {
		case constants.QP_COLUMNS:
			splits := strings.Split(val[0], ",")
			for _, split := range splits {
				if !utils.FindItemInList(constants.SUGGEST_CONSTANTS, split) {
					return false, errors.ApiError{
						Message: errors.ErrorMap["InvalidColumn"] + split,
					}
				}
			}

		case constants.QP_START, constants.QP_END:
			matched, _ := regexp.MatchString(constants.START_END_REGEX, val[0])
			if !matched {
				return false, errors.ApiError{
					Message: errors.ErrorMap["InvalidTimestamp"] + key,
				}
			}

		case constants.QP_FILTER:
			matched, _ := regexp.MatchString(constants.FILTER_REGEX2, val[0])
			if !matched {
				return false, errors.ApiError{
					Message: errors.ErrorMap["InvalidFilter"] + val[0],
				}
			}

		}

	}

	return true, errors.ApiError{}
}
