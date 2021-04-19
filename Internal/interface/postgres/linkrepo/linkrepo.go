package linkrepo

import (
	"URezL/Internal/domain"
	"URezL/Internal/domain/link"
	"database/sql"
)

type Postgres struct{
	conn *sql.DB
}

func New(conn *sql.DB) *Postgres{
	return &Postgres{conn: conn}
}

const queryAddLink = `
	INSERT INTO oldLinkByNewLink(
		newLink,
		oldLink
	) VALUES ($1, $2)
`

const queryAddLinkWithUser = `
	INSERT INTO oldLinkByNewLink(
		newLink,
		oldLink,
	    lifetime,
		accountId
	) VALUES ($1, $2, $3, $4)
`

const queryAddUserLink = `
	INSERT INTO linksByAccountId(
		accountId,
		newLink,
		oldLink
	) VALUES ($1, $2, $3)
`

func (p *Postgres) AddLink(lnk link.Link) (link.Link, error) {
	if lnk.AccountId == nil {
		row := p.conn.QueryRow(queryAddLink, lnk.ShortenLink, lnk.Link)
		err := row.Scan()
		if err != sql.ErrNoRows {
			return link.Link{}, domain.ErrAlreadyExist
		}
	} else {
		row := p.conn.QueryRow(queryAddLinkWithUser, lnk.ShortenLink, lnk.Link, nil, lnk.AccountId)
		err := row.Scan()
		if err != sql.ErrNoRows {
			return link.Link{}, domain.ErrAlreadyExist
		}
	}

	if lnk.AccountId != nil {
		row := p.conn.QueryRow(queryAddUserLink, lnk.AccountId, lnk.ShortenLink, lnk.Link)
		err := row.Scan()
		if err != sql.ErrNoRows {
			return link.Link{}, err
		}
	}
	return lnk, nil
}

const queryGetLink = `
	SELECT 
		newLink,
		oldLink,
	    lifetime,
	    accountId
	FROM oldLinkByNewLink
	WHERE newLink = $1
`

func (p *Postgres) GetLinkByShorten(lnk string) (link.Link, error) {
	l := link.Link{}
	row := p.conn.QueryRow(queryGetLink, lnk)
	err := row.Scan(&l.ShortenLink, &l.Link, &l.Lifetime, &l.AccountId)
	if err != nil{
		return link.Link{}, err
	}
	return l, nil
}

const queryGetLinksByAccount = `
	SELECT 
		newLink,
	    oldLink
	FROM linksByAccountId
	WHERE accountId = $1
`

func (p *Postgres) GetLinksByAccountId(accountId string) ([]link.Link, error) {
	var links = make([]link.Link, 0)
	row, err := p.conn.Query(queryGetLinksByAccount, accountId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var newLink string
		var oldLink string
		if err := row.Scan(&newLink, &oldLink); err != nil {
			return nil, err
		}
		links = append(links, link.Link{Link: oldLink, ShortenLink: newLink})
	}
	return links, nil
}

const queryCheckLink = `
	SELECT 1 
	FROM oldLinkByNewLink
	WHERE newLink = $1
`

func (p *Postgres) CheckLinkExists(lnk string) bool {
	row := p.conn.QueryRow(queryCheckLink, lnk)
	err := row.Scan()
	if err == sql.ErrNoRows {
		return false
	}
	return true
}

const queryDeleteLink1 = `
	DELETE FROM oldLinkByNewLink
	WHERE newLink = $1 AND accountId = $2
`
const queryDeleteLink2 = `
	DELETE FROM linksByAccountId
	WHERE accountId = $2 AND newLink = $1
`

func (p *Postgres) DeleteLink(lnk string, accountId string) error {
	row := p.conn.QueryRow(queryDeleteLink1, lnk, accountId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		return err
	}
	row = p.conn.QueryRow(queryDeleteLink2, lnk, accountId)
	err = row.Scan()
	if err != sql.ErrNoRows {
		return err
	}
	return nil
}
const queryChangeLink1 = `
	UPDATE oldLinkByNewLink
	SET newLink = $2
	WHERE newLink = $1 and accountId = $3
`
const queryChangeLink2 = `
	UPDATE linksByAccountId
	SET newLink = $2
	WHERE newLink = $1 AND accountId = $3
`

func (p *Postgres) ChangeLink(oldLink string, newLink string, accountId string) error {
	row := p.conn.QueryRow(queryChangeLink1, oldLink, newLink, accountId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		return err
	}
	row = p.conn.QueryRow(queryChangeLink2, oldLink, newLink, accountId)
	err = row.Scan()
	if err != sql.ErrNoRows {
		return err
	}
	return nil
}
