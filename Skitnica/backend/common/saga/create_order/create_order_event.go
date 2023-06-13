package create_order

import "go.mongodb.org/mongo-driver/bson/primitive"

type Color struct {
	Code string
}

type Product struct {
	Id    string
	Color Color
}

type OrderItem struct {
	Product  Product
	Quantity uint16
}

type OrderDetails struct {
	Id      string
	Items   []OrderItem
	Address string
}

type CreateOrderCommandType int8

const (
	UpdateInventory CreateOrderCommandType = iota
	RollbackInventory
	ApproveOrder
	CancelOrder
	ShipOrder
	UnknownCommand
)

type CreateOrderCommand struct {
	Order OrderDetails
	Type  CreateOrderCommandType
}

type CreateOrderReplyType int8

const (
	InventoryUpdated CreateOrderReplyType = iota
	InventoryNotUpdated
	InventoryRolledBack
	OrderShippingScheduled
	OrderShippingNotScheduled
	OrderApproved
	OrderCancelled
	UnknownReply
)

type CreateOrderReply struct {
	Order OrderDetails
	Type  CreateOrderReplyType
}

type User struct {
	Id        primitive.ObjectID
	Username  string
	Password  string
	FirstName string
	LastName  string
	Role      string
	Address   string
}

type DeleteUserCommandType int8

const (
	CheckReservations DeleteUserCommandType = iota
	DeleteUser
	DeleteAccomodations
	Unknown
)

type DeleteUserCommand struct {
	User User
	Type DeleteUserCommandType
}

type DeleteUserReplyType int8

const (
	CanDelete DeleteUserReplyType = iota
	CanNotDelete
	DeletedGuest
	DeletedHost
	AccomodationsDeleted
	UnknownRep
)

type DeleteUserReply struct {
	User User
	Type DeleteUserReplyType
}
