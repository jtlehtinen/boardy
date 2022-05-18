// Package models contains objects, i.e. structs, containing data.
// Subpackages of this package contain "DAOs" providing
// CRUD operations (create/read/update/delete) on the models.
// Each subpackage contains DAOs for a specific data source,
// such as sqlite3 or postgres.
//
// The models should not know anything about DAOs. DAO uses
// the model.
package models
