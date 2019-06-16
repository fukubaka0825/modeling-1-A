package domain

import (
	"time"
)

type MinimumReservedRoom struct {
	ID                   int
	MinimumReservedTime  MinimumReservedTime
	StartDateTime        ReservedTime
	EndDateTime          ReservedTime
	RoomID               int
}

type MinimumReservedRooms []MinimumReservedRoom

type MinimumReservedTime time.Duration

const MINIMUM_RESERVED_TIME MinimumReservedTime = MinimumReservedTime(60 * time.Minute)

func(m MinimumReservedRooms)isEmpty(startDateTime  ReservedTime,endDateTime  ReservedTime,roomID int) bool{
	reservedMinitues := time.Time(endDateTime).Sub(time.Time(startDateTime)).Minutes()
	minimumReservedTimeNumber := reservedMinitues / float64(MINIMUM_RESERVED_TIME)
	for minimumReservedTimeNumber < 1{
		startDateTime = ReservedTime(time.Time(startDateTime).Add(time.Duration(MINIMUM_RESERVED_TIME)))
		endDateTime = ReservedTime(time.Time(startDateTime).Add(time.Duration(MINIMUM_RESERVED_TIME)))

		for _,minimumReservationRoom := range m{
			if minimumReservationRoom.isEmpty(startDateTime,endDateTime,roomID){
				return false
			}
		}

		minimumReservedTimeNumber--
	}
	return true
}

func(m MinimumReservedRoom)isEmpty(startDateTime  ReservedTime,endDateTime  ReservedTime,roomID int) bool {
	if startDateTime == m.StartDateTime && endDateTime == m.EndDateTime && roomID == m.RoomID{
		return false
	}
	return true
}

