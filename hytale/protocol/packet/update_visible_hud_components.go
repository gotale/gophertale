package packet

import (
	"github.com/gotale/gophertale/hytale/protocol"
)

const (
	HudComponentHotbar = iota
	HudComponentStatusIcons
	HudComponentReticle
	HudComponentChat
	HudComponentRequests
	HudComponentNotifications
	HudComponentKillFeed
	HudComponentInputBindings
	HudComponentPlayerList
	HudComponentEventTitle
	HudComponentCompass
	HudComponentObjectivePanel
	HudComponentPortalPanel
	HudComponentBuilderToolsLegend
	HudComponentSpeedometer
	HudComponentUtilitySlotSelector
	HudComponentBlockVariantSelector
	HudComponentBuilderToolsMaterialsSlotSelector
	HudComponentStamina
	HudComponentAmmoIndicator
	HudComponentHealth
	HudComponentMana
	HudComponentOxygen
	HudComponentSleep
)

type UpdateVisibleHudComponents struct {
	VisibleComponents []uint8
}

func (*UpdateVisibleHudComponents) ID() uint32 {
	return IDUpdateVisibleHudComponents
}

func (pk *UpdateVisibleHudComponents) Marshal(io protocol.IO) {
	opts := protocol.NewOptionals(io)
	hasComponents := opts.Has(len(pk.VisibleComponents) > 0)
	opts.Write()

	if hasComponents {
		protocol.FuncSlice(io, &pk.VisibleComponents, 4096000, io.Uint8)
	}
}
