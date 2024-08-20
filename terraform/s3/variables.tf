variable "bucket_name" {
  type        = string
  default     = "kevindev-applications"
  description = "O nome do bucket S3 a ser criado."
}

variable "region" {
  type        = string
  default     = "us-east-1"
  description = "A região da AWS onde o bucket S3 será criado."
}

variable "lifecycle_transition_days" {
  type        = number
  default     = 30
  description = "Número de dias para mover objetos para o armazenamento Standard-IA."
}


variable "deep_archive_transition_days" {
  type        = number
  default     = 90
  description = "Número de dias para mover objetos para o armazenamento Glacier Deep Archive."
}

variable "expiration_days" {
  type        = number
  default     = 365
  description = "Número de dias para expirar objetos no bucket."
}

variable "user_name" {
  type        = string
  default     = "github_s3_user"
  description = "Nome do usuário IAM com acesso ao bucket."
}
