package linkrepo

import (
	"URezL/Internal/domain/account"
	"URezL/Internal/domain/link"
	"sync"
)

type Memory struct {
	oldLinkByNewLink    map[string]link.Link
	linksByAccountId map[string]map[string]link.Link
	mu              *sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		oldLinkByNewLink:    make(map[string]link.Link),
		linksByAccountId: make(map[string]map[string]link.Link),
		mu:              &sync.Mutex{},
	}
}

func (m* Memory) AddLink (lnk link.Link) (link.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.oldLinkByNewLink[lnk.ShortenLink]; ok {
		return link.Link{}, account.ErrAlreadyExist
	}
	m.oldLinkByNewLink[lnk.ShortenLink] = lnk
	if lnk.UserId != nil {
		accountLinks, ok := m.linksByAccountId[*lnk.UserId]
		if !ok {
			accountLinks = make(map[string]link.Link)
		}
		accountLinks[lnk.ShortenLink] = lnk
		m.linksByAccountId[*lnk.UserId] = accountLinks
	}
	return  lnk, nil
}