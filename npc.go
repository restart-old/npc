package npc

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// NPC embeds *player.Player as private so the user doesn't break the NPC system.
// action is the function that's going to be executed when a player hits the NPC.
type NPC struct {
	p          *player.Player
	action     func(*player.Player)
	spawnFunc  func(*NPC)
	yaw, pitch float64
	world.NopViewer
}

// New returns a new *NPC with the information provided.
func New(name, displayName string, skin skin.Skin, pos mgl64.Vec3) *NPC {
	p := player.New(name, skin, pos)

	p.SetNameTag(displayName)

	npc := &NPC{p: p}
	p.Handle(Handler{NPC: npc})
	return npc
}

// SetNameTag sets the name tag of the NPC.
func (s *NPC) SetNameTag(name string) {
	s.p.SetNameTag(name)
}

// WithAction adds the action provided to the NPC and returns a pointer to the NPC.
func (s *NPC) WithAction(action func(*player.Player)) *NPC {
	s.action = action
	return s
}

// WithSpawnFunc adds the function provided to spawnFunc and will run it when the NPC is added to a world.
func (s *NPC) WithSpawnFunc(f func(*NPC)) *NPC {
	s.spawnFunc = f
	return s
}

// WithYawAndPitch sets the yaw and pitch of the NPC.
func (s *NPC) WithYawAndPitch(yaw, pitch float64) *NPC {
	s.yaw = yaw
	s.pitch = pitch
	return s
}

// AddToWorld adds the NPC to the world provided and executes the spawnFunc if any.
func (s *NPC) AddToWorld(w *world.World) {
	if s.spawnFunc != nil {
		go s.spawnFunc(s)
	}
	l := world.NewLoader(16, w, s)
	l.Move(s.p.Position())
	l.Load(1)
	w.AddEntity(s.p)
	s.p.Move(mgl64.Vec3{0, 0, 0}, s.yaw, s.pitch)
}

// RemoveFromWorld removes the NPC from the world provided.
func (s *NPC) RemoveFromWorld(w *world.World) {
	w.RemoveEntity(s.p)
}
