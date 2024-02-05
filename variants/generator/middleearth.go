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
	SoloWinner:        common.SCCountWinner(10),
	SoloSCCount:       func(*state.State) int { return 10 },
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
		Prov("bro").Conn("rhh", godip.Land).Conn("sor", godip.Sea).Conn("srh", godip.Coast...).Conn("esg", godip.Land).Conn("eam", godip.Land).Conn("som", godip.Land).Conn("emy", godip.Land).Conn("dea", godip.Land).Flag(godip.Coast...).
		// Moria
		Prov("mor").Conn("ere", godip.Land).Conn("gla", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lune Valley
		Prov("lun").Conn("eve", godip.Land).Conn("bor", godip.Coast...).Conn("foc", godip.Sea).Conn("mit", godip.Land).Conn("tow", godip.Land).Flag(godip.Coast...).
		// Forchel
		Prov("foc").Conn("gul", godip.Sea).Conn("lun", godip.Sea).Conn("bor", godip.Sea).Conn("foa", godip.Sea).Flag(godip.Sea).
		// Carrock
		Prov("car").Conn("anv", godip.Land).Conn("nom", godip.Land).Conn("eaa", godip.Land).Conn("rhn", godip.Land).Flag(godip.Land).
		// Fornost
		Prov("fos").Conn("tow", godip.Land).Conn("buc", godip.Land).Conn("old", godip.Land).Conn("wea", godip.Land).Conn("bor", godip.Land).Conn("nod", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Vale of Erech
		Prov("val").Conn("lam", godip.Coast...).Conn("anf", godip.Coast...).Conn("bob", godip.Sea).Flag(godip.Coast...).
		// Harad
		Prov("haa").Conn("nea", godip.Land).Conn("hro", godip.Land).Conn("umb", godip.Land).Conn("dee", godip.Land).Flag(godip.Land).
		// Tower Hills
		Prov("tow").Conn("fos", godip.Land).Conn("eve", godip.Land).Conn("lun", godip.Land).Conn("mit", godip.Land).Conn("ths", godip.Land).Conn("buc", godip.Land).Flag(godip.Land).
		// Near Harad
		Prov("nea").Conn("kha", godip.Land).Conn("haa", godip.Land).Flag(godip.Land).
		// South Ithilien
		Prov("soi").Conn("osg", godip.Land).Conn("eth", godip.Sea).Conn("hro", godip.Coast...).Flag(godip.Coast...).
		// Gulf of Lune
		Prov("gul").Conn("nob", godip.Sea).Conn("gwa", godip.Sea).Conn("hal", godip.Sea).Conn("mit", godip.Sea).Conn("frl", godip.Sea).Conn("frl", godip.Sea).Conn("foc", godip.Sea).Flag(godip.Sea).
		// The Shire
		Prov("ths").Conn("sod", godip.Land).Conn("buc", godip.Land).Conn("tow", godip.Land).Conn("mnh", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lefnui Sea
		Prov("lef").Conn("ant", godip.Sea).Conn("sob", godip.Sea).Conn("boh", godip.Sea).Conn("bob", godip.Sea).Conn("anf", godip.Sea).Conn("Ant", godip.Sea).Flag(godip.Sea).
		// Sea of Rhun
		Prov("sor").Conn("srh", godip.Sea).Conn("bro", godip.Sea).Flag(godip.Sea).
		// Anduin Valley
		Prov("anv").Conn("eam", godip.Land).Conn("car", godip.Land).Conn("gla", godip.Land).Conn("lor", godip.Land).Conn("emy", godip.Land).Conn("som", godip.Land).Flag(godip.Land).
		// Unthered Heath
		Prov("unt").Conn("esg", godip.Land).Conn("nrh", godip.Land).Conn("foa", godip.Land).Flag(godip.Land).
		// Lamedon
		Prov("lam").Conn("val", godip.Coast...).Conn("bob", godip.Sea).Conn("bel", godip.Coast...).Conn("leb", godip.Land).Flag(godip.Coast...).
		// Minas Morgul
		Prov("mim").Conn("osg", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Lebennin
		Prov("leb").Conn("osg", godip.Land).Conn("mti", godip.Land).Conn("lam", godip.Land).Conn("bel", godip.Coast...).Conn("eth", godip.Sea).Flag(godip.Coast...).
		// Umbar
		Prov("umb").Conn("eth", godip.Sea).Conn("boh", godip.Sea).Conn("dee", godip.Coast...).Conn("haa", godip.Land).Conn("hro", godip.Coast...).Flag(godip.Coast...).SC(Haradwaith).
		// Osgiliath
		Prov("osg").Conn("dea", godip.Land).Conn("mti", godip.Land).Conn("leb", godip.Land).Conn("soi", godip.Land).Conn("mim", godip.Land).Flag(godip.Land).
		// Forlindon
		Prov("frl").Conn("mit", godip.Coast...).Conn("gul", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Coast...).
		// Borderlands
		Prov("bor").Conn("wea", godip.Land).Conn("rhn", godip.Land).Conn("anm", godip.Land).Conn("foa", godip.Coast...).Conn("foc", godip.Sea).Conn("lun", godip.Coast...).Conn("eve", godip.Land).Conn("nod", godip.Land).Conn("fos", godip.Land).Flag(godip.Coast...).
		// Sea of Nurnen
		Prov("son").Conn("gor", godip.Sea).Conn("Nur", godip.Sea).Flag(godip.Sea).
		// East Mirkwood
		Prov("eam").Conn("som", godip.Land).Conn("bro", godip.Land).Conn("esg", godip.Land).Conn("nom", godip.Land).Conn("anv", godip.Land).Flag(godip.Land).
		// Dead Marshes
		Prov("dea").Conn("udu", godip.Land).Conn("bro", godip.Land).Conn("emy", godip.Land).Conn("thw", godip.Land).Conn("gap", godip.Land).Conn("osg", godip.Land).Flag(godip.Land).
		// Anfalas
		Prov("anf").Conn("dru", godip.Land).Conn("Ant", godip.Coast...).Conn("lef", godip.Sea).Conn("bob", godip.Sea).Conn("val", godip.Coast...).Flag(godip.Coast...).
		// Forodwaith
		Prov("foa").Conn("foc", godip.Sea).Conn("bor", godip.Coast...).Conn("anm", godip.Land).Conn("eaa", godip.Land).Conn("unt", godip.Land).Conn("nrh", godip.Land).Flag(godip.Coast...).
		// Mithlond
		Prov("mit").Conn("tow", godip.Land).Conn("lun", godip.Land).Conn("frl", godip.Coast...).Conn("gul", godip.Sea).Conn("hal", godip.Coast...).Flag(godip.Coast...).
		// Evendim Hills
		Prov("eve").Conn("lun", godip.Land).Conn("tow", godip.Land).Conn("nod", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Druwaith Iaur
		Prov("dru").Conn("wes", godip.Coast...).Conn("agr", godip.Sea).Conn("ant", godip.Sea).Conn("anf", godip.Land).Flag(godip.Coast...).
		// Fangorn
		Prov("fan").Conn("gap", godip.Land).Conn("thw", godip.Land).Conn("lor", godip.Land).Conn("ene", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// North Belegaer
		Prov("nob").Conn("sob", godip.Sea).Conn("ant", godip.Sea).Conn("agr", godip.Sea).Conn("gwa", godip.Sea).Conn("gul", godip.Sea).Flag(godip.Sea).
		// Gap of Rohan
		Prov("gap").Conn("ene", godip.Land).Conn("wes", godip.Land).Conn("mti", godip.Land).Conn("dea", godip.Land).Conn("thw", godip.Land).Conn("fan", godip.Land).Flag(godip.Land).
		// Esgaroth
		Prov("esg").Conn("unt", godip.Land).Conn("nom", godip.Land).Conn("eam", godip.Land).Conn("bro", godip.Land).Conn("srh", godip.Land).Flag(godip.Land).
		// Rhun Hills
		Prov("rhh").Conn("bro", godip.Land).Flag(godip.Land).
		// Barad-Dur
		Prov("bar").Conn("mod", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Minhiriath
		Prov("mnh").Conn("hal", godip.Coast...).Conn("gwa", godip.Sea).Conn("ene", godip.Coast...).Conn("dun", godip.Land).Conn("sod", godip.Land).Conn("ths", godip.Land).Flag(godip.Coast...).SC(godip.Neutral).
		// North Downs
		Prov("nod").Conn("fos", godip.Land).Conn("bor", godip.Land).Conn("eve", godip.Land).Flag(godip.Land).
		// North Mirkwood
		Prov("nom").Conn("car", godip.Land).Conn("eam", godip.Land).Conn("esg", godip.Land).Flag(godip.Land).
		// Khand
		Prov("kha").Conn("nea", godip.Land).Flag(godip.Land).SC(Haradwaith).
		// Eregion
		Prov("ere").Conn("rhn", godip.Land).Conn("sod", godip.Land).Conn("dun", godip.Land).Conn("mor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Mount Doom
		Prov("mod").Conn("bar", godip.Land).Conn("udu", godip.Land).Conn("gor", godip.Land).Flag(godip.Land).
		// Enedwaith
		Prov("ene").Conn("gap", godip.Land).Conn("fan", godip.Land).Conn("dun", godip.Land).Conn("mnh", godip.Coast...).Conn("gwa", godip.Sea).Conn("agr", godip.Sea).Conn("wes", godip.Coast...).Flag(godip.Coast...).
		// Nurn
		Prov("Nur").Conn("son", godip.Sea).Conn("gor", godip.Coast...).Flag(godip.Coast...).
		// Harondor
		Prov("hro").Conn("soi", godip.Coast...).Conn("eth", godip.Sea).Conn("umb", godip.Coast...).Conn("haa", godip.Land).Flag(godip.Coast...).
		// Minas Tirith
		Prov("mti").Conn("osg", godip.Land).Conn("gap", godip.Land).Conn("leb", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Weather Hills
		Prov("wea").Conn("rhn", godip.Land).Conn("bor", godip.Land).Conn("fos", godip.Land).Conn("old", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// South Belegaer
		Prov("sob").Conn("boh", godip.Sea).Conn("lef", godip.Sea).Conn("ant", godip.Sea).Conn("nob", godip.Sea).Flag(godip.Sea).
		// South Mirkwood
		Prov("som").Conn("emy", godip.Land).Conn("bro", godip.Land).Conn("eam", godip.Land).Conn("anv", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Ethir Anduin
		Prov("eth").Conn("bel", godip.Sea).Conn("boh", godip.Sea).Conn("umb", godip.Sea).Conn("hro", godip.Sea).Conn("soi", godip.Sea).Conn("leb", godip.Sea).Flag(godip.Sea).
		// Harlindon
		Prov("hal").Conn("mnh", godip.Coast...).Conn("mit", godip.Coast...).Conn("gul", godip.Sea).Conn("gwa", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Gladden Fields
		Prov("gla").Conn("anv", godip.Land).Conn("mor", godip.Land).Conn("lor", godip.Land).Flag(godip.Land).
		// Andrast
		Prov("Ant").Conn("ant", godip.Sea).Conn("lef", godip.Sea).Conn("anf", godip.Coast...).Flag(godip.Coast...).
		// Gwathlo Sea
		Prov("gwa").Conn("gul", godip.Sea).Conn("nob", godip.Sea).Conn("agr", godip.Sea).Conn("ene", godip.Sea).Conn("mnh", godip.Sea).Conn("hal", godip.Sea).Flag(godip.Sea).
		// Dunland
		Prov("dun").Conn("mnh", godip.Land).Conn("ene", godip.Land).Conn("ere", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).
		// The Wold
		Prov("thw").Conn("dea", godip.Land).Conn("emy", godip.Land).Conn("lor", godip.Land).Conn("fan", godip.Land).Conn("gap", godip.Land).Flag(godip.Land).
		// Emyn Muil
		Prov("emy").Conn("som", godip.Land).Conn("anv", godip.Land).Conn("lor", godip.Land).Conn("thw", godip.Land).Conn("dea", godip.Land).Conn("bro", godip.Land).Flag(godip.Land).
		// Angren Sea
		Prov("agr").Conn("ene", godip.Sea).Conn("gwa", godip.Sea).Conn("nob", godip.Sea).Conn("ant", godip.Sea).Conn("dru", godip.Sea).Conn("wes", godip.Sea).Flag(godip.Sea).
		// Bay of Belfalas
		Prov("bob").Conn("lam", godip.Sea).Conn("val", godip.Sea).Conn("anf", godip.Sea).Conn("lef", godip.Sea).Conn("boh", godip.Sea).Conn("bel", godip.Sea).Flag(godip.Sea).
		// Bay of Harmen
		Prov("boh").Conn("dee", godip.Sea).Conn("umb", godip.Sea).Conn("eth", godip.Sea).Conn("bob", godip.Sea).Conn("lef", godip.Sea).Conn("sob", godip.Sea).Flag(godip.Sea).
		// Angmar
		Prov("anm").Conn("foa", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).
		// South Rhun
		Prov("srh").Conn("nrh", godip.Land).Conn("esg", godip.Land).Conn("bro", godip.Coast...).Conn("sor", godip.Sea).Flag(godip.Coast...).
		// Gorgoroth
		Prov("gor").Conn("bar", godip.Land).Conn("mod", godip.Land).Conn("mim", godip.Land).Conn("Nur", godip.Coast...).Conn("son", godip.Sea).Flag(godip.Coast...).
		// North Rhun
		Prov("nrh").Conn("foa", godip.Land).Conn("unt", godip.Land).Conn("srh", godip.Land).Flag(godip.Land).
		// Deep Harad
		Prov("dee").Conn("haa", godip.Land).Conn("umb", godip.Coast...).Conn("boh", godip.Sea).Flag(godip.Coast...).
		// Lorien
		Prov("lor").Conn("anv", godip.Land).Conn("gla", godip.Land).Conn("fan", godip.Land).Conn("thw", godip.Land).Conn("emy", godip.Land).Flag(godip.Land).
		// Andrast Sea
		Prov("ant").Conn("lef", godip.Sea).Conn("Ant", godip.Sea).Conn("dru", godip.Sea).Conn("agr", godip.Sea).Conn("nob", godip.Sea).Conn("sob", godip.Sea).Flag(godip.Sea).
		// Old Forest
		Prov("old").Conn("wea", godip.Land).Conn("fos", godip.Land).Conn("buc", godip.Land).Conn("sod", godip.Land).Flag(godip.Land).
		// West Rohan
		Prov("wes").Conn("dru", godip.Coast...).Conn("gap", godip.Land).Conn("ene", godip.Coast...).Conn("agr", godip.Sea).Flag(godip.Coast...).SC(godip.Neutral).
		// Buckland
		Prov("buc").Conn("sod", godip.Land).Conn("old", godip.Land).Conn("fos", godip.Land).Conn("tow", godip.Land).Conn("ths", godip.Land).Flag(godip.Land).
		// Rhundaur
		Prov("rhn").Conn("wea", godip.Land).Conn("sod", godip.Land).Conn("ere", godip.Land).Conn("car", godip.Land).Conn("bor", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		// Eastern Angmar
		Prov("eaa").Conn("car", godip.Land).Conn("foa", godip.Land).Flag(godip.Land).
		// South Downs
		Prov("sod").Conn("ths", godip.Land).Conn("mnh", godip.Land).Conn("dun", godip.Land).Conn("ere", godip.Land).Conn("rhn", godip.Land).Conn("wea", godip.Land).Conn("old", godip.Land).Conn("buc", godip.Land).Flag(godip.Land).
		// Belfalas
		Prov("bel").Conn("lam", godip.Coast...).Conn("bob", godip.Sea).Conn("eth", godip.Sea).Conn("leb", godip.Coast...).Flag(godip.Coast...).
		// Udun Vale
		Prov("udu").Conn("mod", godip.Land).Conn("dea", godip.Land).Flag(godip.Land).SC(godip.Neutral).
		Done()
}

var provinceLongNames = map[godip.Province]string{
	"bro": "Brown Lands",
	"mor": "Moria",
	"lun": "Lune Valley",
	"foc": "Forchel",
	"car": "Carrock",
	"fos": "Fornost",
	"val": "Vale of Erech",
	"haa": "Harad",
	"tow": "Tower Hills",
	"nea": "Near Harad",
	"soi": "South Ithilien",
	"gul": "Gulf of Lune",
	"ths": "The Shire",
	"lef": "Lefnui Sea",
	"sor": "Sea of Rhun",
	"anv": "Anduin Valley",
	"unt": "Unthered Heath",
	"lam": "Lamedon",
	"mim": "Minas Morgul",
	"leb": "Lebennin",
	"umb": "Umbar",
	"osg": "Osgiliath",
	"frl": "Forlindon",
	"bor": "Borderlands",
	"son": "Sea of Nurnen",
	"eam": "East Mirkwood",
	"dea": "Dead Marshes",
	"anf": "Anfalas",
	"foa": "Forodwaith",
	"mit": "Mithlond",
	"eve": "Evendim Hills",
	"dru": "Druwaith Iaur",
	"fan": "Fangorn",
	"nob": "North Belegaer",
	"gap": "Gap of Rohan",
	"esg": "Esgaroth",
	"rhh": "Rhun Hills",
	"bar": "Barad-Dur",
	"mnh": "Minhiriath",
	"nod": "North Downs",
	"nom": "North Mirkwood",
	"kha": "Khand",
	"ere": "Eregion",
	"mod": "Mount Doom",
	"ene": "Enedwaith",
	"Nur": "Nurn",
	"hro": "Harondor",
	"mti": "Minas Tirith",
	"wea": "Weather Hills",
	"sob": "South Belegaer",
	"som": "South Mirkwood",
	"eth": "Ethir Anduin",
	"hal": "Harlindon",
	"gla": "Gladden Fields",
	"Ant": "Andrast",
	"gwa": "Gwathlo Sea",
	"dun": "Dunland",
	"thw": "The Wold",
	"emy": "Emyn Muil",
	"agr": "Angren Sea",
	"bob": "Bay of Belfalas",
	"boh": "Bay of Harmen",
	"anm": "Angmar",
	"srh": "South Rhun",
	"gor": "Gorgoroth",
	"nrh": "North Rhun",
	"dee": "Deep Harad",
	"lor": "Lorien",
	"ant": "Andrast Sea",
	"old": "Old Forest",
	"wes": "West Rohan",
	"buc": "Buckland",
	"rhn": "Rhundaur",
	"eaa": "Eastern Angmar",
	"sod": "South Downs",
	"bel": "Belfalas",
	"udu": "Udun Vale",
}
