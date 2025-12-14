package repository

// func TestTotalOrderPerMount(t *testing.T) {
// 	// Mock DB and sqlmock
// 	mock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer mock.Close()

// 	rows := pgxmock.NewRows([]string{"month", "total_order"}).
// 		AddRow("January", 10).
// 		AddRow("February", 20)

// 	mock.ExpectQuery("SELECT (.+) FROM orders WHERE status = \\$1 GROUP BY month ORDER BY MIN").
// 		WithArgs("completed").
// 		WillReturnRows(rows)

// 	repo := NewRepoReport(mock)

// 	result, err := repo.GetReportMonthly("completed")

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)

// 	assert.Equal(t, "January", result[0].Month)
// 	assert.Equal(t, 10, result[0].TotalOrder)

// 	assert.Equal(t, "February", result[1].Month)
// 	assert.Equal(t, 20, result[1].TotalOrder)

// 	err = mock.ExpectationsWereMet()
// 	assert.NoError(t, err)
// }

// func TestTotalOrderPerArea(t *testing.T) {
// 	mock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer mock.Close()

// 	rows := pgxmock.NewRows([]string{"area", "total_order"}).
// 		AddRow("Sukajadi", 10).
// 		AddRow("Antapani", 20)

// 	mock.ExpectQuery(`SELECT.*FROM ride_requests`).
// 		WillReturnRows(rows)

// 	repo := NewRepoReport(mock)

// 	result, err := repo.GetBusyAreas()

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)

// 	assert.Equal(t, "Sukajadi", result[0].Area)
// 	assert.Equal(t, 10, result[0].TotalOrder)

// 	assert.Equal(t, "Antapani", result[1].Area)
// 	assert.Equal(t, 20, result[1].TotalOrder)

// 	assert.NoError(t, mock.ExpectationsWereMet())
// }

// func TestTotalOrderPerHour(t *testing.T) {
// 	mock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer mock.Close()

// 	rows := pgxmock.NewRows([]string{"hour", "total_order"}).
// 		AddRow("06", 10).
// 		AddRow("11", 20)

// 	mock.ExpectQuery(`SELECT.*FROM ride_requests`).
// 		WillReturnRows(rows)

// 	repo := NewRepoReport(mock)

// 	result, err := repo.GetBusyTimes()

// 	assert.NoError(t, err)
// 	assert.Len(t, result, 2)

// 	assert.Equal(t, "06", result[0].Hour)
// 	assert.Equal(t, 10, result[0].TotalOrder)

// 	assert.Equal(t, "11", result[1].Hour)
// 	assert.Equal(t, 20, result[1].TotalOrder)

// 	assert.NoError(t, mock.ExpectationsWereMet())
// }
