// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Sumo Logic and manual
//     changes will be clobbered when the file is regenerated. Do not submit
//     changes to this file.
//
// ----------------------------------------------------------------------------
package sumologic

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSumologicMonitorsLibraryMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicMonitorsLibraryMonitorCreate,
		Read:   resourceSumologicMonitorsLibraryMonitorRead,
		Update: resourceSumologicMonitorsLibraryMonitorUpdate,
		Delete: resourceSumologicMonitorsLibraryMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{

			"version": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: false,
			},

			"modified_at": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"is_system": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"queries": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"row_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"query": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"created_by": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"is_disabled": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: false,
				// Default:  false,
			},

			"is_mutable": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

			"triggers": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trigger_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"threshold": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"threshold_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time_range": {
							Type:     schema.TypeString,
							Required: true,
						},
						"trigger_source": {
							Type:     schema.TypeString,
							Required: true,
						},
						"occurrence_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"detection_method": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"notifications": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"notification": {
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"subject": {
										Type:     schema.TypeString,
										Required: true,
									},
									"recipients": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"message_body": {
										Type:     schema.TypeString,
										Required: true,
									},
									"time_zone": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
						"notification_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"run_for_trigger_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"created_at": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"monitor_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"is_locked": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
			},

			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"modified_by": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"post_request_map": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSumologicMonitorsLibraryMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	if d.Id() == "" {
		monitor := resourceToMonitorsLibraryMonitor(d)
		paramMap := make(map[string]string)
		paramMap["parentId"] = "0000000000000001"
		monitorDefinitionID, err := c.CreateMonitorsLibraryMonitor(monitor, paramMap)
		if err != nil {
			return err
		}

		d.SetId(monitorDefinitionID)
	}
	return resourceSumologicMonitorsLibraryMonitorRead(d, meta)
}

func resourceSumologicMonitorsLibraryMonitorRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	// id, _ := strconv.Atoi(d.Id())
	monitor, err := c.MonitorsRead(d.Id())
	if err != nil {
		return err
	}

	if monitor == nil {
		log.Printf("[WARN] Monitor not found, removing from state: %v - %v", d.Id(), err)
		d.SetId("")
		return nil
	}
	return nil
}

func resourceSumologicMonitorsLibraryMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	monitor := resourceToMonitorsLibraryMonitor(d)
	err := c.UpdateMonitorsLibraryMonitor(monitor)
	if err != nil {
		return err
	}
	return resourceSumologicMonitorsLibraryMonitorRead(d, meta)
}

func resourceSumologicMonitorsLibraryMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)
	monitor := resourceToMonitorsLibraryMonitor(d)
	err := c.DeleteMonitorsLibraryMonitor(monitor.ID)
	if err != nil {
		return err
	}
	return nil
}

func resourceToMonitorsLibraryMonitor(d *schema.ResourceData) MonitorsLibraryMonitor {
	// handle notifications
	rawNotifications := d.Get("notifications").([]interface{})
	notifications := make([]MonitorNotification, len(rawNotifications))
	for i := range rawNotifications {
		notificationDict := rawNotifications[i].(map[string]interface{})
		n := MonitorNotification{}
		notificationAction := EmailNotification{}
		rawNotificationAction := notificationDict["notification"].([]interface{})
		notificationActionDict := rawNotificationAction[0].(map[string]interface{})
		notificationAction.MessageBody = notificationActionDict["message_body"].(string)
		// var recipients []string
		notificationAction.Recipients = notificationActionDict["recipients"].([]interface{})
		notificationAction.Subject = notificationActionDict["subject"].(string)
		notificationAction.ActionType = notificationActionDict["action_type"].(string)
		n.NotificationType = notificationDict["notification_type"].(string)
		n.Notification = notificationAction
		n.RunForTriggerTypes = notificationDict["run_for_trigger_types"].([]interface{})
		// n.RunForTriggerTypes = n.RunForTriggerTypes[0].([]string)
		// n.Notification
		if n.NotificationType == "EmailAction" {
			log.Printf("[DEBUG] Found notification type EmailAction")
		}
		log.Printf("[DEBUG] Notification %v", n)
		notifications[i] = n
	}
	// handle triggers
	rawTriggers := d.Get("triggers").([]interface{})
	triggers := make([]TriggerCondition, len(rawTriggers))
	for i := range rawTriggers {
		triggerDict := rawTriggers[i].(map[string]interface{})
		t := TriggerCondition{}
		t.TriggerType = triggerDict["trigger_type"].(string)
		t.Threshold = triggerDict["threshold"].(float64)
		t.ThresholdType = triggerDict["threshold_type"].(string)
		t.TimeRange = triggerDict["time_range"].(string)
		t.OccurrenceType = triggerDict["occurrence_type"].(string)
		t.TriggerSource = triggerDict["trigger_source"].(string)
		t.DetectionMethod = triggerDict["detection_method"].(string)
		triggers[i] = t
	}
	// handle queries
	rawQueries := d.Get("queries").([]interface{})
	queries := make([]MonitorQuery, len(rawQueries))
	for i := range rawQueries {
		queryDict := rawQueries[i].(map[string]interface{})
		q := MonitorQuery{}
		q.Query = queryDict["query"].(string)
		q.RowID = queryDict["row_id"].(string)
		queries[i] = q
	}

	return MonitorsLibraryMonitor{
		CreatedBy:     d.Get("created_by").(string),
		Name:          d.Get("name").(string),
		ID:            d.Id(),
		CreatedAt:     d.Get("created_at").(string),
		MonitorType:   d.Get("monitor_type").(string),
		Description:   d.Get("description").(string),
		Queries:       queries,
		ModifiedBy:    d.Get("modified_by").(string),
		IsMutable:     d.Get("is_mutable").(bool),
		Version:       d.Get("version").(int),
		Notifications: notifications,
		Type:          d.Get("type").(string),
		ParentId:      d.Get("parent_id").(string),
		ModifiedAt:    d.Get("modified_at").(string),
		Triggers:      triggers,
		ContentType:   d.Get("content_type").(string),
		IsLocked:      d.Get("is_locked").(bool),
		IsSystem:      d.Get("is_system").(bool),
		IsDisabled:    d.Get("is_disabled").(bool),
	}
}
