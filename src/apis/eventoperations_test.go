package apis

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer_AddNewEvent(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		s    Server
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.AddNewEvent(tt.args.c)
		})
	}
}