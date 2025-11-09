package db

import (
	"errors"
	"maps"
	"slices"

	"github.com/pErfEcto2/url_shortener/internal/models"
	"github.com/pErfEcto2/url_shortener/internal/shortener"
)

var users []models.User = []models.User{
	{Username: "u", Password: "$2a$10$BGUqqy078P.YvRzljPjCBuCQVMz.stJmD0ywk8SRthpf1.2Egj0E2"}, // for testing
	{Username: "system"}, // to store shortened urls from the index page
}

func AddUser(user models.User) error {
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

func GetUsers() []models.User {
	return users
}

func isUnique(url string) bool {
	for _, u := range users {
		if slices.Contains(slices.Collect(maps.Values(u.Urls)), url) {
			return false
		}
	}
	return true
}

func HasUrl(url string) bool {
	for _, u := range users {
		if slices.Contains(slices.Collect(maps.Keys(u.Urls)), url) {
			return true
		}
	}
	return false
}

func GetShortenedUrlByUrl(url string) (string, bool) {
	for _, u := range users {
		if v, ok := u.Urls[url]; ok {
			return v, true
		}
	}
	return "", false
}

func AddUrlToUser(url string, user models.User) (string, bool) {
	if !HasUserByUsername(user.Username) {
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
			if !isUnique(shortenedUrl) {
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

func GetUrlsByUsername(username string) (map[string]string, error) {
	for _, u := range users {
		if u.Username == username {
			return u.Urls, nil
		}
	}
	return nil, errors.New("no such user")
}

func GetUserByUsername(username string) models.User {
	for _, u := range users {
		if u.Username == username {
			return u
		}
	}
	return models.User{}
}

func HasUserByUsername(username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}
