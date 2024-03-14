package notify

type Notify struct {
}

func MakeNotify() Notify {
	return Notify{}
}
func (n *Notify) Notify(title, message string) {
	NotifyMessage(title, message)
}
