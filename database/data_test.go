package database

import (
        "testing"
        "github.com/stretchr/testify/assert"
        "rodrigo/restful/configs"
)

func TestReturnFilters(t *testing.T) {
        correctData := []byte(`{
                        "startDate": "2016-01-01",
                        "endDate"  : "2016-01-31",
                        "minCount" : 2300,
                        "maxCount" : 3000
                }`)

        wrongData := []byte(`{
                        "startDate": "2016-01-01"
                        "endDate"  : "2016-01-31",
                        "minCount" : 2300,
                        "maxCount" : 3000
                }`)

        emptyData := []byte(`{}`)

        msgValueMatch := "The value should match"
        noError       := "Error need to be nil"

        filter, err := returnFilters(correctData)
        assert.Equal(t, 0, err, noError)

        assert.Equal(t, "2016-01-01", filter.StartDate, msgValueMatch)
        assert.Equal(t, "2016-01-31", filter.EndDate, msgValueMatch)
        assert.Equal(t, 2300, filter.MinCount, msgValueMatch)
        assert.Equal(t, 3000, filter.MaxCount, msgValueMatch)

        filter, err = returnFilters(wrongData)
        assert.Equal(t, configs.DataErrorCode, err, noError)
        assert.Empty(t, filter, msgValueMatch)

        filter, err = returnFilters(emptyData)
        assert.Equal(t, configs.EmptyBodyCode, err, noError)
        assert.Empty(t, filter, msgValueMatch)
}
