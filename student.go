package main

import "fmt"

// 定义一个结构体描述学生信息
type student struct {
	id    int
	name  string
	class string
}

// 定义student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student //student类型的切片
}

// 这是studentMgr的构造函数
func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100), //初始化切片
	}
}

// 1.添加学员
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.删改学员
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { //当学号相同时，表示找到了对应学生
			s.allStudents[i] = newStu //根据切片的索引直接把新学生赋值进来
			return
		}
		//如果走到这里说明输入的学生没有找到
		fmt.Println("输入的学生信息有误，系统中没有学号是:%d的学生\n", newStu.id)
	}
}

// 3.展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}
