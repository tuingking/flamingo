package product

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/logger"
)

type Service interface {
	GetAllProducts(ctx context.Context) ([]Product, error)

	GenerateProducts(ctx context.Context, n int64) []Product

	// GenerateProductsCsv create random product and save it as csv file
	GenerateProductsCsv(ctx context.Context, n int64) (*os.File, error)

	// BulkCreate insert data from the csv to DB using goroutine
	BulkCreate(ctx context.Context, filename string) error
}

type service struct {
	config  ConfigSvc
	logger  logger.Logger
	product Repository
}

type ConfigSvc struct {
	Worker int
}

func NewService(
	config ConfigSvc,
	logger logger.Logger,
	product Repository,
) Service {
	return &service{
		config:  config,
		logger:  logger,
		product: product,
	}
}

func (s *service) GetAllProducts(ctx context.Context) ([]Product, error) {
	return s.product.FindAll(ctx)
}

func (s *service) GenerateProducts(ctx context.Context, n int64) []Product {
	var products []Product

	for i := int64(0); i < n; i++ {
		// g := namegenerator.NewNameGenerator(time.Now().UTC().UnixNano())

		product := Product{
			Name: fmt.Sprintf("Product-%d", i+1),
			// Name:  g.Generate(),
			Price: float64(1000 + (i+1)*100),
		}

		products = append(products, product)
	}

	return products
}

func (s *service) GenerateProductsCsv(ctx context.Context, n int64) (*os.File, error) {
	now := time.Now()
	defer func() {
		s.logger.Info("GenerateProductsCsv took ", time.Since(now))
	}()
	products := s.GenerateProducts(ctx, n)

	csvfile, err := os.Create("tmp/products.csv")
	if err != nil {
		s.logger.Error(errors.Wrap(err, "create csv file"))
		return nil, errors.Wrap(err, "create csv file")
	}
	defer csvfile.Close()

	csvwritter := csv.NewWriter(csvfile)
	for i, product := range products {
		_ = csvwritter.Write([]string{fmt.Sprintf("%d", i), product.Name, fmt.Sprintf("%f", product.Price)})
	}
	csvwritter.Flush()

	return csvfile, nil
}

// BulkCreate read csv file then insert to BD using worker pool
func (s *service) BulkCreate(ctx context.Context, filename string) error {
	var (
		totalJobDispatched int
		totalJobFinished   int

		// define channel
		jobs   = make(chan Product)
		output = make(chan string)
		wg     = sync.WaitGroup{}
		lock   = sync.Mutex{}
		worker = s.config.Worker
	)

	now := time.Now()
	defer func() {
		s.logger.Info("GenerateProductsCsv took ", time.Since(now))

		s.logger.Info("Total Worker: ", worker)
		s.logger.Info("Total Job: ", totalJobDispatched)
		lock.Lock()
		s.logger.Info("Total Done: ", totalJobFinished)
		if totalJobDispatched == totalJobFinished {
			s.logger.Info("Job Succesfully Done")
		}
		lock.Unlock()
	}()

	// open file
	f, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "open file")
	}
	defer f.Close()

	// spawn worker
	for i := 0; i < worker; i++ {
		go s.workerDispatcher(ctx, i, jobs, output)
	}

	// listen to channel output
	go func() {
		for res := range output {
			wg.Done()
			s.logger.Info(res)
			lock.Lock()
			totalJobFinished++
			lock.Unlock()
		}
	}()

	// read file
	totalRows := 0
	csvreader := csv.NewReader(f)
	for {
		row, err := csvreader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			s.logger.Error(errors.Wrap(err, "read file"))
		}

		totalRows++

		price, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			s.logger.Error(errors.Wrap(err, "parse price to float"))
			return errors.Wrap(err, "parse price to float")
		}

		product := Product{
			Name:  row[1],
			Price: price,
		}

		wg.Add(1)
		jobs <- product // block, need consumer
		totalJobDispatched++
	}
	close(jobs)

	wg.Wait()

	return nil
}

func (s *service) workerDispatcher(ctx context.Context, id int, jobs <-chan Product, output chan<- string) {
	for product := range jobs {
		p, err := s.product.Create(ctx, product)
		if err != nil {
			s.logger.Error(err, " failed create product")
		}
		output <- fmt.Sprintf("worker:%d, product %s__%d ✅", id, product.Name, p.ID)
	}
}
