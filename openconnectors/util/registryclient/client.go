package registryclient

import (
	"fmt"

	"time"

	apiclient "github.com/akrausesap/goexperiments/openconnectors/util/registryclient/client"
	"github.com/akrausesap/goexperiments/openconnectors/util/registryclient/client/services"
	"github.com/akrausesap/goexperiments/openconnectors/util/registryclient/models"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

//RegisterAPIMetadata interacts with the Kyma Application Registry
func RegisterAPIMetadata(tlsEnabled bool, registryHost string, applicationName string, metadataProvider string, metadataName string,
	metadataDescription string, connectorInstanceID string, connectorURL string, apiSpecification interface{},
	authorizationHeader map[string][]string) error {
	targetURI := strfmt.URI(connectorURL)
	var scheme []string
	if tlsEnabled {
		scheme = []string{"https"}
	} else {
		scheme = []string{"http"}
	}
	transport := httptransport.New(registryHost, fmt.Sprintf("/%s/", applicationName), scheme)
	client := apiclient.New(transport, strfmt.Default)
	params := services.NewRegisterServiceParams()
	params.SetTimeout(30 * time.Second)

	body := &models.ServiceDetails{
		Provider:    &metadataProvider,
		Name:        &metadataName,
		Description: &metadataDescription,
		Identifier:  connectorInstanceID,
		API: &models.API{
			TargetURL: &targetURI,
			Spec:      apiSpecification,
			RequestParameters: &models.RequestParameters{
				Headers: authorizationHeader,
			},
			APIType: "REST",
		},
	}

	params.SetBody(body)

	_, err := client.Services.RegisterService(params)

	if err != nil {
		return err
	}

	return nil
}
