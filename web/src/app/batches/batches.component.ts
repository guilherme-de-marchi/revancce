import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BatchComponent } from '../batch/batch.component';

import * as globals from '../globals';

@Component({
  selector: 'app-batches',
  standalone: true,
  imports: [CommonModule, BatchComponent],
  templateUrl: './batches.component.html',
  styleUrls: ['./batches.component.scss']
})
export class BatchesComponent implements OnInit {
  batches: BatchComponent[] = []

  async ngOnInit(): Promise<void> {
      const bresponse = await fetch(
        `${globals.apiAddress}/api/v1/event/batch`,
        {method: "GET", mode: "cors"},
      )
      const bdata = await bresponse.json()

      for (var i = 0; i < bdata.length; i++) {
        let lt = new Date(bdata[i].limit_time)

        let b = new BatchComponent
        b.id = bdata[i].id
        b.ticket = bdata[i].ticket
        b.rawLimitTime = lt
        b.limitTime = `${lt.getFullYear()}/${lt.getMonth()}/${lt.getDay()} às ${lt.getHours()}:${lt.getMinutes()}`
        b.opened = bdata[i].opened
        b.price = bdata[i].price

        const tresponse = await fetch(
          `${globals.apiAddress}/api/v1/event/ticket?id=${b.ticket}`,
          {method: "GET", mode: "cors"},
        )
        const tdata = await tresponse.json()
        if (tdata.length == 0) {
          continue 
        }
        console.log(tdata)
    
        b.ticketName = tdata[0].name

        this.batches.push(b)
      }
  }
}
