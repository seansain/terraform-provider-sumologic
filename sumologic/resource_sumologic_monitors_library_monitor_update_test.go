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
	"fmt"
	"testing"
  "strconv"
  "strings"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)
func TestAccSumologicMonitorsLibraryMonitorUpdate_basic(t *testing.T) {
	var monitorsLibraryMonitorUpdate MonitorsLibraryMonitorUpdate
	testDescription := "J7?~E.V]{%"
  testMonitorType := "30`r9m'A[N"
  testName := "-w^ud,\"tbc"
  testNotifications := []string{"6l6;@J^=Z="}
  testTriggers := []string{"Rx0pkSAD<@"}
  testQueries := []string{"75LMy(C1Qp"}
  testVersion := 0
  testType := "X):N:YmM19"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorsLibraryMonitorUpdateDestroy(monitorsLibraryMonitorUpdate),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicMonitorsLibraryMonitorUpdateConfigImported(testDescription, testMonitorType, testName, testNotifications, testTriggers, testQueries, testVersion, testType),
			},
			{
				ResourceName:      "sumologic_monitors_library_monitor_update.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccMonitorsLibraryMonitorUpdate_create(t *testing.T) {
  var monitorsLibraryMonitorUpdate MonitorsLibraryMonitorUpdate
  testDescription := "d#qqquvb4y"
  testMonitorType := "3MPnGdaV9a"
  testName := "*lXVo-q^sJ"
  testNotifications := []string{"SZd84Qnh\6"}
  testTriggers := []string{"9@??p<?+ij"}
  testQueries := []string{"6,=%m:HIGH"}
  testVersion := 0
  testType := "x96)<DBtR$"
  resource.Test(t, resource.TestCase{
    PreCheck: func() { testAccPreCheck(t) },
    Providers:    testAccProviders,
    CheckDestroy: testAccCheckMonitorsLibraryMonitorUpdateDestroy(monitorsLibraryMonitorUpdate),
    Steps: []resource.TestStep{
      {
        Config: testAccSumologicMonitorsLibraryMonitorUpdate(testDescription, testMonitorType, testName, testNotifications, testTriggers, testQueries, testVersion, testType),
        Check: resource.ComposeTestCheckFunc(
          testAccCheckMonitorsLibraryMonitorUpdateExists("sumologic_monitors_library_monitor_update.test", &monitorsLibraryMonitorUpdate, t),
          testAccCheckMonitorsLibraryMonitorUpdateAttributes("sumologic_monitors_library_monitor_update.test"),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "description", testDescription),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "monitor_type", testMonitorType),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "name", testName),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "notifications.0", strings.Replace(testNotifications[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "triggers.0", strings.Replace(testTriggers[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "queries.0", strings.Replace(testQueries[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "version", strconv.Itoa(testVersion)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "type", testType),
        ),
      },
    },
  })
}

func testAccCheckMonitorsLibraryMonitorUpdateDestroy(monitorsLibraryMonitorUpdate MonitorsLibraryMonitorUpdate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
    for _, r := range s.RootModule().Resources {
      id := r.Primary.ID
		  u, err := client.GetMonitorsLibraryMonitorUpdate(id)
		  if err != nil {
        return fmt.Errorf("Encountered an error: " + err.Error())
		  }
      if u != nil {
        return fmt.Errorf("MonitorsLibraryMonitorUpdate still exists")
      }
    }
		return nil
	}
}
func testAccCheckMonitorsLibraryMonitorUpdateExists(name string, monitorsLibraryMonitorUpdate *MonitorsLibraryMonitorUpdate, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
      //need this so that we don't get an unused import error for strconv in some cases
			return fmt.Errorf("Error = %s. MonitorsLibraryMonitorUpdate not found: %s", strconv.FormatBool(ok), name)
		}

    //need this so that we don't get an unused import error for strings in some cases
		if strings.EqualFold(rs.Primary.ID, "") {
			return fmt.Errorf("MonitorsLibraryMonitorUpdate ID is not set")
		}

		id := rs.Primary.ID
		c := testAccProvider.Meta().(*Client)
		newMonitorsLibraryMonitorUpdate, err := c.GetMonitorsLibraryMonitorUpdate(id)
		if err != nil {
			return fmt.Errorf("MonitorsLibraryMonitorUpdate %s not found", id)
		}
		monitorsLibraryMonitorUpdate = newMonitorsLibraryMonitorUpdate
		return nil
	}
}

func TestAccMonitorsLibraryMonitorUpdate_update(t *testing.T) {
  var monitorsLibraryMonitorUpdate MonitorsLibraryMonitorUpdate
  testDescription := "kw2F}0,Puz"
  testMonitorType := "Wt(5&d&GIY"
  testName := "=t(F(rbNQ&"
  testNotifications := []string{"vf9V_3Cpf("}
  testTriggers := []string{"_atep/^6-m"}
  testQueries := []string{"BJqW\"RVpCC"}
  testVersion := 0
  testType := "2i\"V'!Y[7/"

  testUpdatedDescription := "gZSE,k$&$!Update"
  testUpdatedMonitorType := ".Q)VIJTtlsUpdate"
  testUpdatedName := "U+.7Ju^9'wUpdate"
  testUpdatedNotifications := []string{"6@NMm12;b/"}
  testUpdatedTriggers := []string{"+^>n3iI:i2"}
  testUpdatedQueries := []string{"*fR2jAU67H"}
  testUpdatedVersion := 1
  testUpdatedType := "pVoRL%2uUZUpdate"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorsLibraryMonitorUpdateDestroy(monitorsLibraryMonitorUpdate),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicMonitorsLibraryMonitorUpdate(testDescription, testMonitorType, testName, testNotifications, testTriggers, testQueries, testVersion, testType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorsLibraryMonitorUpdateExists("sumologic_monitors_library_monitor_update.test", &monitorsLibraryMonitorUpdate, t),
					testAccCheckMonitorsLibraryMonitorUpdateAttributes("sumologic_monitors_library_monitor_update.test"),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "description", testDescription),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "monitor_type", testMonitorType),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "name", testName),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "notifications.0", strings.Replace(testNotifications[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "triggers.0", strings.Replace(testTriggers[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "queries.0", strings.Replace(testQueries[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "version", strconv.Itoa(testVersion)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "type", testType),
				),
			},
			{
				Config: testAccSumologicMonitorsLibraryMonitorUpdateUpdate(testUpdatedDescription, testUpdatedMonitorType, testUpdatedName, testUpdatedNotifications, testUpdatedTriggers, testUpdatedQueries, testUpdatedVersion, testUpdatedType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "description", testUpdatedDescription),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "monitor_type", testUpdatedMonitorType),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "name", testUpdatedName),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "notifications.0", strings.Replace(testUpdatedNotifications[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "triggers.0", strings.Replace(testUpdatedTriggers[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "queries.0", strings.Replace(testUpdatedQueries[0], "\"", "", 2)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "version", strconv.Itoa(testUpdatedVersion)),
          resource.TestCheckResourceAttr("sumologic_monitors_library_monitor_update.test", "type", testUpdatedType),
				),
			},
		},
	})
}
func testAccCheckSumologicMonitorsLibraryMonitorUpdateConfigImported(description string, monitorType string, name string, notifications []string, triggers []string, queries []string, version int, type_field string) string {
	return fmt.Sprintf(`
resource "sumologic_monitors_library_monitor_update" "foo" {
      description = "%s"
      monitor_type = "%s"
      name = "%s"
      notifications = %v
      triggers = %v
      queries = %v
      version = %d
      type = "%s"
}
`, description, monitorType, name, notifications, triggers, queries, version, type_field)
}

func testAccSumologicMonitorsLibraryMonitorUpdate(description string, monitorType string, name string, notifications []string, triggers []string, queries []string, version int, type_field string) string {
	return fmt.Sprintf(`
resource "sumologic_monitors_library_monitor_update" "test" {
    description = "%s"
    monitor_type = "%s"
    name = "%s"
    notifications = %v
    triggers = %v
    queries = %v
    version = %d
    type = "%s"
}
`, description, monitorType, name, notifications, triggers, queries, version, type_field)
}

func testAccSumologicMonitorsLibraryMonitorUpdateUpdate(description string, monitorType string, name string, notifications []string, triggers []string, queries []string, version int, type_field string) string {
	return fmt.Sprintf(`
resource "sumologic_monitors_library_monitor_update" "test" {
      description = "%s"
      monitor_type = "%s"
      name = "%s"
      notifications = %v
      triggers = %v
      queries = %v
      version = %d
      type = "%s"
}
`, description, monitorType, name, notifications, triggers, queries, version, type_field)
}

func testAccCheckMonitorsLibraryMonitorUpdateAttributes(name string) resource.TestCheckFunc {
  return func(s *terraform.State) error {
      f := resource.ComposeTestCheckFunc(
        resource.TestCheckResourceAttrSet(name, "description"),
        resource.TestCheckResourceAttrSet(name, "monitor_type"),
        resource.TestCheckResourceAttrSet(name, "name"),
        resource.TestCheckResourceAttrSet(name, "notifications"),
        resource.TestCheckResourceAttrSet(name, "triggers"),
        resource.TestCheckResourceAttrSet(name, "queries"),
        resource.TestCheckResourceAttrSet(name, "version"),
        resource.TestCheckResourceAttrSet(name, "type"),
      )
      return f(s)
   }
}