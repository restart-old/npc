package npc

import (
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

// NPCHandler ...
type NPCHandler struct {
	player.NopHandler
	*NPC
}

// HandleHurt will cancel the hurt action and execute the action of the NPC if there is any
func (s NPCHandler) HandleHurt(ctx *event.Context, dmg *float64, src damage.Source) {
	ctx.Cancel()
	if e, ok := src.(damage.SourceEntityAttack); ok && s.action != nil {
		if p, ok := e.Attacker.(*player.Player); ok {
			s.action(p)
		}
	}
}
