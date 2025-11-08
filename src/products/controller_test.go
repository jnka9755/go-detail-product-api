package products

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	handleerrors "github.com/jnka9755/go-detail-product-api/src/products/handle-errors"
	"github.com/jnka9755/go-detail-product-api/src/products/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock service
type mockService struct {
	mock.Mock
}

func (m *mockService) GetProductsDetail() []models.Product {
	args := m.Called()
	return args.Get(0).([]models.Product)
}

func (m *mockService) GetProductsDetailById(id string) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func setupTestController() (*mockService, *mux.Router) {
	service := new(mockService)
	endpoints := MakeEndpoints(service)

	router := mux.NewRouter()
	router.HandleFunc("/products", endpoints.GetProductsDetail).Methods("GET")
	router.HandleFunc("/product/{id}", endpoints.GetProductsDetailById).Methods("GET")

	return service, router
}

func TestGetProductsDetailEndpoint(t *testing.T) {
	mockService, router := setupTestController()

	tests := []struct {
		name           string
		mockSetup      func()
		expectedStatus int
		checkResponse  func(*testing.T, []byte)
	}{
		{
			name: "success with products",
			mockSetup: func() {
				products := []models.Product{
					{
						ID:    "TEST1",
						Title: "Test Product 1",
					},
					{
						ID:    "TEST2",
						Title: "Test Product 2",
					},
				}
				mockService.On("GetProductsDetail").Return(products)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var products []models.Product
				err := json.Unmarshal(body, &products)
				assert.NoError(t, err)
				assert.Len(t, products, 2)
				assert.Equal(t, "TEST1", products[0].ID)
				assert.Equal(t, "TEST2", products[1].ID)
			},
		},
		{
			name: "success empty list",
			mockSetup: func() {
				mockService.On("GetProductsDetail").Return([]models.Product{})
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var products []models.Product
				err := json.Unmarshal(body, &products)
				assert.NoError(t, err)
				assert.Len(t, products, 0)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService.ExpectedCalls = nil // limpia las expectativas previas
			mockService.Calls = nil
			tt.mockSetup()

			req := httptest.NewRequest("GET", "/products", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			if tt.checkResponse != nil {
				tt.checkResponse(t, body)
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetProductDetailByIdEndpoint(t *testing.T) {
	mockService, router := setupTestController()

	tests := []struct {
		name           string
		productID      string
		mockSetup      func()
		expectedStatus int
		checkResponse  func(*testing.T, []byte)
	}{
		{
			name:      "existing product",
			productID: "TEST1",
			mockSetup: func() {
				product := models.Product{
					ID:    "TEST1",
					Title: "Test Product 1",
				}
				mockService.On("GetProductsDetailById", "TEST1").Return(product, nil)
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, body []byte) {
				var product models.Product
				err := json.Unmarshal(body, &product)
				assert.NoError(t, err)
				assert.Equal(t, "TEST1", product.ID)
			},
		},
		{
			name:      "non-existing product",
			productID: "NONEXISTENT",
			mockSetup: func() {
				mockService.On("GetProductsDetailById", "NONEXISTENT").Return(models.Product{}, handleerrors.ErrNotFound)
			},
			expectedStatus: http.StatusNotFound,
			checkResponse: func(t *testing.T, body []byte) {
				var resp ErrorResponse
				err := json.Unmarshal(body, &resp)
				assert.NoError(t, err)
				assert.Contains(t, resp.Error, "not found")
			},
		},
		{
			name:      "internal error",
			productID: "ERROR",
			mockSetup: func() {
				mockService.On("GetProductsDetailById", "ERROR").Return(models.Product{}, errors.New("internal error"))
			},
			expectedStatus: http.StatusInternalServerError,
			checkResponse: func(t *testing.T, body []byte) {
				var resp ErrorResponse
				err := json.Unmarshal(body, &resp)
				assert.NoError(t, err)
				assert.Contains(t, resp.Error, "internal error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest("GET", "/product/"+tt.productID, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			if tt.checkResponse != nil {
				tt.checkResponse(t, body)
			}
		})
	}

	mockService.AssertExpectations(t)
}
