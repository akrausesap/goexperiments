# Kyma / SAP Open Connectors Integration

## About

This repo provides some PoC code to add a CRD and a Controller to a Kyma (kyma-project.io) (Kubernetes) cluster that will connect to SAP Open Connectors and fetch instantiated Connectors to create their corresponding respresentations as Applications in the Application Connector (https://kyma-project.io/docs/components/application-connector/).

It uses Kubebuilder (https://github.com/kubernetes-sigs/kubebuilder) and Go Swagger https://github.com/go-swagger/go-swagger as a foundation.

## Installation

There is currently no ready made installer or helm char available. The only safe way is to clone the repository and run

```
make install
make run
```

for local testing.

And run:

```
make docker-build docker-push IMG=<some-registry>/controller
make deploy
```

for testing in cluster.

If you want to test w/o full kyma deployed. First run `kubectl apply -f config/kyma/application.yaml`. Otherwise Kyma 1.2 is a prerequisite.

## Development

Connectivity to SAP CP Open connectors is based on `util/connectorclient/open_connectors_instances_swagger.json`. A client is generated using Go Swagger. To regenerate, use `make swagger`. 

Periodic polling is scheduled using the Watch function (`controllers/connectorinstance_controller.go`):

```
return ctrl.NewControllerManagedBy(mgr).
		Watches(&source.Channel{Source: r.Events},
			&handler.EnqueueRequestForObject{}).
		For(&openconnectors.ConnectorInstance{}).
		Complete(r)
```

`r.Events` is fed with periodic pings from `util/reconciletrigger/reconciler.go`.

## Sample

The controller is bound to the `ConnectorInstance` CRD. Just apply it using your own credentials and host. `organizationSecret` and `userSecret` can be retrieved from within the Open Connectors UI (see: https://blogs.sap.com/2019/02/20/sap-api-management-managing-3rd-party-services-via-open-connector-instances/ (Pre-requisite section) and https://help.openconnectors.ext.hana.ondemand.com/home/organization-secret-and-user-secret (Finding Your Secrets and Token))

```
apiVersion: openconnectors.incubator.kyma-project.io/v1
kind: ConnectorInstance
metadata:
  name: connectorinstance-sample
spec:
  displayName: "Open Connectors Trial"
  refreshIntervalSeconds: 30
  hostName: "api.openconnectors.ext.hanatrial.ondemand.com"
  organizationSecret: "your secret
  userSecret: "your secret"

```

For further details on Open COnnectors see https://blogs.sap.com/2018/09/19/part-1-enable-sap-cloud-platform-open-connectors-in-trial/