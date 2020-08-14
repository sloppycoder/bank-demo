package app

import (
	"context"
	"dashboard/api"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDashboard(t *testing.T) {
	os.Setenv("USE_CUST_SVC", "false")
	os.Setenv("USE_CASA_SVC", "false")
	serverCtx := newServerContext()

	server := &Server{serverCtx}
	req := &api.GetDashboardRequest{LoginName: "10001000"}
	resp, err := server.GetDashboard(context.Background(), req)

	assert.Nil(t, err, "GetDashboard returned error")
	assert.Equal(t, resp.Customer.LoginName, "skip")
}
