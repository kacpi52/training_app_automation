variable "prefix" {
  description = "prefix for resources in aws"
  default     = "ta"
}

variable "project" {
  description = "project name for tagging resources "
  default     = "training-app"
}

variable "contact" {
  description = "contact email for tagging reosurces"
  default     = "kacdevtest@gmail.com"
}

variable "db_username" {
  description = "username for farmacy app api database"
  default     = "admin"
}
# this var will be injected from git secrets - added in compose 
variable "db_password" {
  description = "password for terraform db "

}


