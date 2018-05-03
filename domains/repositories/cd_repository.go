package repositories

type CDRepository interface {
	GetByID(id int64) (*models.CD, error)
	Create(cd *models.CD) error
}
