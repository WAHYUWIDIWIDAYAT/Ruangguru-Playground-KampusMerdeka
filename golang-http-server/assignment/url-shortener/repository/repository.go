package repository

import (
	"sync"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/url-shortener/entity"
)

type URLRepository struct {
	mu   sync.Mutex
	Data map[string]string
}

func NewMapRepository() URLRepository {
	data := make(map[string]string, 0)
	return URLRepository{
		Data: data,
	}
}

func (r *URLRepository) Get(path string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	url := r.Data[path]

	if url == "" {
		return nil, entity.ErrURLNotFound
	}

	return &entity.URL{
		LongURL:  url,
		ShortURL: path,
	}, entity.ErrCustomURLIsExists // TODO: replace this
}

func (r *URLRepository) Create(longURL string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	shortUrl := entity.GetRandomShortURL(longURL)
	r.Data[shortUrl] = longURL
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: shortUrl,
	}, nil // TODO: replace this
}

func (r *URLRepository) CreateCustom(longURL, customPath string) (*entity.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Data[customPath] = longURL
	return &entity.URL{
		LongURL:  longURL,
		ShortURL: customPath,
	}, nil // TODO: replace this
}
