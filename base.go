package upcloud

import (
	api_result "github.com/wunderkraut/radi-api/result"
)

/**
 * Some base structs which other UpCloud Handler and
 * Operation implementations can include
 */

// Base handler with an upcloud service
type BaseUpcloudServiceHandler struct {
	factory         UpcloudFactory
	builderSettings *UpcloudBuilderSettings
}

// Constructor for BaseUpcloudServiceHandler
func New_BaseUpcloudServiceHandler(factory UpcloudFactory, builderSettings *UpcloudBuilderSettings) *BaseUpcloudServiceHandler {
	return &BaseUpcloudServiceHandler{
		factory:         factory,
		builderSettings: builderSettings,
	}
}

// Get the service
func (base *BaseUpcloudServiceHandler) BaseUpcloudServiceOperation() *BaseUpcloudServiceOperation {
	return New_BaseUpcloudServiceOperation(base.factory, base.builderSettings)
}

// Initialize and activate the Handler
func (base *BaseUpcloudServiceHandler) Validate() api_result.Result {
	return api_result.MakeSuccessfulResult()
}

// Get the factory
func (base *BaseUpcloudServiceHandler) Factory() UpcloudFactory {
	return base.factory
}

// Get the settings
func (base *BaseUpcloudServiceHandler) BuilderSettings() *UpcloudBuilderSettings {
	return base.builderSettings
}

/**
 * Base operations for Upcloud operations, which
 * allow sharing of Upcloud service across instances
 */

// Base operation with an upcloud service
type BaseUpcloudServiceOperation struct {
	factory         UpcloudFactory
	builderSettings *UpcloudBuilderSettings
}

// Constructor for BaseUpcloudServiceOperation
func New_BaseUpcloudServiceOperation(factory UpcloudFactory, builderSettings *UpcloudBuilderSettings) *BaseUpcloudServiceOperation {
	return &BaseUpcloudServiceOperation{
		factory:         factory,
		builderSettings: builderSettings,
	}
}

// Set the service
func (base *BaseUpcloudServiceOperation) ServiceWrapper() *UpcloudServiceWrapper {
	return base.factory.ServiceWrapper()
}

// Get the service
func (base *BaseUpcloudServiceOperation) ServerDefinitions() *ServerDefinitions {
	defs := base.factory.ServerDefinitions()
	return &defs
}

// Get the settings
func (base *BaseUpcloudServiceOperation) BuilderSettings() *UpcloudBuilderSettings {
	return base.builderSettings
}
