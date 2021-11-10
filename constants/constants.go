package constants

const (
	TIME     = "time"
	METRIC   = "metric"
	HOSTNAME = "hostname"
	VALUE    = "value"
)

const (
	STATUS_PATH       = "/"
	SUGGEST_PATH      = "/suggest"
	AGGREGATIONS_PATH = "/aggregations"
	QUERY_PATH        = "/query"
)

var SUGGEST_CONSTANTS = []string{TIME, METRIC, HOSTNAME, VALUE}
var AGGREGATE_FUNCTIONS = []string{"avg", "sum"}

const FILTER_REGEX1 = "^{([a-z0-9A-Z,=]+)}$"
const FILTER_REGEX2 = "^{([a-z]+)=([a-z0-9]+)}$"

const QP_COLUMNS = "columns"
const QP_FILTER = "filter"
const QP_START = "start"
const QP_END = "end"

var ALLOWED_QUERYPARAMS = []string{QP_COLUMNS, QP_FILTER, QP_START, QP_END}
var MANDATORY_QUERYPARAMS = []string{QP_COLUMNS, QP_START, QP_END}

const START_END_REGEX = "^[0-9]{13}$"
