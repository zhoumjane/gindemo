package model

import (
	"fmt"
	"ginblog1/utils/errmsg"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Code string `gorm:"type:varchar(100);not null" json:"code"`
	Name string `gorm:"unique;type:varchar(100);not null" json:"name"`
	Books []Book `gorm:"many2many:student_books;" json:"books"`
}

type Book struct {
	gorm.Model
	Name string `gorm:"unique;type:varchar(100);not null" json:"name"`
	Students []Student `gorm:"many2many:student_books;" json:"students"`
}

type studentBooks struct {
	BookId int `json:"book_id"`
	StudentId int `json:"student_id"`
}

func CheckStudentIsExist(studentID []int) int {
	var students  []Student
	db.Where(studentID).Find(&students)
	fmt.Println(len(students))
	fmt.Println(len(studentID))
	if len(students) != len(studentID){
		return errmsg.ErrorStudentNameNotExist
	}
	return errmsg.SUCCESS
}

func CreateForeignKeyStudent(studentID []int, bookID int) {
	var studentBookSlice []studentBooks
	for _, v := range studentID {
		student := studentBooks{
			BookId:    bookID,
			StudentId: v,
		}
		studentBookSlice = append(studentBookSlice, student)
	}
	db.Create(&studentBookSlice)

}

// 创建学生
func CreateStudent(data *Student) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ErrorStudentNameUsed
	}
	return errmsg.SUCCESS
}
// 创建书籍
func CreateBook(data *Book) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ErrorBookNameUsed
	}
	return errmsg.SUCCESS
}

// 查询学生有多少书籍
func GetStudentBooks(id int) (Student, int) {
	var books []Book
	var booksId []int
	var studentBookSlice []studentBooks
	var student Student
	db.Where("id = ?", id).First(&student)
	db.Where("student_id = ?", id).Find(&studentBookSlice)
	for _, v := range studentBookSlice {
		booksId = append(booksId, v.BookId)
	}
	db.Where(booksId).Find(&books)
	student.Books = books
	return student, errmsg.SUCCESS
}

// 查询一本书籍对应多少学生

