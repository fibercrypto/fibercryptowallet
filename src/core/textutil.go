package core

// PasswordReader secure retrieval of passwords from users
type PasswordReader func(string) (string, error)
