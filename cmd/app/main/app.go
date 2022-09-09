package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/config"
	business_line2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/business_line"
	business_line "github.com/IgorTkachuk/cartridge_accounting/internal/domain/business_line/db"
	cartridge_model3 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_model"
	cartridge_model "github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_model/db"
	doc_type2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/doc_type"
	doc_type "github.com/IgorTkachuk/cartridge_accounting/internal/domain/doc_type/db"
	employee2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/employee"
	employee "github.com/IgorTkachuk/cartridge_accounting/internal/domain/employee/db"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/ou"
	ou2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/ou/db"
	prnt2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/prnt"
	prnt "github.com/IgorTkachuk/cartridge_accounting/internal/domain/prnt/db"
	user2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	user "github.com/IgorTkachuk/cartridge_accounting/internal/domain/user/db"
	vndr2 "github.com/IgorTkachuk/cartridge_accounting/internal/domain/vndr"
	vndr "github.com/IgorTkachuk/cartridge_accounting/internal/domain/vndr/db"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers/auth"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers/bisiness_line"
	cartridge_model2 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/cartridge_model"
	doc_type3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/doc_type"
	employee3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/employee"
	ou3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/ou"
	prnt3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/prnt"
	user3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/user"
	vndr3 "github.com/IgorTkachuk/cartridge_accounting/internal/handlers/vndr"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/cache/freecache"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/client/postgresql"
	http2 "github.com/IgorTkachuk/cartridge_accounting/pkg/http"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/shutdown"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"os"
	"syscall"
	"time"
)

func main() {
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	//cfg := db.NewPgConfig("postgres", "mg0208", "localhost", "5432", "ctr")
	cli, _ := postgresql.NewClient(context.Background(), 3, 5*time.Second, cfg.Storage)
	r := user.NewRepository(cli, logger)
	svc := user2.NewService(r, logger)

	userHandler := user3.Handler{
		UserService: svc,
	}

	RTCache := freecache.NewCacheRepo(10)
	jwtHelper := jwt.NewHelper(RTCache, *logger)

	authHandler := auth.Handler{
		Logger:      *logger,
		UserService: svc,
		JWTHelper:   jwtHelper,
	}

	vendorsRepo := vndr.NewRepository(cli, logger)
	vendorSvc := vndr2.NewService(vendorsRepo, logger)
	vendorHandler := vndr3.Handler{
		VendorService: vendorSvc,
	}

	printersRepo := prnt.NewRepository(cli, logger)
	printersSvc := prnt2.NewService(printersRepo, logger)
	printersHandler := prnt3.Handler{
		PrinterService: printersSvc,
	}

	ctrModelsRepo := cartridge_model.NewRepository(cli, logger)
	ctrModelsSvc := cartridge_model3.NewService(ctrModelsRepo, logger)
	ctrModelsHandler := cartridge_model2.Handler{
		CartridgeModelSvc: ctrModelsSvc,
	}

	ouRepo := ou2.NewRepository(cli, logger)
	ouSvc := ou.NewService(ouRepo, logger)
	ouHandler := ou3.Handler{
		OuService: ouSvc,
	}

	blRepo := business_line.NewRepository(cli, logger)
	blSvc := business_line2.NewService(blRepo, logger)
	blHandler := bisiness_line.Handler{
		BusinessLineSvc: blSvc,
	}

	employeeRepo := employee.NewRepository(cli, logger)
	employeeSvc := employee2.NewService(employeeRepo, logger)
	employeeHandler := employee3.Handler{
		EmployeeService: employeeSvc,
	}

	docTypeRepo := doc_type.NewRepository(cli, logger)
	docTypeSvc := doc_type2.NewService(docTypeRepo, logger)
	docTypeHandler := doc_type3.Handler{
		DocTypeService: docTypeSvc,
	}

	logger.Info("create approuter")
	approuter := httprouter.New()

	logger.Info("register user handler")
	userHandler.Register(approuter)

	logger.Info("register auth handler")
	authHandler.Register(approuter)

	logger.Info("register vendor handler")
	vendorHandler.Register(approuter)

	logger.Info("register printers handler")
	printersHandler.Register(approuter)

	logger.Info("register cartridge models handler")
	ctrModelsHandler.Register(approuter)

	logger.Info("register organizational unit handler")
	ouHandler.Register(approuter)

	logger.Info("Register business line handler")
	blHandler.Register(approuter)

	logger.Info("Register employee handler")
	employeeHandler.Register(approuter)

	logger.Info("Register doc type handler")
	docTypeHandler.Register(approuter)

	logger.Info("apply CORS settings")
	corsSettings := http2.CorsSettings()

	router := corsSettings.Handler(approuter)

	start(router)

	//users, _ := r.GetAll(context.Background())
	//
	//for _, u := range users {
	//	fmt.Println(u)
	//}
}

func start(router http.Handler) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var server *http.Server
	var listener net.Listener
	var listenerErr error

	listener, listenerErr = net.Listen("tcp", fmt.Sprintf("%s:%s", "0.0.0.0", "3001"))
	if listenerErr != nil {
		logger.Fatal(listenerErr)
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
