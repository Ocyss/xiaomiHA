package main

import (
	"github.com/Ocyss/xiaomiHA/internal/feishu"
	"github.com/Ocyss/xiaomiHA/internal/task"
	"github.com/Ocyss/xiaomiHA/utils"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	go feishu.StartWsClient()

	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt:               utils.C.Redis,
			PeriodicTaskConfigProvider: task.T,           // this provider object is the interface to your config source
			SyncInterval:               30 * time.Second, // this field specifies how often sync should happen
			SchedulerOpts: &asynq.SchedulerOpts{
				Location: time.Local,
			},
		})

	if err != nil {
		log.Fatal(err)
	}

	if err := mgr.Run(); err != nil {
		log.Fatal(err)
	}

	srv := asynq.NewServer(
		utils.C.Redis,
		asynq.Config{Concurrency: 10},
	)

	// NOTE: We'll cover what this `handler` is in the section below.
	if err := srv.Run(task.T.Mux); err != nil {
		log.Fatal(err)
	}

	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: utils.C.Redis,
	})

	// Note: We need the tailing slash when using net/http.ServeMux.
	http.Handle(h.RootPath()+"/", h)

	// Go to http://localhost:8088/monitoring to see asynqmon homepage.
	log.Fatal(http.ListenAndServe(":8088", nil))
}
