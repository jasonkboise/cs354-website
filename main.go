package main

import 
(
	"fmt"
	"os"
)

func showMenu() 
{
	fmt.Println("Welcome")
	fmt.Println
	(`
		1.View Student Information
		2.Add Student Information
		3.Delete Student Information
		4.Modify Student Information
		5.Exit
	`)
}

var smr studentMgr

func main() 
{
	var smr = studentMgr
	{
		allStudent: make(map[int64]student, 100),
	}
	for 
	{

		showMenu()
		fmt.Print("Enter The Number:")
		
		var choice int
		fmt.Scanln(&choice)
		
		switch choice 
		{
			case 1:
				smr.showAllStudent()
			case 2:
				smr.addStudent()
			case 3:
				smr.deleteStudent()
			case 4:
				smr.editStudent()
			case 5:
				os.Exit(1)
			default:
				fmt.Println("Please Enter The Number In 1~5")
		}
	}
}
