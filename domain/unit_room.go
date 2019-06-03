package domain

import "github.com/golang/protobuf/ptypes/timestamp"

type UnitRoom struct{
    Id       int
    UnitTime timestamp.Timestamp
    RoomId   int

}
