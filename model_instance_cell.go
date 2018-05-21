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

type InstanceCell struct {
	ValidationStatusCode int32 `json:"validationStatusCode,omitempty"`
	DataPointName string `json:"dataPointName,omitempty"`
	InstanceId int32 `json:"instanceId"`
	DataPointId int32 `json:"dataPointId,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
}
