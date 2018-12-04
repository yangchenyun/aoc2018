package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

const (
	Begin      = 0
	Wakeup     = 1
	Fallasleep = 2
)

type Event struct {
	T         time.Time
	GuardID   int
	EventType int
}

func parseInput(filename string) []Event {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	result := make([]Event, 0)

	for _, l := range bytes.Split(dat, []byte{'\n'}) {
		if len(l) == 0 {
			continue
		}
		parts := bytes.Split(l, []byte{']', ' '})
		layout := "2006-01-02 15:04"
		t, err := time.Parse(layout, string(parts[0][1:]))
		if err != nil {
			panic(err)
		}
		evt := string(parts[1])

		var evtT int
		var id int

		if evt == "wakes up" {
			evtT = Wakeup
		} else if evt == "falls asleep" {
			evtT = Fallasleep
		} else {
			evtT = Begin
			_, err := fmt.Sscanf(evt, "Guard #%d begins shift", &id)
			if err != nil {
				panic(err)
			}
		}
		result = append(result, Event{t, id, evtT})
	}

	// Sort the events
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].T.Before(result[j].T)
	})

	// Filling the guard
	var lastGuardID int
	for i := range result {
		evt := &result[i]
		if evt.EventType == Begin {
			lastGuardID = evt.GuardID
		} else {
			if evt.GuardID != 0 {
				panic(fmt.Errorf("Expect event %v to have null guardID", evt))
			} else {
				evt.GuardID = lastGuardID
			}
		}
	}

	return result
}

type SleepSession struct {
	start int
	end   int
}
type Guard struct {
	ID   int
	evts []Event // Must be sorted
}

func (g *Guard) String() string {
	return fmt.Sprintf("<Guard #%d: %v>", g.ID, g.evts)
}

func (g *Guard) ComputeSleepSessions() []*SleepSession {
	var lastShiftDate time.Time
	var lastsleepMin int
	var session *SleepSession

	result := make([]*SleepSession, 0)

	for _, e := range g.evts {
		switch e.EventType {
		case Begin:
			{
				// closing the last sleep session
				if session != nil && session.start != 0 {
					session.end = 59
					result = append(result, session)
				}

				// reset for each shift
				lastsleepMin = -1
				lastShiftDate = e.T
				session = &SleepSession{}
			}
		case Wakeup:
			{
				if e.T == lastShiftDate {
					panic(fmt.Errorf("Wakeup before having a shift."))
				}
				if lastsleepMin == -1 {
					panic(fmt.Errorf("Wakeup before falling asleep."))
				}
				session.end = e.T.Minute()
				result = append(result, session)
				session = &SleepSession{} // reset sessions
			}

		case Fallasleep:
			{
				if e.T == lastShiftDate {
					panic(fmt.Errorf("Fallasleep before having a shift."))
				}
				lastsleepMin = e.T.Minute()
				session.start = e.T.Minute()
			}
		}
	}

	if session.start != 0 && session.end == 0 {
		panic(fmt.Errorf("unfinished sleep session"))
	}

	return result
}

func (g *Guard) TotalSleepTime() int {
	sessions := g.ComputeSleepSessions()
	result := 0
	for _, s := range sessions {
		result += s.end - s.start
	}
	return result
}

func (g *Guard) MostSleptMin() int {
	m := g.MostSleptTimes()
	dis := g.SleepDistribution()

	for i := range dis {
		if dis[i] == m {
			return i
		}
	}
	return -1
}

func (g *Guard) MostSleptTimes() int {
	dis := g.SleepDistribution()
	sort.Ints(dis)
	return dis[len(dis)-1]
}

// Returns a slice of sleep distribution over 00-59 min.
func (g *Guard) SleepDistribution() []int {
	sessions := g.ComputeSleepSessions()
	result := make([]int, 60)

	// for every session
	for _, s := range sessions {
		// for every minute
		for i := range result {
			// NOTE: that guards count as asleep on the minute they
			// fall asleep, and they count as awake on the minute
			// they wake up.
			if i >= s.start && i < s.end {
				result[i]++
			}
		}
	}

	return result
}

func main() {
	evts := parseInput("input.txt")
	guardMap := make(map[int]*Guard)
	for _, evt := range evts {
		if _, ok := guardMap[evt.GuardID]; !ok {
			guardMap[evt.GuardID] = &Guard{ID: evt.GuardID, evts: make([]Event, 0)}
		}
		guardMap[evt.GuardID].evts = append(guardMap[evt.GuardID].evts, evt)
	}

	guards := make([]*Guard, 0)
	for _, guard := range guardMap {
		guards = append(guards, guard)
	}

	// Strategy 1: Find the guard that has the most minutes asleep. What
	// minute does that guard spend asleep the most?
	sort.SliceStable(guards, func(i, j int) bool {
		return guards[i].TotalSleepTime() > guards[j].TotalSleepTime()
	})

	lazyGuard := guards[0]
	fmt.Println(lazyGuard.ID)
	fmt.Println(lazyGuard.MostSleptMin())
	fmt.Println(lazyGuard.ID * lazyGuard.MostSleptMin())
}
