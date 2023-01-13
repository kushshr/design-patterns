package main

type IObserver interface {
	update()
}

type TemperatureStation interface {
	add(IObserver)
	remove(IObserver)
	notify()
	getTemperature()
}

type WeatherStation struct {
	observers   []IObserver
	temperature float32
}

func (ws WeatherStation) newStation(temperature float32) WeatherStation {
	ws = WeatherStation{}
	ws.temperature = 20.5
	ws.observers = []IObserver{}
	return ws
}

func (ws WeatherStation) add(ob IObserver) {
	ws.observers = append(ws.observers, ob)
	return
}

func (ws WeatherStation) remove(ob IObserver) {
	for i, o := range ws.observers {
		if o == ob {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
		}
	}
}

func (ws WeatherStation) getTemperature() float32 {
	return ws.getTemperature()
}

func (ws WeatherStation) notify() {
	for _, ob := range ws.observers {
		ob.update()
	}
}

type PhoneObserver struct {
	id string
}

func (po PhoneObserver) update() {
	print("Updated phone observer")
}

type WallObserver struct {
	id string
}

func (wo WallObserver) update() {
	print("Updated wall observer")
}

func main() {
	ws := WeatherStation{}
	ws.observers = []IObserver{}
	ws.temperature = 20.5

	po := PhoneObserver{}

	ws.add(po)
	ws.notify()

	wo := WallObserver{}
	ws.add(wo)
	ws.notify()
}
