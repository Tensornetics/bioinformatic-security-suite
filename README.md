# bioinformatic-security-suite

```
bioinformatic-security-suite
├── README.md
├── LICENSE
├── docs
│   ├── architecture.md
│   ├── design.md
│   ├── installation.md
│   ├── usage.md
│   └── troubleshooting.md
├── chaincode
│   ├── chaincode.go
│   ├── chaincode_test.go
│   ├── fabcar.go
│   ├── fabcar_test.go
│   ├── metadata.go
│   └── metadata_test.go
├── microservices
│   ├── edge-service
│   │   ├── edge-service.py
│   │   ├── edge-service_test.py
│   │   ├── requirements.txt
│   │   └── deployment.yaml
│   └── cloud-service
│       ├── cloud-service.py
│       ├── cloud-service_test.py
│       ├── requirements.txt
│       └── deployment.yaml
├── webgui
│   ├── index.html
│   ├── style.css
│   ├── app.js
│   ├── login.js
│   ├── webgui_test.js
│   └── deployment.yaml
├── k8s
│   ├── cluster-config
│   ├── istio-config
│   ├── kustomization.yaml
│   └── cert-manager-config
└── infra
    ├── terraform
    │   ├── main.tf
    │   ├── variables.tf
    │   ├── outputs.tf
    │   └── terraform.tfvars
    └── ansible
        ├── playbook.yml
        ├── inventory
        └── group_vars
```
