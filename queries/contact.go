package queries

const (
	ContactCreate = `
	INSERT INTO contacts (name,email,subject,message) VALUES (:name,:email,:subject,:message)
	`
)
