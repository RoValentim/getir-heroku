package database

import (
        "time"
        "testing"

        "github.com/stretchr/testify/assert"
        "go.mongodb.org/mongo-driver/bson/primitive"

        "rodrigo/restful/configs"
)

func TestApplyFilterByCountsSum(t *testing.T) {
        min       := 3100
        max       := 4000
        key       := "ZpoHRnZT"
        created   := time.Now()
        createdAt := created.Format("2006-01-02T15:04:05.000Z")

        msgValueMatch := "The value should match"
        noError       := "Error need to be nil"

        id, err := primitive.ObjectIDFromHex("1a2319ef96fb765873a24bae")
        assert.Equal(t, nil, err, noError)

        data := []register{
                register{Id: id, Counts: []int{1000}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000, 100}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000, 100, 200}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000, 100, 200, 300}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000, 100, 300, 300, 400}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
                register{Id: id, Counts: []int{1000, 2000, 100, 300, 300, 400, 500}, CreatedAt: created, Key: key, Value: "ZpoHRnZTZpoHRnZT"},
        }

        records := applyFilterByCountsSum( min, max, data )
        assert.Equal(t, 3, len(records), msgValueMatch)

        assert.Equal(t, key,       records[0].Key,        msgValueMatch)
        assert.Equal(t, createdAt, records[0].CreatedAt,  msgValueMatch)
        assert.Equal(t, 3100,      records[0].TotalCount, msgValueMatch)

        assert.Equal(t, key,       records[1].Key,        msgValueMatch)
        assert.Equal(t, createdAt, records[1].CreatedAt,  msgValueMatch)
        assert.Equal(t, 3300,      records[1].TotalCount, msgValueMatch)

        assert.Equal(t, key,       records[2].Key,        msgValueMatch)
        assert.Equal(t, createdAt, records[2].CreatedAt,  msgValueMatch)
        assert.Equal(t, 3600,      records[2].TotalCount, msgValueMatch)
}

func TestGetDataByDateByFilter(t *testing.T) {
        wrongDate := datum{
                StartDate: "2016-01-31",
                EndDate  : "2016-01-01",
                MinCount : 2300,
                MaxCount : 3000,
        }

        wrongCount := datum{
                StartDate: "2016-01-01",
                EndDate  : "2016-01-31",
                MinCount : 3000,
                MaxCount : 2300,
        }

        msgValueMatch := "The value should match"
        noError       := "Error need to be nil"

        response, err := getDataByDateByFilter(&wrongDate)
        assert.Equal(t, configs.EndDatePreviousStartDateCode, err, noError)
        assert.Equal(t, 0, len(response.Records), msgValueMatch)

        response, err = getDataByDateByFilter(&wrongCount)
        assert.Equal(t, configs.MinHigherThanMaxCode, err, noError)
        assert.Equal(t, 0, len(response.Records), msgValueMatch)
}
