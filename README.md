# cocoon
cocoon

type CocoonHalder interface {
    HandleLetter(lt *Letter) (uint8, []byte, error)
}

Cocoon
    Citizen (Taistra, Justice)
    Contact (All)
    Letter
        Unprocessed Letter
        Unsent Letter
        Processed Letter
        Sent Letter
        Refer Letter

type Cocoon struct {
    UnprocessedLetterChan chan *Letter
    UnsentLetterChan chan *Letter
    ProcessLetterChan chan *Letter
    SendLetterChan chan *Letter
    TerminateChan chan *boolean
    CocoonHalder
}

package main

func (ccn *Cocoon) HandleLetter(lt *Letter) (uint8, []byte, error) {
    return Handle(lt)
}

ccn := cocoon.New()

cnn.Start()

funct (c *Cocoon) Start() {
    go ccn.ListenCocoon()
    go ccn.ListenContactGate()
    go ccn.HandleUnpreocessedLetter()
    ...
    select {
    case <- c.TerminateChan
        fmt.Fatalf("exit")
    }
}

Cocoon:
    cocoon:
        struct Cocoon
            Handle Struct
            Handle CVC
        struct Letter
        struct Contact
        struct Citizen
    internal:
        struct Hidden
        struct Handle
        struct

letter -> Gate | Keeper -> Process | Handler -> Response


Suspicious Letter Status:
    1: SYSTEM_ERROR
    2: ...
    3: ...
    4: ...

Process Letter Status:
    1: Wating
    2: Keep
    3: Done
    4: Error

Sent Letter Status:
    1: Sent
    2: Later

Cocoon
|- Mailily < Cocoon
|- Taska < Cocoon


Rule:
    (letter) -> Send Response (set status = 2) | Send Request (set status = 1) | auto
    After creae or update recored point need to update some field to suitable with context

Ability:
    Citizens Id, Gate,....
    Contact CitizenId, Permission,....

No Taistra:
- Peer to Peer
- Freedom

cocoon.db | safety.dat | 

safety: use to safe error data when it cannot connect database


DB store:
    Unprocessed Letter
    Processed Letter
    Sent Letter


Unprocessed Letter:
    status:
        system_error
        letter_error
Processed Letter:
    status:
        waiting
        keep
        pending
        done
        system_error
        message_error
Sent Letter:
    status:
        sent
        system_error

Safty Mechanisms
