package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphone int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}
func (t TapePlayer) playList(songs []string) {
	for _, song := range songs {
		t.Play(song)
	}
	t.Stop()
}
func (t *TapeRecorder) playList(songs []string) {
	for _, song := range songs {
		t.Play(song)
	}
	t.Stop()
}

type MyInterface interface {
	playList(songs []string)
	Play(song string)
	Stop()
}
type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}

}
func main() {
	//player := TapePlayer{}
	mixtape := []string{"Jessie a Girl", "Whip It", "9 to 5"}
	var value MyInterface
	value = TapePlayer{}
	//value.Play("yyyyyy")
	value.playList(mixtape)
	value = &TapeRecorder{}
	value.playList(mixtape)
	TryOut(TapePlayer{})

}
