package domain

import (
    "github.com/labstack/gommon/log"
    "github.com/pkg/errors"
    "time"
)

type Reservation struct {
	ID            int
	StartDateTime ReservedTime
	EndDateTime   ReservedTime
	RoomID        int
	UserID        int
}

type ReservedTime time.Time

const START_RESERVED_HOUR = 9
const END_RESERVED_HOUR = 18



func (r ReservedTime) validate() error{
    meetingHour := time.Time(r).Hour()
    if START_RESERVED_HOUR >= meetingHour && END_RESERVED_HOUR <= meetingHour {
        err := errors.New(`INVALID  ReservedTime`)
        log.Error(err)
        return err
    }
    return nil
}
