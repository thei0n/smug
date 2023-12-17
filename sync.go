package main

import(
	"github.com/go-git/go-git/v5"
	"os"
	"syscall"
)

func sync(packageName string) int {

	const baseUrl = "https://aur.archlinux.org/"
	const suffix = ".git"

	var home = os.Getenv("HOME")
	var finalUrl = baseUrl+packageName+suffix
	var args = []string{"makepkg", "-si"}
	var parentPath = home + "/.cache/smug/"
	var directoryPath = parentPath + packageName

	os.RemoveAll(directoryPath);

	errDir := os.MkdirAll(directoryPath, 0755)

	if errDir != nil {
		panic("Error: creating directory")
	}

	_, errGit := git.PlainClone(directoryPath, false, &git.CloneOptions{
	    URL:      finalUrl,
	    Progress: os.Stdout,
	})

	if errGit != nil{
		panic ("Error: git clone unsuccessful")
	}
	
	os.Chdir(directoryPath);
	syscall.Exec("/usr/bin/makepkg", args, os.Environ())

return 0;
}
