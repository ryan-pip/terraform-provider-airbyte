// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSftpAuthenticationWildcardPasswordAuthentication struct {
	AuthMethod       types.String `tfsdk:"auth_method"`
	AuthUserPassword types.String `tfsdk:"auth_user_password"`
}
