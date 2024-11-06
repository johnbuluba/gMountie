package service

import "gmountie/pkg/client"

// AppService is a service for managing the application
type AppService interface {
	// GetContext returns the app context
	GetContext() *client.AppContext
	// SetContext sets the app context
	SetContext(*client.AppContext)
}

// AppServiceImpl is the implementation of the AppService
type AppServiceImpl struct {
	appCtx *client.AppContext
}

// NewAppService creates a new AppService
func NewAppService() AppService {
	return &AppServiceImpl{}
}

// GetContext returns the app context
func (a *AppServiceImpl) GetContext() *client.AppContext {
	return a.appCtx
}

// SetContext sets the app context
func (a *AppServiceImpl) SetContext(appCtx *client.AppContext) {
	a.appCtx = appCtx
}
