package main

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Companionship struct {
	Id int
	PetId *int
	GuardianId *int
	Pet Creature
	Guardian Creature
}

type Creature struct {
	Id int
	Name string
	CreatureTypeId *int
	CreatureType CreatureType
	CreatureHobbies []*CreatureHobby
	CreatureFoods []*CreatureFood
}

type CreatureType struct {
	Id int
	CreatureType string
}

type CreatureHobby struct {
	Creature *Creature
	CreatureId *int
	Hobby *Hobby
	HobbyId *int
	FavoriteHobby bool
}

type CreatureFood struct {
	Creature *Creature
	CreatureId *int
	Food *Food
	FoodId *int
	FavoriteFood bool
}

type Hobby struct {
	Id int
	Hobby string
}

type Food struct {
	Id int
	Food string
}

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=companionships sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()


	catCreatureType := CreatureType{CreatureType: "Cat"}
	db.Where(catCreatureType).Find(&catCreatureType)

	gorillaCreatureType := CreatureType{CreatureType: "Gorilla"}
	db.Where(gorillaCreatureType).Find(&gorillaCreatureType)

	hobbyClimbing := Hobby{Hobby: "Climbing"}
	hobbyPurring := Hobby{Hobby: "Purring"}

	foodTuna := Food{Food: "Tuna"}
	foodBamboo := Food{Food: "Bamboo Shoots"}


	guardianCreature := Creature{
				Name: "Koko",
				CreatureType: gorillaCreatureType,
				CreatureHobbies: []*CreatureHobby{{Hobby: &hobbyClimbing, FavoriteHobby: true}},
				CreatureFoods: []*CreatureFood{{Food: &foodBamboo, FavoriteFood: false}},
			}
	petCreature := Creature{
				Name: "All Ball",
				CreatureType: catCreatureType,
				CreatureHobbies: []*CreatureHobby{{Hobby: &hobbyPurring, FavoriteHobby: false}},
				CreatureFoods: []*CreatureFood{{Food: &foodTuna, FavoriteFood: true}},
			}
	companionship := Companionship{
				Pet: petCreature,
				Guardian: guardianCreature,
			}
	dbc := db.Create(&companionship)
	if dbc.Error != nil {
		fmt.Println("error saving companionship", dbc.Error)
	}


	//Read some data
	companions := Companionship{}
	dbc = db.Preload("Guardian").Preload("Pet").Preload("Guardian.CreatureType").Preload("Pet.CreatureType").Preload("Pet.CreatureHobbies").Preload("Guardian.CreatureHobbies").Preload("Guardian.CreatureHobbies.Hobby").Where(Companionship{Id: 1}).Find(&companions)
	if dbc.Error != nil {
		fmt.Println("error reading:", dbc.Error)
	}
	printCompanionship(companions)
}

func printCompanionship(companionship Companionship) {
      fmt.Println("Companions:")
      fmt.Println("Pet:", companionship.Pet.Name)
      fmt.Println("Guardian:", companionship.Guardian.Name)
      fmt.Println("Pet Hobbies:", companionship.Pet.CreatureHobbies[0])
      fmt.Println("Guardian Hobbies:", companionship.Guardian.CreatureHobbies[0])
}
