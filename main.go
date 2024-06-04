package main

import (
	"github.com/PhotoManager/notification"
	"github.com/PhotoManager/tui"
)

func main() {
	notification.NewSlackClient()
	tui.CreateOptionList()
}
