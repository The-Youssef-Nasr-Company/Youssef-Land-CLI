import (
        "fmt"
        "time"
)

func main() {
  fmt.Println("Zzz...");
  time.Sleep(2);
  fmt.Println("Huh?");
  time.Sleep(3);
  fmt.Println("Hey, ya wakin' me up?!");
  time.Sleep(2);
  fmt.Println("OK, I'm lettin' you in...");
}

main();
