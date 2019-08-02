package headers

import (
	"bytes"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"hash/fnv"
	"strconv"
)

// AddETag adds response header
func AddETag(c *gin.Context, body interface{}) error {
	h := fnv.New32a()

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(body)
	if err != nil {
		return err
		// c.JSON(http.StatusInternalServerError, bad{Errors: []clienterr{{Title: "Failed to calculate ETag", Status: http.StatusInternalServerError}}})
	}
	h.Write([]byte(network.Bytes()))
	etag := h.Sum32()
	c.Writer.Header().Set("ETag", strconv.FormatUint(uint64(etag), 10))
	c.Header("ETag", strconv.FormatUint(uint64(etag), 10))
	return nil
}
