# Design

The bioinformatic-security-suite project is designed to securely collect, store, and share biodata from wearables using a blockchain network and a mesh architecture with Kubernetes and Istio.

## Wearables

The wearables are devices that collect biodata, this data can be collected in different formats and protocols, to keep the project flexible we chose a protocol that is widely adopted in the wearables industry, Bluetooth Low Energy (BLE).

## Edge-service

The edge-service is implemented using Python, it receives data from wearables and sends it to the cloud-service. It validates the data, ensures that it is properly formed, and verifies its authenticity by checking it against the blockchain network, it also authenticates the sender.

## Cloud-service

The cloud-service is also implemented using Python, it receives data from the edge-service and stores it in Azure Govcloud. The cloud-service also provides an API for authorized users to access the data through a web-based GUI.

## Web-based GUI

The web-based GUI is implemented using JavaScript, HTML, and CSS. It allows authorized users to view and interact with the biodata stored in the cloud database. The GUI is designed to be simple and intuitive, with a focus on usability.

## Blockchain network

The blockchain network is implemented using Hyperledger Fabric, a popular open-source platform for building enterprise blockchain applications. The chaincode implements the chain-of-custody and the validation of the data

## Kubernetes and Istio

Kubernetes is used to manage the deployment and scaling of the edge-service and cloud-service, while Istio is used to provide a service mesh for the edge-service and cloud-service pods. Istio also provides features such as traffic management, load balancing, and service-to-service authentication.

## Data storage

The biodata is stored in a cloud database, Azure Govcloud, it provides a secure and compliant solution for storing sensitive data. It's designed to meet the compliance requirements of the health industry like HIPAA

## Authentication and Authorization

Authentication and Authorization are handled through Istio's ingress gateway, it uses JSON Web Tokens (JWT) to authenticate the user, and it's based on OAuth2 for authorization.

## Monitoring and logging

The Kubernetes cluster and Istio service mesh are instrumented with monitoring and logging systems to provide visibility into the health and performance of the system, and facilitate troubleshooting and debugging.

## Security

The data is encrypted in transit and at rest, security protocols were adopted as AES for transit and at-rest data encryption and RSA for key encryption. A certification manager was used to issue certificates for the mTLS.
