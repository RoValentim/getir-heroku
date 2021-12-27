package internal

import (
        "reflect"
        "encoding/json"
        "net/http"
)

func isDataNil(data interface{}) bool {
        if data == nil {
                return true
        }

        switch reflect.TypeOf(data).Kind() {
                case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
                        return reflect.ValueOf(data).IsNil()
        }

        return false
}

func ReturnData(w http.ResponseWriter, data interface{}, statusCode int ){
        if statusCode >= 200 && statusCode < 300 && isDataNil(data) {
                w.WriteHeader( http.StatusNoContent )
                return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader( statusCode )
        json.NewEncoder(w).Encode(data)
}

func Info(w http.ResponseWriter, data interface{}, statusCode int ){
        if statusCode >= 200 && statusCode < 300 && isDataNil(data) {
                w.WriteHeader( http.StatusNoContent )
                return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader( statusCode )
        json.NewEncoder(w).Encode(data)
}
