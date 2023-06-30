package main

import (
	"context"
	"fmt"
	"log"
)

type State string
type Event string

type Node struct {
	State
	Transitions map[Event]*Transition
}

type Transition struct {
	*Node
	Action
}

type Action func(ctx context.Context) error

type StateMachine struct {
	initialNode *Node
	CurrentNode *Node
}

func (m *StateMachine) getCurrentNode() *Node {
	return m.CurrentNode
}

func (m *StateMachine) getNextNode(event Event) (*Node, error) {
	if m.CurrentNode == nil {
		return nil, fmt.Errorf("nowhere to go anymore!\n")
	}

	transition, ok := m.CurrentNode.Transitions[event]
	if !ok {
		return nil, fmt.Errorf("invalid event: %v", event)
	}

	return transition.Node, nil
}

func (m *StateMachine) Transition(ctx context.Context, event Event) (*Node, error) {
	node, err := m.getNextNode(event)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = m.CurrentNode.Transitions[event].Action(ctx)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	m.CurrentNode = node

	return m.CurrentNode, nil
}

func NewStateMachine(initialNode *Node) *StateMachine {
	if initialNode == nil {
		return &StateMachine{}
	}

	return &StateMachine{
		initialNode: initialNode,
		CurrentNode: initialNode,
	}
}

func main() {
	var (
		cartPageNode,
		checkoutProcessingNode,
		paymentProcessingNode,
		doneNode,
		failedNode Node
	)

	cartPageNode = Node{
		State: "cartPage",
		Transitions: map[Event]*Transition{
			"checkout_requested": {
				Node: &checkoutProcessingNode,
				Action: func(ctx context.Context) error {
					fmt.Println("cart page -> checkout_requested -> checkout processing")
					return nil
				},
			},
		},
	}

	checkoutProcessingNode = Node{
		State: "checkoutProcessing",
		Transitions: map[Event]*Transition{
			"payment_requested": {
				Node: &paymentProcessingNode,
				Action: func(ctx context.Context) error {
					fmt.Println("checkout processing -> payment_requested -> payment processing")
					return nil
				},
			},
		},
	}

	paymentProcessingNode = Node{
		State: "paymentProcessing",
		Transitions: map[Event]*Transition{
			"success": {
				Node: &doneNode,
				Action: func(ctx context.Context) error {
					fmt.Println("payment processing -> success -> done")
					return nil
				},
			},
			"timed_out": {
				Node: &cartPageNode,
				Action: func(ctx context.Context) error {
					fmt.Println("payment processing -> timed_out -> cart page")
					return nil
				},
			},
			"failed": {
				Node: &failedNode,
				Action: func(ctx context.Context) error {
					fmt.Println("payment processing -> failed -> failed")
					return nil
				},
			},
		},
	}

	machine := NewStateMachine(&cartPageNode)

	fmt.Printf("0. initial: %#v\n\n", machine.getCurrentNode())

	nextNode, _ := machine.Transition(context.TODO(), "checkout_requested")
	fmt.Printf("1. next state for event checkout requested: %#v\n\n", nextNode)

	nextNode, err := machine.Transition(context.TODO(), "gibberish")
	fmt.Printf("2. next state for event gibberish: %#v, error: %#v\n\n", nextNode, err)

	nextNode, _ = machine.Transition(context.TODO(), "payment_requested")
	fmt.Printf("3. next state for event payment requested: %#v\n\n", nextNode)

	nextNode, _ = machine.Transition(context.TODO(), "timed_out")
	fmt.Printf("4. next state for event timed out: %#v\n\n", nextNode)

	nextNode, err = machine.Transition(context.TODO(), "success")
	fmt.Printf("5. next state for event success: %#v, error: %#v\n\n", nextNode, err)

	_, err = machine.Transition(context.TODO(), "checkout_requested")
	if err != nil {
		log.Fatal(err)
	}
	_, err = machine.Transition(context.TODO(), "payment_requested")
	if err != nil {
		log.Fatal(err)
	}
	nextNode, _ = machine.Transition(context.TODO(), "success")
	fmt.Printf("6. next state for event new success: %#v\n\n", nextNode)
}
