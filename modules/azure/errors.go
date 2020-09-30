package azure

import "fmt"

// SubscriptionIDNotFound is an error that occurs when the Azure Subscription ID could not be found or was not provided
type SubscriptionIDNotFound struct{}

func (err SubscriptionIDNotFound) Error() string {
	return fmt.Sprintf("Could not find an Azure Subscription ID in expected environment variable %s and one was not provided for this test.", AzureSubscriptionID)
}

// ResourceGroupNameNotFound is an error that occurs when the target Azure Resource Group name could not be found or was not provided
type ResourceGroupNameNotFound struct{}

func (err ResourceGroupNameNotFound) Error() string {
	return fmt.Sprintf("Could not find an Azure Resource Group name in expected environment variable %s and one was not provided for this test.", AzureResGroupName)
}

// NotFoundError is returned when an expected object is not found in the search spa
type NotFoundError struct {
	objectType  string
	objectID    string
	searchSpace string
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("Object of type %s with id %s not found in %s", err.objectType, err.objectID, err.searchSpace)
}

// NewNotFoundError creates a new not found error when an expected object is not found in the search space
func NewNotFoundError(objectType string, objectID string, region string) NotFoundError {
	return NotFoundError{objectType, objectID, region}
}
