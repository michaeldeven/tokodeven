package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/michaeldeven/tokodeven/api/router"
	"github.com/michaeldeven/tokodeven/auto"
	"github.com/michaeldeven/tokodeven/config"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tlistening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
