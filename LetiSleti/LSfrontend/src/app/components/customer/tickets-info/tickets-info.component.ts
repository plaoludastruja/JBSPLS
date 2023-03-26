import { Component, OnInit } from '@angular/core';
import { TicketDTO } from 'src/app/shared/model/DTO/ticketDTO';
import { TicketService } from 'src/app/shared/services/ticket.service';

@Component({
  selector: 'app-tickets-info',
  templateUrl: './tickets-info.component.html',
  styleUrls: ['./tickets-info.component.css']
})
export class TicketsInfoComponent implements OnInit {

  tickets: TicketDTO[] = []
  displayedColumns: string[] = ['from', 'to', 'depart', 'return', 'price'];

  constructor(public ticketService: TicketService) { }

  ngOnInit(): void {
    this.ticketService.getAll().subscribe( data => {
      this.tickets = data
    })
  }

}
