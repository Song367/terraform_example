package subdocprovider

import (
	"context"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AliClient struct {
	AccessKey string
	SecretKey string
	Region    string
}

func AlibabaCloudManage() *schema.Resource {
	return &schema.Resource{
		CreateContext: CreateServer,
		Schema: map[string]*schema.Schema{
			"bucket": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": &schema.Schema{
				Type: schema.TypeMap,
			},
		},
	}
}

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"alicloud_oss_bucket_manage": AlibabaCloudManage(),
		},
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, r *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	region := r.Get("region").(string)
	access_key := r.Get("access_key").(string)
	secret_key := r.Get("secret_key").(string)
	var cli = AliClient{
		AccessKey: access_key,
		SecretKey: secret_key,
		Region:    region,
	}

	return cli, diags
}

func ServerProvider(ali AliClient, bucket string, tags string) error {
	c, err := sdk.NewClientWithAccessKey(ali.Region, ali.AccessKey, ali.SecretKey)
	if err != nil {
		return err
	}
	request := requests.NewCommonRequest()
	request.Scheme = "https"
	request.ApiName = "CreateBucketProvider"
	request.Version = "2015-12-15"
	request.QueryParams["Bucket"] = bucket
	request.QueryParams["tags"] = tags
	rsp, err := c.ProcessCommonRequest(request)
	if err != nil {
		return err
	}
	if code := rsp.GetHttpStatus(); code != 200 {
		return err
	}
	return nil
}

func CreateServer(c context.Context, r *schema.ResourceData, i interface{}) diag.Diagnostics {
	ali := i.(AliClient)
	bucket := r.Get("bucket").(string)
	tag := r.Get("tags").(string)
	err := ServerProvider(ali, bucket, tag)
	if err != nil {
		diag.FromErr(err)
	}
	var diagd diag.Diagnostics
	r.SetId(ali.Region + bucket)
	return diagd
}
