package core

type PasswordReader func(message string) (string, error)
