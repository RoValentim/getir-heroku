package configs

import (
        "net/http"
)

// Validation errors
const EndDatePreviousStartDateCode,
      EndDatePreviousStartDateError,
      EndDatePreviousStartDateStatusCode = 1, "End date need to be the same or greater than start date", http.StatusUnprocessableEntity

const MinHigherThanMaxCode,
      MinHigherThanMaxError,
      MinHigherThanMaxStatusCode         = 5, "minCount need to be greater than maxCount", http.StatusUnprocessableEntity

const KeyNotFoundCode,
      KeyNotFoundError,
      KeyNotFoundStatusCode              = 10, "Key value not found", http.StatusUnprocessableEntity

const TooManyKeysFoundCode,
      TooManyKeysFoundError,
      TooManyKeysFoundStatusCode         = 15, "Key value should be unique", http.StatusUnprocessableEntity


// Request errors
const MessageNotFound        = "Message not found"

const DataConflictCode,
      DataConflict,
      DataConflictStatusCode = 100, "Can not add the same key again", http.StatusConflict

const DataErrorCode,
      DataError,
      DataErrorStatusCode    = 105, "Incorrect json message", http.StatusUnprocessableEntity

const EmptyBodyCode,
      EmptyBodyError,
      EmptyBodyStatusCode    = 110, "Json message not found", http.StatusUnprocessableEntity


func ReturnMessage( code int ) (int, string) {
        switch code {
                // Validation errors
                case EndDatePreviousStartDateCode:
                        return EndDatePreviousStartDateStatusCode, EndDatePreviousStartDateError;
                case MinHigherThanMaxCode:
                        return MinHigherThanMaxStatusCode, MinHigherThanMaxError;
                case KeyNotFoundCode:
                        return KeyNotFoundStatusCode, KeyNotFoundError;
                case TooManyKeysFoundCode:
                        return TooManyKeysFoundStatusCode, TooManyKeysFoundError;

                // Request errors
                case DataConflictCode:
                        return DataConflictStatusCode, DataConflict;
                case DataErrorCode   :
                        return DataErrorStatusCode, DataError;
                case EmptyBodyCode   :
                        return EmptyBodyStatusCode, EmptyBodyError;

                // Database errors
                case DBClientFailedCode:
                        return DBClientFailedStatusCode, DBClientFailedError;
                case DBConnectionFailedCode:
                        return DBConnectionFailedStatusCode, DBConnectionFailedError;
                case DBCursorErrorCode:
                        return DBCursorErrorStatusCode, DBCursorErrorError;
        }

        return http.StatusInternalServerError, MessageNotFound
}
