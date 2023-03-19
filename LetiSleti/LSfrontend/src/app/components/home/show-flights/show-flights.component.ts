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
  displayedColumns: string[] = ['start', 'startPlace', 'end', 'endPlace', 'pricePerPlace', 'freePlaces'];

  constructor(
    public flightService: FlightService) {}

  ngOnInit(): void {
    this.flightService.getAll().subscribe(data => this.flights=data);
  }

  search(startPlace: string){
    const searchCriteria: SearchDto = {startPlace:startPlace}
    this.flightService.search(searchCriteria).subscribe((data) => {
      this.flights = data;
      console.log('view:', this.flights);
    });
  }

}
