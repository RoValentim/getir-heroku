package memory

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestGetDataById(t *testing.T) {
        id := "ZpoHRnZT"

        memoryData = append(memoryData, memoryDatum{Key: "0", Value: "Value 0"})
        memoryData = append(memoryData, memoryDatum{Key: id, Value: "Value ID"})
        memoryData = append(memoryData, memoryDatum{Key: "A", Value: "Value A"})

        msgValueMatch := "The value should match"

        data := getDataById(id)

        assert.NotEqual(t, memoryData[0].Key,   data.Key,   msgValueMatch)
        assert.NotEqual(t, memoryData[0].Value, data.Value, msgValueMatch)

        assert.Equal(t, memoryData[1].Key,   data.Key,   msgValueMatch)
        assert.Equal(t, memoryData[1].Value, data.Value, msgValueMatch)

        data = getDataById("NoMatch")
        assert.Equal(t, "", data.Key,   msgValueMatch)
        assert.Equal(t, "", data.Value, msgValueMatch)

}
