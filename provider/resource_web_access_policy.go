package provider

import (
	"bitbucket.org/accezz-io/terraform-provider-symcsc/service"
	"bitbucket.org/accezz-io/terraform-provider-symcsc/service/dto"
	"bitbucket.org/accezz-io/terraform-provider-symcsc/utils"
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

func LuminateWebAccessPolicy() *schema.Resource {
	webSchema := LuminateAccessPolicyBaseSchema()

	conditionsResource := webSchema["conditions"].Elem.(*schema.Resource)

	conditionsResource.Schema["managed_device"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "Indicate whatever to restrict access to managed devices only",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"opswat": {
					Type:         schema.TypeBool,
					Optional:     true,
					Default:      false,
					Description:  "Indicate whatever to restrict access to Opswat MetaAccess",
					ValidateFunc: utils.ValidateBool,
				},
				"symantec_cloudsoc": {
					Type:         schema.TypeBool,
					Optional:     true,
					Default:      false,
					Description:  "Indicate whatever to restrict access to symantec cloudsoc",
					ValidateFunc: utils.ValidateBool,
				},
				"symantec_web_security_service": {
					Type:         schema.TypeBool,
					Optional:     true,
					Default:      false,
					Description:  "Indicate whatever to restrict access to symantec web security service",
					ValidateFunc: utils.ValidateBool,
				},
			},
		},
	}

	conditionsResource.Schema["unmanaged_device"] = &schema.Schema{
		Type:         schema.TypeBool,
		Optional:     true,
		Default:      false,
		Description:  "Indicate whatever to restrict access to unmanaged devices only",
		ValidateFunc: utils.ValidateBool,
	}

	return &schema.Resource{
		Schema: webSchema,
		Create: resourceCreateWebAccessPolicy,
		Read:   resourceReadWebAccessPolicy,
		Update: resourceUpdateWebAccessPolicy,
		Delete: resourceDeleteAccessPolicy,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCreateWebAccessPolicy(d *schema.ResourceData, m interface{}) error {
	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	accessPolicy := extractWebAccessPolicy(d)

	createdAccessPolicy, err := client.AccessPolicies.CreateAccessPolicy(accessPolicy)
	if err != nil {
		return err
	}

	setAccessPolicyBaseFields(d, createdAccessPolicy)

	return resourceReadWebAccessPolicy(d, m)
}

func resourceReadWebAccessPolicy(d *schema.ResourceData, m interface{}) error {
	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	accessPolicy, err := client.AccessPolicies.GetAccessPolicy(d.Id())
	if err != nil {
		return err
	}

	if accessPolicy == nil {
		d.SetId("")
		return nil
	}

	setAccessPolicyBaseFields(d, accessPolicy)

	return nil
}

func resourceUpdateWebAccessPolicy(d *schema.ResourceData, m interface{}) error {
	client, ok := m.(*service.LuminateService)
	if !ok {
		return errors.New("unable to cast Luminate service")
	}

	accessPolicy := extractWebAccessPolicy(d)
	accessPolicy.Id = d.Id()

	accessPolicy, err := client.AccessPolicies.UpdateAccessPolicy(accessPolicy)
	if err != nil {
		return err
	}

	setAccessPolicyBaseFields(d, accessPolicy)

	return resourceReadWebAccessPolicy(d, m)
}

func extractWebAccessPolicy(d *schema.ResourceData) *dto.AccessPolicy {
	accessPolicy := extractAccessPolicyBaseFields(d)
	accessPolicy.TargetProtocol = "HTTP"

	return accessPolicy
}