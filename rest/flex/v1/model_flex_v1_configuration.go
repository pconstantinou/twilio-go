/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// FlexV1Configuration struct for FlexV1Configuration
type FlexV1Configuration struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// An object that contains application-specific data
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Whether call recording is enabled
	CallRecordingEnabled *bool `json:"call_recording_enabled,omitempty"`
	// The call recording webhook URL
	CallRecordingWebhookUrl *string `json:"call_recording_webhook_url,omitempty"`
	// Flex Conversations channels' attachments configurations
	ChannelConfigs *[]map[string]interface{} `json:"channel_configs,omitempty"`
	// The SID of the chat service this user belongs to
	ChatServiceInstanceSid *string `json:"chat_service_instance_sid,omitempty"`
	// An object that contains the CRM attributes
	CrmAttributes *map[string]interface{} `json:"crm_attributes,omitempty"`
	// The CRM Callback URL
	CrmCallbackUrl *string `json:"crm_callback_url,omitempty"`
	// Whether CRM is present for Flex
	CrmEnabled *bool `json:"crm_enabled,omitempty"`
	// The CRM Fallback URL
	CrmFallbackUrl *string `json:"crm_fallback_url,omitempty"`
	// The CRM Type
	CrmType *string `json:"crm_type,omitempty"`
	// The ISO 8601 date and time in GMT when the Configuration resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Configuration resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Setting to enable Flex UI redirection
	FlexInsightsDrilldown *bool `json:"flex_insights_drilldown,omitempty"`
	// Object that controls workspace reporting
	FlexInsightsHr *map[string]interface{} `json:"flex_insights_hr,omitempty"`
	// The SID of the Flex service instance
	FlexServiceInstanceSid *string `json:"flex_service_instance_sid,omitempty"`
	// URL to redirect to in case drilldown is enabled.
	FlexUrl *string `json:"flex_url,omitempty"`
	// A list of objects that contain the configurations for the Integrations supported in this configuration
	Integrations *[]map[string]interface{} `json:"integrations,omitempty"`
	// Configurable parameters for Markdown
	Markdown *map[string]interface{} `json:"markdown,omitempty"`
	// The SID of the Messaging service instance
	MessagingServiceInstanceSid *string `json:"messaging_service_instance_sid,omitempty"`
	// Configurable parameters for Notifications
	Notifications *map[string]interface{} `json:"notifications,omitempty"`
	// The list of outbound call flows
	OutboundCallFlows *map[string]interface{} `json:"outbound_call_flows,omitempty"`
	// The plugin service attributes
	PluginServiceAttributes *map[string]interface{} `json:"plugin_service_attributes,omitempty"`
	// Whether the plugin service enabled
	PluginServiceEnabled *bool `json:"plugin_service_enabled,omitempty"`
	// The list of public attributes
	PublicAttributes *map[string]interface{} `json:"public_attributes,omitempty"`
	// Configurable parameters for Queues Statistics
	QueueStatsConfiguration *map[string]interface{} `json:"queue_stats_configuration,omitempty"`
	// The URL where the Flex instance is hosted
	RuntimeDomain *string `json:"runtime_domain,omitempty"`
	// The list of serverless service SIDs
	ServerlessServiceSids *[]string `json:"serverless_service_sids,omitempty"`
	// The Flex Service version
	ServiceVersion *string `json:"service_version,omitempty"`
	// The status of the Flex onboarding
	Status *string `json:"status,omitempty"`
	// The TaskRouter SID of the offline activity
	TaskrouterOfflineActivitySid *string `json:"taskrouter_offline_activity_sid,omitempty"`
	// The Skill description for TaskRouter workers
	TaskrouterSkills *[]map[string]interface{} `json:"taskrouter_skills,omitempty"`
	// The SID of the TaskRouter Target TaskQueue
	TaskrouterTargetTaskqueueSid *string `json:"taskrouter_target_taskqueue_sid,omitempty"`
	// The SID of the TaskRouter target Workflow
	TaskrouterTargetWorkflowSid *string `json:"taskrouter_target_workflow_sid,omitempty"`
	// The list of TaskRouter TaskQueues
	TaskrouterTaskqueues *[]map[string]interface{} `json:"taskrouter_taskqueues,omitempty"`
	// The TaskRouter Worker attributes
	TaskrouterWorkerAttributes *map[string]interface{} `json:"taskrouter_worker_attributes,omitempty"`
	// The TaskRouter default channel capacities and availability for workers
	TaskrouterWorkerChannels *map[string]interface{} `json:"taskrouter_worker_channels,omitempty"`
	// The SID of the TaskRouter Workspace
	TaskrouterWorkspaceSid *string `json:"taskrouter_workspace_sid,omitempty"`
	// The object that describes Flex UI characteristics and settings
	UiAttributes *map[string]interface{} `json:"ui_attributes,omitempty"`
	// The object that defines the NPM packages and versions to be used in Hosted Flex
	UiDependencies *map[string]interface{} `json:"ui_dependencies,omitempty"`
	// The primary language of the Flex UI
	UiLanguage *string `json:"ui_language,omitempty"`
	// The Pinned UI version
	UiVersion *string `json:"ui_version,omitempty"`
	// The absolute URL of the Configuration resource
	Url *string `json:"url,omitempty"`
}
