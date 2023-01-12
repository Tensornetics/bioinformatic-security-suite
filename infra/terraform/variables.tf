variable "region" {
  description = "The AWS region to deploy the resources in"
  default     = "us-east-1"
}
variable "vpc_cidr" {
  description = "The CIDR block to use for the VPC"
  default     = "10.0.0.0/16"
}
variable "subnet_cidr" {
  description = "The CIDR block to use for the public and private subnets"
  default     = "10.0.0.0/24"
}
variable "key_name" {
  description = "The name of an EC2 key pair that will be used to connect to instances"
}
variable "db_username" {
  description = "The username for the database user"
  default     = "root"
}
variable "db_password" {
  description = "The password for the database user"
  default     = "supersecretpassword"
}
variable "db_instance_type" {
  description = "The type of EC2 instance to use for the database"
  default     = "t2.micro"
}
variable "db_volume_size" {
  description = "The size of the EBS volume in GB to use for the database"
  default     = 8
}
variable "db_port" {
  description = "The port that the database should listen on"
  default     = "3306"
}
