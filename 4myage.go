package main

//importing standard package fmt and time
import(
  "fmt"
  "time"
)

//CalculateAge takes the time now and the birth date as an input
//CalculateAge returns the age in Year, Month, and Day (in this particular order)
func CalculateAge(now time.Time, birth time.Time) (int,int,int) {

  //short declaration for time now and just taking in the Year, Month and Day
  nYear,nMonth,nDay:=now.Date()

  //short declaration for birth date and just taking in the Year, Month and Day
  bYear,bMonth,bDay:=birth.Date()

  //short declaration for the age result in Year, Month and Day
  rYear := nYear-bYear
  rMonth := int(nMonth-bMonth)
  rDay := nDay-bDay

  //adding condition if the result is less than 0
  if rMonth < 0 {
    rMonth += 12
    rYear--
  }

  if rDay < 0 {
    rDay += birth.YearDay()-now.YearDay()
    rMonth--
  }

  //just printing to check the numbers
  fmt.Println(rYear,rMonth,rDay)

  return rYear,rMonth,rDay
}

func main(){
  const shortFormat = "02 Jan 2006"
  var ageYear,ageMonth,ageDay int
  var myBirth time.Time
  current:=time.Now()
  fmt.Println("Today is",current)
  myBirth,_ = time.Parse(shortFormat,"31 Dec 1985")
  fmt.Println("The birthday is",myBirth)
  fmt.Println(time.Now().Date())
  ageYear,ageMonth,ageDay = CalculateAge(current,myBirth)
  fmt.Printf("My age is %d years %d months %d days\n", ageYear,ageMonth,ageDay)
}
