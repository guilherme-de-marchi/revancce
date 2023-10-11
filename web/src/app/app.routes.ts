import { Routes } from '@angular/router';
import { EventPageComponent } from './event-page/event-page.component'
import { HomeComponent } from './home/home.component';

export const routes: Routes = [
    {path: '', component: HomeComponent},
    {path: 'evento/vibracoes-urbanas', component: EventPageComponent}
];
