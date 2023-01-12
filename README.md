# Bioinformatic Security Suite

This security suite is designed to protect bioinformatic systems from malicious attacks. It provides a user-friendly interface for managing users, monitoring system activity, and implementing security controls.

## Getting Started

### Prerequisites

- Docker
- Python 3.8
- Node.js

### Installation

1. Clone the repository
```
git clone https://github.com/Tensornetics/bioinformatic-security-suite.git
```
2. Install the microservices
```
cd bioinformatic-security-suite/microservices/edge-service
pip install -r requirements.txt

cd ../cloud-service
pip install -r requirements.txt
```
3. Install webgui dependencies
```
cd ../../webgui
npm install
```
4. Run the webgui
```
npm start
```

## Usage

1. Register a new user through the webgui
2. Log in with the new user
3. Manage users and security settings through the webgui

## Contributing

We welcome contributions from the community. Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to submit contributions.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

- [OpenAI](https://openai.com/) for providing the language model.



```
bioinformatic-security-suite
├── README.md
├── LICENSE
├── docs
│   ├── architecture.md  (updated to explain the overall structure and design of the security suite)
│   ├── design.md (updated to detail the specific design decisions made during development)
│   ├── installation.md  (added to explain how to set up and install the security suite)
│   ├── usage.md (added to explain how to use the security suite, including examples)
│   ├── troubleshooting.md (added to explain how to diagnose and fix common issues that may arise)
│   ├── security_best_practices.md (added to explain the security practices used in the project)
│   ├── biometric_authentication.md (added to explain the biometric authentication feature and it's usage)
│   └── regulatory_compliance.md (added to explain the regulatory compliance in regards to the sensitive data)
├── chaincode
│   ├── chaincode.go
│   ├── chaincode_test.go (added to test the chaincode)
│   ├── fabcar.go
│   ├── fabcar_test.go (added to test the fabcar code)
│   ├── metadata.go
│   └── metadata_test.go (added to test the metadata code)
├── microservices
│   ├── edge-service
│   │   ├── edge-service.py
│   │   ├── edge-service_test.py (added to test the edge-service code)
│   │   ├── requirements.txt
│   │   └── deployment.yaml
│   └── cloud-service
│       ├── cloud-service.py
│       ├── cloud-service_test.py (added to test the cloud-service code)
│       ├── requirements.txt
│       └── deployment.yaml
├── webgui
│   ├── index.html
│   ├── style.css
│   ├── app.js
│   ├── login.js
│   ├── webgui_test.js (added to test the webgui code)
│   ├── biometric_authentication.js (added for the biometric authentication feature)
│   └── deployment.yaml
├── k8s
│   ├── cluster-config
│   ├── istio-config
│   ├── kustomization.yaml
│   ├── cert-manager-config
│   └── monitoring_logging.yaml (added to monitor and log the k8s resources)
└── infra
    ├── terraform
    │   ├── main.tf
    │   ├── variables.tf
    │   ├── outputs.tf
    │   └── terraform.tfvars
    ├── ansible
    │    ├── playbook.yml
    │    ├── inventory
    │    └── group_vars
    ├── ci_cd_pipeline  (added a folder for the CI/CD pipeline)
        ├── .travis.yml (added for travis CI)
        ├── Jenkinsfile (added for Jenkins pipeline)
        ├── Jenkinsfile.tests (added for testing pipeline)
        ├── Jenkinsfile.build (added for build pipeline)
        └── Jenkinsfile.deploy (added for deploy pipeline)
```
