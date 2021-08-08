package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Kubeitron/soda-api/cmd/api/models"
	"github.com/Kubeitron/soda-api/cmd/api/store"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	VegetableHandler struct {
		store *store.VegetableStore
	}
)

func NewVegetableHandler(store *store.VegetableStore) (h *VegetableHandler) {
	return &VegetableHandler{
		store,
	}
}

func (h *VegetableHandler) GetVegetables(c echo.Context) error {
	vc := h.store.Db.Conn.Database(h.store.Db.DbName).Collection(models.VegetableCollectionName)
	findOptions := options.Find()
	findOptions.SetLimit(20)
	var resultSet []*models.Vegetable
	cur, err := vc.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem models.Vegetable
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		resultSet = append(resultSet, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	content := c.JSON(http.StatusOK, resultSet)
	return content
}

// func postVegetable(c echo.Context) error {
// 	vc := db.Database(dbName).Collection(VegetableCollection)
// 	weight, _ := strconv.ParseFloat(c.Param("weight"), 64)
// 	cal, _ := strconv.ParseInt(c.Param("calories"), 10, 0)
// 	veg := Vegetable{
// 		Color:    c.Param("color"),
// 		Weight:   weight,
// 		Name:     c.Param("name"),
// 		Vitamins: c.Param("vitamins"),
// 		Calories: int(cal),
// 	}
// 	_, err := vc.InsertOne(context.TODO(), veg)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	content := c.JSON(http.StatusCreated, veg)
// 	return content
// }
