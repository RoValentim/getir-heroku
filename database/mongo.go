package database

import (
        "context"
        "time"

        "io/ioutil"
        "net/http"

        "go.mongodb.org/mongo-driver/bson"
        "go.mongodb.org/mongo-driver/bson/primitive"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"

        "rodrigo/restful/configs"
        "rodrigo/restful/internal"
)

func returnError(w http.ResponseWriter, code int) {
        var resp response

        statusCode, msg := configs.ReturnMessage(code)

        resp.Code = code
        resp.Msg = msg

        internal.ReturnData(w, resp, statusCode)
}

func applyFilterByCountsSum( min, max int, data []register ) []records {
        var results []records

        for _, value := range data {
                i := 0
                for _, v := range value.Counts {
                        i += v
                }

                if (i >= min) && (i <= max) {
                        results = append( results, records{Key:value.Key, CreatedAt:value.CreatedAt.Format("2006-01-02T15:04:05.000Z"), TotalCount:i} )
                }
         }

        return results
}

func getDataByDateByFilter( filter *datum ) (response, int) {
        var response response

        dateLayout   := "2006-01-02"
        startDate, _ := time.Parse(dateLayout, filter.StartDate)
        endDate, _   := time.Parse(dateLayout, filter.EndDate  )

        if endDate.Before(startDate) {
                return response, configs.EndDatePreviousStartDateCode
        }

        if filter.MinCount > filter.MaxCount {
                return response, configs.MinHigherThanMaxCode
        }

        client, err := mongo.NewClient(options.Client().ApplyURI(configs.DbUrl))
        if err != nil {
                return response, configs.DBClientFailedCode
        }

        ctx, _ := context.WithTimeout(context.Background(), configs.DbTimeout*time.Second)
        err = client.Connect(ctx)
        if err != nil {
                return response, configs.DBConnectionFailedCode
        }
        defer client.Disconnect(ctx)

        collection := client.Database(configs.DbCollection).Collection(configs.DbTable)
        cur, currErr := collection.Find(ctx, bson.M{ "createdAt": bson.M{"$gte":primitive.NewDateTimeFromTime(startDate),"$lt":primitive.NewDateTimeFromTime(endDate.AddDate(0,0,1))} })

        if currErr != nil {
                return response, configs.DBCursorErrorCode
        }
        defer cur.Close(ctx)

        var results []register
        if err = cur.All(ctx, &results); err != nil {
                return response, configs.DBCursorErrorCode
        }

        response.Code = 0
        response.Msg = "Success"
        response.Records = applyFilterByCountsSum(filter.MinCount, filter.MaxCount, results)

        return response, 0
}

func GetData(w http.ResponseWriter, r *http.Request) {
        reqBody, _ := ioutil.ReadAll(r.Body)

        filter, errCode := returnFilters(reqBody)
        if errCode != 0 {
                returnError(w, errCode)
                return
        }

        data, errCode := getDataByDateByFilter( filter )
        if errCode != 0 {
                returnError(w, errCode)
                return
        }

        internal.ReturnData(w, data, http.StatusOK)
}
