/*
	Copyright (C) 2023  patapancakes <maru@myyahoo.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Lists referenced from https://github.com/hrydgard/ppsspp

package main

var (
	productLinks = map[string]string{
		// Ace Combat X2 - Joint Assault
		"ULES01408": "ULUS10511",
		"NPJH50263": "ULUS10511",

		// Armored Core 3 Portable
		"ULJM05492": "NPUH10023",

		// BlazBlue - Continuum Shift 2
		"NPJH50401": "ULUS10579",

		// Blood Bowl
		"ULES01230": "ULUS10516",

		// Bomberman
		"ULJM05034": "ULUS10121",
		"ULES00469": "ULUS10121",
		"ULJM05316": "ULUS10121",

		// Bomberman Land
		"ULJM05181": "ULUS10319",
		"ULJM05319": "ULUS10319",
		"ULES00959": "ULUS10319",

		// Call of Duty - Roads to Victory
		"ULES00643": "ULUS10218",

		// Dissidia 012 Duodecim Final Fantasy
		"ULES01505": "ULUS10566",
		"NPJH50377": "ULUS10566",

		// Dissidia Final Fantasy
		"ULES01270": "ULUS10437",
		"ULJM05262": "ULUS10437",

		// Dragon Ball Z - Shin Budokai
		"ULJS00049": "ULUS10081",
		"ULKS46085": "ULUS10081",
		"ULES00309": "ULUS10081",

		// Dragon Ball Z - Shin Budokai 2
		"ULJS00107": "ULUS10234",
		"ULES00789": "ULUS10234",

		// Dragon Ball Z - Tenkaichi Tag Team
		"ULES01456": "ULUS10537",

		// Dungeon Siege - Throne of Agony
		"ULES00569": "ULUS10177",

		// Everybody's Tennis
		"UCJS10101": "UCUS98701",
		"UCES01420": "UCUS98701",

		// Fat Princess - Fistful of Cake
		"UCES01312": "UCUS98740",
		"NPHG00025": "UCUS98740",

		// God Eater Burst
		"ULES01519": "ULUS10563",
		"NPJH50352": "ULUS10563",

		// Gran Turismo
		"UCES01245": "UCUS98632",
		"UCES00543": "UCUS98645",

		// Gundam VS Gundam - Next Plus
		"ULJS00250": "NPJH50107",
		"ULJS19048": "NPJH50107",

		// Hatsune Miku - Project Diva Extend
		"NPJH50465": "ULJM05933",

		// Hot Pixel
		"ULES00642": "ULUS10298",

		// Lord of Arcana
		"ULJM05767": "ULES01507",
		"ULUS10479": "ULES01507",

		// M.A.C.H. - Modified Air Combat Heroes
		"ULES00565": "ULUS10180",
		"ULES00566": "ULUS10180",
		"ULJM05202": "ULUS10180",

		// Metal Gear Solid - Peace Walker
		"ULES01372": "NPJH50045",
		"ULUS10509": "NPJH50045",

		// Metal Gear Solid - Portable Ops
		"ULES00645": "ULUS10202",
		"ULJM05193": "ULUS10202",

		// Metal Gear Solid - Portable Ops +
		"ULES01003": "ULUS10290",
		"ULJM05261": "ULUS10290",

		// Midnight Club - LA Remix
		"ULES01144": "ULUS10383",
		"ULJS00180": "ULUS10383",

		// Mod Nation Racers
		"UCES01327": "UCUS98741",
		"UCJS10112": "UCUS98741",
		"UCAS40306": "UCUS98741",

		// Monster Hunter Freedom
		"ULJM05066": "ULUS10084",
		"ULES00318": "ULUS10084",

		// Monster Hunter Freedom 2
		"ULJM05156": "ULUS10266",
		"ULES00851": "ULUS10266",

		// Monster Hunter Freedom Unite
		"ULES01213": "ULUS10391",
		"ULJM05500": "ULUS10391",

		// N+
		"ULES01026": "ULUS10340",

		// Need for Speed - Undercover
		"ULJM05403": "ULUS10376",
		"ULJM05612": "ULUS10376",
		"ULES01145": "ULUS10376",

		// Outrun 2006 - Coast 2 Coast
		"ULES00262": "ULUS10064",

		// Pangya! - Fantasy Golf
		"ULJM05440": "ULUS10438",
		"ULKS46164": "ULUS10438",

		// PRO Evolution Soccer 2012
		"ULES01540": "ULUS10586",
		"ULES01541": "ULUS10586",
		"ULES01542": "ULUS10586",
		"ULAS42289": "ULUS10586",

		// Patapon 2
		"UCJS10089": "UCUS98732",
		"PSPJ30000": "UCUS98732",
		"UCES01177": "UCUS98732",
		"UCJS18036": "UCUS98732",

		// Patapon 3
		"UCES01421": "UCUS98751",
		"NPJG00122": "UCUS98751",

		// Phantasy Star Portable
		"ULJM05309": "ULUS10410",
		"ULES01218": "ULUS10410",
		"ULJM08023": "ULUS10410",

		// Phantasy Star Portable 2
		"ULJM05493": "ULUS10529",
		"ULJM08030": "ULUS10529",
		"ULES01439": "ULUS10529",

		// Resistance - Retribution
		"UCES01184": "UCJS10090",
		"UCUS98668": "UCJS10090",

		// Rocky Balboa
		"ULUS10233": "ULES00670",

		// SOCOM - Fireteam Bravo
		"UCES00038": "UCUS98615",
		"UCJS10102": "UCUS98615",

		// SOCOM - Fireteam Bravo 3
		"UCES01242": "UCUS98716",
		"NPJG00035": "UCUS98716",

		// Shrek - Smash and Crash Racing
		"ULES00618": "ULUS10194",

		// Smash Court Tennis 3
		"ULJS00098": "UCES00758",
		"ULUS10269": "UCES00758",

		// Soul Calibur - Broken Destiny
		"ULES01298": "ULUS10457",
		"ULJS00202": "ULUS10457",

		// Split Second - Velocity
		"ULES01402": "ULUS10513",
		"ULJM05812": "ULUS10513",

		// Street Fighter Alpha 3 MAX
		"ULJM05082": "ULUS10062",
		"ULES00235": "ULUS10062",
		"ULJM05225": "ULUS10062",

		// Taiko no Tatsujin Portable DX"
		"ULJS00383": "NPJH50426",

		// Tekken 6
		"ULES01376": "ULUS10466",
		"NPJH50184": "ULUS10466",
		"ULJS00224": "ULUS10466",

		// TRON - Evolution
		"ULES01495": "ULUS10548",

		// Untold Legends - Brotherhood of the Blade
		"ULES00046": "ULUS10003",
		"ULJM05087": "ULUS10003",
		"ULKS46015": "ULUS10003",

		// Untold Legends - The Warrior's Code
		"ULES00301": "ULUS10086",
		"ULJM05179": "ULUS10086",
		"ULKS46069": "ULUS10086",

		// Virtua Tennis 3
		"ULES00763": "ULUS10246",

		// World Series of Poker 2008 - Battle for the Bracelets
		"ULES00991": "ULUS10321",

		// Worms Battle Islands
		"NPEH00019": "NPUH10045",

		// Worms Open Warfare
		"ULES00268": "ULUS10065",

		// Worms Open Warfare 2
		"ULES00819": "ULUS10260",

		// Yu-Gi-Oh! 5D's Tag Force 5
		"ULUS10555": "ULJM05734",
		"ULES01474": "ULJM05734",
	}

	productNames = map[string]string{
		"ULUS10511": "Ace Combat X2 - Joint Assault",
		"ULUS10245": "Alien Syndrome",
		"NPUH10023": "Armored Core 3 Portable",
		"ULES00719": "Asphalt - Urban GT 2",
		"ULUS10579": "BlazBlue - Continuum Shift 2",
		"ULUS10519": "BlazBlue Calamity Trigger",
		"UCJS10110": "Bleach Heat The Soul 7",
		"ULUS10516": "Blood Bowl",
		"ULUS10121": "Bomberman",
		"ULUS10319": "Bomberman Land",
		"ULES00703": "Burnout Dominator",
		"ULES00125": "Burnout Legends",
		"ULJM05538": "Busou Shinki - Battle Masters",
		"ULUS10057": "Bust A Move Deluxe",
		"ULUS10218": "Call of Duty - Roads to Victory",
		"ULUS10351": "Code Lyoko - Quest for Infinity",
		"NPJH50583": "Conception - Please have my children!",
		"ULUS10044": "Crash Tag Team Racing",
		"ULUS10100": "Def Jam Fight For NY - The Takeover",
		"NPJH50588": "Digimon World Re:Digitize",
		"ULUS10566": "Dissidia 012 Duodecim Final Fantasy",
		"ULUS10437": "Dissidia Final Fantasy",
		"ULUS10081": "Dragon Ball Z - Shin Budokai",
		"ULUS10234": "Dragon Ball Z - Shin Budokai 2",
		"ULUS10537": "Dragon Ball Z - Tenkaichi Tag Team",
		// maybe we can crosslinks this 2 region to ULUS10537 not having the game to test
		"ULJS00311": "Dragon Ball Z - Tenkaichi Tag Team",
		"NPJH90135": "Dragon Ball Z - Tenkaichi Tag Team",
		"ULJM05127": "Dragon Quest & Final Fantasy in Itadaki Street Special",
		"ULES00847": "Dungeon Explorer - Warriors of Ancient Arts",
		"ULUS10177": "Dungeon Siege - Throne of Agony",
		"ULUS10170": "Dynasty Warrior 2",
		// looks like can be crosslinked too
		"ULES01221": "Dynasty Warriors - Strike Force",
		"ULUS10416": "Dynasty Warriors - Strike Force",
		"UCUS98701": "Everybody's Tennis",
		"UCUS98740": "Fat Princess - Fistful of Cake",
		"ULJM05360": "Fate Tiger Colosseum Upper",
		"ULUS10297": "Final Fantasy Tactics - The War of the Lions",
		"ULES00850": "Final Fantasy Tactics - War of the Lions",
		"NPJH50443": "Final Fantasy Type 0",
		"NPJH50468": "Frontier Gate",
		"NPJH50721": "Frontier Gate Boost+",
		"ULES01432": "Full Metal Alchemist - Brotherhood",
		"ULUS10490": "GTA Chinatown Wars",
		"ULUS10160": "GTA Vice City Stories",
		"ULUS10210": "Ghost Rider",
		"ULJS00237": "God Eater",
		"NPJH50832": "God Eater 2",
		"ULUS10563": "God Eater Burst",
		"UCUS98632": "Gran Turismo",
		"NPJH50107": "Gundam VS Gundam - Next Plus",
		"ULJM05933": "Hatsune Miku - Project Diva Extend",
		"ULUS10298": "Hot Pixel",
		"ULJM05709": "K-ON! Houkago Live",
		"NPJH50221": "Kateikyoushi Hitman Reborn! Kizuna no Tag Battle",
		"ULJS00165": "Kidou Senshi Gundam - Gundam vs. Gundam",
		"UCUS98646": "Killzone Liberation",
		"ULJM05775": "Kingdom Hearts - Birth by Sleep Final Mix",
		"ULUS10487": "LEGO Indiana Jones 2",
		"NPJH50503": "Lord of Apocalypse",
		"ULES01507": "Lord of Arcana",
		"ULUS10180": "M.A.C.H. - Modified Air Combat Heroes",
		"UCUS98758": "MLB11 - The Show",
		"ULUS10581": "Madden NFL 12",
		"ULJS00385": "Mahou Shoujo Nanoha A's Portable - The Gears of Destiny",
		"ULUS10408": "Mana Khemia Student Alliance",
		"ULUS10141": "Medal Of Honor Heroes",
		"NPJH50045": "Metal Gear Solid - Peace Walker",
		"ULUS10202": "Metal Gear Solid - Portable Ops",
		"ULUS10290": "Metal Gear Solid - Portable Ops +",
		"ULUS10154": "Metal Slug Anthology",
		"ULUS10495": "Metal Slug XX",
		"ULES01429": "Metal Slug XX",
		"ULES00368": "Micro Machines V4",
		"ULUS10383": "Midnight Club - LA Remix",
		"UCUS98741": "Mod Nation Racers",
		"ULUS10084": "Monster Hunter Freedom",
		"ULUS10266": "Monster Hunter Freedom 2",
		"ULUS10391": "Monster Hunter Freedom Unite",
		"ULJM05800": "Monster Hunter Portable 3rd",
		"ULJM06097": "Musou Orochi 2 Special",
		"ULUS10340": "N+",
		"ULES01578": "NBA 2K13",
		"ULUS10598": "NBA 2K13",
		"ULUS10349": "Naruto - Ultimate Ninja Heroes 2",
		"ULUS10518": "Naruto - Ultimate Ninja Heroes 3",
		"ULJS00236": "Naruto - Accel 3",
		"ULUS10582": "Naruto Shippuden - Ultimate Ninja Impact",
		"ULES01537": "Naruto Shippuden - Ultimate Ninja Impact",
		"ULUS10571": "Naruto Shippuden - Kizuna Drive",
		"ULES00196": "Need For Speed - Most Wanted",
		"ULUS10036": "Need For Speed - Most Wanted",
		"ULUS10376": "Need for Speed - Undercover",
		"ULKS46004": "Need for Speed - Underground Rivals",
		"ULES01340": "Obscure - The Aftermath",
		"ULUS10064": "Outrun 2006 - Coast 2 Coast",
		"ULUS10586": "PRO Evolution Soccer 2012",
		"ULUS10149": "Pac Man - World Rally",
		"ULUS10438": "Pangya! - Fantasy Golf",
		"UCUS98732": "Patapon 2",
		"UCUS98751": "Patapon 3",
		"UCES01422": "Patapon 3 Overhaul", // mod
		"ULUS10410": "Phantasy Star Portable",
		"ULUS10529": "Phantasy Star Portable 2",
		// looks like this japan version can crosslink to ULUS10529
		"NPJH50332": "Phantasy Star Portable 2",
		"ULJM05732": "Phantasy Star Portable 2 - Infinity",
		"ULES01596": "Pro Evolution Soccer 2014",
		"ULES01595": "Pro Evolution Soccer 2015",
		"NPJH50520": "Pro Yakyuu Spirits 2012",
		"NPJH50838": "Pro Yakyuu Spirits 2014",
		"NPJH50492": "Puyo Puyo!! 20th Anniversary",
		"ULUS10292": "Renegrade Squadron",
		"UCJS10090": "Resistance - Retribution",
		"ULES00670": "Rocky Balboa",
		"ULJS00360": "Rurouni Kenshin - Meiji Kenkaku Romantan Saisen",
		"UCUS98615": "SOCOM - Fireteam Bravo",
		"UCUS98645": "SOCOM - Fireteam Bravo 2",
		"UCUS98716": "SOCOM - Fireteam Bravo 3",
		"NPJH50460": "Sengoku Basara - Chronicles Heroes",
		"ULJM05436": "Sengoku Basara - Battle Heroes",
		"ULJM05637": "Shin Sangoku Musou - Multi Raid 2",
		"ULJM05035": "Shinobido - Tales of the Ninja",
		"ULUS10194": "Shrek - Smash and Crash Racing",
		"UCES00758": "Smash Court Tennis 3",
		"ULUS10195": "Sonic Rivals",
		"ULUS10457": "Soul Calibur - Broken Destiny",
		"ULUS10513": "Split Second - Velocity",
		"ULES00183": "Star Wars Battle Front 2",
		"ULUS10062": "Street Fighter Alpha 3 MAX",
		"NPUH10020": "Strikers 1945 Plus Portable",
		"ULUS10548": "TRON - Evolution",
		"NPJH50426": "Taiko no Tatsujin Portable DX",
		"ULUS10466": "Tekken 6",
		"NPJH50691": "Tokusatsu University",
		// looks like can be crosslinked
		"ULUS10445": "Tom Clancy's Ghost Recon - Predator",
		"ULES01350": "Tom Clancy's Ghost Recon - Predator",
		"NPJH50789": "Toukiden",
		"NPJH50878": "Toukiden - Kiwami",
		"UCUS98601": "Twisted Metal - Head On",
		"ULUS10508": "UFC Undisputed 2010",
		"ULJS00069": "Ultraman Fighting Evo Zero",
		"ULUS10003": "Untold Legends - Brotherhood of the Blade",
		"ULUS10086": "Untold Legends - The Warrior's Code",
		"ULUS10515": "Valkryia Chronicles 2",
		"ULUS10087": "Viewtiful Joe",
		"ULUS10246": "Virtua Tennis 3",
		"ULUS82741": "WWE 2K14",
		"ULUS10543": "WWE Smackdown vs. Raw 2011",
		"ULUS10423": "Warriors Orochi 2",
		"ULJM05553": "Warship Gunner 2 Portable",
		"ULJS00155": "Way Of The Samurai",
		"UCES00465": "Wipeout Pulse",
		"ULUS10321": "World Series of Poker 2008 - Battle for the Bracelets",
		"NPUH10045": "Worms Battle Islands",
		"ULUS10065": "Worms Open Warfare",
		"ULUS10260": "Worms Open Warfare 2",
		"ULJM05734": "Yu-Gi-Oh! 5D's Tag Force 5",
		"ULJM05940": "Yu-Gi-Oh! 5D's Tag Force 6",
		"NPJH00142": "Yu-Gi-Oh! Arc-V Tag Force",
		"ULJM05151": "Yu-Gi-Oh! GX Tag Force",
		"ULJM05373": "Yu-Gi-Oh! GX Tag Force 3",
		"NPUG80086": "flOw",
	}
)
