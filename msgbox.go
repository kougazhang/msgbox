package msgbox

type Msg interface {
	Send(msg any) error
}
