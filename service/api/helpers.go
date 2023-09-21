// package api contains API functionalities including the multiverse service.
package api

// Import necessary third-party packages.
import (
	"fmt"
	httpIface "github.com/taubyte/http"
	"github.com/taubyte/tau/libdream/common"
)

// getUniverse is a method for the multiverseService struct. It fetches a common.Universe
// object based on the "universe" variable present in the HTTP request's context.
func (srv *multiverseService) getUniverse(ctx httpIface.Context) (common.Universe, error) {
	// Retrieve the "universe" string variable from the HTTP context.
	name, err := ctx.GetStringVariable("universe")
	if err != nil {
		// Return nil and wrap the error if fetching the variable fails.
		return nil, fmt.Errorf("failed getting name with: %w", err)
	}

	// Check if a universe with the provided name exists.
	exist := srv.Exist(name)
	if exist {
		// Return the common.Universe object corresponding to the provided name.
		return srv.Universe(name), nil
	} else {
		// Return nil and an error if the universe does not exist.
		return nil, fmt.Errorf("universe %s does not exist", name)
	}
}
