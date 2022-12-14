package main

import (
	"rt-msg-carrier/models/gen/dal/model"
	"rt-msg-carrier/tools"

	"gorm.io/gen"
	"gorm.io/gorm"
)

const mytableSQL = "CREATE TABLE IF NOT EXISTS `customers` (" +
	"    `ID` int(11) NOT NULL," +
	"    `username` varchar(16) DEFAULT NULL," +
	"    `age` int(8) NOT NULL," +
	"    `phone` varchar(11) NOT NULL," +
	"    INDEX `idx_username` (`username`)" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"

func prepare(db *gorm.DB) {
	db.Exec(mytableSQL)
}

func init() {
	_db := tools.GetDB()

	prepare(_db) // prepare table for generate
}

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		// if you want generate field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// ???????????????model??????crud?????????????????????????????????model struct ?????????model.User{}
	// ??????????????????????????????model???crud??????????????????????????????????????????g.GenerateModel("company")
	// ??????????????????????????????????????????struct?????????/????????????/tag??????????????????opt?????????g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	// g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	g.ApplyBasic(model.Customer{})

	// apply diy interfaces on structs or table models
	// ???????????????????????????model?????????????????????????????????ApplyInterface????????????????????????????????????????????????DIY??????????????????
	// g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	// execute the action of code generation
	g.Execute()
}
