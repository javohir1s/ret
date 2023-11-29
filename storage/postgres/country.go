package postgres

import (
	"database/sql"
	"essy_travel/models"

	"github.com/google/uuid"
)

type CountryRepo struct {
	db *sql.DB
}

func NewCountryRepo(db *sql.DB) *CountryRepo {
	return &CountryRepo{
		db: db,
	}
}

func (p *CountryRepo) Create(req models.CreateCountry) (*models.Country, error) {
	_, err := p.db.Exec(`INSERT INTO countries(guid, title, code, continent) VALUES ($1, $2, $3, $4)`, uuid.New().String(), req.Title, req.Code, req.Continent)
	if err != nil {
		return nil, err
	}
	return &models.Country{}, nil
}

func (c *CountryRepo) GetById(req models.CountryPrimaryKey) (*models.Country, error) {
	var country = models.Country{}
	err := c.db.QueryRow(`SELECT guid, title, code, continent FROM countries WHERE guid = $1`, req.Id).
		Scan(
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Continent,
		)
	if err != nil {
		return nil, err
	}

	return &country, nil
}

func (c *CountryRepo) GetList(req models.GetListCountryRequest) (*models.GetListCountryResponse, error) {
	var countries = models.GetListCountryResponse{}
	offset := req.Offset
	limit := req.Limit

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 {
		limit = 10
	}

	rows, err := c.db.Query(`SELECT COUNT(*) OVER(), guid, title, code, continent FROM countries`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err = rows.Scan(
			&countries.Count,
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Continent,
		)
		if err != nil {
			return nil, err
		}
		countries.Countries = append(countries.Countries, country)
	}

	return &countries, nil
}

func (c *CountryRepo) Update(req models.UpdateCountry) (*models.Country, error) {
	_, err := c.db.Exec(`UPDATE countries SET guid=$1, title=$2, code=$3, continent=$4 WHERE guid = $5`, req.Guid, req.Title, req.Code, req.Continent, req.Guid)
	if err != nil {
		return nil, err
	}

	return &models.Country{}, nil
}

func (c *CountryRepo) Delete(req models.CountryPrimaryKey) (string, error) {
	_, err := c.db.Exec(`DELETE FROM countries WHERE guid = $1`, req.Id)
	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CountryRepo) Search(req models.Country) (*models.GetListCountryResponse, error) {
	var countries models.GetListCountryResponse
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(), guid, title, code, continent FROM countries WHERE title ILIKE $1`, "%"+req.Title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var country models.Country
		err := rows.Scan(
			&countries.Count,
			&country.Guid,
			&country.Title,
			&country.Code,
			&country.Continent,
		)
		if err != nil {
			return nil, err
		}
		countries.Countries = append(countries.Countries, country)
	}

	return &countries, nil
}
