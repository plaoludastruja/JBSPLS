import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateFlightComponent } from './components/home/create-flight/create-flight.component';
import { TicketsInfoComponent } from './components/home/tickets-info/tickets-info.component';
import { HomepageComponent } from './components/home/homepage/homepage.component';
import { LoginComponent } from './components/home/login/login.component';
import { RegisterComponent } from './components/home/register/register.component';

const routes: Routes = [
  { path: 'home', component: HomepageComponent},
  { path: 'login', component: LoginComponent},
  { path: 'register', component: RegisterComponent},
  { path: 'tickets-info', component: TicketsInfoComponent},
  { path: 'create-flight', component: CreateFlightComponent},

  // ova linija mora biti zadnja, biti zadnja
  { path: '', redirectTo: '/home', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
