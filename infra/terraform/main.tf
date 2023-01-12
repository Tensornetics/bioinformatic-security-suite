provider "aws" {
  region = var.aws_region
}

resource "aws_vpc" "main" {
  cidr_block = var.vpc_cidr_block
}

resource "aws_subnet" "main" {
  vpc_id            = aws_vpc.main.id
  cidr_block        = var.subnet_cidr_block
  availability_zone = var.availability_zone
}

resource "aws_security_group" "main" {
  vpc_id = aws_vpc.main.id
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_kubernetes_cluster" "main" {
  name = var.cluster_name
  role_arn = var.role_arn
  vpc_config {
    security_group_ids = [aws_security_group.main.id]
    subnet_ids         = [aws_subnet.main.id]
  }
}

resource "aws_eks_node_group" "main" {
  cluster_name    = aws_kubernetes_cluster.main.name
  node_group_name = var.node_group_name
  node_role_arn   = var.node_role_arn
  subnet_ids      = [aws_subnet.main.id]
}

output "kubeconfig" {
  value = aws_kubernetes_cluster.main.kubeconfig
}
