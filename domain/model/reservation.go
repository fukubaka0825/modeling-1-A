package domain

import (
    "github.com/labstack/gommon/log"
    "github.com/pkg/errors"
    "time"
    "./enum/Priority"
)

// Reservation 予約
type Reservation struct {
	id int
	reservationPeriod ReservationPeriod
	roomID int
    userID int
    priority Priority
}

// Reservations 予約リスト
type Reservations struct {
    reservations := [...]Reservation
}

// Period 期間
type Period struct {
    StartDataTime time.Time
    EndDataTime time.Time
}

// ReservationPeriod 予約期間
type ReservationPeriod struct {
    const StartHour = 9
    const EndHour = 18
    period Period
}

func newReservationPeriod(startAt time.Time, endAd time.Time) (*ReservationPeriod, error) {
    if startAt > endAt && StartHour > startAt && EndHour < endAt {
        return nill, errors.New('invalid Param.')
    }
    reservationPeriod := new(ReservationPeriod)
    reservationPeriod.period := Period{statAt, endAt}
    return reservationPeriod, nill
}

func (r ReservedTime) validate() error{
    meetingHour := time.Time(r).Hour()
    if START_RESERVED_HOUR >= meetingHour && END_RESERVED_HOUR <= meetingHour {
        err := errors.New(`INVALID  ReservedTime`)
        log.Error(err)
        return err
    }
    return nil
}
