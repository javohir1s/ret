package postgres

import (
	"database/sql"
	"essy_travel/models"
	"essy_travel/pkg/helpers"
	"fmt"

	"github.com/google/uuid"
)

type CityRepo struct {
	db *sql.DB
}

func NewCityRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (p *CityRepo) Create(req models.CreateCity) (*models.City, error) {

	query := `
		INSERT INTO cities(
			guid,
			title,
			country_id,
			city_code,
			latitude,
			longitude,
			"offset",
			timezone_id,
			country_name
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	id := uuid.New().String()
	_, err := p.db.Exec(query,
		id,
		req.Title,
		helpers.NewNullString(req.CountryId),
		req.CityCode,
		req.Latitude,
		req.Longitude,
		req.Offset,
		helpers.NewNullString(req.TimezoneId),
		req.CountryName,
	)
	if err != nil {
		return &models.City{}, err
	}

	return p.GetById(models.CityPrimaryKey{Id: id})
}

func (c *CityRepo) GetById(req models.CityPrimaryKey) (*models.City, error) {
	query := `
		SELECT
			guid,
			title,
			country_id,
			city_code,
			latitude,
			longitude,
			"offset",
			timezone_id,
			country_name
		FROM cities
		WHERE guid = $1
	`

	var city models.City
	err := c.db.QueryRow(query, req.Id).Scan(
		&city.Guid,
		&city.Title,
		&city.CountryId,
		&city.CityCode,
		&city.Latitude,
		&city.Longitude,
		&city.Offset,
		&city.TimezoneId,
		&city.CountryName,
	)
	if err != nil {
		return nil, err
	}

	return &city, nil
}

func (c *CityRepo) GetList(req models.GetListCityRequest) (*models.GetListCityResponse, error) {
	var (
		resp   = models.GetListCityResponse{}
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			guid,
			title,
			country_id,
			city_code,
			latitude,
			longitude,
			"offset",
			timezone_id,
			country_name
		FROM cities
	`
	query += where + limit + offset

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		err = rows.Scan(
			&resp.Count,
			&city.Guid,
			&city.Title,
			&city.CountryId,
			&city.CityCode,
			&city.Latitude,
			&city.Longitude,
			&city.Offset,
			&city.TimezoneId,
			&city.CountryName,
		)
		if err != nil {
			return nil, err
		}

		resp.Cities = append(resp.Cities, city)
	}

	return &resp, nil
}

func (c *CityRepo) Update(req models.UpdateCity) (*models.City, error) {

	_, err := c.db.Exec(`UPDATE cities SET guid=$1, title=$2, country_id=$3, city_code=$4, latitude=$5, longitude=$6, "offset"=$7, timezone_id=$8, country_name=$ WHERE guid = $10`, req.Guid)
	if err != nil {
		return nil, err
	}

	return &models.City{}, nil
}

func (c *CityRepo) Delete(req models.CityPrimaryKey) (string, error) {

	_, err := c.db.Exec(`DELETE FROM cities WHERE guid = $1`, req.Id)

	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (c *CityRepo) Search(req models.City) (*models.GetListCityResponse, error) {
	var cities models.GetListCityResponse
	rows, err := c.db.Query(`SELECT COUNT(*) OVER(), guid, title, country_id, city_code, latitude, longitude, "offset", timezone_id, country_nam FROM cities WHERE title = $1`, req.Title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		err := rows.Scan(
			&cities.Count,
			&city.Guid,
			&city.Title,
			&city.CountryId,
			&city.CityCode,
			&city.Latitude,
			&city.Longitude,
			&city.Offset,
			&city.TimezoneId,
			&city.CountryName,
		)
		if err != nil {
			return nil, err
		}
		cities.Cities = append(cities.Cities, city)
	}

	return &cities, nil
}
