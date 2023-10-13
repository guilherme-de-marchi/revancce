import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

import { NavigationMenuComponent } from '../navigation-menu/navigation-menu.component'
import { FooterComponent } from '../footer/footer.component'
import { BatchesComponent } from '../batches/batches.component'

@Component({
  selector: 'app-event-page',
  standalone: true,
  imports: [CommonModule, NavigationMenuComponent, FooterComponent, BatchesComponent],
  templateUrl: './event-page.component.html',
  styleUrls: ['./event-page.component.scss']
})
export class EventPageComponent {

}
