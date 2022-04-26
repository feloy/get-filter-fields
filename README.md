# get-filter-fields

Kubernetes provides a way to filter resources using *field selectors*. Available fields are, for all resources, `metadata.name` and, for all namespaced resources, `metadata.namespace`.

Other fields are available depending on resources, but the list of available fields is not documented (to my knowledge).

**get-filter-fields** is a program that will list the possible fields for different resources. For non listed resources, only `metadata.name` and `metadata.namespace` will be available.

The list of fields is not obtained dynamically, but by linking with a specific version of the Kubernetes sources, at this time the `1.23.6` version.

To get the list of the fields for a different version of Kubernetes, you can cleanup the `go.mod` and `go.sum` files (by removing the `go.sum` file, and by keeping only `module` and `go` lines in `go.mod`), then running the script `./download-deps.sh` with the Kubernetes version as parameter.

```
$ go run main.go 

"configmap" filterable fields:
- metadata.name
- metadata.namespace

"event" filterable fields:
- involvedObject.apiVersion
- involvedObject.fieldPath
- involvedObject.kind
- involvedObject.name
- involvedObject.namespace
- involvedObject.resourceVersion
- involvedObject.uid
- metadata.name
- metadata.namespace
- reason
- reportingComponent
- source
- type

"namespace" filterable fields:
- metadata.name
- name
- status.phase

"node" filterable fields:
- metadata.name
- spec.unschedulable

"persistentvolume" filterable fields:
- metadata.name
- name

"persistentvolumeclaim" filterable fields:
- metadata.name
- metadata.namespace
- name

"pod" filterable fields:
- metadata.name
- metadata.namespace
- spec.nodeName
- spec.restartPolicy
- spec.schedulerName
- spec.serviceAccountName
- status.nominatedNodeName
- status.phase
- status.podIP

"replicationcontroller" filterable fields:
- metadata.name
- metadata.namespace
- status.replicas

"secret" filterable fields:
- metadata.name
- metadata.namespace
- type

"replicaset" filterable fields:
- metadata.name
- metadata.namespace
- status.replicas

"job" filterable fields:
- metadata.name
- metadata.namespace
- status.successful

"certificatesigningrequest" filterable fields:
- metadata.name
- spec.signerName
```
