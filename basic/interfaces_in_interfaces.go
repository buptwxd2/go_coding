package main

import "fmt"

type AuthorDetails interface {

	details()
}

type AuthorArticles interface {

	articles()
}

type FinalDetails interface {
	AuthorArticles
	AuthorDetails
}

// Structure
type author struct {

	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d",
		a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

// Main value
func main() {

	// Assigning values
	// to the structure
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the methods of
	// the interface 1 and 2
	// Using FinalDetails interface
	var f FinalDetails = values
	f.details()
	f.articles()
}


