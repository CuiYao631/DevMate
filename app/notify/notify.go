package notify

type Notify struct {
}

func MakeNotify() Notify {
	return Notify{}
}
func (n *Notify) Notify(title, message string) {
	//sysType := runtime.GOOS
	//
	//if sysType == "windows" {
	//	// windows系统
	//	WindowsNotify(title, message)
	//}
	//if sysType == "darwin" {
	//	// MAC系统
	//	MacNotify(title, message)
	//}
	NotifyMessage(title, message)
}
