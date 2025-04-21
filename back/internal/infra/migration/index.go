package main

import (
	// go modで記載しているmodule名/importするpackage名

	"fmt"
	"gin_todo/internal/infra/database"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	dbConn := database.NewDB()
	defer fmt.Println("Successfully migrated")
	defer database.CloseDB(dbConn)

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/entity/query",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		FieldNullable:     true,
	})

	g.UseDB(dbConn)
	tasksModel := g.GenerateModel("tasks")
	usersModel := g.GenerateModel("users", gen.FieldRelate(field.HasMany, "Tasks", tasksModel, &field.RelateConfig{
		GORMTag: field.GormTag{"foreignKey": []string{"UserID"}},
	}))

	g.ApplyBasic(tasksModel, usersModel)

	g.Execute()
}
