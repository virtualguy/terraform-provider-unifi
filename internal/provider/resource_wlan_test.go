package provider

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// TODO: test changing security, see https://github.com/paultyng/terraform-provider-unifi/issues/32

// there is a max of 4 SSID's at once, and if you are running this on a
// controller with existing SSID's, you may want to limit the concurrency.
var wlanConcurrency chan struct{}

func init() {
	wcs := os.Getenv("UNIFI_ACC_WLAN_CONCURRENCY")
	if wcs == "" {
		// default concurrent SSIDs
		wcs = "1"
	}
	wc, err := strconv.Atoi(wcs)
	if err != nil {
		panic(err)
	}
	wlanConcurrency = make(chan struct{}, wc)
}

func wlanPreCheck(t *testing.T) {
	if cap(wlanConcurrency) == 0 {
		t.Skip("concurrency for WLAN testing set to 0")
	}

	wlanConcurrency <- struct{}{}
}

func TestAccWLAN_wpapsk(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_wpapsk(subnet, vlan, "disabled"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_open(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_open(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_open_mac_filter(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_open(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_change_security_and_pmf(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_wpapsk(subnet, vlan, "disabled"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_open(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_wpapsk(subnet, vlan, "optional"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_wpapsk(subnet, vlan, "required"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_wpapsk(subnet, vlan, "disabled"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_schedule(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_schedule(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			// remove schedule
			{
				Config: testAccWLANConfig_open(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_wpaeap(t *testing.T) {
	subnet, vlan := getTestVLAN(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_wpaeap(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_wlan_band(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_wlan_band(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_no2ghz_oui(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_no2ghz_oui(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_uapsd(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_uapsd(subnet, vlan),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_wpa3(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			preCheckMinVersion(t, controllerVersionWPA3)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_wpa3(subnet, vlan, false, "required"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_wpa3(subnet, vlan, true, "optional"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_wpa3(subnet, vlan, false, "required"),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func TestAccWLAN_minimum_data_rate(t *testing.T) {
	subnet, vlan := getTestVLAN(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			preCheck(t)
			wlanPreCheck(t)
		},
		ProviderFactories: providerFactories,
		CheckDestroy: func(*terraform.State) error {
			// TODO: actual CheckDestroy

			<-wlanConcurrency
			return nil
		},
		Steps: []resource.TestStep{
			{
				Config: testAccWLANConfig_minimum_data_rate(subnet, vlan, 5500, 18000),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_minimum_data_rate(subnet, vlan, 1000, 18000),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_minimum_data_rate(subnet, vlan, 0, 0),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_minimum_data_rate(subnet, vlan, 6000, 9000),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
			{
				Config: testAccWLANConfig_minimum_data_rate(subnet, vlan, 18000, 6000),
				Check:  resource.ComposeTestCheckFunc(
				// testCheckNetworkExists(t, "name"),
				),
			},
			importStep("unifi_wlan.test"),
		},
	})
}

func testAccWLANConfig_wpapsk(subnet *net.IPNet, vlan int, pmf string) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-wpapsk"
	network_id    = unifi_network.test.id
	passphrase    = "12345678"
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "wpapsk"
	
	multicast_enhance = true

	pmf_mode = %[3]q
}
`, subnet, vlan, pmf)
}

func testAccWLANConfig_wpaeap(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

data "unifi_radius_profile" "default" {}

resource "unifi_setting_radius" "this" {
	enabled = true
	secret  = "securepw"
}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-wpapsk"
	network_id    = unifi_network.test.id
	passphrase    = "12345678"
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "wpaeap"

	radius_profile_id = data.unifi_radius_profile.default.id
}
`, subnet, vlan)
}

func testAccWLANConfig_open(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-open"
	network_id    = unifi_network.test.id
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "open"
}
`, subnet, vlan)
}

func testAccWLANConfig_schedule(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-open-schedule"
	network_id    = unifi_network.test.id
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "open"

	schedule {
		day_of_week = "mon"
		start_hour  = 3
		duration    = 60*6
	}

	schedule {
		day_of_week  = "wed"
		start_hour   = 13
		start_minute = 30
		duration     = (60*3)+30
		name         = "minute"
	}

	schedule {
		day_of_week = "thu"
		start_hour  = 19
		duration    = 60*1
	}

	schedule {
		day_of_week = "fri"
		start_hour  = 19
		duration    = 60*1
	}
}
`, subnet, vlan)
}

func testAccWLANConfig_open_mac_filter(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-open"
	network_id    = unifi_network.test.id
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "open"

	mac_filter_enabled = true
	mac_filter_list    = ["ab:cd:ef:12:34:56"]
	mac_filter_policy  = "allow"
}
`, subnet, vlan)
}

func testAccWLANConfig_wlan_band(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name              = "tfacc-wpapsk"
	network_id        = unifi_network.test.id
	passphrase        = "12345678"
	ap_group_ids      = [data.unifi_ap_group.default.id]
	user_group_id     = data.unifi_user_group.default.id
	security          = "wpapsk"
	wlan_band         = "5g"
	multicast_enhance = true
}
`, subnet, vlan)
}

func testAccWLANConfig_no2ghz_oui(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name              = "tfacc-wpapsk"
	network_id        = unifi_network.test.id
	passphrase        = "12345678"
	ap_group_ids      = [data.unifi_ap_group.default.id]
	user_group_id     = data.unifi_user_group.default.id
	security          = "wpapsk"
	no2ghz_oui        = false
	multicast_enhance = true
}
`, subnet, vlan)
}

func testAccWLANConfig_uapsd(subnet *net.IPNet, vlan int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-wpapsk"
	network_id    = unifi_network.test.id
	passphrase    = "12345678"
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "wpapsk"
	uapsd         = true
}
`, subnet, vlan)
}

func testAccWLANConfig_wpa3(subnet *net.IPNet, vlan int, wpa3Transition bool, pmf string) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-wpapsk"
	network_id    = unifi_network.test.id
	passphrase    = "12345678"
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "wpapsk"

	wpa3_support    = true
	wpa3_transition = %[3]t
	pmf_mode        = %[4]q
}
`, subnet, vlan, wpa3Transition, pmf)
}

func testAccWLANConfig_minimum_data_rate(subnet *net.IPNet, vlan int, min2g int, min5g int) string {
	return fmt.Sprintf(`
data "unifi_ap_group" "default" {}

data "unifi_user_group" "default" {}

resource "unifi_network" "test" {
	name    = "tfacc"
	purpose = "corporate"
	subnet  = "%[1]s"
	vlan_id = %[2]d
}

resource "unifi_wlan" "test" {
	name          = "tfacc-wpapsk"
	network_id    = unifi_network.test.id
	passphrase    = "12345678"
	ap_group_ids  = [data.unifi_ap_group.default.id]
	user_group_id = data.unifi_user_group.default.id
	security      = "wpapsk"

	multicast_enhance = true

	minimum_data_rate_2g_kbps = %[3]d
	minimum_data_rate_5g_kbps = %[4]d
}
`, subnet, vlan, min2g, min5g)
}
