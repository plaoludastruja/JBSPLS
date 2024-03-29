import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TicketService {

  apiHost: string = 'http://localhost:8082/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  create(ticket: any): Observable<any> {
    return this.http.post<any>(this.apiHost + 'ticket', ticket, {headers: this.headers});
  }

  getAll(): Observable<any> {
    return this.http.get<any>(this.apiHost + 'user/ticket/getAll', {headers: this.headers});
  }

}