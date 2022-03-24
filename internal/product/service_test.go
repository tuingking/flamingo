package product_test

import (
	"context"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/product"
	mockproduct "github.com/tuingking/flamingo/mock/product"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	mockProductRepo *mockproduct.MockRepository

	log        logger.Logger
	productSvc product.Service
)

func setup(t *testing.T) func() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	ctrl := gomock.NewController(t)
	cfg := product.ConfigSvc{Worker: 3}
	log = logger.New(logger.Config{})
	mockProductRepo = mockproduct.NewMockRepository(ctrl)
	productSvc = product.NewService(cfg, log, mockProductRepo)

	return func() {
		ctrl.Finish()
	}
}

func Test_GenerateProducts(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	Convey("Given n=10", t, func() {
		products := productSvc.GenerateProducts(context.Background(), 10)
		Convey("When generated should return 10 products", func() {
			So(len(products), ShouldEqual, 10)
		})
	})
}

func Test_GenerateProductsCsv(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	deleteAfter := false

	f, _ := productSvc.GenerateProductsCsv(context.Background(), 1000)
	defer func() {
		if deleteAfter {
			log.Info("file removed 😘")
			os.Remove(f.Name())
		}
	}()
}

// func Benchmark_GenerateProductsCsv(b *testing.B) {
// 	ctrl := gomock.NewController(b)
// 	log = logger.New(logger.Config{})
// 	mockProductRepo = mockproduct.NewMockRepository(ctrl)
// 	productSvc = product.NewService(log, mockProductRepo)

// 	for i := 0; i < b.N; i++ {
// 		productSvc.GenerateProductsCsv(context.Background(), 100)
// 	}

// }

func Test_BulkCreate(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	mockProductRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(product.Product{}, nil).AnyTimes()
	_ = productSvc.BulkCreate(context.Background(), "tmp/products.csv")
}
