package game

import (
	"net"
	"sync"
	"log"
)

var laddr = &net.UDPAddr {
	Port: 11001,
	IP: net.ParseIP("127.0.0.1"),
}

type Player struct {
	ID string
	Conn *net.UDPConn
	Addr *net.UDPAddr
	X int
	Y int
}

type Game struct {
	l sync.Mutex
	Players map[string]*Player
}

func Start() *Game {
	g := new(Game)
	
	g.Players = make(map[string]*Player)
	
	return g
}

func (g *Game) AddPlayer(addr *net.UDPAddr, x, y int) {
	// conn, err := net.DialUDP("udp", laddr, addr)
	// if err != nil {
	// 	log.Println("Could not create a udp connection to the new player with addr: ", addr.String())
	// 	return	
	// }
	
	p := &Player {
		addr.String(), // ID
		nil, // This should be the udp connection
		addr,
		x,
		y,
	}
	
	g.l.Lock()
	defer g.l.Unlock()
	g.Players[p.ID] = p 
}

func (g *Game) UpdatePlayer(addr *net.UDPAddr, x, y int) {
	g.l.Lock()
	defer g.l.Unlock()
	p, ok := g.Players[addr.String()]
	
	if !ok {
		log.Println("No player with the ID: ", addr.String())
		return
	}
	
	p.X = x
	p.Y = y
}

func (g *Game) UpdateAll() {
	g.l.Lock()
	defer g.l.Unlock()
	for _, p := range g.Players {
		// _, err := p.Conn.Write([]byte("Hello from server!"))
		// 
		// if err != nil {
		// 	log.Println("Error writing to: ", p.Addr.String())
		// 	log.Println(err)
		// }
		
		log.Printf("The player with ID: '%s' is at X: %d, Y: %d", p.ID, p.X, p.Y)
	}
}

func (g *Game) IsNewPlayer(addr *net.UDPAddr) bool {
	g.l.Lock()
	defer g.l.Unlock()
	_, ok := g.Players[addr.String()]
	
	return !ok
}

