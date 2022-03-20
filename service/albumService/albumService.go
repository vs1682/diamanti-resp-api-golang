package albumService

import (
	"errors"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums() (*[]Album, error) {
	return &albums, nil
}

func (a *Album) GetAlbum() (*Album, error) {
	for _, alb := range albums {
		if alb.ID == a.ID {
			return &alb, nil
		}
	}

	return nil, errors.New("Provided album id not found")
}

func (a *Album) PostAlbum() (*Album, error) {
	for _, alb := range albums {
		if alb.ID == a.ID {
			return nil, errors.New("Provided id already exits")
		}
	}

	albums = append(albums, *a)

	return a, nil
}

func (a *Album) UpdateAlbum() (*Album, error) {
	for i, alb := range albums {
		if alb.ID == a.ID {
			if a.Artist != "" {
				albums[i].Artist = a.Artist
			}

			if a.Price > 0 {
				albums[i].Price = a.Price
			}

			if a.Title != "" {
				albums[i].Title = a.Title
			}

			return &albums[i], nil
		}
	}

	return nil, errors.New("Provided album id not found")
}

func (a *Album) DeleteAlbum() (*[]Album, error) {
	idx := -1
	for i, alb := range albums {
		if alb.ID == a.ID {
			idx = i
			break
		}
	}

	if idx > -1 {
		albums[idx] = albums[len(albums)-1]
		albums = albums[:len(albums)-1]

		return &albums, nil
	}

	return nil, errors.New("Provided album id not found")
}
