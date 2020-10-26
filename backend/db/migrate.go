package db

import (
	"fmt"
)

//MigrateDB migracion de tablas
func MigrateDB() {

	//db := GetConnection()

	fmt.Println("Migrating models....")
	// Automigrate se encarga de migrar la base de datos sí no se ha migrado, y lo hace a partir del modelo

}
