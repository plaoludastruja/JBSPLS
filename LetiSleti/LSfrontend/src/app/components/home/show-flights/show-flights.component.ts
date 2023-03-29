import { Component, OnInit } from '@angular/core';
import { IFlight } from 'src/app/shared/model/Flight';
import { FlightService } from 'src/app/shared/services/flight.service';
import { TicketDTO } from 'src/app/shared/model/DTO/ticketDTO';
import { TicketService } from 'src/app/shared/services/ticket.service';
import { User } from 'src/app/shared/model/user';
import { AuthService } from 'src/app/shared/services/auth.service';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-show-flights',
  templateUrl: './show-flights.component.html',
  styleUrls: ['./show-flights.component.css']
})
export class ShowFlightsComponent implements OnInit{

  flights: IFlight[] = []
  displayedColumns: string[] = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'remaining'];
  user: User  = new User 
  isSearchDisplayed: boolean = true;
  count: number = 1;

  constructor(
    public flightService: FlightService, public ticketService: TicketService, public authService: AuthService, public userService: UserService) {}

  ngOnInit(): void {
    this.flightService.getAll().subscribe(data => this.flights=data);
    let userRole = this.authService.decodeToken()?.role
    if(userRole === 'USERROLE')
      this.displayedColumns =  ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'remaining', 'buy']
    else if(userRole === 'ADMINROLE') {
      this.isSearchDisplayed = false
      this.displayedColumns =  ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'remaining', 'delete']
    }
  }

  search(startPlace: string, endPlace: string, numberOfPlaces: string, date:string){
    this.displayedColumns = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'remaining','totalPrice', 'buy', 'delete'];
    this.count = parseInt(numberOfPlaces);
    var numberOfPlacesConv = parseInt(numberOfPlaces);
    var dateConverted: Date = new Date(); 
    var searchCriteria;
    if(date != ""){
      var d = new Date(date);
      dateConverted = new Date(d.getFullYear(), d.getMonth(), d.getDate(), 0, 0, 0, 0)
      searchCriteria = {startPlace:startPlace,
        endPlace:endPlace,
        numberOfPlaces: numberOfPlacesConv,
        date:d.toISOString()}
    } else{
      searchCriteria = {startPlace:startPlace,
        endPlace:endPlace,
        numberOfPlaces: numberOfPlacesConv}
    }
    
    this.flightService.search(searchCriteria).subscribe((data) => {
      this.flights = data;
      console.log('view:', this.flights);
    });
  }

  deleteFlight(flight: IFlight){
    //const deleteId: DeleteDto = {id:id}
    this.flightService.delete(flight.id).subscribe(() => {
      console.log('super');
      this.flightService.getAll().subscribe(data => this.flights=data);
    });
  }

createTicket(flight : IFlight){
  if(flight.remaining > 0) {
    this.userService.getByEmail(this.authService.decodeToken()?.email).subscribe(res => {
      this.user = res
      let ticket: TicketDTO = {
        Start: flight.start,
        StartPlace: flight.startPlace,
        End: flight.end,
        EndPlace: flight.endPlace,
        Price: flight.pricePerPlace,
        FirstName: this.user.firstName,
        LastName: this.user.lastName,
        Email: this.user.email,
        Count: this.count
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
  else alert("All tickets are sold out!") 
    
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