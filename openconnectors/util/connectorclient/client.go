package connectorclient

import (
	"fmt"
	"strconv"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	apiclient "github.com/akrausesap/goexperiments/openconnectors/util/connectorclient/client"
	instances "github.com/akrausesap/goexperiments/openconnectors/util/connectorclient/client/instances"
)

//ConnectorInstance describes all necessary parameters to connect to a Connector Insatnace of SAP CP Open Connectors
type ConnectorInstance struct {
	ID            string
	Name          string
	APIKey        string
	ConnectorName string
}

func createAuthorizationHeader(userSecret string, orgSecret string) string {
	return fmt.Sprintf("User %s, Organization %s", userSecret, orgSecret)
}

// GetConnectorAPISpecification retrieves the Swagger File from Open Connectors
func GetConnectorAPISpecification(host string, userSecret string, orgSecret string, connectorInstanceID string) (apiSpecification interface{}, err error) {
	transport := httptransport.New(host, "/elements/api-v2/", []string{"https"})
	client := apiclient.New(transport, strfmt.Default)

	params := instances.NewGetInstancesDocs2Params()
	var connectorInstanceIDInt64 int64
	if connectorInstanceIDInt64, err = strconv.ParseInt(connectorInstanceID, 10, 64); err != nil {
		return
	}

	params.SetTimeout(30 * time.Second)
	params.SetID(connectorInstanceIDInt64)
	params.SetAuthorization(createAuthorizationHeader(userSecret, orgSecret))

	instancedocs, err := client.Instances.GetInstancesDocs2(params)

	if err != nil {
		return
	}
	apiSpecification = instancedocs.Payload

	return
}

// GetConnectorInstances retrieves a list of registered Connector Instances for an SAP CP Open Connectors
// Tenant
func GetConnectorInstances(host string, userSecret string, orgSecret string, tags []string) (connectorInstances []ConnectorInstance, err error) {

	transport := httptransport.New(host, "/elements/api-v2/", []string{"https"})
	client := apiclient.New(transport, strfmt.Default)

	params := instances.NewGetInstancesParams()

	params.SetAuthorization(createAuthorizationHeader(userSecret, orgSecret))
	params.SetTags(tags)
	params.SetTimeout(30 * time.Second)

	elementInstanceList, err := client.Instances.GetInstances(params)

	if err != nil {
		return
	}

	connectorInstances = make([]ConnectorInstance, len(elementInstanceList.Payload))

	for i, elementInstance := range elementInstanceList.Payload {

		connectorInstance := ConnectorInstance{
			ID:            fmt.Sprintf("%.0f", elementInstance.ID),
			Name:          elementInstance.Name,
			APIKey:        fmt.Sprint(elementInstance.Token),
			ConnectorName: elementInstance.Element.Name,
		}
		connectorInstances[i] = connectorInstance
	}
	return
}
