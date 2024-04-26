package cocoon

// TODO: logic for err case
func (c *Cocoon) SuspiciousLetterChanHandler() {
  for lt := range c.SuspiciousLetterChan {
    _, err := c.LetterService.CreateSuspiciousLetter(lt)
    if err != nil {
      // save to safety, current to pass it
    }
  }
}

// TODO: logic for err case
func (c *Cocoon) ProcessRequestLetterChanHandler() {
  for lt := range c.ProcessRequestLetterChan {
    ss, msg, err := c.MessageHandler(lt.Message)

    if err != nil {
      lt.Message = append([]byte{0}, []byte(err.Error())...)
      c.SendRequestLetterChan <-lt
      return
    }

    lt.Status = ss
    _, err = c.LetterService.UpdateProcessedLetter(lt)
    if err != nil {
      // save to must be true
    }

    lt.Message = msg
    c.SendRequestLetterChan <-lt
  }
}

// TODO: logic
func (c *Cocoon) ProcessResponseLetterChanHandler() {
  for lt := range c.ProcessRequestLetterChan {
    rq, err := c.LetterService.GetSentLetterByCodeAndCommitTimeAndForeignId(lt.Code, lt.CommitTime, lt.ForeignId)
    if err != nil {
      lt.Status = SystemErrorStatus
      _, err = c.LetterService.UpdateProcessedLetter(lt)
      if err != nil {
        // save task to safety
      }
      return
    }

    if rq == nil {
      lt.Status = LetterErrorStatus
      _, err := c.LetterService.UpdateProcessedLetter(lt)
      if err != nil {
        // save task to safety
      }
    }

    lt.Status = DoneStatus

    // ...
    // find refer
    // if exist ---> check done all ----> set done | keep ---> send response
    // else keep
  }
}

// TODO: logic
func (c *Cocoon) SendResponseLetterChanHandler() {
  for lt := range c.SendResponseLetterChan {
    err := c.SendLetter(lt)
    lt.Status = SentStatus
    if err != nil {
      lt.Status = SystemErrorStatus
    }

    _, err = c.LetterService.CreateSentLetter(lt)

    if err != nil {
      // safety
    }
  }
}

// TODO: logic
func (c *Cocoon) SendRequestLetterChanHandler() {
  for lt := range c.SendRequestLetterChan {
    err := c.SendLetter(lt)
    lt.Status = SentStatus
    if err != nil {
      lt.Status = SystemErrorStatus
    }

    _, err = c.LetterService.CreateSentLetter(lt)

    if err != nil {
      // safety
    }
  }
}
