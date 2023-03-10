package dto

import "dewetour/models"

// membuat struct untuk digunakan sebagai tipe data response trip
type TripResponse struct {
	ID             int                    `json:"id"`
	Title          string                 `json:"title"`
	Country        models.CountryResponse `json:"country"`
	Accomodation   string                 `json:"accomodation"`
	Transportation string                 `json:"transportation"`
	Eat            string                 `json:"eat"`
	Day            int                    `json:"day"`
	Night          int                    `json:"night"`
	DateTrip       string                 `json:"dateTrip"`
	Price          int                    `json:"price"`
	TotalQuota     int                    `json:"totalQuota"`
	QuotaAvailable int                    `json:"quotaAvailable"`
	Description    string                 `json:"description"`
	Images         []string               `json:"images"`
}

// membuat struct untuk digunakan sebagai tipe data request saat menambahkan country
type AddTripRequest struct {
	// ID             int             `json:"id" validate:"required"`
	Title          string   `json:"title" validate:"required" form:"title"`
	CountryID      int      `json:"id_country" validate:"required" form:"id_country"`
	Accomodation   string   `json:"accomodation" validate:"required" form:"accomodation"`
	Transportation string   `json:"transportation" validate:"required" form:"transportation"`
	Eat            string   `json:"eat" validate:"required" form:"eat"`
	Day            int      `json:"day" validate:"required" form:"day"`
	Night          int      `json:"night" validate:"required" form:"night"`
	DateTrip       string   `json:"dateTrip" validate:"required" form:"dateTrip"`
	Price          int      `json:"price" validate:"required" form:"price"`
	TotalQuota     int      `json:"quota" validate:"required" form:"quota"`
	Description    string   `json:"description" validate:"required" form:"description"`
	Images         []string `json:"image" validate:"required" form:"images"`
}

// membuat struct untuk digunakan sebagai tipe data request saat mengupdate country
type UpdateTripRequest struct {
	// ID             int      `json:"id"`
	Title          string   `json:"title" form:"title"`
	CountryID      int      `json:"country" form:"id_country"`
	Accomodation   string   `json:"accomodation" form:"accomodation"`
	Transportation string   `json:"transportation" form:"transportation"`
	Eat            string   `json:"eat" form:"eat"`
	Day            int      `json:"day" form:"day"`
	Night          int      `json:"night" form:"night"`
	DateTrip       string   `json:"dateTrip" form:"dateTrip"`
	Price          int      `json:"price" form:"price"`
	TotalQuota     int      `json:"quota" form:"quota"`
	Description    string   `json:"description" form:"description"`
	Images         []string `json:"image" form:"images"`
}
