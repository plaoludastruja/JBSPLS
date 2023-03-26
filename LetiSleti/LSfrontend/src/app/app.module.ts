import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './shared/material/material.module';
import { FooterComponent } from './components/home/footer/footer.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { LoginComponent } from './components/home/login/login.component';
import { RegisterComponent } from './components/home/register/register.component';
import { HomepageComponent } from './components/home/homepage/homepage.component';
import { HomepageMenuComponent } from './components/home/homepage-menu/homepage-menu.component';
import { CorouselComponent } from './components/home/corousel/corousel.component';
import { AdminModuleModule } from './components/admin-module/admin-module.module';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { ShowFlightsComponent } from './components/home/show-flights/show-flights.component';
import { MatTableModule } from '@angular/material/table';
import { TicketsInfoComponent } from './components/customer/tickets-info/tickets-info.component';
import { AuthInterceptor } from 'auth.interceptor';

@NgModule({
  declarations: [
    AppComponent,
    FooterComponent,
    LoginComponent,
    RegisterComponent,
    HomepageComponent,
    HomepageMenuComponent,
    CorouselComponent,
    ShowFlightsComponent,
    TicketsInfoComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MaterialModule,
    NgbModule,
    AdminModuleModule,
    HttpClientModule,
    FormsModule,
    MatTableModule 
  ],
  providers: [
    {
      provide : HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi   : true,
    }
  ],
  bootstrap: [AppComponent],
  exports: [HomepageMenuComponent]
})
export class AppModule { }
