package products

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	handleerrors "github.com/jnka9755/go-detail-product-api/src/products/handle-errors"
	"github.com/jnka9755/go-detail-product-api/src/products/models"
	"github.com/stretchr/testify/assert"
)

func setupTestRepo(t *testing.T) (Repository, func()) {
	// Create a temporary directory
	tmpDir, err := ioutil.TempDir("", "product-api-test")
	if err != nil {
		t.Fatal(err)
	}

	// Copy test files to temp directory
	files := []string{"products.json", "categories.json", "sellers.json"}
	for _, f := range files {
		content, err := ioutil.ReadFile(filepath.Join("testdata", f))
		if err != nil {
			t.Fatal(err)
		}
		if err := ioutil.WriteFile(filepath.Join(tmpDir, f), content, 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Create test logger that writes to null
	logger := log.New(ioutil.Discard, "", 0)

	// Change working directory to temp dir
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}

	repo := NewRepository(logger, ".")

	// Return cleanup function
	cleanup := func() {
		os.Chdir(oldWd)
		os.RemoveAll(tmpDir)
	}

	return repo, cleanup
}

func TestGetProductsDetail(t *testing.T) {
	repo, cleanup := setupTestRepo(t)
	defer cleanup()

	products := repo.GetProductsDetail()
	assert.Len(t, products, 2, "Should return 2 products")

	// Verify first product
	assert.Equal(t, "TEST1", products[0].ID)
	assert.Equal(t, "Test Product 1", products[0].Title)
	assert.Equal(t, 100.0, products[0].Price)

	// Verify category enrichment
	assert.Equal(t, "Test Category 1", products[0].Category.Name)

	// Verify seller enrichment
	assert.Equal(t, "Test Seller 1", products[0].Seller.Name)
	assert.Equal(t, 4.5, products[0].Seller.Reputation)
}

func TestGetProductsDetailById(t *testing.T) {
	repo, cleanup := setupTestRepo(t)
	defer cleanup()

	tests := []struct {
		name    string
		id      string
		wantErr error
		check   func(*testing.T, models.Product)
	}{
		{
			name: "existing product",
			id:   "TEST1",
			check: func(t *testing.T, p models.Product) {
				assert.Equal(t, "TEST1", p.ID)
				assert.Equal(t, "Test Product 1", p.Title)
				assert.Equal(t, "Test Category 1", p.Category.Name)
				assert.Equal(t, "Test Seller 1", p.Seller.Name)
			},
		},
		{
			name:    "non-existing product",
			id:      "NONEXISTENT",
			wantErr: handleerrors.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := repo.GetProductsDetailById(tt.id)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			assert.NoError(t, err)
			if tt.check != nil {
				tt.check(t, product)
			}
		})
	}
}

func TestMissingFiles(t *testing.T) {
	// Create empty temp dir
	tmpDir, err := ioutil.TempDir("", "product-api-test-missing")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	oldWd, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldWd)

	logger := log.New(ioutil.Discard, "", 0)
	repo := NewRepository(logger, ".")

	// Should return empty list when files are missing
	products := repo.GetProductsDetail()
	assert.Len(t, products, 0, "Should return empty list when files are missing")
}

func TestMalformedJSON(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "product-api-test-malformed")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Write malformed JSON
	if err := ioutil.WriteFile(filepath.Join(tmpDir, "products.json"), []byte("{malformed"), 0644); err != nil {
		t.Fatal(err)
	}

	oldWd, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(oldWd)

	logger := log.New(ioutil.Discard, "", 0)
	repo := NewRepository(logger, ".")

	// Should handle malformed JSON gracefully
	products := repo.GetProductsDetail()
	assert.Len(t, products, 0, "Should return empty list for malformed JSON")
}
