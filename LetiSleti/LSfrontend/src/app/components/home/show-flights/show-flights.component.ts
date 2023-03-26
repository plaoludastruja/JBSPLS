import { Component, OnInit } from '@angular/core';
import { IFlight } from 'src/app/shared/material/Flight';
import { FlightService } from 'src/app/shared/services/flight.service';
import { TicketDTO } from 'src/app/shared/model/DTO/ticketDTO';
import { TicketService } from 'src/app/shared/services/ticket.service';
import { User } from 'src/app/shared/model/user';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-show-flights',
  templateUrl: './show-flights.component.html',
  styleUrls: ['./show-flights.component.css']
})
export class ShowFlightsComponent implements OnInit{

  flights: IFlight[] = []
  displayedColumns: string[] = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'remaining', 'buy'];
  user: User  = new User 

  constructor(
    public flightService: FlightService, public ticketService: TicketService, public userService: UserService) {}

  ngOnInit(): void {
    this.flightService.getAll().subscribe(data => this.flights=data);
    
  }

createTicket(flight : IFlight){
    this.userService.getByEmail(this.userService.decodeToken()?.email).subscribe(res => {
      this.user = res
      let ticket: TicketDTO = {
        Start: flight.start,
        StartPlace: flight.startPlace,
        End: flight.end,
        EndPlace: flight.endPlace,
        Price: flight.pricePerPlace,
        FirstName: this.user.firstName,
        LastName: this.user.lastName,
        Email: this.user.email
      }
      console.log(ticket)
      this.ticketService.create(ticket).subscribe(res => {
        alert("Ticket is successfully created.")
        this.flightService.changePlacesLeft(flight.id).subscribe(res => {
          window.location.reload()
        })
      });
    })
  }

splitDate(date: string){
    var forShowing = '';
    var splittedOnT = date.split("T");
    console.log(splittedOnT);
    console.log(splittedOnT[0]);
    var splittedDate = splittedOnT[0].split('-');
    console.log(splittedDate);
    console.log(splittedDate[0]);
    forShowing= forShowing + splittedDate[2] + '.' + splittedDate[1] + '.' + splittedDate[0] +'.';
    var splittedTime = splittedOnT[1].split(':');
    forShowing.concat(" ");
    forShowing= forShowing + splittedTime[0] + ":" + splittedTime[1];
    return forShowing;
  }
}