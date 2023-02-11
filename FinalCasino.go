package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

var (
	choise         string
	plyName        string
	plyNumber      string
	plyTotalAmount float32
	plyFinalAmount float32
	plyWinStatus   string
	plyLuckyNumber int
	random         int
	plyBetAmount   float32
	check          int
	ID             int
	casinoEarn     float32
	casinoLose     float32
)

func main() {

	CallClear()
	fmt.Println("\n************************************")
	fmt.Println("*    Wellcome In Friend Casino     *")
	fmt.Println("************************************")
	fmt.Println("*     Press 1 for visit only       *")
	fmt.Println("*     Press 2 for play Games       *")
	fmt.Println("*     Press 3 for Exit Games       *")
	fmt.Println("************************************")
	fmt.Print("         Enter your Choice: ")
	fmt.Scan(&choise)
	if choise == "1" {
		visitOnely()
	} else if choise == "2" {
		playGame()
	} else if choise == "3" {
		fmt.Println("Thanks for commming Allah Shony ny Hawaly")
		time.Sleep(3 * time.Second)
	} else {
		fmt.Println("Wrong Input!")
		CallClear()
		main()
	}

}

func visitOnely() {
	fmt.Println("Thanks for comming")
}

func playGame() {

	CallClear()
	fmt.Print("\nInsert you Name: ")
	fmt.Scan(&plyName)
	fmt.Print("\nInsert you Mobile Number: ")
	fmt.Scan(&plyNumber)
	fmt.Print("\nInsert your Total amount: $")
	fmt.Scan(&plyTotalAmount)
	temp := plyTotalAmount

	for check != 2 {
		fmt.Print("\nPlease Enter Your Bet amount $")
		fmt.Scan(&plyBetAmount)
		fmt.Print("\nPlease Enter your lucky numnber Between 0 - 30 : ")
		fmt.Scan(&plyLuckyNumber)
		rand.Seed(time.Now().UnixNano())
		random = rand.Intn(30)

		if plyLuckyNumber == random {
			fmt.Print("\n\nCongratulation you win successfully")
			plyTotalAmount = plyTotalAmount + (plyBetAmount * 10)
			fmt.Print("\nNow your total amount is : ", plyTotalAmount)
			plyFinalAmount = plyTotalAmount

		} else {
			fmt.Print("\n\nSorry you loss the game\nThe lucky number was : ", random)
			plyTotalAmount = plyTotalAmount - plyBetAmount
			fmt.Print("\nNow your total amount is : ", plyTotalAmount)
			plyFinalAmount = plyTotalAmount
		}
		fmt.Print("\n\nPress 1 for play again \nPress 2 for exit \n    Press :")
		fmt.Scan(&check)

	}
	check = 1
	if plyFinalAmount > temp {
		plyWinStatus = "Win"
		casinoLose = plyFinalAmount - temp
		casinoEarn = 0.00
	} else if plyFinalAmount == temp {
		plyWinStatus = "Same"
		casinoEarn = 0.00
		casinoLose = 0.00
	} else {
		plyWinStatus = "Lose"
		casinoEarn = temp - plyFinalAmount
		casinoLose = 0.00
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/Casino")
	if err != nil {
		log.Fatal("Error in database opening \n", err)
	}
	quary := "INSERT INTO `playersdata` (`Name`, `Mobile`, `totalAmount`, `finalAmount`, `winStatus`,`casinoWin`,`casinoLose`, `Day`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	insertResult, err := db.ExecContext(context.Background(), quary, plyName, plyNumber, temp, plyFinalAmount, plyWinStatus, casinoEarn, casinoLose, time.Now())
	if err != nil {
		log.Fatal("Error in data Inserting ", err)
	}
	ID, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatal("Imposible to retrive last result ID ", err)
	}

	fmt.Printf("Your id is : %d", ID)
	time.Sleep(8 * time.Second)
	defer db.Close()
	main()
}
