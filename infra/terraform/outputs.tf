output "bioinformatic_security_suite_load_balancer_endpoint" {
  value = "${aws_elb.bioinformatic_security_suite.dns_name}"
}
output "bioinformatic_security_suite_load_balancer_hostname" {
  value = "${aws_elb.bioinformatic_security_suite.hostname}"
}
output "bioinformatic_security_suite_load_balancer_arn" {
  value = "${aws_elb.bioinformatic_security_suite.arn}"
}
output "bioinformatic_security_suite_db_address" {
  value = "${aws_db_instance.bioinformatic_security_suite.address}"
}
output "bioinformatic_security_suite_db_port" {
  value = "${aws_db_instance.bioinformatic_security_suite.port}"
}
output "bioinformatic_security_suite_db_username" {
  value = "${var.db_username}"
}
output "bioinformatic_security_suite_db_password" {
  value = "${var.db_password}"
}
