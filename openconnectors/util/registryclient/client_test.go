package registryclient

import (
	"fmt"
	"testing"

	"time"

	apiclient "github.com/akrausesap/goexperiments/openconnectors/util/registryclient/client"
	"github.com/akrausesap/goexperiments/openconnectors/util/registryclient/client/services"
	"github.com/akrausesap/goexperiments/openconnectors/util/registryclient/models"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func stringToAddress(i string) *string {
	return &i
}

func Test_Registry(t *testing.T) {
	transport := httptransport.New("registry-dummy.kyma.local", "/test/", []string{"https"})
	client := apiclient.New(transport, strfmt.Default)
	params := services.NewV1MetadataServicesPostParams()
	params.SetTimeout(30 * time.Second)

	body := &models.ServiceDetails{
		Provider:    stringToAddress("SAP CP Open Connectors"),
		Name:        stringToAddress("Connector Instance - Connector Type - Instance Name"),
		Description: stringToAddress("SAP CP Open Connectors instance for Twitter"),
		Identifier:  "connectorId",
		API: &models.API{
			TargetURL:        stringToAddress("https://www.google.com"),
			SpecificationURL: "https://raw.githubusercontent.com/kyma-project/kyma/release-1.2/components/application-registry/docs/api/api.yaml",
			Headers: map[string]interface{}{
				"Authorization": []string{"Dummy"},
			},
			APIType: "REST",
		},
		/*Documentation: &models.Documentation{

		},*/
	}

	params.SetBody(body)

	result, err := client.Services.V1MetadataServicesPost(params)

	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	fmt.Println(result)
	t.Log(result)

}
