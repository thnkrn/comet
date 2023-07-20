package handler

type DBListReponse struct {
	DBName   string `json:"dbName"`
	OpenMode string `json:"openMode"`
}
