package views

import (
	"testing"

	mdl "github.com/briggysmalls/archie/internal/model"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestLandscapeElements(t *testing.T) {
	// Create a simple model
	m := mdl.NewModel()

	// Create the items we'll be testing
	one := mdl.NewItem("One", nil)
	oneChild := mdl.NewItem("OneChild", nil)
	two := mdl.NewItem("Two", nil)
	twoChild := mdl.NewItem("TwoChild", nil)

	// Add the items, and their relationships to the model
	m.AddRootElement(one)
	m.AddElement(oneChild, one)
	m.AddRootElement(two)
	m.AddElement(twoChild, two)

	// Link the children together
	m.AddAssociation(oneChild, twoChild, "")

	// Create the landscape view
	l := NewLandscapeView(&m)

	// Check elements are correct
	assert.Assert(t, is.Contains(l.Elements, one))
	assert.Assert(t, is.Contains(l.Elements, two))
	assert.Assert(t, is.Len(l.Elements, 2))
	assert.Assert(t, is.Len(l.Children(one), 0))
	assert.Assert(t, is.Len(l.Children(two), 0))

	// Check relationships are correct
	assert.Assert(t, is.Len(l.Associations, 1))
	rel := l.Associations[0]
	assert.Equal(t, rel.Source(), one)
	assert.Equal(t, rel.Destination(), two)
}
