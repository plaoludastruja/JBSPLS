import { IFlight } from 'src/app/shared/model/Flight';
import { DeleteDto } from './../model/DTO/deleteDTO';
import { SearchDto } from './../model/DTO/searchDTO';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FlightService {

  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  register(flight: any): Observable<any> {
    return this.http.post<any>(this.apiHost + 'admin/flight/register', flight);
  }

  getAll(): Observable<any> {
    return this.http.get<any>(this.apiHost + 'flight/getAll');
  }

  search(searchCriteria: any): Observable<any> {
    return this.http.post<any>(this.apiHost + 'flight/search', searchCriteria);
  }

  delete(flightId: string): any{
    return this.http.delete(this.apiHost + `admin/flight/${flightId}`);
  }
  changePlacesLeft(flightId: any): Observable<any> {
    return this.http.get<any>(this.apiHost + 'flight/change-places-left/' + flightId);
  }

}