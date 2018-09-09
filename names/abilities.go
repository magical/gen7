package names

func Ability(n int) string {
	if 0 <= n && n < len(abilityNames) {
		return abilityNames[n]
	}
	return ""
}

var abilityNames = []string{
	"—",
	"Stench",
	"Drizzle",
	"Speed Boost",
	"Battle Armor",
	"Sturdy",
	"Damp",
	"Limber",
	"Sand Veil",
	"Static",
	"Volt Absorb",
	"Water Absorb",
	"Oblivious",
	"Cloud Nine",
	"Compound Eyes",
	"Insomnia",
	"Color Change",
	"Immunity",
	"Flash Fire",
	"Shield Dust",
	"Own Tempo",
	"Suction Cups",
	"Intimidate",
	"Shadow Tag",
	"Rough Skin",
	"Wonder Guard",
	"Levitate",
	"Effect Spore",
	"Synchronize",
	"Clear Body",
	"Natural Cure",
	"Lightning Rod",
	"Serene Grace",
	"Swift Swim",
	"Chlorophyll",
	"Illuminate",
	"Trace",
	"Huge Power",
	"Poison Point",
	"Inner Focus",
	"Magma Armor",
	"Water Veil",
	"Magnet Pull",
	"Soundproof",
	"Rain Dish",
	"Sand Stream",
	"Pressure",
	"Thick Fat",
	"Early Bird",
	"Flame Body",
	"Run Away",
	"Keen Eye",
	"Hyper Cutter",
	"Pickup",
	"Truant",
	"Hustle",
	"Cute Charm",
	"Plus",
	"Minus",
	"Forecast",
	"Sticky Hold",
	"Shed Skin",
	"Guts",
	"Marvel Scale",
	"Liquid Ooze",
	"Overgrow",
	"Blaze",
	"Torrent",
	"Swarm",
	"Rock Head",
	"Drought",
	"Arena Trap",
	"Vital Spirit",
	"White Smoke",
	"Pure Power",
	"Shell Armor",
	"Air Lock",
	"Tangled Feet",
	"Motor Drive",
	"Rivalry",
	"Steadfast",
	"Snow Cloak",
	"Gluttony",
	"Anger Point",
	"Unburden",
	"Heatproof",
	"Simple",
	"Dry Skin",
	"Download",
	"Iron Fist",
	"Poison Heal",
	"Adaptability",
	"Skill Link",
	"Hydration",
	"Solar Power",
	"Quick Feet",
	"Normalize",
	"Sniper",
	"Magic Guard",
	"No Guard",
	"Stall",
	"Technician",
	"Leaf Guard",
	"Klutz",
	"Mold Breaker",
	"Super Luck",
	"Aftermath",
	"Anticipation",
	"Forewarn",
	"Unaware",
	"Tinted Lens",
	"Filter",
	"Slow Start",
	"Scrappy",
	"Storm Drain",
	"Ice Body",
	"Solid Rock",
	"Snow Warning",
	"Honey Gather",
	"Frisk",
	"Reckless",
	"Multitype",
	"Flower Gift",
	"Bad Dreams",
	"Pickpocket",
	"Sheer Force",
	"Contrary",
	"Unnerve",
	"Defiant",
	"Defeatist",
	"Cursed Body",
	"Healer",
	"Friend Guard",
	"Weak Armor",
	"Heavy Metal",
	"Light Metal",
	"Multiscale",
	"Toxic Boost",
	"Flare Boost",
	"Harvest",
	"Telepathy",
	"Moody",
	"Overcoat",
	"Poison Touch",
	"Regenerator",
	"Big Pecks",
	"Sand Rush",
	"Wonder Skin",
	"Analytic",
	"Illusion",
	"Imposter",
	"Infiltrator",
	"Mummy",
	"Moxie",
	"Justified",
	"Rattled",
	"Magic Bounce",
	"Sap Sipper",
	"Prankster",
	"Sand Force",
	"Iron Barbs",
	"Zen Mode",
	"Victory Star",
	"Turboblaze",
	"Teravolt",
	"Aroma Veil",
	"Flower Veil",
	"Cheek Pouch",
	"Protean",
	"Fur Coat",
	"Magician",
	"Bulletproof",
	"Competitive",
	"Strong Jaw",
	"Refrigerate",
	"Sweet Veil",
	"Stance Change",
	"Gale Wings",
	"Mega Launcher",
	"Grass Pelt",
	"Symbiosis",
	"Tough Claws",
	"Pixilate",
	"Gooey",
	"Aerilate",
	"Parental Bond",
	"Dark Aura",
	"Fairy Aura",
	"Aura Break",

	// Added in OR/AS
	"Primordial Sea",
	"Desolate Land",
	"Delta Stream",
}
