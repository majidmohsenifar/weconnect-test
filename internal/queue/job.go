package queue

type Job struct {
	LineNumber      int
	SeriesReference string
	Period          string
	DataValue       string
	Suppressed      string
	Status          string
	Units           string
	Magnitude       string
	Subject         string
	Group           string
	SeriesTitle1    string
	SeriesTitle2    string
	SeriesTitle3    string
	SeriesTitle4    string
	SeriesTitle5    string
}

func newJob() Job {
	return Job{}
}
