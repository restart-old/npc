package main

import (
	"fmt"
	"time"

	"github.com/dragonfly-on-steroids/npc"
	"github.com/sandertv/gophertunnel/query"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
)

func practiceSpawnFunc(np *npc.NPC) {
	go func() {
		for {
			var newTag string
			q, err := query.Do("nitrofaction.fr:19132")
			if err != nil {
				fmt.Println(err)
				newTag = "§9Practice\n§cOFFLINE"
			} else {
				newTag = fmt.Sprintf("§9Practice\n§a%v/%v", q["numplayers"], q["maxplayers"])
			}
			np.SetNameTag(newTag)
			time.Sleep(3 * time.Second)
		}
	}()
}
func main() {
	c := server.DefaultConfig()
	c.Players.SaveData = false
	s := server.New(&c, nil)
	s.Start()
	s.CloseOnProgramEnd()
	for {
		if p, err := s.Accept(); err != nil {
			return
		} else {
			n := npc.New("test", "test 1", p.Skin(), p.Position()).WithAction(func(p *player.Player) {
				p.Message("You clicked on a slapper!")
			}).WithSpawnFunc(practiceSpawnFunc).WithYawAndPitch(90, 0)
			n.AddToWorld(p.World())
		}
	}
}
