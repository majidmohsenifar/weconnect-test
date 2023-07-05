package queue

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"we-connect-test/internal/financial"

	"go.uber.org/zap"
)

type Manager struct {
	jobCollector     chan *Job
	errCollector     chan workerErr
	workers          []*worker
	financialService *financial.Service
	logger           *zap.Logger
	quit             chan bool
}

type workerErr struct {
	Err        error
	LineNumber int
}

func (m *Manager) Run(ctx context.Context, filePath string, workerCount int) error {
	for i := 0; i < workerCount; i++ {
		worker := newWorker(m.jobCollector, m.errCollector, i, m.financialService)
		m.workers = append(m.workers, worker)
		go worker.start(ctx)
	}
	go m.collectErrors(ctx)
	return m.startDispatcher(ctx, filePath)
}

func (m *Manager) collectErrors(ctx context.Context) {
	for workErr := range m.errCollector {
		m.logger.Error("worker error",
			zap.Error(workErr.Err),
			zap.Int("lineNumber", workErr.LineNumber),
		)
	}
}

func (m *Manager) startDispatcher(ctx context.Context, filePath string) error {
	if filePath == "" {
		return fmt.Errorf("filePath is empty")
	}
	csvfile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	errChan := make(chan error)
	go func(chan error) {
		i := 0
		for {
			record, err := reader.Read()
			if err == io.EOF {
				err = nil
				break
			} else if err != nil {
				break
			}
			i++
			if i == 1 {
				continue
			}
			m.jobCollector <- &Job{
				LineNumber:      i,
				SeriesReference: record[0],
				Period:          record[1],
				DataValue:       record[2],
				Suppressed:      record[3],
				Status:          record[4],
				Units:           record[5],
				Magnitude:       record[6],
				Subject:         record[7],
				Group:           record[8],
				SeriesTitle1:    record[9],
				SeriesTitle2:    record[10],
				SeriesTitle3:    record[11],
				SeriesTitle4:    record[12],
				SeriesTitle5:    record[13],
			}
		}
		close(m.jobCollector) // close chan to signal workers that no more job are incoming.
		errChan <- err
	}(errChan)
	err = <-errChan
	return err
}

func NewManager(financialService *financial.Service, logger *zap.Logger) *Manager {
	collector := make(chan *Job, 1000)
	errChan := make(chan workerErr, 1000)
	return &Manager{
		jobCollector:     collector,
		errCollector:     errChan,
		financialService: financialService,
		logger:           logger,
		quit:             make(chan bool),
	}
}
