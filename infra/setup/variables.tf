variable "tf_state_bucket" {
  description = "name of s3 bucket in AWS for storing tf state"
  default     = "devops-training-app-tf-state"
}

variable "tf_state_lock_table" {
  description = "name of the dynamoDb table for tf state locking "
  default     = "devops-training-app-tf-lock"
}

variable "project" {
  description = "project name for tagging resources"
  default     = "training-app"
}

variable "contact" {
  description = "contact name for tagging resources"
  default     = "kacdevtest@gmail.com"
}

