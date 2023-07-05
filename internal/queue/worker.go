package queue

import (
	"context"
	"we-connect-test/internal/financial"
)

type worker struct {
	ID               int
	jobChan          chan *Job
	errChan          chan workerErr
	financialService *financial.Service
}

func (w *worker) start(ctx context.Context) {
	for job := range w.jobChan {
		_, err := w.financialService.CreateFinancialData(ctx, financial.FinancialModel{
			SeriesReference: job.SeriesReference,
			Period:          job.Period,
			DataValue:       job.DataValue,
			Suppressed:      job.Suppressed,
			Status:          job.Status,
			Units:           job.Units,
			Magnitude:       job.Magnitude,
			Subject:         job.Subject,
			Group:           job.Group,
			SeriesTitle1:    job.SeriesTitle1,
			SeriesTitle2:    job.SeriesTitle2,
			SeriesTitle3:    job.SeriesTitle3,
			SeriesTitle4:    job.SeriesTitle4,
			SeriesTitle5:    job.SeriesTitle5,
		})
		if err != nil {
			w.errChan <- workerErr{
				Err:        err,
				LineNumber: job.LineNumber,
			}
		}
	}
}

func newWorker(jobChan chan *Job, errChan chan workerErr, workerID int, financialService *financial.Service) *worker {
	return &worker{
		ID:               workerID,
		jobChan:          jobChan,
		errChan:          errChan,
		financialService: financialService,
	}
}
