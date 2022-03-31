package product_test

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"path"
// 	"runtime"
// 	"testing"
// 	"time"

// 	"github.com/tuingking/flamingo/config"
// 	"github.com/tuingking/flamingo/infra/logger"
// 	"github.com/tuingking/flamingo/infra/mysql"
// 	"github.com/tuingking/flamingo/internal/product"

// 	. "github.com/smartystreets/goconvey/convey"
// )

// type DI struct {
// 	cfg  config.Config
// 	log  logger.Logger
// 	sql  mysql.MySQL
// 	repo product.Repository
// 	svc  product.Service
// }

// var (
// 	di DI
// )

// func TestMain(m *testing.M) {
// 	// change dir
// 	_, filename, _, _ := runtime.Caller(0)
// 	dir := path.Join(path.Dir(filename), "..", "..")
// 	err := os.Chdir(dir)
// 	if err != nil {
// 		panic(err)
// 	}

// 	di = DI{}

// 	// init config
// 	di.cfg = config.Init(
// 		config.WithConfigFile("config"),
// 		config.WithConfigType("yaml"),
// 	)

// 	fmt.Printf("[DEBUG] MaxOpenConn: %+v\n", di.cfg.MySQL.MaxOpenConn)
// 	fmt.Printf("[DEBUG] MaxIdleConn: %+v\n", di.cfg.MySQL.MaxIdleConn)

// 	// init dependency
// 	di.log = logger.New(logger.Config{Level: "error"})
// 	di.sql = mysql.New(di.cfg.MySQL)
// 	di.repo = product.NewRepository(di.sql)
// 	di.svc = product.NewService(di.cfg.Product.Service, di.log, di.repo)

// 	defer func() {
// 		di.sql.Close()
// 	}()

// 	exitVal := m.Run()

// 	os.Exit(exitVal)
// }

// func doPrecondition(totalProduct int) (f *os.File, err error) {
// 	return di.svc.GenerateProductsCsv(context.Background(), int64(totalProduct))
// }

// func Test_CreateBulk_WorkerVariation(t *testing.T) {
// 	var (
// 		totalProduct  = 10000
// 		tWorker       = []int{1, 5, 10, 30, 40, 50}
// 		results       = make(map[int]time.Duration)
// 		minTime       = 10 * time.Minute
// 		optimumWorker = tWorker[0]
// 	)

// 	Convey(fmt.Sprintf("Test bulk create with %d product", totalProduct), t, func() {
// 		Convey("do precondition", func() {
// 			f, err := doPrecondition(totalProduct)
// 			Convey("precondition should be success", func() {
// 				So(err, ShouldBeNil)
// 			})
// 			defer os.Remove(f.Name())

// 			for _, worker := range tWorker {
// 				if worker >= totalProduct {
// 					continue
// 				}

// 				Convey(fmt.Sprintf("Then insert %d task with %d worker", totalProduct, worker), func() {
// 					start := time.Now()
// 					defer func() {
// 						elapsed := time.Since(start)
// 						fmt.Print("took: ", elapsed)

// 						if elapsed < minTime {
// 							optimumWorker = worker
// 							minTime = elapsed
// 						}

// 						results[optimumWorker] = elapsed
// 					}()

// 					svc := product.NewService(product.ConfigSvc{Worker: worker}, di.log, di.repo)
// 					err := svc.BulkCreate(context.Background(), f.Name())
// 					Convey("Product should be created", func() {
// 						So(err, ShouldBeNil)
// 					})
// 				})
// 			}
// 		})
// 	})

// 	fmt.Printf("🥳🥳🥳 The most optimum to handle %d product is %d worker which took %v\n", totalProduct, optimumWorker, minTime)
// }

// func Test_CreateBulk_TaskVariation(t *testing.T) {
// 	var (
// 		tProduct = []int{5}
// 		tWorker  = 10
// 	)

// 	for _, tproduct := range tProduct {
// 		Convey(fmt.Sprintf("Test bulk create with %d product with %d worker", tproduct, tWorker), t, func() {
// 			Convey("do precondition", func() {
// 				f, err := doPrecondition(tproduct)
// 				Convey("precondition should be success", func() {
// 					So(err, ShouldBeNil)
// 				})
// 				defer os.Remove(f.Name())

// 				Convey(fmt.Sprintf("Then insert %d product", tproduct), func() {
// 					start := time.Now()
// 					defer func() {
// 						elapsed := time.Since(start)
// 						fmt.Print("took: ", elapsed)
// 					}()

// 					svc := product.NewService(product.ConfigSvc{Worker: tWorker}, di.log, di.repo)
// 					err := svc.BulkCreate(context.Background(), f.Name())
// 					Convey(fmt.Sprintf("%d product should be created", tproduct), func() {
// 						So(err, ShouldBeNil)
// 					})
// 				})
// 			})
// 		})
// 	}
// }
