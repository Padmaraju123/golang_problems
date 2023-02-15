package main
import "fmt"

func main(){
	//embedded struct
	type Contact struct{
		email,address string
		status bool
	}

	type employee struct{
		name string
		salary int
		contactinfo Contact
	}

	obj := employee{
		name: "raju",
		salary: 10000,
		contactinfo: Contact{         //contactinfo act as object and contact as class 
			email:"asdadada",
			address: "@32323",
			status:true,
		},
	}


	//to access the vlaues
	fmt.Println(obj.contactinfo.email)
	//to update the values
	obj.contactinfo.email = "padmaraju1232"
	fmt.Println(obj.contactinfo.email)
}