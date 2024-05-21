package models

import "gorm.io/gorm"

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Register struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=3,max=32"`
	LineID   string `json:"line_id" validate:"required,min=3,max=32"`
	Phone    string `json:"phone" validate:"required,min=3,max=32"`
	Company  string `json:"company" validate:"required,min=3,max=32"`
	Website  string `json:"website" validate:"required,min=3,max=32"`
}

type Register2 struct {
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
	Name     string `json:"name" validate:"required,min=3,max=32,regexp=^[a-zA-Z0-9_ -]+$"`
	Password string `json:"password" validate:"required,min=6,max=20,regexp=^[a-zA-Z0-9_-]+$"`
	LineID   string `json:"line_id" validate:"required,min=6,max=20,regexp=^[a-zA-Z0-9_-]+$"`
	Phone    string `json:"phone" validate:"required,min=8,max=10"`
	Company  string `json:"company" validate:"required,min=3,max=32,regexp=^[a-zA-Z0-9_ -]+$"`
	Website  string `json:"website" validate:"required,min=3,max=32,regexp=^[a-zA-Z0-9_ - .]+$"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

//7.0 ย้ายStructที่อยู่ใน func GetDogsJson ไปไว้ในmodels แล้วเรียกใช้ผ่านmodelแทน

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data       []DogsRes `json:"data"`
	Name       string    `json:"name"`
	Count      int       `json:"count"`
	SumRed     int       `json:"sumred"`
	SumPink    int       `json:"sumpink"`
	SumGreen   int       `json:"sumgreen"`
	SumNoColor int       `json:"sumnocolor"`
}

//7.0.1 สร้างตารางcompany โดยใช้AutoMigrate โดยที่โครงสร้างcompanyควรจะมีอะไรบ้างใส่มาตามความเหมาะสม และGroupเพิ่มCRUD

type Company struct {
	gorm.Model
	Company_id int    `json:"company_id"`
	Name       string `json:"name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Website    string `json:"website"`
}

type ProFile struct {
	gorm.Model
	Name        string `json:"name"`
	Employee_id string `json:"employee"`
	LastName    string `json:"lastName"`
	Birthday    string `json:"birthday"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	Tel         string `json:"tel"`
}
type ProfileRes struct {
	Name      string `json:"name"`
	ProfileID int    `json:"profile_id"`
	Type      string `json:"type"`
}

type ResultProfileData struct {
	Data            []ProfileRes `json:"data"`
	Name            string       `json:"name"`
	Count           int          `json:"count"`
	SumGenZ         int          `json:"sum_gen_z"`
	SumGenY         int          `json:"sum_gen_y"`
	SumGenX         int          `json:"sum_gen_x"`
	SumBabyBoomer   int          `json:"sum_baby_boomer"`
	SumGIGeneration int          `json:"sum_gi_generation"`
}
