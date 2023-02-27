package go_jwt

type ClaimsI interface {
	Void()
	GetExpirationTime() int64
}
