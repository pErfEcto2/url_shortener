package memory

import (
	"errors"
	"maps"
	"slices"

	"github.com/pErfEcto2/url_shortener/internal/models"
	"github.com/pErfEcto2/url_shortener/internal/shortener"
)

var users []models.User = []models.User{
	{Username: "system"}, // to store shortened urls from the index page
}

type MemoryDB struct{}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}

func (m *MemoryDB) OriginalURLByShortened(shortenedURL string) (string, error) {
	for _, u := range users {
		for k, v := range u.Urls {
			if v == shortenedURL {
				return k, nil
			}
		}
	}
	return "", errors.New("no such shortened url")
}

func (m *MemoryDB) AddUser(user models.User) error {
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("user already exists")
		}
	}

	if user.Username == "" || user.Password == "" {
		return errors.New("invalid user")
	}

	users = append(users, user)

	return nil
}

func (m *MemoryDB) Users() []models.User {
	return users
}

func (m *MemoryDB) isUnique(url string) bool {
	for _, u := range users {
		if slices.Contains(slices.Collect(maps.Values(u.Urls)), url) {
			return false
		}
	}
	return true
}

func (m *MemoryDB) HasUrl(url string) bool {
	for _, u := range users {
		if slices.Contains(slices.Collect(maps.Keys(u.Urls)), url) {
			return true
		}
	}
	return false
}

func (m *MemoryDB) ShortenedUrlByUrl(url string) (string, bool) {
	for _, u := range users {
		if v, ok := u.Urls[url]; ok {
			return v, true
		}
	}
	return "", false
}

func (m *MemoryDB) AddUrlToUser(url string, user models.User) (string, bool) {
	if !m.HasUserByUsername(user.Username) {
		return "", false
	}

	var shortenedUrl string
	for i, u := range users {
		if u.Username != user.Username {
			continue
		}
		if _, ok := u.Urls[url]; ok {
			return "", false
		}

		shortenedUrl = shortener.ShortenUrl(url)
		for {
			if !m.isUnique(shortenedUrl) {
				shortenedUrl = shortener.ShortenUrl(url)
				continue
			}
			break
		}

		if users[i].Urls == nil {
			users[i].Urls = make(map[string]string)
		}

		users[i].Urls[url] = shortenedUrl

		break
	}

	return shortenedUrl, true
}

func (m *MemoryDB) UrlsByUsername(username string) (map[string]string, error) {
	for _, u := range users {
		if u.Username == username {
			return u.Urls, nil
		}
	}
	return nil, errors.New("no such user")
}

func (m *MemoryDB) UserByUsername(username string) models.User {
	for _, u := range users {
		if u.Username == username {
			return u
		}
	}
	return models.User{}
}

func (m *MemoryDB) HasUserByUsername(username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}
