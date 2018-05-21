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

type OpsNoteServiceScopeV1 struct {
	FullPath string `json:"fullPath,omitempty"`
	GroupId int32 `json:"groupId,omitempty"`
	Type_ string `json:"type"`
	ServiceId int32 `json:"serviceId,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}
