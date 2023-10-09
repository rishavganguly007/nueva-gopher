package main

import "fmt"

func  main()  {
	lan := make(map[string]string)

	lan["js"] = "javascript"
	lan["go"] = "golang"
	lan["py"] = "python"

	fmt.Println("List lan: ", lan["py"])

	delete(lan, "js")
	fmt.Println("new langs: ", lan)
}