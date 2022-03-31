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
	GenerateRandomProducts(ctx context.Context, n int) []Product
	GenerateProductsCsv(ctx context.Context, n int) (*os.File, error)
	Seed(ctx context.Context, n int) error
}

type service struct {
	config  ConfigSvc
	logger  logger.Logger
	product Repository
}

type ConfigSvc struct{}

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

func (s *service) GenerateRandomProducts(ctx context.Context, n int) []Product {
	var products []Product

	for i := 0; i < n; i++ {
		product := Product{
			Name:  fmt.Sprintf("Product-%d", i+1),
			Price: float64(1000 + (i+1)*100),
		}

		products = append(products, product)
	}

	return products
}

func (s *service) GenerateProductsCsv(ctx context.Context, n int) (*os.File, error) {
	now := time.Now()
	defer func() {
		s.logger.Info("GenerateProductsCsv took ", time.Since(now))
	}()
	products := s.GenerateRandomProducts(ctx, n)

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

func (s *service) Seed(ctx context.Context, n int) error {
	var (
		totalJobFinished int

		// define channel
		jobs   = make(chan Product)
		output = make(chan string)
		wg     = sync.WaitGroup{}
		lock   = sync.Mutex{}
		worker = 10
	)

	defer func() {
		s.logger.Warnf("success: %d/%d", totalJobFinished, n)
		if totalJobFinished != int(n) {
			failed := n - totalJobFinished
			s.logger.Warnf("failed: %d", failed)
		}
	}()

	// generate random product
	csvf, err := s.GenerateProductsCsv(ctx, n)
	if err != nil {
		return errors.Wrap(err, "GenerateProductsCsv")
	}

	f, err := os.Open(csvf.Name())
	if err != nil {
		return errors.Wrap(err, "open csv file")
	}
	defer f.Close()

	// spawn worker
	for i := 0; i < worker; i++ {
		go s.workerDispatcher(ctx, i, jobs, output)
	}

	// listen to channel output
	go func() {
		for res := range output {
			if res != "ERR" {
				lock.Lock()
				totalJobFinished++
				lock.Unlock()
				s.logger.Info(res)
			}
			wg.Done()
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
	}
	close(jobs)

	wg.Wait()

	return nil
}

func (s *service) workerDispatcher(ctx context.Context, id int, jobs <-chan Product, output chan<- string) {
	for product := range jobs {
		p, err := s.product.Create(ctx, product)
		if err != nil {
			output <- "ERR"
		}
		output <- fmt.Sprintf("worker:%d, product %s__%d ✅", id, product.Name, p.ID)
	}
}
