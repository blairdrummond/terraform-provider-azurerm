package client

import (
	// "github.com/Azure/azure-sdk-for-go/services/preview/authorization/mgmt/2020-04-01-preview/authorization"
	"github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2020-05-01-preview/billing"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	RoleAssignmentsClient *billing.RoleAssignmentsClient
	RoleDefinitionsClient *billing.RoleDefinitionsClient
}

func NewClient(o *common.ClientOptions) *Client {
	roleAssignmentsClient := billing.NewRoleAssignmentsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleAssignmentsClient.Client, o.ResourceManagerAuthorizer)

	roleDefinitionsClient := billing.NewRoleDefinitionsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&roleDefinitionsClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		RoleAssignmentsClient: &roleAssignmentsClient,
		RoleDefinitionsClient: &roleDefinitionsClient,
	}
}
