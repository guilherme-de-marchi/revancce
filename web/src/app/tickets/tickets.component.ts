import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import {NgPipesModule} from 'ngx-pipes'
import { BatchComponent } from '../batch/batch.component';

import * as globals from '../globals';

@Component({
  selector: 'app-tickets',
  standalone: true,
  imports: [CommonModule, BatchComponent, NgPipesModule],
  templateUrl: './tickets.component.html',
  styleUrls: ['./tickets.component.scss']
})
export class TicketsComponent implements OnInit {
  tickets!: any[]

  async ngOnInit(): Promise<void> {
    const resp = await fetch(
      `${globals.apiAddress}/api/v1/event/ticket`,
      {method: "GET", mode: "cors"},
    )
    this.tickets = await resp.json()

    for (let i=0; i < this.tickets.length; i++) {
      const resp = await fetch(
        `${globals.apiAddress}/api/v1/event/batch?ticket=${this.tickets[i].id}`,
        {method: "GET", mode: "cors"},
      )
      this.tickets[i].batches = await resp.json()
    
      this.tickets[i].hidden = true
      this.tickets[i].buttonClass = "batch-button-closed"
      this.tickets[i].arrowClass = "arrow-closed"
    }
  }

  toggleBatchInfo(ticket: any) {
    ticket.hidden = !ticket.hidden
    ticket.buttonClass = ticket.hidden ? "batch-button-closed" : "batch-button-opened"
    ticket.arrowClass = ticket.hidden ? "arrow-closed" : "arrow-opened"
  }
}
