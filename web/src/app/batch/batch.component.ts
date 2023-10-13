import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';

import * as globals from '../globals';

@Component({
  selector: 'app-batch',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './batch.component.html',
  styleUrls: ['./batch.component.scss']
})
export class BatchComponent {
  hidden = true
  buttonClass = "batch-button-closed"
  arrowClass = "arrow-closed"
  infoClass = "batch-info-closed"
  
  @Input() id = ""
  @Input() ticket = ""
  @Input() ticketName = ""
  @Input() rawLimitTime = new Date
  @Input() limitTime = ""
  @Input() opened = false
  @Input() price = 0


  async purchase() {
    if (!this.opened || this.rawLimitTime.getTime() >= Date.now()) {
      return
    }

    const response = await fetch(
      `${globals.apiAddress}/api/v1/client/ticket/purchase`,
      {
        method: "POST", 
        mode: "cors", 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          batch: this.id,
        }),
      },
    )
    const data = await response.json()

    window.open(data.PaymentLinkURL)
  }

  toggleBatchInfo() {
    this.hidden = !this.hidden
    this.buttonClass = this.hidden ? "batch-button-closed" : "batch-button-opened"
    this.arrowClass = this.hidden ? "arrow-closed" : "arrow-opened"
    this.infoClass = this.opened ? "batch-info-opened" : "batch-info-closed"
  }
}
