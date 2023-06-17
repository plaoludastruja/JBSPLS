import { Component, OnInit } from '@angular/core';
import { TicketDTO } from 'src/app/shared/model/DTO/ticketDTO';
import { TicketService } from 'src/app/shared/services/ticket.service';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-tickets-info',
  templateUrl: './tickets-info.component.html',
  styleUrls: ['./tickets-info.component.css']
})
export class TicketsInfoComponent implements OnInit {

  tickets: TicketDTO[] = []
  displayedColumns: string[] = ['from', 'to', 'depart', 'return', 'price', 'count', 'total'];
  email: string=""
  isDurable: boolean = false;


  constructor(public ticketService: TicketService, public userService: UserService) { }

  ngOnInit(): void {
    this.ticketService.getAll().subscribe( data => {
      this.tickets = data
    })
  }

  onCheckBoxChange(event: any){
    this.isDurable = event.checked
  }

  generateApiKey(){
    this.userService.generateApiKey(this.email, this.isDurable).subscribe(res => {
      alert("You have generated API KEY successfully")
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
