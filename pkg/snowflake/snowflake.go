package snowflake

import "github.com/sony/sonyflake/v2"

type Generator interface {
	NextID() (int64, error)
}

type generatorImpl struct {
	sf *sonyflake.Sonyflake
}


func NewGenerator(sf *sonyflake.Sonyflake) Generator {
	return &generatorImpl{sf}
}

func (p *generatorImpl) NextID() (int64, error) {
	id, err := p.sf.NextID()
	if err != nil {
		return 0, err
	}

	return id, nil
}