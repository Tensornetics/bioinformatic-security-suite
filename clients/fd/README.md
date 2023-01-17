# FD Client

This folder contains the configuration and deployment files for the FD client of the Bioinformatic Security Suite.

## Version

This is version `vX.X` of the FD client files.

## Usage

To use these files, follow the instructions in the installation guide for deploying the Bioinformatic Security Suite to a cluster. Be sure to update the configuration files with the appropriate settings for the FD client.

## Changes Made

- List of changes made to the files in this version, such as new features or bug fixes.
- This can include the relevant issue or pull request numbers if the project is using a issue tracking or code review tool.

## Contact

If you have any questions or issues with these files, please contact the FD client team at `admin@tensornetics.com`.

```
To Do for FD Pipeline
├── README.md
├── LICENSE
├── docs
│   ├── architecture.md
│   ├── design.md
│   ├── installation.md
│   ├── usage.md
│   ├── troubleshooting.md
│   ├── security_best_practices.md
│   ├── biometric_authentication.md
│   └── regulatory_compliance.md
├── config
│   ├── config.yaml
│   ├── k8s-config.yaml
│   ├── terraform.tfvars
│   └── ansible-vars.yaml
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
│   ├── biometric_authentication.js
│   └── deployment.yaml
├── k8s
│   ├── cluster-config
│   ├── istio-config
│   ├── kustomization.yaml
│   ├── cert-manager-config
│   └── monitoring_logging.yaml
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
