# :monkey_face: :email: Go Mail With SMTP Google
    First you must be less secure app access is ON. chek this link https://myaccount.google.com/lesssecureapps
## Install Your Need Package.
    go get github.com/gin-gonic/gin
## Set Your Account Google in config/send.go
    sender := NewSender("<Your Email>", "<Your Password>")

## Change Your Message
    1. send with html 
       sender.SendMail(Receiver, Subject, bodyMessage)
    2. send with Plain Text
       sender.WritePlainEmail(Receiver, Subject, bodyMessage)
    
## And Run
    go run main.go