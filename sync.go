package main

import(
	// "strings"
	"github.com/go-git/go-git/v5"
	"os"
	"os/exec"
)

func sync(packageName string) int {

	//get the link name
	const baseUrl = "https://aur.archlinux.org/"
	const suffix = ".git"
	finalUrl := baseUrl+packageName+suffix

	//create the directory
	/*
	*Remove later
	*/
	const parentPath = "/home/theion/.cache/smug/"
	directoryPath := parentPath + packageName
	os.RemoveAll("/home/theion/.cache/smug");
	err := os.MkdirAll(directoryPath, 0755)

	if err != nil {
		panic("Error: creating directory")
	}

	_, err2 := git.PlainClone(directoryPath, false, &git.CloneOptions{
	    URL:      finalUrl,
	    Progress: os.Stdout,
	})

	if err2 != nil{
		panic ("Error: git clone unsuccessful")
	}
	//change directory
	os.Chdir(directoryPath);

	//makepkg -Si
	cmd:=exec.Command("makepkg","-si")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run();
	//os.RemoveAll("/home/theion/.cache/smug");

return 0;
}
