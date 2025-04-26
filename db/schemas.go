package db

type schema struct {
	name  string
	query string
}

var schemas = []schema{
	{
		name: "Services",
		query: `
			CREATE TABLE IF NOT EXISTS services(
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			image VARCHAR(100) NOT NULL,
			title VARCHAR(50) NOT NULL,
			description TEXT NOT NULL,
			created DATETIME DEFAULT NOW(),
			updated DATETIME DEFAULT NOW() ON UPDATE NOW());`,
	},
	{
		name: "Comments",
		query: `
			CREATE TABLE IF NOT EXISTS comments(
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			avatar VARCHAR(100) NOT NULL,
			title VARCHAR(50) NOT NULL,
			description TEXT NOT NULL,
			created DATETIME DEFAULT NOW(),
			updated DATETIME DEFAULT NOW() ON UPDATE NOW());`,
	},
	{
		name: "Team",
		query: `
			CREATE TABLE IF NOT EXISTS team(
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			avatar VARCHAR(100) NOT NULL,
			created DATETIME DEFAULT NOW(),
			updated DATETIME DEFAULT NOW() ON UPDATE NOW());`,
	},
	{
		name: "Contacts",
		query: `
			CREATE TABLE IF NOT EXISTS contacts(
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL,
			subject VARCHAR(50) NOT NULL,
			message TEXT NOT NULL,
			created DATETIME DEFAULT NOW(),
			updated DATETIME DEFAULT NOW() ON UPDATE NOW());`,
	},
}
