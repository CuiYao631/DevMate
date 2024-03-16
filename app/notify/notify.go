package notify

type Notify struct {
	Appid string
}

func MakeNotify() Notify {
	appId := "DevMate"
	//powershellCommand := "Get-StartApps -Name DevMate"
	//cmd := exec.Command("powershell.exe", "-Command", powershellCommand)
	//var stdout bytes.Buffer
	//cmd.Stdout = &stdout
	//err := cmd.Run()
	//if err != nil {
	//	log.Println(err)
	//}
	//if len(stdout.String()) > 0 {
	//	ss := strings.Fields(stdout.String())
	//	appId = ss[len(ss)-1]
	//} else {
	//	appId = "DevMate"
	//}
	return Notify{
		Appid: appId,
	}
}
func (n *Notify) Notify(title, message string) {
	NotifyMessage(n.Appid, title, message)
}
