package database

import (
        "time"
        "encoding/json"
        "go.mongodb.org/mongo-driver/bson/primitive"
        "rodrigo/restful/configs"
)

type datum struct {
        StartDate   string   `json:"startDate"`
        EndDate     string   `json:"endDate"`
        MinCount    int      `json:"minCount"`
        MaxCount    int      `json:"maxCount"`
}

type records struct {
        Key          string   `json:"key"`
        CreatedAt    string   `json:"createdAt"`
        TotalCount   int      `json:"totalCount"`
}

type response struct {
        Code      int         `json:"code"`
        Msg       string      `json:"msg"`
        Records   []records   `json:"records"`
}

type register struct {
        Id          primitive.ObjectID   `bson:"_id"`
        Counts      []int                `bson:"counts"`
        CreatedAt   time.Time            `bson:"createdAt"`
        Key         string               `bson:"key"`
        Value       string               `bson:"value"`
}

func returnFilters(data []byte) (*datum, int) {
        var filter datum

        if err := json.Unmarshal(data, &filter); err != nil {
                if err.Error() == "unexpected end of JSON input" {
                        return nil, configs.EmptyBodyCode
                }

                return nil, configs.DataErrorCode
        }

        if filter == (datum{}) {
                return nil, configs.EmptyBodyCode
        }

        return &filter, 0
}
