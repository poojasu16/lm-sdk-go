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

type AlertTrendsReport struct {
	LastmodifyUserId int32 `json:"lastmodifyUserId,omitempty"`
	Delivery string `json:"delivery,omitempty"`
	UserPermission string `json:"userPermission,omitempty"`
	LastGenerateOn int64 `json:"lastGenerateOn,omitempty"`
	ReportLinkNum int32 `json:"reportLinkNum,omitempty"`
	GroupId int32 `json:"groupId"`
	Format string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
	LastGenerateSize int64 `json:"lastGenerateSize,omitempty"`
	CustomReportTypeId int32 `json:"customReportTypeId,omitempty"`
	Type_ string `json:"type,omitempty"`
	LastGeneratePages int32 `json:"lastGeneratePages,omitempty"`
	Schedule string `json:"schedule,omitempty"`
	Recipients []ReportRecipient `json:"recipients,omitempty"`
	CustomReportTypeName string `json:"customReportTypeName,omitempty"`
	Name string `json:"name"`
	EnableViewAsOtherUser bool `json:"enableViewAsOtherUser,omitempty"`
	LastmodifyUserName string `json:"lastmodifyUserName,omitempty"`
	Id int32 `json:"id,omitempty"`
	DateRange string `json:"dateRange,omitempty"`
	Metrics []AlertTrendsMetric `json:"metrics"`
}
