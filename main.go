package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/alimohseni/fem_project/internal/app"
	"github.com/alimohseni/fem_project/internal/routes"
)

func main()	{

		var port int
		flag.IntVar(&port, "port", 8080, "go backend server port")
		flag.Parse()

		app, err := app.NewApplication()

		if err != nil {
			panic(err)
		}

		r := routes.SetupRoutes(app)
		server := &http.Server{
						Addr: fmt.Sprintf(":%d", port),
						Handler: r,
						ReadTimeout: 10 * time.Second,
						WriteTimeout: 30 *time.Second,
		}

		app.Logger.Printf("We are running on port %d\n", port)

		err = server.ListenAndServe()
		if err != nil {
			app.Logger.Fatal(err)
		}
}
