package helpers

import (
	"log"

	"runtime/debug"

	"github.com/DimensionDataResearch/go-dd-cloud-compute/compute"
	"github.com/DimensionDataResearch/packer-plugins-ddcloud/artifacts"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

// ForStateBag creates a new `State` helper for the specified multistep.StateBag.
func ForStateBag(stateBag multistep.StateBag) State {
	return State{
		Data: stateBag,
	}
}

// State is the helper for working with `multistep` state data.
type State struct {
	// The state data.
	Data multistep.StateBag
}

// Get retrieves the state data with the specified key.
func (state State) Get(key string) (value interface{}) {
	return state.Data.Get(key)
}

// GetOk retrieves the state data with the specified key, if it exists.
func (state State) GetOk(key string) (value interface{}, exists bool) {
	return state.Data.GetOk(key)
}

// Set updates the state data with the specified key and value.
func (state State) Set(key string, value interface{}) {
	state.Data.Put(key, value)
}

// GetUI gets a reference to the Packer UI from the state data.
func (state State) GetUI() packer.Ui {
	value, ok := state.Data.GetOk("ui")
	if !ok || value == nil {
		log.Printf("helpers.State.GetUI: Warning - UI not available.\n%s",
			debug.Stack(),
		)

		return nil
	}

	return value.(packer.Ui)
}

// SetUI updates the reference to the Packer UI in the state data.
func (state State) SetUI(ui packer.Ui) {
	state.Data.Put("ui", ui)
}

// GetHook gets a reference to the Packer extensibility hook from the state data.
func (state State) GetHook() *packer.Hook {
	value, ok := state.Data.GetOk("hook")
	if !ok || value == nil {
		log.Printf("helpers.State.GetHook: Warning - Hook not available.\n%s",
			debug.Stack(),
		)

		return nil
	}

	return value.(*packer.Hook)
}

// SetHook updates the reference to the Packer extensibility hook in the state data.
func (state State) SetHook(hook *packer.Hook) {
	state.Data.Put("hook", hook)
}

// GetConfig gets the plugin configuration from the state data.
func (state State) GetConfig() PluginConfig {
	value, ok := state.Data.GetOk("config")
	if !ok || value == nil {
		return nil
	}

	return value.(PluginConfig)
}

// SetConfig updates the plugin configuration in the state data.
func (state State) SetConfig(config PluginConfig) {
	state.Data.Put("config", config)
}

// GetClient gets the CloudControl API client from the state data.
func (state State) GetClient() *compute.Client {
	value, ok := state.Data.GetOk("client")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.Client)
}

// SetClient updates the CloudControl API client in the state data.
func (state State) SetClient(client *compute.Client) {
	state.Data.Put("client", client)
}

// GetNetworkDomain gets the target network domain from the state data.
func (state State) GetNetworkDomain() *compute.NetworkDomain {
	value, ok := state.Data.GetOk("network_domain")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.NetworkDomain)
}

// SetNetworkDomain updates the target network domain in the state data.
func (state State) SetNetworkDomain(networkDomain *compute.NetworkDomain) {
	state.Data.Put("network_domain", networkDomain)
}

// GetVLAN gets the target VLAN from the state data.
func (state State) GetVLAN() *compute.VLAN {
	value, ok := state.Data.GetOk("vlan")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.VLAN)
}

// SetVLAN updates the target VLAN in the state data.
func (state State) SetVLAN(vlan *compute.VLAN) {
	state.Data.Put("vlan", vlan)
}

// GetServer gets the target server from the state data.
func (state State) GetServer() *compute.Server {
	value, ok := state.Data.GetOk("server")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.Server)
}

// SetServer updates the target server in the state data.
func (state State) SetServer(server *compute.Server) {
	state.Data.Put("server", server)
}

// GetNATRule gets the NAT rule from the state data.
func (state State) GetNATRule() *compute.NATRule {
	value, ok := state.Data.GetOk("nat_rule")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.NATRule)
}

// SetNATRule updates the NAT rule in the state data.
func (state State) SetNATRule(natRule *compute.NATRule) {
	state.Data.Put("nat_rule", natRule)
}

// GetFirewallRule gets the firewall rule from the state data.
func (state State) GetFirewallRule() *compute.FirewallRule {
	value, ok := state.Data.GetOk("firewall_rule")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.FirewallRule)
}

// SetFirewallRule updates the firewall rule in the state data.
func (state State) SetFirewallRule(firewallRule *compute.FirewallRule) {
	state.Data.Put("firewall_rule", firewallRule)
}

// GetSourceImage gets the source image from the state data.
func (state State) GetSourceImage() compute.Image {
	value, ok := state.Data.GetOk("source_image")
	if !ok || value == nil {
		return nil
	}

	return value.(compute.Image)
}

// SetSourceImage updates the source image in the state data.
func (state State) SetSourceImage(image compute.Image) {
	state.Data.Put("source_image", image)
}

// GetSourceImageArtifact gets the source image artifact from the state data.
func (state State) GetSourceImageArtifact() *artifacts.Image {
	value, ok := state.Data.GetOk("source_image_artifact")
	if !ok || value == nil {
		return nil
	}

	return value.(*artifacts.Image)
}

// SetSourceImageArtifact updates the source image artifact in the state data.
func (state State) SetSourceImageArtifact(imageArtifact *artifacts.Image) {
	state.Data.Put("source_image_artifact", imageArtifact)
}

// GetTargetImage gets the target image from the state data.
func (state State) GetTargetImage() *compute.CustomerImage {
	value, ok := state.Data.GetOk("target_image")
	if !ok || value == nil {
		return nil
	}

	return value.(*compute.CustomerImage)
}

// SetTargetImage updates the target image in the state data.
func (state State) SetTargetImage(image *compute.CustomerImage) {
	state.Data.Put("target_image", image)
}

// GetTargetImageArtifact gets the target image artifact from the state data.
func (state State) GetTargetImageArtifact() *artifacts.Image {
	value, ok := state.Data.GetOk("target_image_artifact")
	if !ok || value == nil {
		return nil
	}

	return value.(*artifacts.Image)
}

// SetTargetImageArtifact updates the target image artifact in the state data.
func (state State) SetTargetImageArtifact(imageArtifact *artifacts.Image) {
	state.Data.Put("target_image_artifact", imageArtifact)
}