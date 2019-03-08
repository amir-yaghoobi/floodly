package main

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/amir-yaghoobi/floodly/generator"
	"github.com/amir-yaghoobi/floodly/reporter"
	"github.com/amir-yaghoobi/floodly/user"
	"github.com/amir-yaghoobi/floodly/user/repository"

	_ "github.com/herenow/go-crate"
)

type runner struct {
	generatorSrv generator.Generator
	reportSrv    reporter.Reporter
}

func (r *runner) setupWorkers(total int, concurrency int, results chan<- reporter.Result) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go r.generatorSrv.Generate(total/concurrency, results, &wg)
	}

	wg.Wait()
	close(results)
}

func (r *runner) Run(total int, concurrency int) {
	results := make(chan reporter.Result, total)

	go r.setupWorkers(total, concurrency, results)

	r.reportSrv.Analyze(results)
}

func main() {
	db, err := sql.Open("crate", "http://localhost:4200/")
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewCrateRepository(db)
	err = userRepository.Migrate(true)
	if err != nil {
		log.Fatal(err)
	}

	genSrv := user.NewUserGenerator(*userRepository)
	repSrv := reporter.NewStandardReporter(os.Stdout, time.Now())

	srvRunner := &runner{
		generatorSrv: genSrv,
		reportSrv:    repSrv,
	}

	srvRunner.Run(5000, 500)
}
