package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	ID   uint64
	Name string
}

func (b BasicMember) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s | Plan: Basic", b.ID, b.Name)
}

type PremiumMember struct {
	ID   uint64
	Name string
}

func (p PremiumMember) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s | Plan: Premium (All Access)", p.ID, p.Name)
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() *Gym {
	return &Gym{Members: make(map[uint64]Member)}
}

func (g *Gym) AddMember(m Member, id uint64) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	fmt.Println("Gym Members List:")
	for _, m := range g.Members {
		fmt.Println(m.GetDetails())
	}
}
