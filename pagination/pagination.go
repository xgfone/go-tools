// It is usually used to compute the web pagination.
package pagination

import "math"

type Pagination struct {
	// The number of the current page
	Current int

	// The number of the elements in every page
	Limit int

	// The total number of the pages
	Total int

	// The number of the first page in the page group.
	Start int

	// The number of the last page in the page group.
	End int

	// The number of the previous page outside the page group.
	// It's 0 if there is no previous page.
	Prev int

	// The number of the next page outside the page group.
	// It's 0 if there is no next page.
	Next int

	// Whether there is the previous page.
	HasPrev bool

	// Whether there is the next page.
	HasNext bool

	// The page group, which contains all the produced pages.
	Pages []Page
}

type Page struct {
	// Whether this page is the current page.
	Active bool

	// The number of this page.
	Number int
}

// Create a new pagination.
//
// total is the total number of all the elements. currentPage is the current
// page number, that's, which page the current is. limit is the number of the
// elements in every page. number is the total number of the pages to be
// produced.
//
// For example, there are 100 elements in total, and there are 10 elements in
// every page. The current is on page one, and we want to produce 5 pages.
//
//  p := NewPagination(100, 1, 10, 5)
func NewPagination(total, currentPage, limit, number int) Pagination {
	p := Pagination{Current: currentPage, Limit: limit}

	if currentPage < 1 {
		currentPage = 1
	}

	if limit < 1 {
		limit = 10
	}

	p.Total = int(math.Ceil(float64(total) / float64(limit)))
	if p.Total >= 1 {
		// Compute the Start ant the End page number.
		if p.Total <= number {
			p.Start = 1
			p.End = p.Total
		} else {
			p.Start = currentPage - int(math.Floor(float64(number)/float64(2)))
			p.End = currentPage + int(math.Floor(float64(number)/float64(2)))
			if p.Start < 1 {
				p.End += int(math.Abs(float64(p.Start))) + 1
				p.Start = 1
			} else if p.End > p.Total {
				p.Start -= (p.End - number)
				p.End = number
			}
		}

		// Compute whether there are the previous and the next pages.
		if p.Start > 1 {
			p.HasPrev = true
			p.Prev = p.Start - 1
		}
		if p.End < p.Total {
			p.HasNext = true
			p.Next = p.End + 1
		}

		// Compute the number of each page from start to end.
		for i := p.Start; i <= p.End; i++ {
			page := Page{Number: i}
			if i == currentPage {
				page.Active = true
			}
			p.Pages = append(p.Pages, page)
		}
	} else {
		p.Total = 0
	}

	return p
}
