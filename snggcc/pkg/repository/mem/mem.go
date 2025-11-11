package mem

import (
	"encoding/json"
	"github.com/derbylock/snggcc/pkg/model/github.com/derbylock/snggcc"
)

type StoryRepository struct {
	Stories      map[string][]*snggcc.Story       `json:"stories,omitempty" yaml:"stories,omitempty"`
	Characters   map[string][]*snggcc.Character   `json:"characters,omitempty" yaml:"characters,omitempty"`
	Locations    map[string][]*snggcc.Location    `json:"locations,omitempty" yaml:"locations,omitempty"`
	Passages     map[string][]*snggcc.Passage     `json:"passages,omitempty" yaml:"passages,omitempty"`
	Environments map[string][]*snggcc.Environment `json:"environments,omitempty" yaml:"environments,omitempty"`
	Actions      map[string][]*snggcc.Action      `json:"actions,omitempty" yaml:"actions,omitempty"`
	Dialogues    map[string][]*snggcc.Dialogue    `json:"dialogues,omitempty" yaml:"dialogues,omitempty"`
}

func NewStoryRepository() *StoryRepository {
	return &StoryRepository{
		Stories:      make(map[string][]*snggcc.Story),
		Characters:   make(map[string][]*snggcc.Character),
		Locations:    make(map[string][]*snggcc.Location),
		Passages:     make(map[string][]*snggcc.Passage),
		Environments: make(map[string][]*snggcc.Environment),
		Actions:      make(map[string][]*snggcc.Action),
		Dialogues:    make(map[string][]*snggcc.Dialogue),
	}
}

func (r *StoryRepository) ConsumeCharacter(o *snggcc.Character) error {
	current, ok := r.Characters[o.Id]
	if !ok {
		current = make([]*snggcc.Character, 0)
	}
	current = append(current, o)
	r.Characters[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumeLocation(o *snggcc.Location) error {
	current, ok := r.Locations[o.Id]
	if !ok {
		current = make([]*snggcc.Location, 0)
	}
	current = append(current, o)
	r.Locations[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumePassage(o *snggcc.Passage) error {
	current, ok := r.Passages[o.Id]
	if !ok {
		current = make([]*snggcc.Passage, 0)
	}
	current = append(current, o)
	r.Passages[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumeEnvironment(o *snggcc.Environment) error {
	current, ok := r.Environments[o.Id]
	if !ok {
		current = make([]*snggcc.Environment, 0)
	}
	current = append(current, o)
	r.Environments[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumeAction(o *snggcc.Action) error {
	current, ok := r.Actions[o.Id]
	if !ok {
		current = make([]*snggcc.Action, 0)
	}
	current = append(current, o)
	r.Actions[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumeDialogue(o *snggcc.Dialogue) error {
	current, ok := r.Dialogues[o.Id]
	if !ok {
		current = make([]*snggcc.Dialogue, 0)
	}
	current = append(current, o)
	r.Dialogues[o.Id] = current
	return nil
}

func (r *StoryRepository) ConsumeStory(o *snggcc.Story) error {
	current, ok := r.Stories[o.Id]
	if !ok {
		current = make([]*snggcc.Story, 0)
	}
	current = append(current, o)
	r.Stories[o.Id] = current
	return nil
}

func (r *StoryRepository) ExportToJSON() ([]byte, error) {
	return json.Marshal(r)
}
