package linkrepo

import (
	"URezL/Internal/domain"
	"URezL/Internal/domain/link"
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

const queryAddLink = `
	INSERT INTO oldLinkByNewLink(
		newLink,
		oldLink
	) VALUES ($1, $2)
`

const queryAddUserLink = `
	INSERT INTO linksByAccountId(
		accountId,
		link
	) VALUES ($1, $2)
`

func (p *Postgres) AddLink(lnk link.Link) (link.Link, error) {
	row := p.conn.QueryRow(queryAddLink, lnk.ShortenLink, lnk.Link)
	err := row.Scan()
	if err != nil {
		return link.Link{}, domain.ErrAlreadyExist
	}
	if lnk.AccountId != nil {
		row := p.conn.QueryRow(queryAddUserLink, lnk.ShortenLink, lnk.AccountId)
		err := row.Scan()
		if err != nil {
			return link.Link{}, domain.ErrAlreadyExist
		}
	}
	return lnk, nil
}

func (Postgres) GetLinkByShorten(link string) (link.Link, error) {
	panic("implement me")
}

func (Postgres) GetLinksByAccountId(accountId string) ([]link.Link, error) {
	panic("implement me")
}

func (Postgres) CheckLinkExists(link string) bool {
	panic("implement me")
}

func (Postgres) DeleteLink(link string, accountId string) error {
	panic("implement me")
}

func (Postgres) ChangeLink(oldLink string, newLink string, accountId string) error {
	panic("implement me")
}
