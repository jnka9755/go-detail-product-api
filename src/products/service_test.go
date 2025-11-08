package products

import (
	"errors"
	"io/ioutil"
	"log"
	"testing"

	handleerrors "github.com/jnka9755/go-detail-product-api/src/products/handle-errors"
	"github.com/jnka9755/go-detail-product-api/src/products/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) GetProductsDetail() []models.Product {
	args := m.Called()
	return args.Get(0).([]models.Product)
}

func (m *mockRepo) GetProductsDetailById(id string) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func setupTestService() (*mockRepo, Service) {
	logger := log.New(ioutil.Discard, "", 0)
	repo := new(mockRepo)
	service := NewService(logger, repo)
	return repo, service
}

func TestServiceGetProductsDetail(t *testing.T) {
	mockRepo, service := setupTestService()

	testProducts := []models.Product{
		{
			ID:    "TEST1",
			Title: "Test Product 1",
		},
		{
			ID:    "TEST2",
			Title: "Test Product 2",
		},
	}

	// Setup expectations
	mockRepo.On("GetProductsDetail").Return(testProducts)

	// Call service
	products := service.GetProductsDetail()

	// Verify
	assert.Len(t, products, 2)
	assert.Equal(t, testProducts, products)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetProductsDetailById(t *testing.T) {
	mockRepo, service := setupTestService()

	tests := []struct {
		name    string
		id      string
		mock    func()
		wantErr error
		check   func(*testing.T, models.Product)
	}{
		{
			name: "existing product",
			id:   "TEST1",
			mock: func() {
				product := models.Product{
					ID:    "TEST1",
					Title: "Test Product 1",
				}
				mockRepo.On("GetProductsDetailById", "TEST1").Return(product, nil)
			},
			check: func(t *testing.T, p models.Product) {
				assert.Equal(t, "TEST1", p.ID)
				assert.Equal(t, "Test Product 1", p.Title)
			},
		},
		{
			name: "non-existing product",
			id:   "NONEXISTENT",
			mock: func() {
				mockRepo.On("GetProductsDetailById", "NONEXISTENT").Return(models.Product{}, handleerrors.ErrNotFound)
			},
			wantErr: handleerrors.ErrNotFound,
		},
		{
			name: "repository error",
			id:   "ERROR",
			mock: func() {
				mockRepo.On("GetProductsDetailById", "ERROR").Return(models.Product{}, errors.New("internal error"))
			},
			wantErr: errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			product, err := service.GetProductsDetailById(tt.id)

			if tt.wantErr != nil {
				assert.Error(t, err)
				if tt.wantErr == handleerrors.ErrNotFound {
					assert.ErrorIs(t, err, handleerrors.ErrNotFound)
				} else {
					assert.Equal(t, tt.wantErr.Error(), err.Error())
				}
				return
			}

			assert.NoError(t, err)
			if tt.check != nil {
				tt.check(t, product)
			}
		})
	}

	mockRepo.AssertExpectations(t)
}
