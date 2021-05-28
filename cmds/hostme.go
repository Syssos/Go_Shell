package cmds

import (
	"os"
	"log"
	"fmt"
	"time"
	"errors"
	"syscall"
	"context"
	"net/http"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/Syssos/Go_Shell/color"
)
// Holds hostme command data such as arguments
type HostMe_cmd struct {
	Args []string
}
// Creates a lightweight server, handles CTRL+C signal to prevent closing shell
func (hmc HostMe_cmd) Run() error{
	if len(hmc.Args) > 0 {

	    router := mux.NewRouter()
		router.HandleFunc("/", hmc.serveFiles).Methods("GET")

		srv := &http.Server{
			Addr:    ":3000",
			Handler: router,
		}

		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
		log.Print("Server Started")

		<-done
		log.Print("Server Stopped")
		
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			// extra handling here
			cancel()
		}()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown Failed:%+v", err)
		}
		log.Print("Server Exited Properly")
		signal.Stop(done)
		return nil

	} else {
		return errors.New("No file to serve")
	}
	return nil
}
// Prints usage message for command
func (hmc HostMe_cmd) Usage() {
	fmt.Println(color.Yellow + "\n\thostme - This command will be used to host a single html file on a locally hosted server\n\n\t\thostme <filename>\n" + color.Reset)
}
// Function responsible for serving the html file
func (hmc HostMe_cmd)serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = hmc.Args[0]
    }
    http.ServeFile(w, r, p)
}
