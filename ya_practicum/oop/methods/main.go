package main

import "fmt"

type DeliveryState string

const (
	DeliveryStatePending   DeliveryState = "pending"
	DeliveryStateAck       DeliveryState = "acknowledged"
	DeliveryStateProcessed DeliveryState = "processed"
	DeliveryStateCanceled  DeliveryState = "canceled"
)

func main() {
	if err := HandleMsgDeliveryStatus(DeliveryState("fake")); err != nil {
		panic(err)
	}
}

func (s DeliveryState) IsValid() bool {
	switch s {
	case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
		return true
	default:
		return false
	}
}

func (s DeliveryState) String() string {
	return string(s)
}

func HandleMsgDeliveryStatus(status DeliveryState) error {
	if !status.IsValid() {
		return fmt.Errorf("status: invalid")
	}

	return nil
}
