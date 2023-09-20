package sdkv2provider

import (
    "fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareKeylessCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: consts.ZoneIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"bundle_method": {
        	Type:         schema.TypeString,
        	Optional:     true,
        	ForceNew:    true,
        	Default:      "ubiquitous",
        	Description:  fmt.Sprintf("A ubiquitous bundle has the highest probability of being verified everywhere, even by clients using outdated or unusual trust stores. An optimal bundle uses the shortest chain and newest intermediates. And the force bundle verifies the chain, but does not otherwise modify it. %s", renderAvailableDocumentationValuesStringSlice([]string{"ubiquitous", "optimal", "force"})),
        },
		"certificate": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The zone's SSL certificate or SSL certificate and intermediate(s).",
		},
		"host": {
			Type:        schema.TypeString,
			Required:     true,
			Description: "The keyless SSL host",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The keyless SSL name",
		},
		"port": {
			Type:         schema.TypeInt,
			Optional:     true,
			ValidateFunc: validation.IntBetween(0, 65535),
			Default:      24008,
			Description:  "The keyless SSL port used to communicate between Cloudflare and the client's Keyless SSL server.",
		},
		"enabled": {
		    Type:        schema.TypeBool,
			Optional:     true,
			Description: "Whether or not the Keyless SSL is on or off.",
		},
	}
}