// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1

package querytest

import ()

type Bar struct {
	ID    int32
	Owner string
}

type Foo struct {
	Barid int32
}