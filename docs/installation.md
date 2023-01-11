# Installation

The bioinformatic-security-suite project can be installed in a local development environment or in a production environment. The following instructions will guide you through the installation process.

## Prerequisites

- A Kubernetes cluster
- [Istio](https://istio.io/docs/setup/)  installed in the cluster
- [Azure Govcloud](https://azure.microsoft.com/en-us/services/government/) account
- [Docker](https://www.docker.com/) installed on the local development environment

## Install the project

1. Clone the bioinformatic-security-suite repository

```
git clone https://github.com/tensornetics/bioinformatic-security-suite
```

2. Install the required components by running the provided ansible playbooks

```
cd infra/ansible
ansible-playbook -i inventory --ask-vault-pass playbook.yml
```

3. Create the k8s resources using the provided kustomization files

```
cd k8s
kubectl apply -k .
```

4. Create the Microsoft Azure GovCloud resources using terraform

```
cd infra/terraform
terraform init
terraform apply
```

5. Build and push the Docker images for edge-service, cloud-service, and webgui

```
cd microservices/edge-service
docker build -t edge-service .
docker push edge-service
```

```
cd microservices/cloud-service
docker build -t cloud-service .
docker push cloud-service
```

```
cd webgui
docker build -t webgui .
docker push webgui
```

6. Configure the edge-service and cloud-service to use mTLS and ingress gateway by applying the provided Istio config files

```
cd k8s
istioctl apply -f istio-config
```

The installation process should now be complete. To verify that the system is running correctly, you can check the status of the pods and services using the kubectl command.

```
kubectl get pods
kubectl get services
```
