package pci

import "fmt"

type Class struct {
	label      string
	subclasses map[uint8]Subclass
}

type Subclass struct {
	label   string
	progifs map[uint8]ProgIf
}

type ProgIf struct {
	label string
}

var Classes = map[uint8]Class{
	0x0: {
		label: "Unclassified",
		subclasses: map[uint8]Subclass{
			0x0: {
				label:   "Non-VGA-Compatible Unclassified Device",
				progifs: map[uint8]ProgIf{},
			},
			0x1: {
				label:   "VGA-Compatible Unclassified Device",
				progifs: map[uint8]ProgIf{},
			},
		},
	},
}

func ClassNames(class, subclass, progif *uint8) (*string, *string, *string, error) {
	if class == nil {
		return nil, nil, nil, fmt.Errorf("a class must be provided")
	}
	c, ok := Classes[*class]
	if !ok {
		return nil, nil, nil, fmt.Errorf("no class name found for '%d'", class)
	}
	if subclass == nil {
		return &c.label, nil, nil, nil
	}
	s, ok := c.subclasses[*subclass]
	if !ok {
		return &c.label, nil, nil, fmt.Errorf("no subclass '%d' found in class '%d'", subclass, class)
	}
	if progif == nil {
		return &c.label, &s.label, nil, nil
	}
	p, ok := s.progifs[*progif]
	if !ok {
		return &c.label, &s.label, nil, fmt.Errorf("no programming interface '%d' found in subclass '%d' of class '%d'", progif, subclass, class)
	}
	return &c.label, &s.label, &p.label, nil
}
