package accountrepo

import (
	"URezL/Internal/domain/account"
	"database/sql"
)

type Postgres struct{
	conn *sql.DB
}

func New(connStr string) *Postgres{
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return &Postgres{conn: conn}
}

const queryCreateAccount = `
	INSERT INTO accounts(
		login,
		password
	) VALUES ($1, $2)
	RETURNING id
`

func (p *Postgres) CreateAccount(cred account.Credentials) (account.Account, error) {
	a := account.Account{Credentials: cred}
	row := p.conn.QueryRow(queryCreateAccount, cred.Login, cred.Password)
	err := row.Scan(&a.Id)
	if err != nil{
		return account.Account{}, err
	}
	return a, nil
}

const queryGetAccountById = `
	SELECT 
		id, 
		login,
		password
	FROM accounts
	WHERE id = $1
`

func (p *Postgres) GetAccountById(id string) (account.Account, error) {
	a := account.Account{}
	row := p.conn.QueryRow(queryGetAccountById, id)
	err := row.Scan(&a.Id, &a.Login, &a.Password)
	if err != nil{
		return account.Account{}, err
	}
	return a, nil
}

const queryGetAccountByLogin = `
	SELECT 
		id, 
		login,
		password
	FROM accounts
	WHERE login = $1
`
func (p *Postgres) GetAccountByLogin(login string) (account.Account, error) {
	a := account.Account{}
	row := p.conn.QueryRow(queryGetAccountByLogin, login)
	err := row.Scan(&a.Id, &a.Login, &a.Password)
	if err != nil{
		return account.Account{}, err
	}
	return a, nil
}
