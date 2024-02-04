package middleearth

import (
	"github.com/zond/godip"
	"github.com/zond/godip/graph"
	"github.com/zond/godip/state"
	"github.com/zond/godip/variants/classical"
	"github.com/zond/godip/variants/common"
)

const (
	Haradwaith godip.Nation = "Haradwaith"
	Mordor     godip.Nation = "Mordor"
)

var Nations = []godip.Nation{Haradwaith, Mordor}

var MiddleEarthVariant = common.Variant{
	Name:              "MiddleEarth",
	Graph:             func() godip.Graph { return MiddleEarthGraph() },
	Start:             MiddleEarthStart,
	Blank:             MiddleEarthBlank,
	Phase:             classical.NewPhase,
	Parser:            classical.Parser,
	Nations:           Nations,
	PhaseTypes:        classical.PhaseTypes,
	Seasons:           classical.Seasons,
	UnitTypes:         classical.UnitTypes,
	SoloWinner:        common.SCCountWinner(2),
	SoloSCCount:       func(*state.State) int { return 2 },
	ProvinceLongNames: provinceLongNames,
	SVGMap: func() ([]byte, error) {
		return Asset("svg/middleearthmap.svg")
	},
	SVGVersion: "1",
	SVGUnits: map[godip.UnitType]func() ([]byte, error){
		godip.Army: func() ([]byte, error) {
			return classical.Asset("svg/army.svg")
		},
		godip.Fleet: func() ([]byte, error) {
			return classical.Asset("svg/fleet.svg")
		},
	},
	CreatedBy:   "",
	Version:     "",
	Description: "",
	Rules:       "",
}

func MiddleEarthBlank(phase godip.Phase) *state.State {
	return state.New(MiddleEarthGraph(), phase, classical.BackupRule, nil, nil)
}

func MiddleEarthStart() (result *state.State, err error) {
	startPhase := classical.NewPhase(3018, godip.Spring, godip.Movement)
	result = MiddleEarthBlank(startPhase)
	if err = result.SetUnits(map[godip.Province]godip.Unit{
		"umb": godip.Unit{godip.Fleet, Haradwaith},
		"kha": godip.Unit{godip.Army, Haradwaith},
		"nur": godip.Unit{godip.Army, Mordor},
	}); err != nil {
		return
	}
	result.SetSupplyCenters(map[godip.Province]godip.Nation{
		"umb": Haradwaith,
		"kha": Haradwaith,
	})
	return
}

func MiddleEarthGraph() *graph.Graph {
	return graph.New().
		// Gorgoroth
		Prov("gor").Conn("SBE", godip.Sea).Conn("SBE", godip.Sea).Conn("SBE", godip.Sea).Conn("SBE", godip.Sea).Conn("nur", godip.Coast...).Conn("ims", godip.Sea).Flag(godip.Coast...).
		// Umbar
		Prov("umb").Conn("hao", godip.Coast...).Conn("SBE", godip.Sea).Conn("dee", godip.Coast...).Conn("haa", godip.Land).Flag(godip.Coast...).SC(Haradwaith).
		// Deep Harad
		Prov("dee").Conn("haa", godip.Land).Conn("umb", godip.Coast...).Conn("SBE", godip.Sea).Flag(godip.Coast...).
		// Khand
		Prov("kha").Conn("nea", godip.Land).Flag(godip.Land).SC(Haradwaith).
		// Harondor
		Prov("hao").Conn("SBE", godip.Sea).Conn("SBE", godip.Sea).Conn("umb", godip.Coast...).Conn("haa", godip.Land).Flag(godip.Coast...).
		// Harad
		Prov("haa").Conn("nea", godip.Land).Conn("hao", godip.Land).Conn("umb", godip.Land).Conn("dee", godip.Land).Flag(godip.Land).
		// Near Harad
		Prov("nea").Conn("kha", godip.Land).Conn("haa", godip.Land).Flag(godip.Land).
		// Impassable Sea
		Prov("ims").Conn("gor", godip.Sea).Conn("nur", godip.Sea).Flag(godip.Sea).
		// South Belegaer
		Prov("SBE").Conn("dee", godip.Sea).Conn("umb", godip.Sea).Conn("hao", godip.Sea).Conn("hao", godip.Sea).Conn("gor", godip.Sea).Conn("gor", godip.Sea).Conn("gor", godip.Sea).Conn("gor", godip.Sea).Flag(godip.Sea).
		// Nurn
		Prov("nur").Conn("ims", godip.Sea).Conn("gor", godip.Coast...).Flag(godip.Coast...).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"gor": "Gorgoroth",
	"umb": "Umbar",
	"dee": "Deep Harad",
	"kha": "Khand",
	"hao": "Harondor",
	"haa": "Harad",
	"nea": "Near Harad",
	"ims": "Impassable Sea",
	"SBE": "South Belegaer",
	"nur": "Nurn",
}
