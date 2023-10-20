import { Component, Input, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

import * as globals from '../globals';

@Component({
  selector: 'app-batch',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './batch.component.html',
  styleUrls: ['./batch.component.scss']
})
export class BatchComponent implements OnInit {
  @Input() batch!: any
  @Input() ticket!: any

  infoClass = "batch-info-closed"
  rawLimitTime!: Date
  limitTime!: string

  async ngOnInit() {
    this.infoClass = this.batch.opened ? "batch-info-opened" : "batch-info-closed"
    this.rawLimitTime = new Date(this.batch.limit_time)
    this.limitTime = `${this.rawLimitTime.getFullYear()}/${this.rawLimitTime.getMonth()}/${this.rawLimitTime.getDay()} às ${this.rawLimitTime.getHours()}:${this.rawLimitTime.getMinutes()}`
  }

  async purchase() {
    if (!this.batch.opened || this.rawLimitTime.getTime() >= Date.now()) {
      return
    }

    // const response = await fetch(
    //   `${globals.apiAddress}/api/v1/client/ticket/purchase`,
    //   {
    //     method: "POST", 
    //     mode: "cors", 
    //     headers: {
    //       'Content-Type': 'application/json'
    //     },
    //     body: JSON.stringify({
    //       batch: this.batch.id,
    //     }),
    //   },
    // )
    // const data = await response.json()

    // window.open(data.PaymentLinkURL)

    console.log(this.ticket)

    const resp = await fetch(
      `${globals.apiAddress}/api/v1/event?id=${this.ticket.event}`,
      {method: "GET", mode: "cors"},
    )
    let events = await resp.json()

    if (events.length == 0) {
      return
    }

    let text = `Olá. Gostaria de comprar um ingresso para o evento ${events[0].name}, tipo '${this.ticket.name}' e lote ${this.batch.number}.` 
    window.open(`https://api.whatsapp.com/send?phone=5511940452258&text=${text}`)
  }
}
