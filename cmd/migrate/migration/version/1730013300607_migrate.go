package version

import (
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"gorm.io/gorm"
	"runtime"
	"smart-api/app/smart/models"
	"smart-api/cmd/migrate/migration"
	models2 "smart-api/cmd/migrate/migration/models"
	common "smart-api/common/models"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1730013300607Test)
}

// 添加自定义的表，为系统表
func _1730013300607Test(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		// 先进行自动迁移，创建表
		err := tx.Migrator().AutoMigrate(
			new(models.OrderCategory),
			new(models.OrderItems),
			new(models.ExecMachine),
			new(models.ExecutionHistoryTask),
			new(models.FlowManage),
			new(models.FlowTemplates),
			new(models.OperationHistory),
			new(models.OrderComment),
			new(models.OrderRating),
			new(models.OrderTask),
			new(models.OrderWorks),
			new(models.UserFavorites),
			new(models.WorksNotify),
		)
		if err != nil {
			return err
		}
		// 确保初始化数据库模型
		if err := models2.InitDb(tx); err != nil {
			return err
		}

		return tx.Create(&common.Migration{
			Version: version,
		}).Error

	})
}
