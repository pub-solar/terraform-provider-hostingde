package hostingde

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRecordResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test" {
  zone_id = hostingde_zone.test.id
  name = "test.example2.test"
  type = "CNAME"
  content = "www.example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name attribute.
					resource.TestCheckResourceAttr("hostingde_record.test", "name", "test.example2.test"),
					// Verify type attribute.
					resource.TestCheckResourceAttr("hostingde_record.test", "type", "CNAME"),
					// Verify email attribute.
					resource.TestCheckResourceAttr("hostingde_zone.test", "email", "hostmaster@example2.test"),
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test", "content", "www.example.com"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("hostingde_record.test", "id"),
					resource.TestCheckResourceAttrSet("hostingde_record.test", "zone_id"),
				),
			},
			// Create and read MX testing
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test_mx" {
  zone_id = hostingde_zone.test.id
  name = "example2.test"
  type = "MX"
  content = "mail.example2.test"
  priority = 10
  comments = "Example Comment"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "name", "example2.test"),
					// Verify type attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "type", "MX"),
					// Verify priority attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "priority", "10"),
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "content", "mail.example2.test"),
					// Verify email attribute.
					resource.TestCheckResourceAttr("hostingde_zone.test", "email", "hostmaster@example2.test"),
					// Verify comments attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "comments", "Example Comment"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("hostingde_record.test_mx", "id"),
				),
			},
			// Create and read TXT testing
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test_dkim" {
  zone_id = hostingde_zone.test.id
  name = "default._domainkey.example2.test"
  type = "TXT"
  content = "v=DKIM1;k=rsa;p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyla9hW3TvoXvZQxwzaJ4SZ9ict1HU3E6+FwLWniGe6TiPtcYrjTIsiudQb8tltibOXiS+qqbxzI+quI3aGU6osy2rIv0eWo8+oOOqOD9pERftc/aqe51cXuv4kPqwvpXEBwrXFWVM+VxivEubUJ7eKkFyXJpelv0LslXv/MmYbUyed6dF+reOGZCsvnbiRv74qdxbAL/25j62E8WrnxzJwhUtx/JhdBOjsHBvuw9hy6rZsVJL9eXayWyGRV6qmsLRzsRSBs+mDrgmKk4dugADd11+A03ics3i8hplRoWDkqnNKz1qy4f5TsV6v9283IANrAzRfHwX8EvNiFsBz+ZCQIDAQAB"
  ttl = 300
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify name attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_dkim", "name", "default._domainkey.example2.test"),
					// Verify type attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_dkim", "type", "TXT"),
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_dkim", "content", "v=DKIM1;k=rsa;p=MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyla9hW3TvoXvZQxwzaJ4SZ9ict1HU3E6+FwLWniGe6TiPtcYrjTIsiudQb8tltibOXiS+qqbxzI+quI3aGU6osy2rIv0eWo8+oOOqOD9pERftc/aqe51cXuv4kPqwvpXEBwrXFWVM+VxivEubUJ7eKkFyXJpelv0LslXv/MmYbUyed6dF+reOGZCsvnbiRv74qdxbAL/25j62E8WrnxzJwhUtx/JhdBOjsHBvuw9hy6rZsVJL9eXayWyGRV6qmsLRzsRSBs+mDrgmKk4dugADd11+A03ics3i8hplRoWDkqnNKz1qy4f5TsV6v9283IANrAzRfHwX8EvNiFsBz+ZCQIDAQAB"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("hostingde_record.test_dkim", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "hostingde_zone.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test" {
  zone_id = hostingde_zone.test.id
  name = "test.example2.test"
  type = "CNAME"
  content = "www2.example.com"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test", "content", "www2.example.com"),
				),
			},
			// Update and Read testing for MX records
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test_mx" {
  zone_id = hostingde_zone.test.id
  name = "mail.example2.test"
  type = "MX"
  content = "mail2.example2.test"
  priority = 20
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "content", "mail2.example2.test"),
					// Verify priority attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_mx", "priority", "20"),
				),
			},
			// Update and Read testing for TXT records
			{
				Config: providerConfig + `
resource "hostingde_zone" "test" {
  name = "example2.test"
  type = "NATIVE"
  email = "hostmaster@example2.test"
}
resource "hostingde_record" "test_dkim" {
  zone_id = hostingde_zone.test.id
  name = "default._domainkey.example2.test"
  type = "TXT"
  content = "v=DKIM1;k=rsa;p=MiibiJanbGKQHKIg9W0baqefaaocaq8amiibcGkcaqeaYLA9Hw3tVOxVzqXWZAj4sz9ICT1hu3e6+fWlwNIgE6tIpTCyRJtiSIUDqB8TLTIBoxIs+QQBXZi+QUi3Agu6OSY2RiV0EwO8+OooQod9PerFTC/AQE51CxUV4KpQWVPxebWRxfwvm+vXIVeUBuj7EkKfYxjPELV0lSLxV/mMyBuYED6Df+REogzcSVNBIrV74QDXBal/25J62e8wRNXzJwhUtx/JhdBOjsHBvuw9hy6rZsVJL9eXayWyGRV6qmsLRzsRSBs+mDrgmKk4dugADd11+A03ics3i8hplRoWDkqnNKz1qy4f5TsV6v9283IANrAzRfHwX8EvNiFsBz+ZCQIDAQAB"
  ttl = 300
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify content attribute.
					resource.TestCheckResourceAttr("hostingde_record.test_dkim", "content", "v=DKIM1;k=rsa;p=MiibiJanbGKQHKIg9W0baqefaaocaq8amiibcGkcaqeaYLA9Hw3tVOxVzqXWZAj4sz9ICT1hu3e6+fWlwNIgE6tIpTCyRJtiSIUDqB8TLTIBoxIs+QQBXZi+QUi3Agu6OSY2RiV0EwO8+OooQod9PerFTC/AQE51CxUV4KpQWVPxebWRxfwvm+vXIVeUBuj7EkKfYxjPELV0lSLxV/mMyBuYED6Df+REogzcSVNBIrV74QDXBal/25J62e8wRNXzJwhUtx/JhdBOjsHBvuw9hy6rZsVJL9eXayWyGRV6qmsLRzsRSBs+mDrgmKk4dugADd11+A03ics3i8hplRoWDkqnNKz1qy4f5TsV6v9283IANrAzRfHwX8EvNiFsBz+ZCQIDAQAB"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
