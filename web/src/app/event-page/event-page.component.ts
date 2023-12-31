import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NavigationMenuComponent } from '../navigation-menu/navigation-menu.component'
import { FooterComponent } from '../footer/footer.component'
import { TicketsComponent } from '../tickets/tickets.component'
import { MapComponent } from '../map/map.component'

@Component({
  selector: 'app-event-page',
  standalone: true,
  imports: [CommonModule, NavigationMenuComponent, FooterComponent, TicketsComponent, MapComponent],
  templateUrl: './event-page.component.html',
  styleUrls: ['./event-page.component.scss']
})
export class EventPageComponent {

}
