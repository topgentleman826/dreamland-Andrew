// package api contains the core functionalities for the multiverse service, including HTTP routes.
package api

// Import the necessary third-party package.
import (
	"fmt"
	httpIface "github.com/taubyte/http"
)

// UniverseInfo is a struct that holds an ID field represented as a string in JSON format.
type UniverseInfo struct {
	Id string `json:"id"`
}

// idHttp sets up the HTTP GET route for fetching UniverseInfo.
func (srv *multiverseService) idHttp() {
	// Define the HTTP GET route.
	srv.rest.GET(&httpIface.RouteDefinition{
		// Path specifies the URL pattern for this route.
		Path: "/id/{universe}",
		// Vars holds the variables that are part of the URL pattern.
		Vars: httpIface.Variables{
			Required: []string{"universe"},
		},
		// Handler is the function executed when this route is accessed.
		Handler: func(ctx httpIface.Context) (interface{}, error) {
			// Retrieve the "universe" variable from the HTTP context.
			universeName, err := ctx.GetStringVariable("universe")
			if err != nil {
				// Return nil and the error if fetching the variable fails.
				return nil, err
			}

			// Check if a universe with the given name exists.
			exists := srv.Exist(universeName)
			if !exists {
				// Return nil and an error message if the universe doesn't exist.
				return nil, fmt.Errorf("universe `%s` does not exit", universeName)
			}

			// Fetch the universe object using its name.
			u := srv.Universe(universeName)
			// Return the UniverseInfo containing the universe's ID.
			return UniverseInfo{Id: u.Id()}, nil
		},
	})
}
