package postgres

import (
	"errors"
	"fmt"
	"github.com/odpf/siren/store/model"
	"gorm.io/gorm"
)

// TemplateRepository talks to the store to read or insert data
type Repository struct {
	db *gorm.DB
}

// NewTemplateRepository returns repository struct
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r Repository) List() ([]*model.Namespace, error) {
	var namespaces []*model.Namespace
	selectQuery := "select * from namespaces"
	result := r.db.Raw(selectQuery).Find(&namespaces)
	if result.Error != nil {
		return nil, result.Error
	}

	return namespaces, nil
}

func (r Repository) Create(namespace *model.Namespace) (*model.Namespace, error) {
	var newNamespace model.Namespace
	result := r.db.Create(namespace)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Where(fmt.Sprintf("id = %d", namespace.Id)).Find(&newNamespace)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newNamespace, nil
}

func (r Repository) Get(id uint64) (*model.Namespace, error) {
	var namespace model.Namespace
	result := r.db.Where(fmt.Sprintf("id = %d", id)).Find(&namespace)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &namespace, nil
}

func (r Repository) Update(namespace *model.Namespace) (*model.Namespace, error) {
	var newNamespace, existingNamespace model.Namespace
	result := r.db.Where(fmt.Sprintf("id = %d", namespace.Id)).Find(&existingNamespace)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("namespace doesn't exist")
	} else {
		result = r.db.Where("id = ?", namespace.Id).Updates(namespace)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	result = r.db.Where(fmt.Sprintf("id = %d", namespace.Id)).Find(&newNamespace)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newNamespace, nil
}

func (r Repository) Delete(id uint64) error {
	var namespace model.Namespace
	result := r.db.Where("id = ?", id).Delete(&namespace)
	return result.Error
}

func (r Repository) Migrate() error {
	err := r.db.AutoMigrate(&model.Namespace{})
	if err != nil {
		return err
	}
	return nil
}
