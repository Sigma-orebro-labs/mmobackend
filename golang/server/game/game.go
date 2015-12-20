package game

import (
	"net"
	"sync"
	"log"
)

type Player struct {
	ID string
	Addr *net.UDPAddr
	X int
	Y int
}

type Game struct {
	l sync.Mutex
	Players []*Player
}

func Start() *Game {
	return new(Game)
}

func (g *Game) AddPlayer(p *Player) {
	g.l.Lock()
	defer g.l.Unlock()
	g.Players = append(g.Players, p) 
}

func (g *Game) UpdateAll() {
	addr := net.UDPAddr {
		Port: 11001,
		IP: net.ParseIP("127.0.0.1"),
	}
	
	g.l.Lock()
	defer g.l.Unlock()
	for _, p := range g.Players {
		conn, err := net.DialUDP("udp", &addr, p.Addr)
		
		if err != nil {
			log.Println("Error updating: ", p.Addr.String())
			log.Println(err)
			continue			
		}
		
		_, err = conn.Write([]byte("Hello from server!"))
		
		if err != nil {
			log.Println("Error writing to: ", p.Addr.String())
			log.Println(err)
		}
	}
}

