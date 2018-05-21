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

type AlertFilters struct {
	Severity string `json:"severity,omitempty"`
	Sdted string `json:"sdted,omitempty"`
	Chain string `json:"chain,omitempty"`
	Instance string `json:"instance,omitempty"`
	DataPoint string `json:"dataPoint,omitempty"`
	Host string `json:"host,omitempty"`
	Rule string `json:"rule,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	DataSource string `json:"dataSource,omitempty"`
	Acked string `json:"acked,omitempty"`
	Cleared string `json:"cleared,omitempty"`
	Group string `json:"group,omitempty"`
}
