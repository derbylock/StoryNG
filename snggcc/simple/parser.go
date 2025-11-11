package simple

import (
	"github.com/derbylock/snggcc/pkg/model/github.com/derbylock/snggcc"
	"io"
)

type StoryParts struct {
	Characters   map[string]snggcc.Character
	Locations    map[string]snggcc.Location
	Passages     map[string]snggcc.Passage
	Environments map[string]snggcc.Item
	Actions      map[string]snggcc.Action
}

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (parser *Parser) Parse(reader io.Reader) {

}

func (parser *Parser) ParseFile(reader io.Reader) {

}
