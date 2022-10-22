package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Employee []Employee `gorm:"foreignKey:UserID"`
}

type Gender struct {
	gorm.Model
	Gender   string
	Employee []Employee `gorm:"foreignKey:GenderID"`
}

type Position struct {
	gorm.Model
	Position string
	Employee []Employee `gorm:"foreignKey:PositionID"`
}

type Education struct {
	gorm.Model
	Education string
	Employee  []Employee `gorm:"foreignKey:EducationID"`
}

type Employee struct {
	gorm.Model
	StartJob    time.Time
	FirstName   string
	LastName    string
	Telephone   string
	Email       string `gorm:"uniqueIndex"`
	Slary       int
	Password    string `json:"-"`
	UserID      *uint
	User        User
	GenderID    *uint
	Gender      Gender
	PositionID  *uint
	Position    Position
	EducationID *uint
	Education   Education
}

// --------------------------------------------------------------
type Province struct {
	gorm.Model
	Name    string
	Members []Member `gorm:"foreignKey:ProvinceID"`
}

type Member struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int

	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	Date_Of_Birth time.Time

	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	Telephone string `gorm:"UniqueIndex"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
}

// -------------------------------------------------------------
type Typeproduct struct {
	gorm.Model
	Typeproduct_name string
	Products         []Product `gorm:"foreignKey:TypeproductID"`
}

type Manufacturer struct {
	gorm.Model
	Manufacturer_name string
	Products          []Product `gorm:"foreignKey:ManufacturerID"`
}

type Product struct {
	gorm.Model
	Product_name   string
	Product_price  float32
	TypeproductID  *uint
	Typeproduct    Typeproduct `gorm:"references:id"`
	ManufacturerID *uint
	Manufacturer   Manufacturer `gorm:"references:id"`
	EmployeeID     *uint
	Employee       Employee `gorm:"references:id"`
}

// -------------------------------------------------------------------
type Lot struct {
	gorm.Model
	Lot_number string
	Expired    time.Time
	Stock      []Stock `gorm:"foreignKey:LotID"`
}

type Shelf struct {
	gorm.Model
	Shelf_name string
	Stock      []Stock `gorm:"foreignKey:ShelfID"`
}

type Stock struct {
	gorm.Model
	ProductID  *uint `gorm:"references:id"`
	Product    Product
	Quantity   float64
	LotID      *uint `gorm:"references:id"`
	Lot        Lot
	ShelfID    *uint `gorm:"references:id"`
	Shelf      Shelf
	EmployeeID *uint `gorm:"references:id"`
	Employee   Employee
}

// ============================================================
type Cart struct {
	gorm.Model
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
	MemberID   *uint
	Member     Member `gorm:"references:id"`

	//Receipt []Receipt `gorm:"foreignKey:CartID"`
}

type Order struct {
	gorm.Model
	Quantity  int
	ProductID *uint
	Product   Product `gorm:"references:id"`
	CartID    *uint
	Cart      Cart `gorm:"references:id"`

	//Receipts      []Receipt `gorm:"foreignKey:PaymentTypeID`
}

// ============================================================
type Payment_Type struct {
	gorm.Model
	Payment_Type string
	Receipt      []Receipt `gorm:"foreignKey:PaymentTypeID"`
}

// ตารางหลักของ ระบบชำระเงิน
type Receipt struct {
	gorm.Model
	Receipt_time           time.Time
	Receipt_sum            float32
	Receipt_payment_amount float32
	Receipt_change         float32

	PaymentTypeID *uint
	PaymentType   Payment_Type `gorm:"references:id"`

	CartID *uint
	Cart   Cart `gorm:"references:id"`

	// EmployeeID *uint
	// Employee    Employee `gorm:"references:id"`

	// MemberID *uint
	// Member    Member `gorm:"references:id"`
}
