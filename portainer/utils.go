package portainer

import (
	"errors"
	"fmt"

	"github.com/acottais/portctl/portainerapi/client/stacks"

	"github.com/acottais/portctl/dockerapi/models"
	"github.com/go-openapi/runtime"

	"github.com/acottais/portctl/dockerapi/client/swarm"

	dockerapi "github.com/acottais/portctl/dockerapi/client"
	portainerapi "github.com/acottais/portctl/portainerapi/client"
	"github.com/acottais/portctl/portainerapi/client/endpoints"
	portainertype "github.com/acottais/portctl/portainerapi/models"
)

// EndpointLookup look up for the endpoint based on the name and returns the Id
func EndpointLookup(cli *portainerapi.Portainer, name string) (*portainertype.Endpoint, error) {
	endpointList, err := cli.Endpoints.EndpointList(nil, nil)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(endpointList.Payload); i++ {
		if endpointList.Payload[i].Name == name {
			params := endpoints.NewEndpointInspectParams()
			params.ID = endpointList.Payload[i].ID
			inspect, err := cli.Endpoints.EndpointInspect(params, nil)
			if err != nil {
				return nil, err
			}
			return inspect.Payload, nil
		}
	}
	return nil, nil
}

// StackLookup returns the stack with the name provided for the specified endpoint
func StackLookup(cli *portainerapi.Portainer, swarmID, name string) (*portainertype.Stack, error) {
	params := stacks.NewStackListParams()
	filter := fmt.Sprintf(`{"SwarmId":"%s"}`, swarmID)
	params.Filters = &filter
	stackList, err := cli.Stacks.StackList(params, nil)
	if err != nil {
		_, err = CheckPortainerError(err)
		return nil, err
	}
	for i := 0; i < len(stackList.Payload); i++ {
		if stackList.Payload[i].Name == name {
			params := stacks.NewStackInspectParams()
			params.ID = stackList.Payload[i].ID
			inspect, err := cli.Stacks.StackInspect(params, nil)
			if err != nil {
				return nil, err
			}
			return inspect.Payload, nil
		}
	}
	return nil, nil
}

// GetDockerEndpoint returns the docker endpoint for a given portainer endpoint id
func GetDockerEndpoint(path string, id int64) string {
	return fmt.Sprintf("%s/endpoints/%d/docker", path, id)
}

// SwarmLookup look up for the swarm information of the given engine
func SwarmLookup(cli *dockerapi.DockerEngine) (*models.Swarm, error) {
	params := swarm.NewSwarmInspectParams()
	swarm, err := cli.Swarm.SwarmInspect(params)
	if err != nil {
		return nil, err
	}
	return swarm.Payload, nil

}

type portainererror struct {
	err     string
	details string
}

// CheckPortainerError check for common error like server error or authentication failed.
// It returns 0 when the error is handled or the http error code otherwise
func CheckPortainerError(err error) (int, error) {
	if err != nil {

		if err, ok := err.(*runtime.APIError); ok {
			if err.Code == 401 {

				return err.Code, errors.New("Error: Portainer server requires login. Use portctl login")
			}
			if err.Code == 403 {

				return err.Code, errors.New("Error: access denied. You may not have the correct authorization")
			}
			if err.Code >= 500 {

				return err.Code, errors.New("Error: Portainer Server Error")
			}
			var message string
			switch t := err.Response.(type) {
			default:
				message = fmt.Sprintf("unexpected type %T\n", t) // %T prints whatever type t has
			case portainererror:
				message = t.details // t has type bool
			}
			return err.Code, errors.New(message)
		}
		return 0, err
	}
	return 0, nil
}
