package types

// ContextKey is a custom type that helps avoid
// context keys collission.
type ContextKey string

const APIURLKey ContextKey = "APIURL"
