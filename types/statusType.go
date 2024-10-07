package types

type StatusType string

const (
	Pending   StatusType = "В ожидании"
	Paid      StatusType = "Оплачен"
	Shipped   StatusType = "Отправлен"
	Delivered StatusType = "В доставке"
	Canceled  StatusType = "Отменен"
	Returned  StatusType = "Возврат"
)
