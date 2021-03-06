/*
Copyright 2018 The v3io Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v3ioc

import (
	"context"
	"time"
)

// ControlplaneSession allows operations over a controlplane session
type Session interface {

	// CreateUser creates a user (blocking)
	CreateUserSync(*CreateUserInput) (*CreateUserOutput, error)

	// DeleteUserSync deletes a user (blocking)
	DeleteUserSync(*DeleteUserInput) error

	// CreateContainer creates a container (blocking)
	CreateContainerSync(*CreateContainerInput) (*CreateContainerOutput, error)

	// DeleteUserSync deletes a user (blocking)
	DeleteContainerSync(*DeleteContainerInput) error

	// UpdateClusterInfo updates a cluster info record (blocking)
	UpdateClusterInfoSync(*UpdateClusterInfoInput) (*UpdateCluserInfoOutput, error)
}

type ControlPlaneInput struct {
	ID        string
	IDNumeric int
	Ctx       context.Context
	Timeout   time.Duration
}

type ControlPlaneOutput struct {
	ID        string
	IDNumeric int
	Ctx       context.Context
}

// NewSessionInput specifies how to create a session
type NewSessionInput struct {
	ControlPlaneInput
	Endpoints []string
	SessionAttributes
}

// CreateUserInput specifies how to create a user
type CreateUserInput struct {
	ControlPlaneInput
	UserAttributes
}

// CreateUserOutput holds the response from creating a user
type CreateUserOutput struct {
	ControlPlaneOutput
	UserAttributes
}

// DeleteUserInput specifies how to delete a user
type DeleteUserInput struct {
	ControlPlaneInput
}

// CreateContainerInput specifies how to create a container
type CreateContainerInput struct {
	ControlPlaneInput
	ContainerAttributes
}

// CreateContainerOutput holds the response from creating a container
type CreateContainerOutput struct {
	ControlPlaneOutput
	ContainerAttributes
}

// DeleteContainerInput specifies how to delete a container
type DeleteContainerInput struct {
	ControlPlaneInput
}

// UpdateClusterInfoInput specifies how to update a cluster info record
type UpdateClusterInfoInput struct {
	ControlPlaneInput
	ClusterInfoAttributes
}

// UpdateCluserInfoOutput holds the response from updating a cluster info
type UpdateCluserInfoOutput struct {
	ControlPlaneOutput
	ClusterInfoAttributes
}
