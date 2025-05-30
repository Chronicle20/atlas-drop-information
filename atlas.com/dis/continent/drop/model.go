package drop

import "github.com/google/uuid"

type Model struct {
	tenantId        uuid.UUID
	id              uint32
	continentId     int32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	questId         uint32
	chance          uint32
}

func (d Model) ContinentId() int32 {
	return d.continentId
}

func (d Model) ItemId() uint32 {
	return d.itemId
}

func (d Model) MinimumQuantity() uint32 {
	return d.minimumQuantity
}

func (d Model) MaximumQuantity() uint32 {
	return d.maximumQuantity
}

func (d Model) QuestId() uint32 {
	return d.questId
}

func (d Model) Chance() uint32 {
	return d.chance
}

func (d Model) Id() uint32 {
	return d.id
}

func (d Model) TenantId() uuid.UUID {
	return d.tenantId
}

type builder struct {
	tenantId        uuid.UUID
	id              uint32
	continentId     int32
	itemId          uint32
	minimumQuantity uint32
	maximumQuantity uint32
	questId         uint32
	chance          uint32
}

func NewContinentDropBuilder(tenantId uuid.UUID, id uint32) *builder {
	return &builder{tenantId: tenantId, id: id}
}

func (m *builder) SetContinentId(continentId int32) *builder {
	m.continentId = continentId
	return m
}

func (m *builder) SetItemId(itemId uint32) *builder {
	m.itemId = itemId
	return m
}

func (m *builder) SetMinimumQuantity(minimumQuantity uint32) *builder {
	m.minimumQuantity = minimumQuantity
	return m
}

func (m *builder) SetMaximumQuantity(maximumQuantity uint32) *builder {
	m.maximumQuantity = maximumQuantity
	return m
}

func (m *builder) SetChance(chance uint32) *builder {
	m.chance = chance
	return m
}

func (m *builder) SetQuestId(questId uint32) *builder {
	m.questId = questId
	return m
}

func (m *builder) Build() Model {
	return Model{
		tenantId:        m.tenantId,
		id:              m.id,
		continentId:     m.continentId,
		itemId:          m.itemId,
		minimumQuantity: m.minimumQuantity,
		maximumQuantity: m.maximumQuantity,
		questId:         m.questId,
		chance:          m.chance,
	}
}
