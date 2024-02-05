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
	SoloWinner:        common.SCCountWinner(7),
	SoloSCCount:       func(*state.State) int { return 7 },
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
		"Nur": godip.Unit{godip.Army, Mordor},
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
		// Brown Lands
		Prov("bro").Conn("rhh", godip.Land).Conn("sea", godip.Sea).Conn("sor", godip.Coast...).Conn("esg", godip.Land).Conn("eam", godip.Land).Conn("som", godip.Land).Conn("emy", godip.Land).Conn("dea", godip.Land).Flag(godip.Coast...).
		// North Downs
		Prov("nod").Conn("fon", godip.Land).Conn("bor", godip.Land).Conn("eve", godip.Land).Flag(godip.Land).
		// North Mirkwood
		Prov("nom").Conn("car", godip.Land).Conn("eam", godip.Land).Conn("esg", godip.Land).Flag(godip.Land).
		// Khand
		Prov("kha").Conn("nea", godip.Land).Flag(godip.Land).SC(Haradwaith).
		// Lune Valley
		Prov("lun").Conn("eve", godip.Land).Conn("bor", godip.Coast...).Conn("sob", godip.Sea).Conn("mit", godip.Coast...).Conn("tow", godip.Land).Flag(godip.Coast...).
		// Osgiliath
		Prov("osg").Conn("dea", godip.Land).Conn("mti", godip.Land).Conn("leb", godip.Land).Conn("soi", godip.Land).Conn("mim", godip.Land).Flag(godip.Land).
		// Forlindon
		Prov("fol").Conn("mit", godip.Coast...).Conn("sob", godip.Sea).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// Rhun Hills
		Prov("rhh").Conn("bro", godip.Land).Flag(godip.Land).
		// Eregion
		Prov("ere").Conn("sod", godip.Land).Conn("dun", godip.Land).Conn("mor", godip.Land).Conn("rhn", godip.Land).Flag(godip.Land).
		// Umbar
		Prov("umb").Conn("sob", godip.Sea).Conn("dee", godip.Coast...).Conn("haa", godip.Land).Conn("hao", godip.Coast...).Flag(godip.Coast...).SC(Haradwaith).
		// Mount Doom
		Prov("mod").Conn("bar", godip.Land).Conn("udu", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).
		// Angmar
		Prov("ang").Conn("foo", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).
		// Nurnen
		Prov("nur").Conn("gor", godip.Sea).Conn("Nur", godip.Sea).Flag(godip.Sea).
		// Nurn
		Prov("Nur").Conn("nur", godip.Sea).Conn("gor", godip.Coast...).Flag(godip.Coast...).
		// Deep Harad
		Prov("dee").Conn("haa", godip.Land).Conn("umb", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// Borderlands
		Prov("bor").Conn("wea", godip.Land).Conn("rhn", godip.Land).Conn("ang", godip.Land).Conn("foo", godip.Coast...).Conn("sob", godip.Sea).Conn("lun", godip.Coast...).Conn("eve", godip.Land).Conn("nod", godip.Land).Conn("fon", godip.Land).Flag(godip.Coast...).
		// East Mirkwood
		Prov("eam").Conn("som", godip.Land).Conn("bro", godip.Land).Conn("esg", godip.Land).Conn("nom", godip.Land).Conn("anv", godip.Land).Flag(godip.Land).
		// Dead Marshes
		Prov("dea").Conn("udu", godip.Land).Conn("bro", godip.Land).Conn("emy", godip.Land).Conn("thw", godip.Land).Conn("gap", godip.Land).Conn("osg", godip.Land).Flag(godip.Land).
		// Gorgoroth
		Prov("gor").Conn("bar", godip.Land).Conn("mod", godip.Land).Conn("mim", godip.Land).Conn("Nur", godip.Coast...).Conn("nur", godip.Sea).Flag(godip.Coast...).
		// Fornost
		Prov("fon").Conn("tow", godip.Land).Conn("buc", godip.Land).Conn("old", godip.Land).Conn("wea", godip.Land).Conn("bor", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).
		// South Rhun
		Prov("sor").Conn("nrh", godip.Land).Conn("esg", godip.Land).Conn("bro", godip.Coast...).Conn("sea", godip.Sea).Flag(godip.Coast...).
		// Moria
		Prov("mor").Conn("ere", godip.Land).Conn("gla", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Forodwaith
		Prov("foo").Conn("sob", godip.Sea).Conn("bor", godip.Coast...).Conn("ang", godip.Land).Conn("eaa", godip.Land).Conn("unt", godip.Land).Conn("nrh", godip.Land).Flag(godip.Coast...).
		// Anfalas
		Prov("anf").Conn("dru", godip.Coast...).Conn("anr", godip.Coast...).Conn("sob", godip.Sea).Conn("val", godip.Coast...).Flag(godip.Coast...).
		// Buckland
		Prov("buc").Conn("sod", godip.Land).Conn("old", godip.Land).Conn("fon", godip.Land).Conn("tow", godip.Land).Conn("ths", godip.Land).Flag(godip.Land).
		// Harad
		Prov("haa").Conn("nea", godip.Land).Conn("hao", godip.Land).Conn("umb", godip.Land).Conn("dee", godip.Land).Flag(godip.Land).
		// Tower Hills
		Prov("tow").Conn("fon", godip.Land).Conn("eve", godip.Land).Conn("lun", godip.Land).Conn("mit", godip.Land).Conn("ths", godip.Land).Conn("buc", godip.Land).Flag(godip.Land).
		// Emyn Muil
		Prov("emy").Conn("som", godip.Land).Conn("anv", godip.Land).Conn("lor", godip.Land).Conn("thw", godip.Land).Conn("dea", godip.Land).Conn("bro", godip.Land).Flag(godip.Land).
		// Mithlond
		Prov("mit").Conn("tow", godip.Land).Conn("lun", godip.Coast...).Conn("fol", godip.Coast...).Conn("sob", godip.Sea).Conn("mrb", godip.Coast...).Flag(godip.Coast...).
		// Near Harad
		Prov("nea").Conn("kha", godip.Land).Conn("haa", godip.Land).Flag(godip.Land).
		// Eastern Angmar
		Prov("eaa").Conn("car", godip.Land).Conn("foo", godip.Land).Flag(godip.Land).
		// Mountain Range b
		Prov("mrb").Conn("mih", godip.Coast...).Conn("mit", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// Gladden Fields
		Prov("gla").Conn("anv", godip.Land).Conn("mor", godip.Land).Conn("lor", godip.Land).Flag(godip.Land).
		// North Rhun
		Prov("nrh").Conn("foo", godip.Land).Conn("unt", godip.Land).Conn("sor", godip.Land).Flag(godip.Land).
		// Enedwaith
		Prov("ene").Conn("gap", godip.Land).Conn("fan", godip.Land).Conn("dun", godip.Land).Conn("mih", godip.Coast...).Conn("sob", godip.Sea).Conn("wes", godip.Coast...).Flag(godip.Coast...).
		// Harondor
		Prov("hao").Conn("soi", godip.Coast...).Conn("sob", godip.Sea).Conn("umb", godip.Coast...).Conn("haa", godip.Land).Flag(godip.Coast...).
		// Lorien
		Prov("lor").Conn("anv", godip.Land).Conn("gla", godip.Land).Conn("fan", godip.Land).Conn("thw", godip.Land).Conn("emy", godip.Land).Flag(godip.Land).
		// Evendim Hills
		Prov("eve").Conn("lun", godip.Land).Conn("tow", godip.Land).Conn("nod", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Minas Tirith
		Prov("mti").Conn("osg", godip.Land).Conn("gap", godip.Land).Conn("leb", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Weather Hills
		Prov("wea").Conn("rhn", godip.Land).Conn("bor", godip.Land).Conn("fon", godip.Land).Conn("old", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).
		// Vale of Erech
		Prov("val").Conn("lam", godip.Coast...).Conn("anf", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// South Mirkwood
		Prov("som").Conn("emy", godip.Land).Conn("bro", godip.Land).Conn("eam", godip.Land).Conn("anv", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Druwaith Iaur
		Prov("dru").Conn("wes", godip.Coast...).Conn("sob", godip.Sea).Conn("anf", godip.Coast...).Flag(godip.Coast...).
		// Old Forest
		Prov("old").Conn("wea", godip.Land).Conn("fon", godip.Land).Conn("buc", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).
		// Fangorn
		Prov("fan").Conn("gap", godip.Land).Conn("thw", godip.Land).Conn("lor", godip.Land).Conn("ene", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// West Rohan
		Prov("wes").Conn("dru", godip.Coast...).Conn("gap", godip.Land).Conn("ene", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Minhiriath
		Prov("mih").Conn("mrb", godip.Coast...).Conn("sob", godip.Sea).Conn("ene", godip.Coast...).Conn("dun", godip.Land).Conn("sod", godip.Land).Conn("ths", godip.Land).Flag(godip.Coast...).
		// Sea of Rhun
		Prov("sea").Conn("sor", godip.Sea).Conn("bro", godip.Sea).Flag(godip.Sea).
		// Rhundaur
		Prov("rhn").Conn("wea", godip.Land).Conn("sod", godip.Land).Conn("ere", godip.Land).Conn("car", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Andrast
		Prov("anr").Conn("sob", godip.Sea).Conn("anf", godip.Coast...).Flag(godip.Coast...).
		// South Ithilien
		Prov("soi").Conn("osg", godip.Land).Conn("sob", godip.Sea).Conn("hao", godip.Coast...).Flag(godip.Coast...).
		// The Shire
		Prov("ths").Conn("mih", godip.Land).Conn("sod", godip.Land).Conn("buc", godip.Land).Conn("tow", godip.Land).Flag(godip.Land).
		// Dunland
		Prov("dun").Conn("mih", godip.Land).Conn("ene", godip.Land).Conn("ere", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).
		// Carrock
		Prov("car").Conn("anv", godip.Land).Conn("nom", godip.Land).Conn("eaa", godip.Land).Conn("rhn", godip.Land).Flag(godip.Land).
		// Gap of Rohan
		Prov("gap").Conn("ene", godip.Land).Conn("wes", godip.Land).Conn("mti", godip.Land).Conn("dea", godip.Land).Conn("thw", godip.Land).Conn("fan", godip.Land).Flag(godip.Land).
		// The Wold
		Prov("thw").Conn("dea", godip.Land).Conn("emy", godip.Land).Conn("lor", godip.Land).Conn("fan", godip.Land).Conn("gap", godip.Land).Flag(godip.Land).
		// Belfalas
		Prov("bel").Conn("lam", godip.Coast...).Conn("sob", godip.Sea).Conn("leb", godip.Coast...).Flag(godip.Coast...).
		// Anduin Valley
		Prov("anv").Conn("eam", godip.Land).Conn("car", godip.Land).Conn("gla", godip.Land).Conn("lor", godip.Land).Conn("emy", godip.Land).Conn("som", godip.Land).Flag(godip.Land).
		// Esgaroth
		Prov("esg").Conn("unt", godip.Land).Conn("nom", godip.Land).Conn("eam", godip.Land).Conn("bro", godip.Land).Conn("sor", godip.Land).Flag(godip.Land).
		// South Downs
		Prov("sod").Conn("mih", godip.Land).Conn("dun", godip.Land).Conn("ere", godip.Land).Conn("rhn", godip.Land).Conn("wea", godip.Land).Conn("old", godip.Land).Conn("buc", godip.Land).Conn("ths", godip.Land).Flag(godip.Land).
		// South Belegaer
		Prov("sob").Conn("dee", godip.Sea).Conn("umb", godip.Sea).Conn("hao", godip.Sea).Conn("soi", godip.Sea).Conn("leb", godip.Sea).Conn("bel", godip.Sea).Conn("lam", godip.Sea).Conn("val", godip.Sea).Conn("anf", godip.Sea).Conn("anr", godip.Sea).Conn("dru", godip.Sea).Conn("wes", godip.Sea).Conn("ene", godip.Sea).Conn("mih", godip.Sea).Conn("mrb", godip.Sea).Conn("mit", godip.Sea).Conn("fol", godip.Sea).Conn("fol", godip.Sea).Conn("lun", godip.Sea).Conn("bor", godip.Sea).Conn("foo", godip.Sea).Flag(godip.Sea).
		// Unthered Heath
		Prov("unt").Conn("esg", godip.Land).Conn("nrh", godip.Land).Conn("foo", godip.Land).Flag(godip.Land).
		// Lamedon
		Prov("lam").Conn("bel", godip.Coast...).Conn("leb", godip.Coast...).Conn("val", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// Barad-Dur
		Prov("bar").Conn("mod", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Minas Morgul
		Prov("mim").Conn("osg", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lebennin
		Prov("leb").Conn("osg", godip.Land).Conn("mti", godip.Land).Conn("lam", godip.Coast...).Conn("bel", godip.Coast...).Conn("sob", godip.Sea).Flag(godip.Coast...).
		// Udun Vale
		Prov("udu").Conn("mod", godip.Land).Conn("dea", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"bro": "Brown Lands",
	"nod": "North Downs",
	"nom": "North Mirkwood",
	"kha": "Khand",
	"lun": "Lune Valley",
	"osg": "Osgiliath",
	"fol": "Forlindon",
	"rhh": "Rhun Hills",
	"ere": "Eregion",
	"umb": "Umbar",
	"mod": "Mount Doom",
	"ang": "Angmar",
	"nur": "Nurnen",
	"Nur": "Nurn",
	"dee": "Deep Harad",
	"bor": "Borderlands",
	"eam": "East Mirkwood",
	"dea": "Dead Marshes",
	"gor": "Gorgoroth",
	"fon": "Fornost",
	"sor": "South Rhun",
	"mor": "Moria",
	"foo": "Forodwaith",
	"anf": "Anfalas",
	"buc": "Buckland",
	"haa": "Harad",
	"tow": "Tower Hills",
	"emy": "Emyn Muil",
	"mit": "Mithlond",
	"nea": "Near Harad",
	"eaa": "Eastern Angmar",
	"mrb": "Mountain Range b",
	"gla": "Gladden Fields",
	"nrh": "North Rhun",
	"ene": "Enedwaith",
	"hao": "Harondor",
	"lor": "Lorien",
	"eve": "Evendim Hills",
	"mti": "Minas Tirith",
	"wea": "Weather Hills",
	"val": "Vale of Erech",
	"som": "South Mirkwood",
	"dru": "Druwaith Iaur",
	"old": "Old Forest",
	"fan": "Fangorn",
	"wes": "West Rohan",
	"mih": "Minhiriath",
	"sea": "Sea of Rhun",
	"rhn": "Rhundaur",
	"anr": "Andrast",
	"soi": "South Ithilien",
	"ths": "The Shire",
	"dun": "Dunland",
	"car": "Carrock",
	"gap": "Gap of Rohan",
	"thw": "The Wold",
	"bel": "Belfalas",
	"anv": "Anduin Valley",
	"esg": "Esgaroth",
	"sod": "South Downs",
	"sob": "South Belegaer",
	"unt": "Unthered Heath",
	"lam": "Lamedon",
	"bar": "Barad-Dur",
	"mim": "Minas Morgul",
	"leb": "Lebennin",
	"udu": "Udun Vale",
}
