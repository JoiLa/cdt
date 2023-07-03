package test

import (
	"fmt"
	"testing"

	"cdt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlData struct {
	ID      int         `gorm:"autoIncrement;primaryKey;column:id;type:integer;not null" json:"id"`
	String  cdt.DataRaw `gorm:"column:string;type:varchar(255)" json:"string" cdt:"type=string"`
	Int     cdt.DataRaw `gorm:"column:int;type:integer" json:"int" cdt:"type=int"`
	Float   cdt.DataRaw `gorm:"column:float;type:double" json:"float" cdt:"type=float"`
	Boolean cdt.DataRaw `gorm:"column:boolean;type:boolean" json:"boolean" cdt:"type=boolean"`
	Null    cdt.DataRaw `gorm:"column:null;type:varchar(50)" json:"null" cdt:"type=null"`
	Array   cdt.DataRaw `gorm:"column:array;type:text" json:"array" cdt:"type=array"`
	Object  cdt.DataRaw `gorm:"column:object;type:text" json:"object" cdt:"type=object"`
}

// TestSqlite test sqlite
func TestSqlite(t *testing.T) {
	storageName := "./data.sqlite"
	sqliteDb, err := gorm.Open(sqlite.Open(storageName), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatalf("open sqlite error: %v", err)
	}
	_ = sqliteDb.AutoMigrate(&sqlData{})
	//
	storeSqlData := new(sqlData)
	var name = "zhang-san"
	storeSqlData.String.Set(&name)
	storeSqlData.Int.Set(18)
	storeSqlData.Float.Set(1.8)
	storeSqlData.Boolean.Set(true)
	storeSqlData.Null.Set(nil)
	storeSqlData.Array.Set([]any{
		"zhang-san",
		"li-si",
		"wang-wu",
	})
	storeSqlData.Object.Set(map[string]any{
		"index": 1,
		"name":  "zhang-san",
	})
	// insert data to sqliteDb
	storeSqlDataErr := sqliteDb.Create(&storeSqlData).Error
	if storeSqlDataErr != nil {
		t.Fatalf("storeSqlData error: %v", storeSqlDataErr)
	}
	// read sql data by last insert id
	readSqlData := new(sqlData)
	readSqlDataErr := sqliteDb.Where("id = ?", storeSqlData.ID).First(&readSqlData).Error
	if readSqlDataErr != nil {
		t.Fatalf("read storeSqlData error: %v", readSqlDataErr)
	}
	// check data
	t.Run("test string is string", func(t *testing.T) {
		if !readSqlData.String.IsString() {
			t.Fatalf("readSqlData name is not string")
		}
	})
	t.Run("test int is int", func(t *testing.T) {
		if !readSqlData.Int.IsInt() {
			t.Fatalf("readSqlData int is not int")
		}
	})
	t.Run("test float is float", func(t *testing.T) {
		if !readSqlData.Float.IsFloat() {
			t.Fatalf("readSqlData float not float")
		}
	})
	t.Run("test int is numeric", func(t *testing.T) {
		if !readSqlData.Int.IsNumeric() {
			t.Fatalf("readSqlData int is not numeric")
		}
	})
	t.Run("test float is numeric", func(t *testing.T) {
		if !readSqlData.Float.IsNumeric() {
			t.Fatalf("readSqlData float not numeric")
		}
	})

	t.Run("test boolean is boolean", func(t *testing.T) {
		if !readSqlData.Boolean.IsBoolean() {
			t.Fatalf("readSqlData boolean not boolean")
		}
	})
	t.Run("test null is null", func(t *testing.T) {
		if !readSqlData.Null.IsNil() {
			t.Fatalf("readSqlData null not null")
		}
	})
	t.Run("test array is array", func(t *testing.T) {
		if !readSqlData.Array.IsArray() {
			t.Fatalf("readSqlData array not array")
		}
	})
	t.Run("test object is map", func(t *testing.T) {
		if !readSqlData.Object.IsMap() {
			t.Fatalf("readSqlData object not object")
		}
	})
}

// TestMysql test mysql
func TestMysql(t *testing.T) {
	// you can `edit/view` $project/docker/docker-compose.yml
	userName := "root"
	userPassword := "nOqbDNpEveiY63Fr"
	dsnPath := "127.0.0.1:53306"
	dbName := "cdt_database"
	dsnConfig := "charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		userName,
		userPassword,
		dsnPath,
		dbName,
		dsnConfig,
	)
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatalf("open sqlite error: %v", err)
	}
	_ = mysqlDb.AutoMigrate(&sqlData{})
	//
	storeSqlData := new(sqlData)
	var name = "zhang-san"
	storeSqlData.String.Set(&name)
	storeSqlData.Int.Set(18)
	storeSqlData.Float.Set(1.8)
	storeSqlData.Boolean.Set(true)
	storeSqlData.Null.Set(nil)
	storeSqlData.Array.Set([]any{
		"zhang-san",
		"li-si",
		"wang-wu",
	})
	storeSqlData.Object.Set(map[string]any{
		"index": 1,
		"name":  "zhang-san",
	})
	// insert data to mysqlDb
	storeSqlDataErr := mysqlDb.Create(&storeSqlData).Error
	if storeSqlDataErr != nil {
		t.Fatalf("storeSqlData error: %v", storeSqlDataErr)
	}
	// read sql data by last insert id
	readSqlData := new(sqlData)
	readSqlDataErr := mysqlDb.Where("id = ?", storeSqlData.ID).First(&readSqlData).Error
	if readSqlDataErr != nil {
		t.Fatalf("read storeSqlData error: %v", readSqlDataErr)
	}
	// check data
	t.Run("test string is string", func(t *testing.T) {
		if !readSqlData.String.IsString() {
			t.Fatalf("readSqlData name is not string")
		}
	})
	t.Run("test int is int", func(t *testing.T) {
		if !readSqlData.Int.IsInt() {
			t.Fatalf("readSqlData int is not int")
		}
	})
	t.Run("test float is float", func(t *testing.T) {
		if !readSqlData.Float.IsFloat() {
			t.Fatalf("readSqlData float not float")
		}
	})
	t.Run("test int is numeric", func(t *testing.T) {
		if !readSqlData.Int.IsNumeric() {
			t.Fatalf("readSqlData int is not numeric")
		}
	})
	t.Run("test float is numeric", func(t *testing.T) {
		if !readSqlData.Float.IsNumeric() {
			t.Fatalf("readSqlData float not numeric")
		}
	})

	t.Run("test boolean is boolean", func(t *testing.T) {
		if !readSqlData.Boolean.IsBoolean() {
			t.Fatalf("readSqlData boolean not boolean")
		}
	})
	t.Run("test null is null", func(t *testing.T) {
		if !readSqlData.Null.IsNil() {
			t.Fatalf("readSqlData null not null")
		}
	})
	t.Run("test array is array", func(t *testing.T) {
		if !readSqlData.Array.IsArray() {
			t.Fatalf("readSqlData array not array")
		}
	})
	t.Run("test object is map", func(t *testing.T) {
		if !readSqlData.Object.IsMap() {
			t.Fatalf("readSqlData object not object")
		}
	})
}
