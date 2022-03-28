package main

import "fmt"

// 定义一个学生结构体
type Student struct {
	id    int
	name  string
	age   int
	class string
}

// NewStudent 是一个创造新学生的构造函数
func NewStudent(id, age int, name, class string) *Student {
	return &Student{
		id:    id,
		name:  name,
		age:   age,
		class: class,
	}
}

// StudentMgr 定义一个学生信息管理的结构体
type StudentMgr struct {
	// 存储所有的学生的学生信息管理系统
	AllStudents []*Student
}

// 增加学生的方法 由于要修改 StudentMgr结构体中的字段 所以要用到指针
func (s *StudentMgr) AddStudent(stu *Student) {
	for _, v := range s.AllStudents {
		if stu.name == v.name {
			fmt.Printf("%s 已存在", stu.name)
			return
		}
	}

	s.AllStudents = append(s.AllStudents, stu)
}

// 修改学生的方法   由于要修改 StudentMgr结构体中的字段 所以要用到指针
func (s *StudentMgr) EditStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if stu.name == v.name {
			s.AllStudents[index] = stu
			fmt.Println("修改成功....")
			return
		}
	}

	fmt.Printf("%s 不存在", stu.name)
}

// 展示学生的方法   由于要修改 StudentMgr结构体中的字段 所以要用到指针
func (s *StudentMgr) ShowStudent() {
	for _, v := range s.AllStudents {
		fmt.Printf("学生id:%d,学生姓名:%s,学生年龄:%d,学生班级:%s", v.id, v.name, v.age, v.class)
	}
}

// 删除学生的方法   由于要修改 StudentMgr结构体中的字段 所以要用到指针
func (s *StudentMgr) DelStudent(stu *Student) {
	for index, v := range s.AllStudents {
		if stu.name == v.name {
			// 从切片中按照索引删除元素
			s.AllStudents = append(s.AllStudents[:index], s.AllStudents[index+1:]...)
			fmt.Println("删除成功....")
			return
		}
	}
	fmt.Printf("%s 不存在", stu.name)
}

// NewStudentMgr 是一个创建新的学生管理系统的构造函数
func NewStudentMgr() *StudentMgr {
	return &StudentMgr{
		AllStudents: make([]*Student, 0, 100),
	}
}
