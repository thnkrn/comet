package domain

import "errors"

type Status string

const (
	Failed    Status = "Failed"
	Ingesting Status = "Ingesting"
	Success   Status = "Success"
	Initiated Status = "Initiated"
)

type Job struct {
	DBName            string
	Source            string
	ApiToken          string
	IngestOnlyLastDir bool
	Status            Status
}

func NewJob(DBName, Source, ApiToken string, IngestOnlyLastDir bool) Job {
	return Job{
		DBName:            DBName,
		Source:            Source,
		ApiToken:          ApiToken,
		IngestOnlyLastDir: IngestOnlyLastDir,
		Status:            Initiated,
	}
}

func (j Job) SetSuccess() {
	j.Status = Success
}

func (j Job) SetFailed() {
	j.Status = Failed
}

func (j Job) SetIngesting() {
	j.Status = Ingesting
}

func (j Job) ShouldIngest() bool {
	return j.Status == Initiated || j.Status == Failed
}

type JobPools map[string]Job

func (jp JobPools) Perform(db string) (Job, error) {

	if job, exist := jp[db]; exist {
		return job, nil
	}

	return Job{}, errors.New("no job match")
}
