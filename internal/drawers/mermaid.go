package drawers

import (
	"fmt"
	"github.com/briggysmalls/archie/internal/types"
	"net/url"
)

func NewMermaidDrawer(linkAddress string) Drawer {
	// Create a new drawer with correct config
	d := newDrawer(MermaidConfig{linkAddress: linkAddress})
	return &d
}

type MermaidConfig struct {
	linkAddress string
}

func (p MermaidConfig) Header(writer Writer) {
	writer.Write("graph TD")
	writer.UpdateIndent(1)
}

func (p MermaidConfig) Footer(writer Writer) {
	// Do nothing
}

func (p MermaidConfig) Element(writer Writer, element *types.Element) {
	writer.Write("%p(%s)", element, element.Name)
	// Add a link if necessary
	if p.linkAddress != "" {
		fullName, err := writer.FullName(element)
		if err != nil {
			panic(err)
		}
		url := fmt.Sprintf("%s%s", p.linkAddress, url.PathEscape(fullName))
		writer.Write("click %p \"%s\"", element, url)
	}
}

func (p MermaidConfig) StartParentElement(writer Writer, element *types.Element) {
	writer.Write("subgraph %s", element.Name)
}

func (p MermaidConfig) EndParentElement(writer Writer, element *types.Element) {
	writer.Write("end")
}

func (p MermaidConfig) Association(writer Writer, association types.Relationship) {
	writer.Write("%p-->%p", association.Source, association.Destination)
}