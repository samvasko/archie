package writers

import (
	"fmt"
	"strings"
)

type PlantUmlStrategy struct {
}

func (p PlantUmlStrategy) Header(scribe Scribe) {
	scribe.WriteLine("@startuml")
}

func (p PlantUmlStrategy) Footer(scribe Scribe) {
	scribe.WriteLine("skinparam shadowing false")
	scribe.WriteLine("skinparam nodesep 10")
	scribe.WriteLine("skinparam ranksep 20")
	scribe.WriteLine("@enduml")
}

func (p PlantUmlStrategy) Element(scribe Scribe, element Element) {
	if element.IsActor() {
		scribe.WriteLine("actor \"%s\" as %s", element.Name(), element.ID())
		return
	}
	if len(element.Tags()) != 0 {
		scribe.WriteLine("rectangle %s [", element.ID())
		scribe.WriteLine("%s", element.Name())
		scribe.WriteLine("<i>%s</i>", strings.Join(element.Tags(), ", "))
		scribe.WriteLine("]")
		return
	}
	scribe.WriteLine("rectangle \"%s\" as %s", element.Name(), element.ID())
}

func (p PlantUmlStrategy) StartParentElement(scribe Scribe, element Element) {
	scribe.WriteLine("package \"%s\" {", element.Name())
	scribe.UpdateIndent(1)
}

func (p PlantUmlStrategy) EndParentElement(scribe Scribe, element Element) {
	scribe.UpdateIndent(-1)
	scribe.WriteLine("}")
}

func (p PlantUmlStrategy) Association(scribe Scribe, association Relationship) {
	linkStr := fmt.Sprintf("%s --> %s", association.Source().ID(), association.Destination().ID())
	if association.Tag() != "" {
		scribe.WriteLine("%s : \"%s\"", linkStr, association.Tag())
	} else {
		scribe.WriteLine(linkStr)
	}
}
