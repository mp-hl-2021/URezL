package token

type Interface interface {
	IssueToken(userId string) (string, error)
	UserIdByToken(token string) (string, error)
}
