package http

import (
	"errors"
	"fmt"
	"github.com/taubyte/dreamland/service/inject"
)

// Inject takes a list of operations of type inject.Injectable and injects them
// into the Universe by calling runInjection on its client.
func (u *Universe) Inject(ops ...inject.Injectable) error {
	// Iterate through each operation in the provided slice
	for _, op := range ops {
		// Execute the injection and get back any errors that occur
		err := u.client.runInjection(u.Name, op)
		if err != nil {
			// Wrap the error with context and return
			return fmt.Errorf("Injection `%s` failed with error: %w", op.Name, err)
		}
	}

	// If we made it through the loop without returning an error, we were successful
	return nil
}

// runInjection takes a universe name and an inject.Injectable object,
// executing the injection operation as specified by the Injectable object.
func (c *Client) runInjection(universe string, op inject.Injectable) (err error) {
	// Initialize Params if it's nil
	if op.Params == nil {
		op.Params = []interface{}{}
	}

	// Create a map to store return values and parameters
	ret := map[string]interface{}{"params": op.Params}

	// If Config is not nil, add it to the return map
	if op.Config != nil {
		ret["config"] = op.Config
	}

	// Perform the injection operation based on the HttpMethod specified
	switch op.Method {
	case inject.POST:
		// Perform POST request
		err = c.post(op.Run(universe), ret, nil)

	default:
		// If the method is not supported, return an error
		err = errors.New("Method not supported " + op.Method.String())
	}

	// Return any errors that occurred
	return
}
