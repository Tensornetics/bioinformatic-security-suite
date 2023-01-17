# AWS Setup

This guide will walk you through the process of setting up a cluster in AWS to run the Bioinformatic Security Suite.

## Prerequisites

- An AWS account
- The AWS CLI installed and configured on your machine
- Familiarity with AWS services, such as EC2 and EKS

## Setting up a Cluster

1. Create an Amazon Elastic Container Service for Kubernetes (EKS) cluster in your AWS account. 

```
aws eks create-cluster --name bioinformatic-security-suite-cluster --role-arn arn:aws:iam::ACCOUNT_ID:role/EKS_role --resources-vpc-config subnetIds=subnet-12345678,securityGroupIds=sg-12345678
```
2. Launch worker nodes and attach them to the cluster.
```
aws eks create-node-group --cluster-name bioinformatic-security-suite-cluster --node-group-name bioinformatic-security-suite-node-group --scaling-config minSize=1,maxSize=5,desiredSize=1 --instance-types t3.medium --ami-version amzn2-ami-hvm-2.0.20210115-x86_64-ebs --remote-access keyName=myKey,ec2SshKey=myKey
```
3. Configure kubectl to connect to your cluster.
```
aws eks update-kubeconfig --name bioinformatic-security-suite-cluster
```
4. Verify that your worker nodes are running and in the `Ready` state.
```
kubectl get nodes
```

## Next Steps

Now that you have set up a cluster in AWS, you can proceed to the next steps of the installation guide, such as deploying the Bioinformatic Security Suite to the cluster and configuring the various components of the system.
