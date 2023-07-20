package handler

type TaskIngestURI struct {
	DB string `uri:"db" binding:"required"`
}

type TaskIngestQuery struct {
	IgnoreLastIngest bool `form:"ignoreLastIngest"`
}
