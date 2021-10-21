// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"reflect"
)

const ComputeServiceAttachmentAssetType string = "compute.googleapis.com/ServiceAttachment"

func resourceConverterComputeServiceAttachment() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeServiceAttachmentAssetType,
		Convert:   GetComputeServiceAttachmentCaiObject,
	}
}

func GetComputeServiceAttachmentCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/serviceAttachments/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeServiceAttachmentApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeServiceAttachmentAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "ServiceAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeServiceAttachmentApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeServiceAttachmentName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeServiceAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	fingerprintProp, err := expandComputeServiceAttachmentFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	connectionPreferenceProp, err := expandComputeServiceAttachmentConnectionPreference(d.Get("connection_preference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("connection_preference"); !isEmptyValue(reflect.ValueOf(connectionPreferenceProp)) && (ok || !reflect.DeepEqual(v, connectionPreferenceProp)) {
		obj["connectionPreference"] = connectionPreferenceProp
	}
	targetServiceProp, err := expandComputeServiceAttachmentTargetService(d.Get("target_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_service"); !isEmptyValue(reflect.ValueOf(targetServiceProp)) && (ok || !reflect.DeepEqual(v, targetServiceProp)) {
		obj["targetService"] = targetServiceProp
	}
	natSubnetsProp, err := expandComputeServiceAttachmentNatSubnets(d.Get("nat_subnets"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("nat_subnets"); ok || !reflect.DeepEqual(v, natSubnetsProp) {
		obj["natSubnets"] = natSubnetsProp
	}
	enableProxyProtocolProp, err := expandComputeServiceAttachmentEnableProxyProtocol(d.Get("enable_proxy_protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_proxy_protocol"); !isEmptyValue(reflect.ValueOf(enableProxyProtocolProp)) && (ok || !reflect.DeepEqual(v, enableProxyProtocolProp)) {
		obj["enableProxyProtocol"] = enableProxyProtocolProp
	}
	consumerRejectListsProp, err := expandComputeServiceAttachmentConsumerRejectLists(d.Get("consumer_reject_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_reject_lists"); ok || !reflect.DeepEqual(v, consumerRejectListsProp) {
		obj["consumerRejectLists"] = consumerRejectListsProp
	}
	consumerAcceptListsProp, err := expandComputeServiceAttachmentConsumerAcceptLists(d.Get("consumer_accept_lists"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("consumer_accept_lists"); ok || !reflect.DeepEqual(v, consumerAcceptListsProp) {
		obj["consumerAcceptLists"] = consumerAcceptListsProp
	}
	regionProp, err := expandComputeServiceAttachmentRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeServiceAttachmentName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConnectionPreference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentTargetService(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("forwardingRules", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_service: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeServiceAttachmentNatSubnets(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: nil")
		}
		f, err := parseRegionalFieldValue("subnetworks", raw.(string), "project", "region", "zone", d, config, true)
		if err != nil {
			return nil, fmt.Errorf("Invalid value for nat_subnets: %s", err)
		}
		req = append(req, f.RelativeLink())
	}
	return req, nil
}

func expandComputeServiceAttachmentEnableProxyProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerRejectLists(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptLists(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedProjectIdOrNum, err := expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(original["project_id_or_num"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedProjectIdOrNum); val.IsValid() && !isEmptyValue(val) {
			transformed["projectIdOrNum"] = transformedProjectIdOrNum
		}

		transformedConnectionLimit, err := expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(original["connection_limit"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnectionLimit); val.IsValid() && !isEmptyValue(val) {
			transformed["connectionLimit"] = transformedConnectionLimit
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsProjectIdOrNum(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentConsumerAcceptListsConnectionLimit(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeServiceAttachmentRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}