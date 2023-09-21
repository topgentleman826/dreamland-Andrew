// package api defines the core functionalities for the multiverse service.
package api

// Import necessary standard and third-party packages.
import (
	"context"
	"time"
	"goHttp "net/http"
	"github.com/pterm/pterm"
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

// multiverseService is a struct that represents the multiverse service.
type multiverseService struct {
	// rest is an interface that represents the RESTful HTTP service.
	rest httpIface.Service
	// Multiverse represents the core functionalities of the multiverse.
	common.Multiverse
}

// BigBang is the function that initializes and starts the multiverse service.
func BigBang() error {
	var err error

	// Create a new multiverse service instance.
	srv := &multiverseService{
		Multiverse: services.NewMultiVerse(),
	}

	// Initialize the RESTful service with configurations.
	srv.rest, err = http.New(srv.Context(), options.Listen(common.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))
	if err != nil {
		return err
	}

	// Set up HTTP routes and start the service.
	srv.setUpHttpRoutes().Start()

	// Create a context that times out after 10 seconds.
	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), 10*time.Second)
	defer waitCtxC()

	// Loop to wait for the service to be ready or encounter an error.
	for {
		select {
		case <-waitCtx.Done():
			// Return any error that the context may have.
			return waitCtx.Err()
		case <-time.After(100 * time.Millisecond):
			// Check for errors in the REST service.
			if srv.rest.Error() != nil {
				pterm.Error.Println("Dreamland failed to start")
				return srv.rest.Error()
			}
			// Perform an HTTP GET request to see if the service is ready.
			_, err := goHttp.Get("http://" + common.DreamlandApiListen)
			if err == nil {
				pterm.Info.Println("Dreamland ready")
				return nil
			}
		}
	}
}
