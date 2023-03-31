import { Component, OnInit } from '@angular/core';
import { IFlight } from 'src/app/shared/model/Flight';
import { FlightService } from 'src/app/shared/services/flight.service';
import { TicketDTO } from 'src/app/shared/model/DTO/ticketDTO';
import { TicketService } from 'src/app/shared/services/ticket.service';
import { User } from 'src/app/shared/model/user';
import { AuthService } from 'src/app/shared/services/auth.service';
import { UserService } from 'src/app/shared/services/user.service';
import { ToastrService } from 'ngx-toastr';

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
    public flightService: FlightService, public ticketService: TicketService, public authService: AuthService, public userService: UserService, private toastr: ToastrService) {}

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
    if(startPlace == '' || endPlace == '' || numberOfPlaces == '' || date == ''){
      this.toastr.error('All criteria must be entered!', '', { closeButton: true, timeOut : 1500  });
      return;
    }
    this.displayedColumns = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace','totalPrice', 'remaining', 'buy'];
    this.count = parseInt(numberOfPlaces);
    var numberOfPlacesConv = parseInt(numberOfPlaces);
    var d = new Date(date);
    //var dateConverted = new Date(d.getFullYear(), d.getMonth(), d.getDate(), 0, 0, 0, 0)
    console.log("datum iz searcha")
    console.log(d)
    //d.setHours(0, 0, 0);
    var searchCriteria = {startPlace:startPlace,
        endPlace:endPlace,
        numberOfPlaces: numberOfPlacesConv,
        date:d.toISOString()}
    
    this.flightService.search(searchCriteria).subscribe((data) => {
      this.flights = data;
      console.log('view:', this.flights);
    });
  }

  deleteFlight(flight: IFlight){
    //const deleteId: DeleteDto = {id:id}
    this.flightService.delete(flight.id).subscribe(() => {
      console.log('super');
      this.toastr.success('Flight deleted successfully', '', { closeButton: true, timeOut : 1500  });
      this.flightService.getAll().subscribe(data => this.flights=data);
    });
  }

createTicket(flight : IFlight){
  if(flight.remaining >= this.count) {
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
        Count: this.count,
        FlightId: flight.id
      }
      console.log(ticket)
      this.ticketService.create(ticket).subscribe(res => {
        this.toastr.success('Ticket created successfully', '', { closeButton: true, timeOut : 1500  });
        this.flightService.changePlacesLeft(flight.id, this.count).subscribe(res => {
          this.flightService.getAll().subscribe(data => this.flights=data);
        })
      });
    })
  } 
  else this.toastr.error('All tickets are sold out!', '', { closeButton: true, timeOut : 1500  });
    
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