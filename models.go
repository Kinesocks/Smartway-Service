package main

type Company struct {
	ID   uint   `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"unique"`
}

type Employee struct {
	ID           uint        `json:"Id"`
	CompanyID    uint        `json:"CompanyId"`
	DepartmentID uint        `json:"DepartmentId"`
	Company      *Company    `json:"-" gorm:"foreignKey:CompanyID;references:ID"`
	Department   *Department `json:"-" gorm:"foreignKey:DepartmentID;references:ID"`
	Name         string      `json:"Name"`
	Surname      string      `json:"Surname"`
	Phone        string      `json:"Phone"`
}

type Passport struct {
	ID         uint      `gorm:"primaryKey;autoIncrement:true"`
	EmployeeID uint      `gorm:"primaryKey;autoIncrement:true"`
	Employee   *Employee `json:"-" gorm:"foreignKey:EmployeeID;references:ID"`
	Type       string    `json:"Type"`
	Number     string    `json:"Number"`
}

type Department struct {
	ID        uint     `json:"ID"`
	CompanyID uint     `json:"CompanyID"`
	Name      string   `json:"Name"`
	Phone     string   `json:"Phone"`
	Company   *Company `json:"-" gorm:"foreignKey:CompanyID;references:ID"`
}

type Model struct {
	ID           uint       `json:"Id"`
	Name         string     `json:"Name"`
	Surname      string     `json:"Surname"`
	Phone        string     `json:"Phone"`
	CompanyID    uint       `json:"CompanyID"`
	DepartmentID uint       `json:"DepartmentID"`
	Passport     Passport   `json:"Passport"`
	Department   Department `json:"Department"`
}
