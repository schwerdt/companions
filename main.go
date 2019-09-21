package main

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Animal struct {
	Id int
	DogId *int
	CatId *int
	Dog Pet
	Cat Pet
}

type Pet struct {
	Id int
	Name string
	PetTypeId *int
	PetType PetType
	PetGuardians []*PetGuardian
}

type Guardian struct {
	Id int
	Name string
}

type PetGuardian struct {
	Pet *Pet
	PetId *int
	Guardian *Guardian
	GuardianId *int
	PrimaryGuardian bool
}

type PetType struct {
	Id int
	PetType string
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=animals sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Write some data
	catType := PetType{PetType: "Cat"}
	db.Where(catType).Find(&catType)

	guardian := Guardian{Name: "Gophie"}
	cat := Pet{
		Name: "Lionini",
		PetType: catType,
		PetGuardians: []*PetGuardian{ {Guardian: &guardian, PrimaryGuardian: true }},
	}

	catAnimal := Animal{Cat: cat}
	dbc := db.Create(&catAnimal)
	if dbc.Error != nil {
		fmt.Println("error saving the cat:", err)
	}

	dogType := PetType{PetType: "Dog"}
	db.Where(dogType).Find(&dogType)

	dog := Pet{
		Name: "Wolfini",
		PetType: dogType,
		PetGuardians: []*PetGuardian{ {Guardian: &guardian, PrimaryGuardian: false }},
	}
	dogAnimal := Animal{Dog: dog}
	dbc = db.Create(&dogAnimal)
	if dbc.Error != nil {
		fmt.Println("error saving the dog:", err)
	}

	// Read some data
	animal := Animal{}
	dbc = db.Preload("Dog").Preload("Cat").Where(Animal{Id: 11}).Find(&animal)
	if dbc.Error != nil {
		fmt.Println("error reading from DB", err)
	}
	fmt.Println(animal)
	printAnimal(animal)
}

func printAnimal(animal Animal) {
	fmt.Println("Id", animal.Id)
	fmt.Println("DogID", animal.DogId)
	fmt.Println("Dog", animal.Dog)
}
