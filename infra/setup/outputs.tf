output "cd_user_access_key_id" {
  description = "aws key id for cicd user"
  value       = aws_iam_access_key.cd.id
}

output "cd_user_access_key_secret" {
  description = "aws access key secret for cicd user"
  value       = aws_iam_access_key.cd.secret
  sensitive   = true
}



