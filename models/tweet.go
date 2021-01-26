package models

import "time"

/*Tweet captura del Body, el mensaje que nos llega*/
type Tweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
