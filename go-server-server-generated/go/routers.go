/*
 * Prometheus HTTP API
 *
 * The current stable HTTP API is reachable under /api/v1 on a Prometheus server. Any non-breaking additions will be added under that endpoint.  # Format overview The API response format is JSON. Every successful API request returns a ```2xx``` status code.  Invalid requests that reach the API handlers return a JSON error object and one of the following HTTP response codes:  ```400 Bad Request``` when parameters are missing or incorrect. ```422 Unprocessable Entity``` when an expression can't be executed ([RFC4918](https://datatracker.ietf.org/doc/html/rfc4918#page-78)). ```503 Service Unavailable``` when queries time out or abort.  Other non-```2xx``` codes may be returned for errors occurring before the API endpoint is reached.  An array of warnings may be returned if there are errors that do not inhibit the request execution. All of the data that was successfully collected will be returned in the data field.  The JSON response envelope format is as follows:  ``` {   \"status\": \"success\" | \"error\",   \"data\": <data>,    // Only set if status is \"error\". The data field may still hold   // additional data.   \"errorType\": \"<string>\",   \"error\": \"<string>\",    // Only if there were warnings while executing the request.   // There will still be data in the data field.   \"warnings\": [\"<string>\"] } ``` # Generic placeholders:  ```<rfc3339 | unix_timestamp>```: Input timestamps may be provided either in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format or as a Unix timestamp in seconds, with optional decimal places for sub-second precision. Output timestamps are always represented as Unix timestamps in seconds.  ```<series_selector>```: Prometheus [time series selectors](https://prometheus.io/docs/prometheus/latest/querying/basics/#time-series-selectors) like ```http_requests_total``` or ```http_requests_total{method=~\"(GET|POST)\"}``` and need to be URL-encoded.  ```<duration>```: [Prometheus duration strings](https://prometheus.io/docs/prometheus/latest/querying/basics/#time_durations). For example, ```5m``` refers to a duration of 5 minutes.  ```<bool>```: boolean values (strings ```true``` and ```false```).  **Note**: Names of query parameters that may be repeated end with ```[]```. 
 *
 * API version: v2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AlertManagersGET",
		strings.ToUpper("Get"),
		"/alertmanagers",
		AlertManagersGET,
	},

	Route{
		"AlertsGET",
		strings.ToUpper("Get"),
		"/alerts",
		AlertsGET,
	},

	Route{
		"MetricMetadataGET",
		strings.ToUpper("Get"),
		"/metadata",
		MetricMetadataGET,
	},

	Route{
		"RulesGET",
		strings.ToUpper("Get"),
		"/rules",
		RulesGET,
	},

	Route{
		"TargetMetadataGET",
		strings.ToUpper("Get"),
		"/targets/metadata",
		TargetMetadataGET,
	},

	Route{
		"TargetsGET",
		strings.ToUpper("Get"),
		"/targets",
		TargetsGET,
	},

	Route{
		"QueryExemplarsGET",
		strings.ToUpper("Get"),
		"/query_exemplars",
		QueryExemplarsGET,
	},

	Route{
		"QueryExemplarsPOST",
		strings.ToUpper("Post"),
		"/query_exemplars",
		QueryExemplarsPOST,
	},

	Route{
		"QueryGET",
		strings.ToUpper("Get"),
		"/query",
		QueryGET,
	},

	Route{
		"QueryPOST",
		strings.ToUpper("Post"),
		"/query",
		QueryPOST,
	},

	Route{
		"QueryRangeGET",
		strings.ToUpper("Get"),
		"/query_range",
		QueryRangeGET,
	},

	Route{
		"QueryRangePOST",
		strings.ToUpper("Post"),
		"/query_range",
		QueryRangePOST,
	},

	Route{
		"LabelNamesGET",
		strings.ToUpper("Get"),
		"/labels",
		LabelNamesGET,
	},

	Route{
		"LabelNamesPOST",
		strings.ToUpper("Post"),
		"/labels",
		LabelNamesPOST,
	},

	Route{
		"LabelValuesGET",
		strings.ToUpper("Get"),
		"/label/{label_name}/values",
		LabelValuesGET,
	},

	Route{
		"SeriesDELETE",
		strings.ToUpper("Delete"),
		"/series",
		SeriesDELETE,
	},

	Route{
		"SeriesGET",
		strings.ToUpper("Get"),
		"/series",
		SeriesGET,
	},

	Route{
		"SeriesPOST",
		strings.ToUpper("Post"),
		"/series",
		SeriesPOST,
	},

	Route{
		"ServeBuildInfo",
		strings.ToUpper("Get"),
		"/status/buildinfo",
		ServeBuildInfo,
	},

	Route{
		"ServeConfig",
		strings.ToUpper("Get"),
		"/status/config",
		ServeConfig,
	},

	Route{
		"ServeFlags",
		strings.ToUpper("Get"),
		"/status/flags",
		ServeFlags,
	},

	Route{
		"ServeRuntimeInfo",
		strings.ToUpper("Get"),
		"/status/runtimeinfo",
		ServeRuntimeInfo,
	},

	Route{
		"ServeTSDBStatus",
		strings.ToUpper("Get"),
		"/status/tsdb",
		ServeTSDBStatus,
	},

	Route{
		"ServeWALReplayStatus",
		strings.ToUpper("Get"),
		"/status/walreplay",
		ServeWALReplayStatus,
	},

	Route{
		"CleanTombstonesPOST",
		strings.ToUpper("Post"),
		"/admin/tsdb/clean_tombstones",
		CleanTombstonesPOST,
	},

	Route{
		"CleanTombstonesPUT",
		strings.ToUpper("Put"),
		"/admin/tsdb/clean_tombstones",
		CleanTombstonesPUT,
	},

	Route{
		"DeleteSeriesPOST",
		strings.ToUpper("Post"),
		"/admin/tsdb/delete_series",
		DeleteSeriesPOST,
	},

	Route{
		"DeleteSeriesPUT",
		strings.ToUpper("Put"),
		"/admin/tsdb/delete_series",
		DeleteSeriesPUT,
	},

	Route{
		"SnapshotPOST",
		strings.ToUpper("Post"),
		"/admin/tsdb/snapshot",
		SnapshotPOST,
	},

	Route{
		"SnapshotPUT",
		strings.ToUpper("Put"),
		"/admin/tsdb/snapshot",
		SnapshotPUT,
	},
}
