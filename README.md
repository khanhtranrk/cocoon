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
