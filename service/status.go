// package http defines the HTTP client and related functionalities.
package http

// Import the external "api" package from the github.com/taubyte/dreamland/service directory.
import "github.com/taubyte/dreamland/service/api"

// Status is a map that associates a string with a UniverseStatus object.
type Status map[string]UniverseStatus

// UniverseStatus holds information about the status of a universe.
type UniverseStatus struct {
	// NodeCount holds the number of nodes in the universe.
	NodeCount int `json:"node-count"`
	// Nodes is a map that associates a string with a slice of strings, presumably to hold node-related information.
	Nodes map[string][]string
}

// Status method for the Client type fetches the status information from the server.
func (c *Client) Status() (Status, error) {
	// Initialize an empty Status map.
	resp := make(Status)
	// Perform a GET request to the "/status" endpoint and populate the resp map with the received data.
	err := c.get("/status", &resp)
	if err != nil {
		// Return nil and the error if the GET request fails.
		return nil, err
	}
	// Return the populated Status map and no error.
	return resp, nil
}

// Status method for the Universe type fetches Echart status information from a specific universe.
func (u *Universe) Status() (resp api.Echart, err error) {
	// Perform a GET request to the "/les/miserables/<Universe Name>" endpoint and populate the resp object with the received data.
	err = u.client.get("/les/miserables/"+u.Name, &resp)
	// Return the populated api.Echart object and any error that may have occurred.
	return
}

// Id method for the Universe type fetches UniverseInfo about a specific universe.
func (u *Universe) Id() (resp api.UniverseInfo, err error) {
	// Perform a GET request to the "/id/<Universe Name>" endpoint and populate the resp object with the received data.
	err = u.client.get("/id/"+u.Name, &resp)
	// Return the populated api.UniverseInfo object and any error that may have occurred.
	return
}
