terraform {
  required_providers {
    random = {
      source  = "hashicorp/random"
      version = "3.5.1"
    }
  }
}

variable "prefix" {
  type = string
  validation {
    condition     = can(regex("^[A-Za-z]+$", var.prefix))
    error_message = "Prefix can contain only alphabetic characters."
  }
}
