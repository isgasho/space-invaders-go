package states

import (
	"fmt"

	"github.com/x-hgg-x/space-invaders-go/lib/loader"

	ecs "github.com/x-hgg-x/goecs"
	"github.com/x-hgg-x/goecsengine/states"
	w "github.com/x-hgg-x/goecsengine/world"

	"github.com/hajimehoshi/ebiten"
)

// MainMenuState is the main menu state
type MainMenuState struct {
	mainMenu  []ecs.Entity
	selection int
}

//
// Menu interface
//

func (st *MainMenuState) getSelection() int {
	return st.selection
}

func (st *MainMenuState) setSelection(selection int) {
	st.selection = selection
}

func (st *MainMenuState) confirmSelection() states.Transition {
	switch st.selection {
	case 0:
		// New game
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	case 1:
		// Exit
		return states.Transition{Type: states.TransQuit}
	}
	panic(fmt.Errorf("unknown selection: %d", st.selection))
}

func (st *MainMenuState) getMenuIDs() []string {
	return []string{"new_game", "exit"}
}

func (st *MainMenuState) getCursorMenuIDs() []string {
	return []string{"cursor_new_game", "cursor_exit"}
}

//
// State interface
//

// OnPause method
func (st *MainMenuState) OnPause(world w.World) {}

// OnResume method
func (st *MainMenuState) OnResume(world w.World) {}

// OnStart method
func (st *MainMenuState) OnStart(world w.World) {
	st.mainMenu = loader.LoadEntities("assets/metadata/entities/ui/main_menu.toml", world)
}

// OnStop method
func (st *MainMenuState) OnStop(world w.World) {
	world.Manager.DeleteEntities(st.mainMenu...)
}

// Update method
func (st *MainMenuState) Update(world w.World, screen *ebiten.Image) states.Transition {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}
	return updateMenu(st, world)
}
