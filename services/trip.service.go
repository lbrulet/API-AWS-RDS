package services

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/models"
)

func GetTrips(c *gin.Context, id string) {
	db := database.DBManager.DB
	var trips []models.Trip
	var rows *sql.Rows
	var err error
	if len(id) > 0 {
		rows, err = db.Query("SELECT * FROM Trip WHERE id_user = " + id)
	} else {
		rows, err = db.Query("SELECT * FROM Trip")
	}
	if err != nil {
		// handle this error better than this
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
	}
	defer rows.Close()
	for rows.Next() {
		var trip models.Trip
		err = rows.Scan(&trip.ID, &trip.StartLat, &trip.StartLng, &trip.EndLat, &trip.EndLng, &trip.IDUser)
		if err != nil {
			// handle this error
			panic(err)
		}
		trips = append(trips, trip)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err})
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": trips})
}

func NewTrip(c *gin.Context, payload models.Trip) {
	db := database.DBManager.DB

	//Insert new user
	_, err := db.Exec(`INSERT INTO Trip (start_lat, start_lng, end_lat, end_lng, id_user) VALUES (?, ?, ?, ?, ?)`, payload.StartLat, payload.StartLng, payload.EndLat, payload.EndLng, payload.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success"})
}
