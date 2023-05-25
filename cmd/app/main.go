package main

import (
	"context"
	"github.com/fionahiklas/simple-go-template/internal/version"
	"github.com/fionahiklas/simple-go-template/pkg/versionhandler"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log := logrus.New()

	log.SetLevel(logrus.DebugLevel)
	
	versionHandler := versionhandler.NewHandler(log, version.CodeVersion(), version.CommitHash())

	mux := http.NewServeMux()
	mux.Handle("/version", versionHandler)

	log.Infof("Starting app version '%s', commit: '%s' on localhost:7777",
		version.CodeVersion(), version.CommitHash())

	if err := http.ListenAndServe("localhost:7777", mux); err != nil {
		log.Error(context.Background(), "Error from HTTP server: %s", err.Error())
	}
}
