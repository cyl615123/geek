package server

import "github.com/cyl615123/geek/graduation_project/api/client"

func Create() {
	client.Create()
}

func Cancel() {
	client.Update()
}

func List() {
	client.GetList()
}

func Detail() {
	client.GetDetail()
}
