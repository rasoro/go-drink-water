package main

type Notification struct {
	Icon  string
	Title string
	Text  string
	Sound string
}

func NewNotification(title, text string) *Notification {
	config := GetConfig()
	return &Notification{
		Title: title,
		Text:  text,
		Icon:  config.IconPath,
		Sound: config.SoundPath,
	}
}

func (n *Notification) Notify() {
	OsNotify(n.Title, n.Text, n.Icon)
	OsSound(n.Sound)
}
