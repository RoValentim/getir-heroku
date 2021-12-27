package configs

import (
        "net/http"
)

const DbUrl = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"
const DbTimeout = 30
const DbCollection = "getir-case-study"
const DbTable = "records"

const DBClientFailedCode,
      DBClientFailedError,
      DBClientFailedStatusCode     = 200, "Server in maintenance at the moment", http.StatusInternalServerError

const DBConnectionFailedCode,
      DBConnectionFailedError,
      DBConnectionFailedStatusCode = 205, "Server in maintenance at the moment", http.StatusInternalServerError

const DBCursorErrorCode,
      DBCursorErrorError,
      DBCursorErrorStatusCode      = 210, "Server in maintenance at the moment", http.StatusInternalServerError
