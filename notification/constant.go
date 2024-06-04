package notification

type MsgAction string

const (
	MsgActionCreate MsgAction = "Photo Manager :mega: `Create`"
	MsgActionDelete MsgAction = "Photo Manager :mega: `Delete`"
)

type StatusColorAction string

const (
	StatusColorActionSuccess = "#22bb33"
	StatusColorActionError   = "#bb2124"
)
