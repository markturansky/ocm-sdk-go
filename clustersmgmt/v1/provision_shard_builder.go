/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ProvisionShardBuilder contains the data and logic needed to build 'provision_shard' objects.
//
// Contains the properties of the provision shard, including AWS and GCP related configurations
type ProvisionShardBuilder struct {
	id                       *string
	href                     *string
	link                     bool
	awsAccountOperatorConfig *ServerConfigBuilder
	awsBaseDomain            *string
	gcpBaseDomain            *string
	gcpProjectOperator       *ServerConfigBuilder
}

// NewProvisionShard creates a new builder of 'provision_shard' objects.
func NewProvisionShard() *ProvisionShardBuilder {
	return new(ProvisionShardBuilder)
}

// ID sets the identifier of the object.
func (b *ProvisionShardBuilder) ID(value string) *ProvisionShardBuilder {
	b.id = &value
	return b
}

// HREF sets the link to the object.
func (b *ProvisionShardBuilder) HREF(value string) *ProvisionShardBuilder {
	b.href = &value
	return b
}

// Link sets the flag that indicates if this is a link.
func (b *ProvisionShardBuilder) Link(value bool) *ProvisionShardBuilder {
	b.link = value
	return b
}

// AWSAccountOperatorConfig sets the value of the 'AWS_account_operator_config' attribute to the given value.
//
// Representation of a server config
func (b *ProvisionShardBuilder) AWSAccountOperatorConfig(value *ServerConfigBuilder) *ProvisionShardBuilder {
	b.awsAccountOperatorConfig = value
	return b
}

// AWSBaseDomain sets the value of the 'AWS_base_domain' attribute to the given value.
//
//
func (b *ProvisionShardBuilder) AWSBaseDomain(value string) *ProvisionShardBuilder {
	b.awsBaseDomain = &value
	return b
}

// GCPBaseDomain sets the value of the 'GCP_base_domain' attribute to the given value.
//
//
func (b *ProvisionShardBuilder) GCPBaseDomain(value string) *ProvisionShardBuilder {
	b.gcpBaseDomain = &value
	return b
}

// GCPProjectOperator sets the value of the 'GCP_project_operator' attribute to the given value.
//
// Representation of a server config
func (b *ProvisionShardBuilder) GCPProjectOperator(value *ServerConfigBuilder) *ProvisionShardBuilder {
	b.gcpProjectOperator = value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *ProvisionShardBuilder) Copy(object *ProvisionShard) *ProvisionShardBuilder {
	if object == nil {
		return b
	}
	b.id = object.id
	b.href = object.href
	b.link = object.link
	if object.awsAccountOperatorConfig != nil {
		b.awsAccountOperatorConfig = NewServerConfig().Copy(object.awsAccountOperatorConfig)
	} else {
		b.awsAccountOperatorConfig = nil
	}
	b.awsBaseDomain = object.awsBaseDomain
	b.gcpBaseDomain = object.gcpBaseDomain
	if object.gcpProjectOperator != nil {
		b.gcpProjectOperator = NewServerConfig().Copy(object.gcpProjectOperator)
	} else {
		b.gcpProjectOperator = nil
	}
	return b
}

// Build creates a 'provision_shard' object using the configuration stored in the builder.
func (b *ProvisionShardBuilder) Build() (object *ProvisionShard, err error) {
	object = new(ProvisionShard)
	object.id = b.id
	object.href = b.href
	object.link = b.link
	if b.awsAccountOperatorConfig != nil {
		object.awsAccountOperatorConfig, err = b.awsAccountOperatorConfig.Build()
		if err != nil {
			return
		}
	}
	object.awsBaseDomain = b.awsBaseDomain
	object.gcpBaseDomain = b.gcpBaseDomain
	if b.gcpProjectOperator != nil {
		object.gcpProjectOperator, err = b.gcpProjectOperator.Build()
		if err != nil {
			return
		}
	}
	return
}