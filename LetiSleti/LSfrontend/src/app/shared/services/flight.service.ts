import { IFlight } from 'src/app/shared/material/Flight';
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
    return this.http.post<any>(this.apiHost + 'flight/register', flight);
  }

  getAll(): Observable<any> {
    return this.http.get<any>(this.apiHost + 'flight/getAll');
  }

  search(searchCriteria: SearchDto): Observable<any> {
    return this.http.post<any>(this.apiHost + 'flight/search', searchCriteria);
  }

  delete(flight: IFlight): any{
    return this.http.post(this.apiHost + 'flight/delete', flight);
  }
}