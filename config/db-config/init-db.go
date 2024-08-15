package dbconfig

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/need/go-backend/queries/branchesquery"
	"github.com/need/go-backend/queries/productquery"
	"github.com/need/go-backend/queries/userquery"
)

func debugDb() {
	_, err := DB.Exec(userquery.UserQueryDebug)
	if err != nil {
		log.Fatal("Error while creating table : ", err)
	}
	// Insert a test value
	_, err = DB.Exec(`INSERT INTO DebugUser (name , date) VALUES ('Started server at' , Now()) ON CONFLICT DO NOTHING;`)
	if err != nil {
		log.Fatal("Error inserting test value: ", err)
	}
}

func initializeEntity() {

	//สร้าง Table ที่จำเป็นก่่อน
	//NOTE - Deapartments
	_, err := DB.Exec(userquery.CreateDepartmentTable)
	if err != nil {
		log.Fatal("Error while creating Deapartments : ", err)
	}
	//NOTE - Users
	_, err = DB.Exec(userquery.CreateUserTable)
	if err != nil {
		log.Fatal("Error while creating Users : ", err)
	}
	//NOTE - Products
	_, err = DB.Exec(productquery.CreateProductsTable)
	if err != nil {
		log.Fatal("Error while creating Products : ", err)
	}

	//NOTE - ProductStock
	_, err = DB.Exec(productquery.CreateProductStock)
	if err != nil {
		log.Fatal("Error while creating ProductStock : ", err)
	}
	//NOTE - Branches
	_, err = DB.Exec(branchesquery.CreateBranchesTable)
	if err != nil {
		log.Fatal("Error while creating Branches : ", err)
	}
	//NOTE - BranchLog
	_, err = DB.Exec(branchesquery.CreateBranchProductQty)
	if err != nil {
		log.Fatal("Error while creating BranchProductQty : ", err)
	}
	//NOTE - ProductLog
	_, err = DB.Exec(productquery.CreateProductLog)
	if err != nil {
		log.Fatal("Error while creating ProductLog : ", err)
	}
	log.Println("Entity created successfully")
}
