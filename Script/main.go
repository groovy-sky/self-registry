//https://gowebexamples.com/password-hashing/

package main

import (
    "fmt"
    "golang.org/x/crypto/ssh/terminal"
    "golang.org/x/crypto/bcrypt"
    "os"
    "bufio"
    "syscall"
    "strings"
    "log"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func credentials() (string, string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter Username: ")
    username, err := reader.ReadString('\n')

    if err != nil || len(username) <= 1 {
        log.Fatal("Incorrect Username")
    }
    
    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err != nil || bytePassword == nil {
        log.Fatal("Incorrect Password") 
    }
    fmt.Println("")
    password := string(bytePassword)

    return strings.TrimSpace(username), strings.TrimSpace(password)
}

func main() {
    username, password := credentials()
    hash, err := HashPassword(password) 
    if err != nil {                                                                    
        log.Fatal("Couldn't generate a hash")                                                                       
    } 
    htpasswd, err := os.Create("htpasswd")
    if err != nil {                                                                                           
        log.Fatal("Couldn't access htpasswd file")                                                                 
    }
    defer htpasswd.Close()
    
    htpasswd.WriteString(username + ":" +hash)
    htpasswd.Sync()

    fmt.Println("\n[INFO] htpasswd was successfully generate\n")
}
