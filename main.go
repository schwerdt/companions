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

	petType := PetType{PetType: "Cat"}
	db.Where(petType).Find(&petType)

	guardian := Guardian{Name: "Christine"}
	cat := Pet{
		Name: "Mittens",
		PetType: petType,
		PetGuardians: []*PetGuardian{ {Guardian: &guardian, PrimaryGuardian: true }},
	}

	animal := Animal{Cat: cat}
	fmt.Println(animal)
	dbc := db.Create(&animal)
	if dbc.Error != nil {
		fmt.Println("there was an error saving:", err)
	}

	dogType := PetType{PetType: "Dog"}
	db.Where(dogType).Find(&dogType)

	dog := Pet{
		Name: "Sunny",
		PetType: dogType,
		PetGuardians: []*PetGuardian{ {Guardian: &guardian, PrimaryGuardian: false }},
	}
	anotherAnimal := Animal{Dog: dog}
	dbc = db.Create(&anotherAnimal)
	if dbc.Error != nil {
		fmt.Println("error saving sunny:", err)
	}


}


