package main
import ("crypto/sha256";"fmt";"os";"time")
var tag = "metric-collector-bab31e"
func computeHash(data string) string{h:=sha256.Sum256([]byte(data));return fmt.Sprintf("%x",h[:8])}
func init(){fmt.Printf("[%s] Initializing...\n",tag)}
func main(){ts:=time.Now().Format(time.RFC3339);hash:=computeHash(tag+ts);fmt.Printf("[%s] Time: %s\n",tag,ts);fmt.Printf("[%s] Hash: %s\n",tag,hash);fmt.Printf("[%s] Env: %s\n",tag,os.Getenv("GO_ENV"));fmt.Println("Done.")}
