package internal

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestIsDataNil(t *testing.T) {
        var interfaceNil interface{} = nil
        var interfaceNotNil interface{} = "value"

        msgValueMatch := "The value should match"

        assert.Equal(t, true,  isDataNil(interfaceNil),    msgValueMatch)
        assert.Equal(t, false, isDataNil(interfaceNotNil), msgValueMatch)
}
