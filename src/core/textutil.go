package core

// PasswordReadeer secure retrieval of passwords from users
type PasswordReader func(message string) (string, error)
