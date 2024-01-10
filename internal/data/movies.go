package data

import(
	"time"
)

// type Movie struct{
// 	ID			int64
// 	CreatedAt	time.Time
// 	Title		string
// 	Year		int32
// 	Runtime		int32
// 	Genres		[]string
// 	Version		int32
// }

type Movie struct {
	ID 			int64 		`json:"id"`
	CreatedAt 	time.Time 	`json:"created_at"`
	Title 		string 		`json:"-"`
	Year 		int32 		`json:"year,omitempty"`
	Runtime 	Runtime 		`json:"runtime"`
	Genres 		[]string 	`json:"genres"`
	Version 	int32 		`json:"version"`
}