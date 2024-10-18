package repository

import (
	"FinalProject/internal/models"
	"FinalProject/pkg/errors"
)

func (r *Repository) AddRoute(route *models.Route) (*models.Route, error) {
	query := `INSERT into routes (from_city,to_city,price,date,driver_id,car_id)
	values(?,?,?,?,?,?)`
	err := r.db.Exec(query, route.FromCity, route.ToCity, route.Price, route.Date, route.DriverId, route.CarId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to insert route, err: ", err)
		return nil, errors.ErrFailedToCreateRoute
	}
	return route, nil
}

func (r *Repository) GetRoutes() ([]models.Route, error) {
	var routes []models.Route
	query := `select * from routes`
	err := r.db.Raw(query).Scan(&routes).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select routes, err: ", err)
		return nil, errors.ErrFailedToGet
	}
	return routes, nil
}

func (r *Repository) GetRouteByID(id int) (*models.Route, error) {
	var route models.Route
	query := `select * from routes where route_id = ?`
	err := r.db.Raw(query, id).Scan(&route).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to select route by id, err: ", err)
		return nil, errors.ErrFailedToGet
	}
	return &route, nil
}

func (r *Repository) FinishRoute(u models.Finish) error {
	query := `update seats set available = true where car_id = ?`
	err := r.db.Exec(query, u.CarId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to update seats, err: ", err)
		return err
	}

	query = `update booking set status_id = 1 where route_id = ?`
	err = r.db.Exec(query, u.CarId).Error
	if err != nil {
		r.logger.Error("Faced an error while tried to update bookings, err: ", err)
		return err
	}

	return nil
}
