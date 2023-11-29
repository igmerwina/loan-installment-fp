package config

const (
	Success             = "00"
	InternalServerError = "X5"
	Failed              = "68"
	DataNotFound        = "12"

	MaxLimit = 10000000
)

var listResponse = map[string]string{
	Success:             "Success",
	InternalServerError: "Internal Server Error",
	Failed:              "Failed",
}

func MessageResponse(responseCode string) string {
	return listResponse[responseCode]
}
