package drop

import (
	"gorm.io/gorm"
)

func BulkCreateMonsterDrop(db *gorm.DB, monsterDrops []Model) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	for _, md := range monsterDrops {
		m := &entity{
			TenantId:        md.TenantId(),
			MonsterId:       md.MonsterId(),
			ItemId:          md.ItemId(),
			MinimumQuantity: md.MinimumQuantity(),
			MaximumQuantity: md.MaximumQuantity(),
			QuestId:         md.QuestId(),
			Chance:          md.Chance(),
		}

		err := tx.Create(m).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
