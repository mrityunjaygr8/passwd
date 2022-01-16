package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/mrtyunjaygr8/passwd/utils"
	"github.com/mrtyunjaygr8/passwd/web/web"
)

func main() {
	fmt.Println("in web")
	config := utils.GetConfig()
	web := web.CreateWebApp(config)

	http.ListenAndServe(fmt.Sprintf("%s:%d", config.HOST, config.PORT), handlers.CombinedLoggingHandler(os.Stdout, web.Router))
}
