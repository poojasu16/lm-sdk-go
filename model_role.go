/*
 * logicmonitor_sdk
 *
 * LogicMonitor is a SaaS-based performance monitoring platform that provides full visibility into complex, hybrid infrastructures, offering granular performance monitoring and actionable data and insights. logicmonitor_sdk enables you to manage your LogicMonitor account programmatically.
 *
 * API version: 1.0.0
 * Contact: sdk@logicmonitor.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package logicmonitor

type Role struct {
	EnableRemoteSessionInCompanyLevel bool `json:"enableRemoteSessionInCompanyLevel,omitempty"`
	CustomHelpLabel string `json:"customHelpLabel,omitempty"`
	CustomHelpURL string `json:"customHelpURL,omitempty"`
	Privileges []Privilege `json:"privileges"`
	AssociatedUserCount int32 `json:"associatedUserCount,omitempty"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Id int32 `json:"id,omitempty"`
	TwoFARequired bool `json:"twoFARequired,omitempty"`
	RequireEULA bool `json:"requireEULA,omitempty"`
	AcctRequireTwoFA bool `json:"acctRequireTwoFA,omitempty"`
}
