package postgres

import (
	"database/sql"
	"essy_travel/models"
	"fmt"

	"github.com/google/uuid"
)

type AirportRepo struct {
	db *sql.DB
}

func NewAirportRepo(db *sql.DB) *AirportRepo {
	return &AirportRepo{
		db: db,
	}
}

func (a *AirportRepo) Create(req models.CreateAirport) (*models.Airport, error) {
	query := `
		INSERT INTO buildings(
			guid,
			title,
			country_id,
			city_id,
			latitude,
			longitude,
			radius,
			image,
			address,
			timezone_id,
			country,
			city,
			search_text,
			code,
			product_count,
			gmt
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`

	id := uuid.New().String()
	_, err := a.db.Exec(query,
		id,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Adress,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
	)
	if err != nil {
		return &models.Airport{}, err
	}

	return a.GetById(models.AirportPrimaryKey{Id: id})
}

func (a *AirportRepo) GetById(req models.AirportPrimaryKey) (*models.Airport, error) {
	query := `
		SELECT
			guid,
			title,
			country_id,
			city_id,
			latitude,
			longitude,
			radius,
			image,
			address,
			timezone_id,
			country,
			city,
			search_text,
			code,
			product_count,
			gmt
		FROM buildings
		WHERE guid = $1
	`

	var airport models.Airport
	err := a.db.QueryRow(query, req.Id).Scan(
		&airport.Id,
		&airport.Title,
		&airport.CountryId,
		&airport.CityId,
		&airport.Latitude,
		&airport.Longitude,
		&airport.Radius,
		&airport.Image,
		&airport.Adress,
		&airport.TimezoneId,
		&airport.Country,
		&airport.City,
		&airport.SearchText,
		&airport.Code,
		&airport.ProductCount,
		&airport.Gmt,
	)
	if err != nil {
		return nil, err
	}

	return &airport, nil
}

func (a *AirportRepo) GetList(req models.GetListAirportRequest) (*models.GetListAirportResponse, error) {
	var (
		resp   = models.GetListAirportResponse{}
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
			city_id,
			latitude,
			longitude,
			radius,
			image,
			address,
			timezone_id,
			country,
			city,
			search_text,
			code,
			product_count,
			gmt
		FROM buildings
	`
	query += where + limit + offset

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airport models.Airport
		err = rows.Scan(
			&resp.Count,
			&airport.Id,
			&airport.Title,
			&airport.CountryId,
			&airport.CityId,
			&airport.Latitude,
			&airport.Longitude,
			&airport.Radius,
			&airport.Image,
			&airport.Adress,
			&airport.TimezoneId,
			&airport.Country,
			&airport.City,
			&airport.SearchText,
			&airport.Code,
			&airport.ProductCount,
			&airport.Gmt,
		)
		if err != nil {
			return nil, err
		}

		resp.Airports = append(resp.Airports, airport)
	}

	return &resp, nil
}

func (a *AirportRepo) Update(req models.UpdateAirport) (*models.Airport, error) {
	_, err := a.db.Exec(`
		UPDATE buildings 
		SET title=$1, country_id=$2, city_id=$3, latitude=$4, longitude=$5, 
		    radius=$6, image=$7, address=$8, timezone_id=$9, country=$10,
		    city=$11, search_text=$12, code=$13, product_count=$14, gmt=$15
		WHERE guid = $16`,
		req.Title,
		req.CountryId,
		req.CityId,
		req.Latitude,
		req.Longitude,
		req.Radius,
		req.Image,
		req.Adress,
		req.TimezoneId,
		req.Country,
		req.City,
		req.SearchText,
		req.Code,
		req.ProductCount,
		req.Gmt,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &models.Airport{}, nil
}

func (a *AirportRepo) Delete(req models.AirportPrimaryKey) (string, error) {
	_, err := a.db.Exec(`DELETE FROM buildings WHERE guid = $1`, req.Id)
	if err != nil {
		return "Does not delete", err
	}

	return "Deleted", nil
}

func (a *AirportRepo) Search(req models.Airport) (*models.GetListAirportResponse, error) {
	var airports models.GetListAirportResponse
	rows, err := a.db.Query(`
		SELECT COUNT(*) OVER(), guid, title, country_id, city_id, latitude, longitude, 
		       radius, image, address, timezone_id, country, city, search_text, code,
		       product_count, gmt
		FROM buildings 
		WHERE title = $1`,
		req.Title,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var airport models.Airport
		err := rows.Scan(
			&airports.Count,
			&airport.Id,
			&airport.Title,
			&airport.CountryId,
			&airport.CityId,
			&airport.Latitude,
			&airport.Longitude,
			&airport.Radius,
			&airport.Image,
			&airport.Adress,
			&airport.TimezoneId,
			&airport.Country,
			&airport.City,
			&airport.SearchText,
			&airport.Code,
			&airport.ProductCount,
			&airport.Gmt,
		)
		if err != nil {
			return nil, err
		}
		airports.Airports = append(airports.Airports, airport)
	}

	return &airports, nil
}
