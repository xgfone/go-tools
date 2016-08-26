package pagination_test

import (
	"fmt"
	"testing"

	"github.com/xgfone/go-tools/pagination"
)

func ExamplePagination() {
	p := pagination.NewPagination(100, 1, 10, 5)
	fmt.Printf("Total: %v\n", p.Total)
	fmt.Printf("Start: %v\n", p.Start)
	fmt.Printf("End  : %v\n", p.End)

	fmt.Printf("HasPrev: %v\n", p.HasPrev)
	fmt.Printf("Prev   : %v\n", p.Prev)
	fmt.Printf("HasNext: %v\n", p.HasNext)
	fmt.Printf("Next   : %v\n", p.Next)

	for _, page := range p.Pages {
		fmt.Printf("Active: %v, Number: %v\n", page.Active, page.Number)
	}

	// Output:
	// Total: 10
	// Start: 1
	// End  : 5
	// HasPrev: false
	// Prev   : 0
	// HasNext: true
	// Next   : 6
	// Active: true, Number: 1
	// Active: false, Number: 2
	// Active: false, Number: 3
	// Active: false, Number: 4
	// Active: false, Number: 5
}

func TestPagination(t *testing.T) {
	if p := pagination.NewPagination(-10, 1, 10, 5); p.Total != 0 {
		t.Fail()
	}

	if p := pagination.NewPagination(10, 1, 10, 5); p.Total != 1 || p.Start != 1 || p.End != 1 {
		t.Fail()
	}

	if p := pagination.NewPagination(5, 1, 10, 5); p.Total != 1 || p.Start != 1 || p.End != 1 {
		t.Fail()
	}

	if p := pagination.NewPagination(100, 2, 10, 5); p.HasPrev || !p.HasNext {
		t.Fail()
	}

	if p := pagination.NewPagination(100, 8, 10, 5); !p.HasPrev || p.HasNext {
		t.Fail()
	}

	if p := pagination.NewPagination(100, 5, 10, 5); !p.HasPrev || !p.HasNext {
		t.Fail()
	}

	if p := pagination.NewPagination(100, 5, 10, 10); p.HasPrev || p.HasNext {
		t.Fail()
	}
}
