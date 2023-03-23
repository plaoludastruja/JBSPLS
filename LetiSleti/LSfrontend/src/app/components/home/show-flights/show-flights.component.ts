import { DeleteDto } from './../../../shared/model/DTO/deleteDTO';
import { SearchDto } from './../../../shared/model/DTO/searchDTO';
import { Component, OnInit } from '@angular/core';
import { IFlight } from 'src/app/shared/material/Flight';
import { FlightService } from 'src/app/shared/services/flight.service';

@Component({
  selector: 'app-show-flights',
  templateUrl: './show-flights.component.html',
  styleUrls: ['./show-flights.component.css']
})
export class ShowFlightsComponent implements OnInit{

  flights: IFlight[] = []
  displayedColumns: string[] = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'freePlaces', 'delete'];

  constructor(
    public flightService: FlightService) {}

  ngOnInit(): void {
    this.flightService.getAll().subscribe(data => this.flights=data);
  }

  search(startPlace: string, endPlace: string){
    const searchCriteria: SearchDto = {startPlace:startPlace,
                                       endPlace:endPlace}
    this.flightService.search(searchCriteria).subscribe((data) => {
      this.flights = data;
      console.log('view:', this.flights);
    });
  }

  deleteFlight(flight: IFlight){
    //const deleteId: DeleteDto = {id:id}
    this.flightService.delete(flight).subscribe(() => {
      console.log('super');
    });
  }

}
