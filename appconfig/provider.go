package appconfig

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{ "key" : &schema.Schema{
			Type: schema.TypeString,
			Required: true,
			DefaultFunc: schema.EnvDefaultFunc("key", nil),
			Description: "Key",
		},
          "value" : &schema.Schema{
			Type: schema.TypeString,
			Required: true,
			DefaultFunc: schema.EnvDefaultFunc("value", nil),
			Description: "value",
		  },
		  "label" : &schema.Schema{
			Type: schema.TypeString,
			Required: false,
			DefaultFunc: schema.EnvDefaultFunc("label", nil),
			Description: "label",
		  }},

		ResourcesMao: map[string]*schema.Resource{},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResiurceData) (interface{}, error) {
	config:= AppconfigValues{
		key: d.Get("key").(string),
		value: d.Get("value").(string),
		label: d.Get("label").(string),

	}

	if err := config.validate(); err != nil {
		return nil, err
	}
	client, err := config.getAppconfigClient()
	if err != nil {
		return nil, err
	}
	return client, nil
}
func (v AppconfigValues) validate() error {

	var err *multierror.Error

	if c.key == "" {
		err = multierror.Append(err, fmt.Errorf("key must be there in appconfigure"))
	}
	if c.value == "" {
		err = multierror.Append(err, fmt.Errorf("value must be there in appconfigure"))
	}
	if c.label == "" {
		err = multierror.Append(err, fmt.Errorf("label must be there in appconfigure"))
	}

	return err.ErrorOrNil()
}

func (c AppconfigValues) getAppconfigClient() (*appconfig.Client, error) {

	client := appconfig.NewValue(c.key, c.value, c.label)
	return client, nil
}

type AppconfigValues struct {
	key string
	value string
	label string
}