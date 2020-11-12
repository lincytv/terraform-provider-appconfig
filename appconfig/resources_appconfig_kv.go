package appconfig

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/lincytv/terraform-provider-appconfig/appconfig"
	"log"
	"os/exec"
	"strconv"
)

func resourceAppconfigKv() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppconfigKvCreate,
		Read: resourceAppconfigKvRead,
		Delete: resourceAppconfigKvDelete,
		Update: resourceAppconfigKvUpdate,
        Schema: map[string]*schema.Schema{
			"id" : &schema.Schema{
			 Type: schema.TypeString,
			 Computed: true,
			},
			"key" : &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"value" : &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"label" : &schema.Schema{
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceAppconfigKvCreate(d schema.ResourceData, meta interface{}) error {
	client := meta.(*appconfig.Client)
	    check := &appconfig.set{
	    	key: d.Get("key").(string),
	    	value: d.Get("value").(string),
	    	label: d.Get("label").(string),
		}
		resp, err := client.Create(check)
		if err != nil {
			return err
		}
		d.SetId(fmt.Sprintf(%d", resp.ID))
	return resourceAppconfigKvRead(d, meta)
}


func resourceAppconfigKvRead(d schema.ResourceData, meta interface{}) error {
	client := meta.(*appconfig.Client)
	checkId, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Error in key", err)
	}
	check, err := clientChecksRead(checkId)
	if err != nil {
		return fmt.Errorf("Error Read", d.Id(), err)
	}
      d.Set("key", check.key)
	  d.Set("value", check.value)
	  d.Set("label", check.label)
	return nil
}
func resourceAppconfigKvDelete(d schema.ResourceData, meta interface{}) error {
	client := meta.(*appconfig.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Error, err")
	}

	log.Printf("[INFO] Deleteing", id)
	_, err = client.Checks.Delete(id)
	if err != nil {
		return fmt.Errorf("Error", err)
	}
	return nil
}

//resourceAppconfigKvShow
func resourceAppconfigKvUpdate(d *schema.ResourceData) (interface{}, error)   {

	excuter := "az appconfig kv "
	key_switch := "--key "
	value_switch := "--value "
	label_switch := "--label "
    method := " list "
	cmd = exec.Command( excuter, method, key_switch, d.key, value_switch, d.value, label_switch, d.label)

    return nil, nil
}
