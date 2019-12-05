package main

import (
	"rank-archive/databases"
	"rank-archive/routers"
)

func main() {
	databases.AutoMigrate()

	router := routers.Load()

	_ = router.Run(":8091")
}
