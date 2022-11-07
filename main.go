package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "time"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "1234"
	dbname = "exam"
)

func main(){
	connstr := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("Error is equaried while connecting to database: %v", err)
	}
	DBmanager := NewDBManager(db)
	// car_id, err := DBmanager.CreateAvtomobile(&Car{
	// 	Name: "Haval",
	// 	Model: "H6",
	// 	Year: time.Date(2020, 9, 3, 13, 34, 00, 00, time.Local),
	// 	Color: "Black",
	// 	HorsePower: 300,
	// 	Km: 0,
	// 	Images: []*Images{
	// 		{
	// 			Sequence_number: 1,
	// 			Image_url: "~/Desktop/car1.jpg",
	// 		},
	// 		{
	// 			Sequence_number: 2,
	// 			Image_url: "~/Desktop/car2.jpg",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("failed to create new car: %v", err)
	// }
	// fmt.Println(car_id)
	// car, err := DBmanager.GetCar(car_id)
	// if err != nil {
	// 	log.Fatalf("failed to get car info: %v", err)
	// }
	// PrintCarInfo(car)
 	// err = DBmanager.UpdateCar(&Car{
	// 	Id: car_id,
	// 	Name: "Captiva",
	// 	Model: "3",
	// 	Year: time.Date(2018, 3, 12, 9, 12, 23, 45, time.Local),
	// 	Color: "White",
	// 	HorsePower: 250,
	// 	Km: 0,
	// 	Images: []*Images{
	// 		{
	// 			Sequence_number: 1,
	// 			Image_url: "~/Desktop/car1.jpg",
	// 		},
	// 		{
	// 			Sequence_number: 2,
	// 			Image_url: "~/Desktop/car2.jpg",
	// 		},
	// 		{
	// 			Sequence_number: 3,
	// 			Image_url: "~/Desktop/captiva3.jpg",
	// 		},
	// 		{
	// 			Sequence_number: 4,
	// 			Image_url: "~/Desktop/captive4.jpg",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("failed to update car: %v", err)
	// }
	// fmt.Println("-------- After updating ---------")
	// car, err = DBmanager.GetCar(car_id)
	// if err != nil {
	// 	log.Fatalf("failed to get car info: %v", err)
	// }
	// PrintCarInfo(car)
	// err = DBmanager.DeleteCar(14)
	// if err != nil {
	// 	log.Fatalf("failed to delete car: %v", err)
	// }
	allcars, err := DBmanager.GetAllCarsInfo(&GetAllParams{
		Limit: 10,
		Page: 1,
	})
	if err != nil {
		log.Fatalf("failed to take all cars info: %v", err)
	}
	for _, v := range allcars.AllCars {
		fmt.Println(*v)
	}
}

func PrintCarInfo(c *Car) {
	fmt.Printf("-------- %v - Car Info --------\n", c.Id)
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Model: ", c.Model)
	fmt.Println("Car Year: ", c.Year)
	fmt.Println("Car Color: ", c.Color)
	fmt.Println("Car Horse Power: ", c.HorsePower)
	fmt.Println("Car Ride Km: ", c.Km)
	PrintImages(c.Images, c.Id)
}

func PrintImages(i []*Images, id int) {
	fmt.Printf("-------- %v - Car Pictures --------\n", id)
	for i, v := range i {
		fmt.Printf("%v - Image: %v\n", i + 1, v.Image_url)
		fmt.Printf("%v - Image Sequence Number: %v\n", i + 1, v.Sequence_number)
		fmt.Println("--------------------------------")
	}
}