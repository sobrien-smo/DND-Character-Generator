package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
)

var (
	run          bool
	actionChoice string
	charType     string
	saveChoice   string
	newFileName  string
	checkFile    string

	classes    = [13]string{"Artificer", "Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"}
	background = [10]string{"Explorer", "Folk Hero", "Noble", "Nomad", "Outlaw", "Peasant", "Sage", "Scholar", "Soldier", "Warrior"}
	races      = [10]string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Goblin", "Half-Elf", "Halfling", "Half-Orc", "Human", "Tiefling"}
	alignment  = [9]string{"Lawful Good", "Neutral Good", "Chaotic Good", "Lawful Neutral", "True Neutral", "Chaotic Neutral", "Lawful Evil", "Neutral Evil", "Chaotic Evil"}
	stats      = [6]string{"Charisma", "Constitution", "Dexterity", "Intelligence", "Strength", "Wisdom"}

	trade           = [7]string{"Sailor", "Smith", "Baker", "City official", "Acrobat", "Cook", "Teacher"}
	talent          = [10]string{"Perfect memory", "Plays a musical instrument", "Can drink anyone in the tavern under the table", "Knows thieves cant", "Speaks several languages", "Unbelievably lucky", "Good with animals", "Master of disguise", "Gifted actor", "Beautiful singer"}
	interaction     = [12]string{"Argumentative", "Arrogant", "Blustering", "Wise", "Curious", "Friendly", "Honest", "Hot tempered", "Irritable", "Ponderous", "Quiet", "Unfocused"}
	mannerisms      = [10]string{"Sings, whistles, or hums quietly", "Whistles when they lie", "Taps fingers or feet", "Fidgets", "Uses colorful oaths/exclamations", "Twirls hair or tugs beard", "Always makes puns", "Sarastic", "Early riser", "Night owl"}
	ideals          = [8]string{"Aspiration", "Discovery", "Glory", "Nation", "Redemption", "Knowledge", "Adventuring", "Friendship"}
	bond            = [7]string{"Dedicated to a personal goal", "Protective of close family members", "Protective of party", "Captivated by romantic interest", "Out for revenge", "Bound to benefactor, patron, or employer", "Always accompanied by a small animal companion"}
	flaws           = [8]string{"Forbidden love, or susceptible to romance", "Arrogant", "Prone to rage", "Has a powerful enemy", "Specific phobia", "Secret crime", "Posession of forbidden lore or artifact", "Foolhardy bravery"}
	relationToParty = [6]string{"Ally?: This NPC could prove useful to the party. ",
		"Adoptable: This character becomes very fond of the party.",
		"BBEG: This character is the campaign's 'Big Bad Evil Guy' in disguise",
		"Nuisance: Always in the wrong place at the wrong time",
		"Lost Puppy: Much to your party's annoyance, this NPC follows your party everywhere",
		"Traveler's Best Friend: This character runs an inn/tavern the party frequents"}
	physicalTrait = [8]string{"Bushy sideburns", "Noticeably crooked teeth", "Freckles", "Extra finger", "Scar from animal attack", "Single braid in beard", "Thin and wiry", "Missing teeth"}
)

type Character struct {
	Gender     string
	Class      string
	Background string
	Race       string
	Alignment  string
	Stats      []Stat
}
type NPC struct {
	Gender          string
	Trade           string
	Talent          string
	Interaction     string
	PhysicalTrait   string
	Mannerisms      string
	Ideals          string
	Bond            string
	Flaws           string
	RelationToParty string
}

type Stat struct {
	Name string
	Roll int
}

func main() {
	run = true
	for run {
		fmt.Println("Would you like to create a new character (1) or a view a saved character (2)?")
		fmt.Scanln(&actionChoice)
		if actionChoice == "1" {
			fmt.Println("Would you like to create a playable character (1) or an NPC (2)?")
			fmt.Scanln(&charType)
			if charType == "1" {
				newChar := newChar()
				fmt.Println("Here is your new character:")
				fmt.Println(newChar)
				fmt.Println("Save character? (y/n)")
				fmt.Scanln(&saveChoice)
				if saveChoice == "y" {
					fmt.Println("Enter a name for your file:")
					fmt.Scanln(&newFileName)
					exists := fileExists(newFileName)
					if exists {
						fmt.Println(newFileName, "Looks like you already created a file with this name")
						readCharFile(newFileName)
					} else {
						writeCharToFile(newChar, newFileName)
					}
				} else if saveChoice == "n" {
					break
				}
			} else if charType == "2" {
				newNPC := newNPC()
				fmt.Println("Here is your new NPC:")
				fmt.Println(newNPC)
				fmt.Println("Save character? (y/n)")
				fmt.Scanln(&saveChoice)
				if saveChoice == "y" {
					fmt.Println("Enter a name for your file:")
					fmt.Scanln(&newFileName)
					exists := fileExists(newFileName)
					if exists {
						fmt.Println(newFileName, "Looks like you already created a file with this name")
						readNPCFile(newFileName)
					} else {
						writeNPCToFile(newNPC, newFileName)
					}
				} else if saveChoice == "n" {
					break
				}
			} else {
				fmt.Println("Please enter either 1 or 2")
			}

		} else if actionChoice == "2" {
			fmt.Println("Enter the file name for the saved characater")
			fmt.Scanln(&checkFile)
			if charType == "1" {
				readCharFile(checkFile)
			} else {
				readNPCFile(checkFile)
			}

		} else {
			fmt.Println("Please enter either 1 or 2")
		}
		fmt.Println("Return to menu? (y/n)")
		fmt.Scanln(&actionChoice)
		if actionChoice != "y" {
			run = false
		}
	}
}

func fileExists(fileName string) bool {
	// Ask os if the stats of the file
	fileInfo, err := os.Stat(fileName)
	// checkErr(err)
	if fileInfo != nil && err == nil {
		return true
	} else {
		return false
	}
}

func newChar() Character {
	gender := getGender()
	class := classes[genRandNum(0, 12)]
	background := background[genRandNum(0, 11)]
	race := races[genRandNum(0, 10)]
	alignment := alignment[genRandNum(0, 8)]
	stats := []Stat{
		{Name: "Strength", Roll: getStatsRoll()},
		{Name: "Dexterity", Roll: getStatsRoll()},
		{Name: "Constitution", Roll: getStatsRoll()},
		{Name: "Intelligence", Roll: getStatsRoll()},
		{Name: "Wisdom", Roll: getStatsRoll()},
		{Name: "Charisma", Roll: getStatsRoll()},
	}
	char := Character{
		Gender:     gender,
		Class:      class,
		Background: background,
		Race:       race,
		Alignment:  alignment,
		Stats:      stats}
	return char
}

func newNPC() NPC {
	gender := getGender()
	trade := trade[genRandNum(0, 6)]
	talent := talent[genRandNum(0, 9)]
	interaction := interaction[genRandNum(0, 11)]
	mannerisms := mannerisms[genRandNum(0, 9)]
	ideals := ideals[genRandNum(0, 8)]
	bond := bond[genRandNum(0, 7)]
	flaws := flaws[genRandNum(0, 8)]
	relationToParty := relationToParty[genRandNum(0, 5)]
	physicalTrait := physicalTrait[genRandNum(0, 7)]
	npc := NPC{
		Gender:          gender,
		Trade:           trade,
		Talent:          talent,
		Interaction:     interaction,
		Mannerisms:      mannerisms,
		Ideals:          ideals,
		Bond:            bond,
		Flaws:           flaws,
		RelationToParty: relationToParty,
		PhysicalTrait:   physicalTrait,
	}
	return npc
}

// func saveCheck(content interface{}) {
// 	fmt.Println("Enter a name for your file:")
// 	fmt.Scanln(&newFileName)
// 	exists := fileExists(newFileName)
// 	if exists {
// 		fmt.Println(newFileName, "Looks like you already created a file with this name")

//			readNPCFile(newFileName)
//		} else {
//			writeCharToFile(newFileName)
//		}
//	}
func createFile(fileName string) {
	file, err := os.Create(fileName)
	checkErr(err)
	fmt.Print("created", file)
}
func writeCharToFile(newChar Character, fileName string) {

	file, _ := json.Marshal(newChar)

	_ = ioutil.WriteFile(fileName, file, 0644)
}

func writeNPCToFile(newNPC NPC, fileName string) {
	file, _ := json.Marshal(newNPC)
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func readCharFile(fileName string) {
	var char Character
	data, _ := ioutil.ReadFile(fileName)
	_ = json.Unmarshal(data, &char)
	fmt.Println("Gender:", char.Gender)
	fmt.Println("Class:", char.Class)
	fmt.Println("Background:", char.Background)
	fmt.Println("Race:", char.Race)
	fmt.Println("Alignment:", char.Alignment)
	fmt.Println("Stats:", char.Stats)

}

func readNPCFile(fileName string) {
	var npc NPC
	data, _ := ioutil.ReadFile(fileName)
	_ = json.Unmarshal(data, &npc)
	fmt.Println("Gender:", npc.Gender)
	fmt.Println("Trade:", npc.Trade)
	fmt.Println("Talent:", npc.Talent)
	fmt.Println("Interaction:", npc.Interaction)
	fmt.Println("Mannerisms:", npc.Mannerisms)
	fmt.Println("Ideals:", npc.Ideals)
	fmt.Println("Bond:", npc.Bond)
	fmt.Println("Flaws:", npc.Flaws)
	fmt.Println("Relation to Party:", npc.RelationToParty)

}

func genRandNum(min, max int) int {
	randNum := rand.Intn(max-min) + 1
	return randNum
}
func getGender() string {
	checkVal := genRandNum(1, 3)
	genderString := ""
	if checkVal == 1 {
		genderString = "Female"
	} else if checkVal == 2 {
		genderString = "Male"
	}
	return genderString
}
func getStatsRoll() int {
	var rolls []int
	total := 0
	for i := 0; i < 4; i++ {
		newRoll := genRandNum(1, 7)
		rolls = append(rolls, newRoll)
		sort.Ints(rolls)
	}
	rolls = rolls[1:]
	for i := 0; i < 3; i++ {
		total = total + rolls[i]
	}
	// fmt.Println("rolls", rolls)
	// fmt.Println("total", total)
	return total
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
