
# Create session
new-session -Ad -c ~/Workspace/Go/src/github.com/CarlosMecha/treasuri -n src -s treasuri '/bin/bash --rcfile envrc'
set-option -t treasuri status-style "bg=colour235"
set-option -a -t treasuri status-style "fg=colour9"
set-option -t treasuri default-command "/bin/bash --rcfile envrc"

# Windows

## 1. git
new-window -n git -t treasuri

## 2. db
new-window -n db -t treasuri

## 3. mvn
new-window -n mvn -t treasuri

## 4. test
new-window -n test -t treasuri

## 5. deploy
new-window -n deploy -t treasuri

## 6. cloud
new-window -n cloud -t treasuri

# Notification
display "Session treasuri created"

