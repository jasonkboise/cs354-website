package main

import "fmt"

type student struct 
{
	id   int64
	name string
}


type studentMgr struct 
{
	allStudent map[int64]student
}


func (s studentMgr) showAllStudent() 
{
	for _, stu := range s.allStudent 
	{
		fmt.Printf("Student ID:%d Student Nmae:%s\n", stu.id, stu.name)
	}
}

func (s studentMgr) addStudent() 
{
	var 
	(
		stuID   int64
		stuName string
	)
	fmt.Print("Please Enter The Student ID:")
	fmt.Scanln(&stuID)
	fmt.Print("Please Enter The Student Name:")
	fmt.Scanln(&stuName)
	
	newStu := student
	{
		id:   stuID,
		name: stuName,
	} 

	s.allStudent[newStu.id] = newStu
}

func (s studentMgr) deleteStudent() 
{
	
	var deleteID int64
	fmt.Print("Please Enter The Student ID That You Want To Delete:")
	fmt.Scanln(&deleteID)
	
	_, ok := s.allStudent[deleteID]
	if !ok 
	{
		fmt.Println("No Such Person Found")
		return
	}
	
	delete(s.allStudent, deleteID)
}


func (s studentMgr) editStudent() 
{
	
	var editID int64
	fmt.Print("Please Enter The Student ID That You Want To Modify:")
	fmt.Scanln(&editID)
	
	stuObj, ok := s.allStudent[editID]
	if !ok 
	{
		fmt.Println("No Such Person Found")
		return
	}
	fmt.Printf("The Inforemation That You Want To Modify Is: Student ID:%d Student Name:%s\n", stuObj.id, stuObj.name)
	
	var newName string
	fmt.Print("Please Enter The New Name For The Student:")
	fmt.Scanln(&newName)
	
	stuObj.name = newName
	s.allStudent[editID] = stuObj
}
