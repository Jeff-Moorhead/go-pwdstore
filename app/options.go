package app

import (
	"flag"
)

type Options struct {
	Init     bool
	File     string
	Key      string
	Master   string
	Title    string
	Password string
	Add      bool
	Set      bool
	Remove   bool
	All      bool
}

func GetOptions() Options {
	init := flag.Bool("init", false, "Initialize a new manager and create a new encryption key")
	file := flag.String("file", "", "Path to the password file, default is ~/.pwdmgr/.store.json")
	key := flag.String("key", "", "Path to the password key file, default is ~/.pwdmgr/.store.key")
	master := flag.String("master", "", "Path to the master password file, default is ~/.pwdmgr/.master")
	title := flag.String("title", "", "Title for this password, e.g. www.facebook.com")
	password := flag.String("password", "", "Password, only used when adding or modifying")
	add := flag.Bool("add", false, "Enables adding a new password")
	set := flag.Bool("set", false, "Enables modifying an existing password")
	remove := flag.Bool("remove", false, "Enables removing a password")
	all := flag.Bool("all", false, "View all titles")
	flag.Parse()

	return Options{
		Init:     *init,
		File:     *file,
		Key:      *key,
		Master:   *master,
		Title:    *title,
		Password: *password,
		Add:      *add,
		Set:      *set,
		Remove:   *remove,
		All:      *all,
	}
}
