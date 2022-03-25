package database

import (
	"context"
	"fmt"
	"github.com/AI1411/go-ecom/ent"
	"github.com/AI1411/go-ecom/ent/user"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

//func ExecEnt(ctx *gin.Context) {
//	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/go_ecom")
//	if err != nil {
//		log.Fatalf("failed opening connection to db: %v", err)
//	}
//	defer client.Close()
//
//	if err := client.Schema.Create(context.Background()); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//
//	if err := client.Schema.Create(ctx); err != nil {
//		log.Fatalf("failed creating schema resources: %v", err)
//	}
//	if _, err = CreateUser(ctx, client); err != nil {
//		log.Fatal(err)
//	}
//	if _, err = QueryUser(ctx, client); err != nil {
//		log.Fatal(err)
//	}
//}

func CreateUser(ctx *gin.Context) {
	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/go_ecom")
	u, err := client.User.
		Create().
		SetFirstName("Akira").
		SetLastName("Ishii").
		SetEmail("test@gmail,com").
		SetPassword("password").
		SetAge(38).
		Save(ctx)

	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.FirstName("Akira")).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Print("User found: ", u)
	return u, nil
}
