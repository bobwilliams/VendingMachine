package main 

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strconv"
  "time"
)

const width int = 6
const height int = 10
const depth int = 20
const minAmount int = 1
var snacks = [...]string{"Candy Bar", "Donuts", "Chips", "Gummies", "Honey Buns", "Crackers", "Popcorn", "Gum", "Pork Rinds", "Beef Jerky"}

type SnackItem struct {
  Name string
  Cost int
  Quantity int
}

type VendingMachine struct {
  Items [width][height]SnackItem
}

func StockRow() string {
  return snacks[rand.Intn(len(snacks))]
}

func Intialize() VendingMachine {
  rand.Seed(time.Now().UnixNano())
  var items = [width][height]SnackItem{}
  for r := 0; r < width; r++ {
    for c := 0; c < height; c++ {
      items[r][c] = SnackItem{StockRow(), r+c+minAmount, depth}
    }
  }
  return VendingMachine{items}
}

func VendHelper(s *bufio.Scanner, p string) int {
  fmt.Printf("%s", p)
  s.Scan()
  r, _ := strconv.Atoi(s.Text())
  return r
}

func Vend(s *bufio.Scanner, m *VendingMachine) {
  amount := VendHelper(s, "amount: ")
  row := VendHelper(s, "row: ")
  col := VendHelper(s, "col: ")
  item := &m.Items[row][col]

  if item.Cost <= amount {
    if item.Quantity > 0 {
      item.Quantity--  
      fmt.Println(item.Name)
    } else {
      fmt.Println("No more items left")
    }
  } else {
    fmt.Println("Not enough money")
  }
}

func Welcome() {
  fmt.Printf("%s\n", "\nWelcome to the Golang Vending Machine model.  Please select one of the commands below.")
  Usage()
}

func Usage() {
  fmt.Printf("%s", "\nv to vend\ni to inventory\nu to see usage\nx to exit\n>> ")
}

func Inventory(m VendingMachine) {
  for r := range m.Items {
    for c := 0; c < height; c++ {
      i := m.Items[r][c]
      fmt.Printf("[%s $%d #%d]", i.Name, i.Cost, i.Quantity)
    }
    fmt.Println()
  }
}

func main() {
  machine := Intialize()
  scanner := bufio.NewScanner(os.Stdin)

  Welcome()

  for scanner.Scan() {
    command := scanner.Text()

    switch command[0] {
    case 'v' :
      Vend(scanner, &machine)
    case 'i' :
      Inventory(machine)
    case 'u' :
      Usage()
    case 'x' :
      os.Exit(1) 
    }

    Usage()
  }
}