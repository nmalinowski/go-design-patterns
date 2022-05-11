package main

import "fmt"

func main() {
	// DONE: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()
	// DONE: Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.ico")
	bldr.SetSubTitle("This is a subtitle")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic Notification with some text in it")
	bldr.SetType("alert")
	// DONE: Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		fmt.Println("Error creating notification", err)
	} else {
		fmt.Printf("Notification: %+v", notif)
	}
	// END EXAMPLE
}
