package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
	"net/http"
	"strings"
	"time"
)

func ResourceK8s() *schema.Resource {
	return &schema.Resource{
		Create: resourceK8sCreate,
		Read:   resourceK8sRead,
		Update: resourceK8sUpdate,
		Delete: resourceK8sDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf(
						"unexpected format of ID (%q), expected ProjectID:VolumeID", d.Id())
				}
				projectID := idParts[0]
				clusterID := idParts[1]
				d.SetId(clusterID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipip_mode": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"k8s_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"enabled_lb": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"master_instance_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_instance_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etcd_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"etcd_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"boot_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"boot_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"docker_volume_size": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"docker_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"calico_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"k8s_network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"security_group": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"min_node_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"max_node_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"auto_scaling": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"auto_healing": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"ingress_controller": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"auto_monitoring": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"node_group_list": {
				Type: schema.TypeList,
				Elem: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
					"node_amount": {
						Type:     schema.TypeInt,
						Required: true,
					},
					"flavor_id": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
			},
			"ssh_key_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_flavor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"master_flavor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_point": {
				Type:     schema.TypeString,
				Computed: true,
			},
			//"master_secgroup_default": {
			//	Computed: true,
			//	Type:     schema.TypeMap,
			//	Elem: map[string]*schema.Schema{
			//		"name": {
			//			Type: schema.TypeString,
			//		},
			//		"rules": {
			//			Type: schema.TypeList,
			//			Elem: map[string]*schema.Schema{
			//				"direction": {
			//					Type: schema.TypeString,
			//				},
			//				"protocol": {
			//					Type: schema.TypeString,
			//				},
			//				"ethertype": {
			//					Type: schema.TypeString,
			//				},
			//				"port_range_max": {
			//					Type: schema.TypeInt,
			//				},
			//				"port_range_min": {
			//					Type: schema.TypeInt,
			//				},
			//				"remote_ip_prefix": {
			//					Type: schema.TypeString,
			//				},
			//				"remote_group_name": {
			//					Type: schema.TypeString,
			//				},
			//			},
			//		},
			//	},
			//},
		},
	}
}

func buildCreateNodeGroupRequests(d *schema.ResourceData) []vserver.NodeGroupRequestModel {
	nodeGroupList := d.Get("node_group_list").([]interface{})
	NodeGroups := make([]vserver.NodeGroupRequestModel, len(nodeGroupList))

	for _, nodeGroup := range nodeGroupList {
		nodeGroup := nodeGroup.(map[string]interface{})
		var request vserver.NodeGroupRequestModel
		request.Name = nodeGroup["name"].(string)
		request.FlavorId = nodeGroup["flavor_id"].(string)
		request.NodeCount = nodeGroup["node_amount"].(int)
		NodeGroups = append(NodeGroups, request)
	}
	return NodeGroups
}

func resourceK8sCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create K8s")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	var secGroupIdList []string
	for _, s := range d.Get("security_group").([]string) {
		secGroupIdList = append(secGroupIdList, s)
	}

	cluster := vserver.CreateClusterRequest{
		AutoHealingEnabled:        d.Get("auto_healing").(bool),
		AutoMonitoringEnabled:     d.Get("auto_monitoring").(bool),
		AutoScalingEnabled:        d.Get("auto_scaling").(bool),
		BootVolumeSize:            d.Get("boot_volume_size").(int32),
		BootVolumeTypeId:          d.Get("boot_volume_type_id").(string),
		CalicoCidr:                d.Get("calico_cidr").(string),
		Description:               d.Get("description").(string),
		DockerVolumeSize:          d.Get("docker_volume_size").(int32),
		DockerVolumeTypeId:        d.Get("docker_volume_type_id").(string),
		EnabledLb:                 d.Get("enabled_lb").(bool),
		EtcdVolumeSize:            d.Get("etcd_volume_size").(int32),
		EtcdVolumeTypeId:          d.Get("etcd_volume_type_id").(string),
		IngressControllerEnabled:  d.Get("ingress_controller").(bool),
		IpipMode:                  d.Get("ipip_mode").(string),
		K8sVersion:                d.Get("k8s_version").(string),
		MasterCount:               d.Get("master_count").(int32),
		MasterInstanceTypeId:      d.Get("master_instance_type_id").(string),
		MaxNodeCount:              d.Get("max_node_count").(int32),
		MinNodeCount:              d.Get("min_node_count").(int32),
		Name:                      d.Get("name").(string),
		NetworkId:                 d.Get("network_id").(string),
		NetworkType:               d.Get("network_type").(string),
		NodeCount:                 d.Get("node_count").(int32),
		NodeFlavorId:              d.Get("node_flavor_id").(string),
		NodeInstanceTypeId:        d.Get("node_instance_type_id").(string),
		SecGroupIds:               secGroupIdList,
		SshKeyId:                  d.Get("ssh_key_id").(string),
		SubnetId:                  d.Get("subnet_id").(string),
		NodeGroupRequestModelList: buildCreateNodeGroupRequests(d),
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.CreateClusterUsingPOST(context.TODO(), cluster, projectId)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsh: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	stateConf := &resource.StateChangeConf{
		Pending:    k8sClusterCreating,
		Target:     k8sClusterCreated,
		Refresh:    resourceK8sClusterStateRefreshFunc(cli, resp.Data.Uuid, projectId),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for create cluster (%s) %s", resp.Data.Uuid, err)
	}

	d.SetId(resp.Data.Uuid)
	return resourceK8sRead(d, m)
}

func resourceK8sRead(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	clusterId := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), projectId, clusterId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	cluster := resp.Data
	d.Set("name", cluster.Name)
	d.Set("master_count", cluster.MasterCount)
	d.Set("node_count", cluster.NodeCount)
	d.Set("master_flavor_name", cluster.MasterFlavorName)
	d.Set("node_flavor_name", cluster.NodeFlavorName)
	d.Set("ssh_key_name", cluster.SshKeyName)
	d.Set("etcd_volume_size", cluster.EtcdVolumeSize)
	d.Set("boot_volume_size", cluster.BootVolumeSize)
	d.Set("docker_volume_size", cluster.DockerVolumeSize)
	d.Set("k8s_network_type", cluster.K8sNetworkType)
	d.Set("k8s_version_name", cluster.K8sVersion)
	d.Set("auto_healing", cluster.AutoHealingEnabled)
	d.Set("auto_monitoring", cluster.AutoMonitoringEnabled)
	d.Set("auto_scaling", cluster.AutoScalingEnabled)
	d.Set("end_point", cluster.Endpoint)
	return nil
}

func resourceK8sUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("node_group_list") {
		//o, n := d.GetChange("node_group_list")
		//return resourceNodeGroupResize(d, m)
	}

	return resourceK8sRead(d, m)
}

func resourceK8sDelete(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.DeleteClusterUsingDELETE(context.TODO(), d.Id(), projectId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsh: %s", responseBody)
		return errorResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    k8sClusterDeleting,
		Target:     k8sClusterDeleted,
		Refresh:    resourceK8sDeleteStateRefreshFunc(cli, d.Id(), projectId),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for delete cluster (%s) : %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

//func resourceNodeGroupResize(d *schema.ResourceData, m interface{}) error {
//
//}

func resourceK8sClusterStateRefreshFunc(cli *client.Client, clusterId string, projectId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpReponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), projectId, clusterId)

		if httpReponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("error describing: %s", GetResponseBody(httpReponse))
		}

		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		cluster := resp.Data
		return cluster, cluster.Status, nil
	}
}

func resourceK8sDeleteStateRefreshFunc(cli *client.Client, clusterId string, projectId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), clusterId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.ClusterDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing instanceL: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		cluster := resp.Data
		return cluster, cluster.Status, nil
	}
}
