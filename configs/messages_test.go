package configs

import (
        "testing"
        "net/http"
        "github.com/stretchr/testify/assert"
)

func TestReturnMessage(t *testing.T) {
        msgValueMatch := "The value should match"

        // Validation errors
        statusCode, message := ReturnMessage( EndDatePreviousStartDateCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, EndDatePreviousStartDateError,  message,      msgValueMatch)

        statusCode, message  = ReturnMessage( MinHigherThanMaxCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, MinHigherThanMaxError,          message,      msgValueMatch)

        statusCode, message  = ReturnMessage( KeyNotFoundCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, KeyNotFoundError,               message,      msgValueMatch)

        statusCode, message  = ReturnMessage( TooManyKeysFoundCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, TooManyKeysFoundError,          message,      msgValueMatch)


        // Request errors
        statusCode, message  = ReturnMessage( 9999 )
        assert.Equal(t, http.StatusInternalServerError, statusCode,   msgValueMatch)
        assert.Equal(t, MessageNotFound,                message,      msgValueMatch)

        statusCode, message  = ReturnMessage( DataConflictCode )
        assert.Equal(t, http.StatusConflict,            statusCode,   msgValueMatch)
        assert.Equal(t, DataConflict,                   message,      msgValueMatch)

        statusCode, message  = ReturnMessage( DataErrorCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, DataError,                      message,      msgValueMatch)

        statusCode, message  = ReturnMessage( EmptyBodyCode )
        assert.Equal(t, http.StatusUnprocessableEntity, statusCode,   msgValueMatch)
        assert.Equal(t, EmptyBodyError,                 message,      msgValueMatch)

        statusCode, message  = ReturnMessage( DBClientFailedCode )
        assert.Equal(t, http.StatusInternalServerError, statusCode,   msgValueMatch)
        assert.Equal(t, DBClientFailedError,            message,      msgValueMatch)


        // Database errors
        statusCode, message  = ReturnMessage( DBConnectionFailedCode )
        assert.Equal(t, http.StatusInternalServerError, statusCode,   msgValueMatch)
        assert.Equal(t, DBConnectionFailedError,        message,      msgValueMatch)

        statusCode, message  = ReturnMessage( DBCursorErrorCode )
        assert.Equal(t, http.StatusInternalServerError, statusCode,   msgValueMatch)
        assert.Equal(t, DBCursorErrorError,             message,      msgValueMatch)
}
