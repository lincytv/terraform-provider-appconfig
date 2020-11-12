package appconfig

import (
	"github.com/hashicorp/terraform/helper/schema"
	"os/exec"
)

func resourceAppconfigKv() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppconfigKvCreate,
		Read: resourceAppconfigKvShow,
		Delete: resourceAppconfigKvDelete,
		Update: resourceAppconfigKvUpdate,

	}
}

//resourceAppconfigKvShow
func resourceAppconfigKvShow(d *schema.ResourceData) (interface{}, error)   {

	excuter := "az appconfig kv "
	key_switch := "--key "
	value_switch := "--value "
	label_switch := "--label "
    method := " list "
	cmd = exec.Command( excuter, method, key_switch, d.key, value_switch, d.value, label_switch, d.label)

    return nil, nil
}
