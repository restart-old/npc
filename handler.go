package npc

import (
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

// Handler is a handler for NPCs.
type Handler struct {
	player.NopHandler
	*NPC
}

// HandleHurt will cancel the hurt action and run the action of the NPC if there is any.
func (s Handler) HandleHurt(ctx *event.Context, dmg *float64, src damage.Source) {
	ctx.Cancel()
	e, ok := src.(damage.SourceEntityAttack)
	if !ok || s.action == nil {
		return
	}
	if p, ok := e.Attacker.(*player.Player); ok {
		s.action(p)
	}
}
