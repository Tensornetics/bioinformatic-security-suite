The cloud-service is deployed as a Kubernetes pod and runs within a Istio service mesh. The service is responsible for check data consistency, authenticate the sender and authenticate the data with the data of the blockchain network.

### Web-based GUI

The web-based GUI allows authorized users to view and interact with the biodata stored in the cloud database. It is deployed as a Kubernetes deployment, and it's accessed through an Istio ingress gateway, which handles authentication and authorization.

## Networking

The biodata is transmitted from the wearables to the edge-service, and then from the edge-service to the cloud-service, through a secure connection. The connection between the edge-service and the cloud-service is set up using mutual TLS (mTLS) which is enabled by Istio.

### Data storage

The biodata is stored in Microsoft Azure GovCloud, a cloud database that provides a secure and compliant solution for storing sensitive data.

### Monitoring and logging

The Kubernetes cluster and Istio service mesh are instrumented with monitoring and logging systems to provide visibility into the health and performance of the system, and facilitate troubleshooting and debugging.

### Authentication and Authorization

Authentication and Authorization are handled through Istio's ingress gateway, it uses JSON Web Tokens (JWT) to authenticate the user, and it's based on OAuth2 for authorization.
