package cocoon

// TODO: logic for err case
func (c *Cocoon) SuspiciousLetterChanHandler() {
  for lt := range c.SuspiciousLetterChan {
    _, err := c.LetterService.CreateSuspiciousLetter(lt)
    if err != nil {
      // save to must be true
    }
  }
}

// TODO: logic for err case
func (c *Cocoon) ProcessRequestLetterChanHandler() {
  for lt := range c.ProcessRequestLetterChan {
    ss, msg, err := c.MessageHandler(lt.Message)

    if err != nil {
      lt.Message = append([]byte{0}, []byte(err.Error())...)
      c.SendResponseLetterChan <-lt
      return
    }

    lt.Status = ss
    _, err = c.LetterService.UpdateProcessedLetter(lt)
    if err != nil {
      // save to must be true
    }


    lt.Message = msg
    c.SendResponseLetterChan <-lt
  }
}

// TODO: logic
func (c *Cocoon) ProcessResponseLetterChanHandler() {
  for lt := range c.ProcessRequestLetterChan {
    // find letter in sent letter
    // set to done
    // find refer
    // if exist ---> check done all ----> set done | keep ---> send response
    // else keep
  }
}

// TODO: logic
func (c *Cocoon) SendResponseLetterChanHandler() {
  for lt := range c.SendResponseLetterChan {
    err := c.SendResponseLetter(lt)
    if err != nil {
      // update db
      return
    }

    // update
  }
}

// TODO: logic
func (c *Cocoon) SendRequestLetterChanHandler() {
  for lt := range c.SendRequestLetterChan {
    if lt.Type == 1 {
      err := c.SendRequestLetter(lt)
      if err != nil {
        // update db
      }
    }

    err := c.SendResponseLetter(lt)
    if err != nil {
      // update db
    }

  }
}
