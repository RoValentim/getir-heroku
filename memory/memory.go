package memory

import (
        "strings"

        "encoding/json"
        "io/ioutil"
        "net/http"

        "rodrigo/restful/configs"
        "rodrigo/restful/internal"
)

type dataError struct {
        Code      int      `json:"code"`
        Message   string   `json:"message"`
}

type memoryDatum struct {
        Key     string   `json:"key"`
        Value   string   `json:"value"`
}

var memoryData []memoryDatum

func returnError(w http.ResponseWriter, code int) {
        var resp dataError

        statusCode, msg := configs.ReturnMessage(code)

        resp.Code    = code
        resp.Message = msg

        internal.ReturnData(w, resp, statusCode)
}

func getDataById( id string ) memoryDatum {
        for _, datum := range memoryData {
                if datum.Key == id {
                        return datum
                }
        }

        return memoryDatum{}
}

func GetData(w http.ResponseWriter, r *http.Request) {
        ctx    := r.Context()
        params := ctx.Value("params").(map[string]string)

        if keys, ok := params["key"]; ok {
                keyArr := strings.Split(keys, ",")

                if len(keyArr) < 1 || keyArr[0] == "" {
                        returnError(w, configs.KeyNotFoundCode)
                        return
                }

                if len(keyArr) > 1 {
                        returnError(w, configs.TooManyKeysFoundCode)
                        return
                }

                data := getDataById( keyArr[0] )
                if data == (memoryDatum{}) {
                        internal.ReturnData(w, nil, http.StatusNoContent)
                        return
                }

                internal.ReturnData(w, data, http.StatusOK)
                return
        }

        // Since it's not required to return all values,
        // it's returning nil instead of memoryData
        internal.ReturnData(w, nil, http.StatusNoContent)
}

func CreateData(w http.ResponseWriter, r *http.Request) {
        reqBody, _ := ioutil.ReadAll(r.Body)

        var data memoryDatum
        if err := json.Unmarshal(reqBody, &data); err != nil {
                returnError(w, configs.DataErrorCode)
                return
        }

        if data == (memoryDatum{}) {
                returnError(w, configs.EmptyBodyCode)
                return
        }

        if data.Key == "" {
                returnError(w, configs.KeyNotFoundCode)
                return
        }

        duplicated := getDataById( data.Key )

        if duplicated != (memoryDatum{}) {
                returnError(w, configs.DataConflictCode)
                return
        }

        memoryData = append(memoryData, data)

        internal.ReturnData(w, data, http.StatusCreated)
}
