package domain

type ClientResponse struct {
	Result     string
	StatusCode int
	Error      string
}

type IngestResponse struct {
	StatusCode int
	Error      string
}

type GetLastIngestResponse struct {
	Result     string
	StatusCode int
	Error      string
}
