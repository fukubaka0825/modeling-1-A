package domain

import "github.com/golang/protobuf/ptypes/timestamp"

type Reservation struct{
    Id            int
    Start_date    timestamp.Timestamp
    End_date      timestamp.Timestamp
}
