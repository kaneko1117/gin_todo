// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTask = "tasks"

// Task mapped from table <tasks>
type Task struct {
	ID        int32     `gorm:"column:id;type:integer;primaryKey" json:"id"`
	Tasks     string    `gorm:"column:tasks;type:text;not null" json:"tasks"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp(0) with time zone;not null" json:"created_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:timestamp(0) with time zone;not null" json:"update_at"`
}

// TableName Task's table name
func (*Task) TableName() string {
	return TableNameTask
}
