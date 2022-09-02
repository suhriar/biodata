package main

import (
	"biodata/handler"
	"fmt"
	"os"
)


func main() {
	if len(os.Args) > 1 {
		args := os.Args
		id := args[1]
		biodata := handler.GetBiodataById(id)

		if biodata.Id != ""{
			fmt.Printf("Id\t\t: %s\nNama\t\t: %s\nAlamat\t\t: %s\nPekerjaan\t: %s\nAlasan\t\t: %s\n", biodata.Id, biodata.Name, biodata.Address, biodata.Job, biodata.Reason)
		} else {
			fmt.Println("Biodata Not Found")
		}
			
	} else {
		fmt.Println("No user arguments")
	}
	
}
	
	