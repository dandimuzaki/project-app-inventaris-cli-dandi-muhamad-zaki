package repository

// func TestGetCategories(t *testing.T) {
// 	// Mock DB and sqlmock
// 	mock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer mock.Close()

// 	rows := pgxmock.NewRows([]string{"ID", "category_name", "description"}).
// 		AddRow(1, "electronics", "alat elektronik").
// 		AddRow(2, "furnitures", "furnitur")

// 	mock.ExpectQuery("SELECT (.+) FROM categories").
// 		WillReturnRows(rows)

// 	repo := NewRepoCategory(mock)

// 	result, err := repo.GetCategories()

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)

// 	assert.Equal(t, 1, result[0].ID)
// 	assert.Equal(t, "electronics", result[0].Name)
// 	assert.Equal(t, "alat elektronik", result[0].Description)

// 	assert.Equal(t, 2, result[0].ID)
// 	assert.Equal(t, "furnitures", result[0].Name)
// 	assert.Equal(t, "furnitur", result[0].Description)

// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }