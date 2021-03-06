package linkrepo

import (
	"URezL/Internal/domain"
	"URezL/Internal/domain/link"
	"sync"
)

type Memory struct {
	oldLinkByNewLink    map[string]link.Link
	linksByAccountId 	map[string]map[string]link.Link
	mu              	*sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		oldLinkByNewLink:   make(map[string]link.Link),
		linksByAccountId: 	make(map[string]map[string]link.Link),
		mu:              	&sync.Mutex{},
	}
}

func (m *Memory) AddLink (lnk link.Link) (link.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.oldLinkByNewLink[lnk.ShortenLink]; ok {
		return link.Link{}, domain.ErrAlreadyExist
	}
	m.oldLinkByNewLink[lnk.ShortenLink] = lnk
	if lnk.AccountId != nil {
		accountLinks, ok := m.linksByAccountId[*lnk.AccountId]
		if !ok {
			accountLinks = make(map[string]link.Link)
		}
		accountLinks[lnk.ShortenLink] = lnk
		m.linksByAccountId[*lnk.AccountId] = accountLinks
	}
	return  lnk, nil
}

func (m *Memory) GetLinkByShorten (lnk string) (link.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.oldLinkByNewLink[lnk]
	if !ok {
		return link.Link{}, domain.ErrNotFound
	}
	return l, nil
}

func (m *Memory) GetLinksByAccountId(accountId string) ([]link.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	lnks, ok := m.linksByAccountId[accountId]
	if !ok {
		return nil, domain.ErrNotFound
	}
	links := make([]link.Link, 0, len(lnks))
	for _, val := range lnks {
		links = append(links, val)
	}
	return links, nil
}

func (m *Memory) CheckLinkExists(lnk string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.oldLinkByNewLink[lnk]
	return ok
}

func (m *Memory) DeleteLink(lnk string, accountId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.oldLinkByNewLink[lnk]
	if !ok {
		return domain.ErrNotFound
	}
	if l.AccountId == nil || *l.AccountId != accountId {
		return domain.ErrPermissionDenied
	}
	delete(m.oldLinkByNewLink, lnk)
	delete(m.linksByAccountId[*l.AccountId], lnk)
	return nil
}

// TODO: think of better implementation
func (m *Memory) ChangeLink(oldLink string, newLink string, accountId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.oldLinkByNewLink[oldLink]
	if !ok {
		return domain.ErrNotFound
	}
	_, ok = m.oldLinkByNewLink[newLink]
	if ok {
		return domain.ErrAlreadyExist
	}
	if l.AccountId == nil || *l.AccountId != accountId {
		return domain.ErrPermissionDenied
	}
	delete(m.oldLinkByNewLink, oldLink)
	delete(m.linksByAccountId[*l.AccountId], oldLink)
	changedLink := link.Link{
		Link:        l.Link,
		ShortenLink: newLink,
		AccountId:   l.AccountId,
		Lifetime:    l.Lifetime,
	}
	m.oldLinkByNewLink[newLink] = changedLink
	accountLinks, ok := m.linksByAccountId[*changedLink.AccountId]
	if !ok {
		accountLinks = make(map[string]link.Link)
	}
	accountLinks[changedLink.ShortenLink] = changedLink
	m.linksByAccountId[*changedLink.AccountId] = accountLinks
	return nil
}

func (m *Memory) SetNotWorking(lnk string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.oldLinkByNewLink[lnk]
	if !ok {
		return domain.ErrNotFound
	}
	m.oldLinkByNewLink[lnk] = link.Link{
		Link:        l.Link,
		ShortenLink: l.ShortenLink,
		AccountId:   l.AccountId,
		Lifetime:    l.Lifetime,
		Working:       true,
	}
	return nil
}
