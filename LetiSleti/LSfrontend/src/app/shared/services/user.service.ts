import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import jwtDecode from 'jwt-decode';
import { Observable } from 'rxjs';
import { Jwt } from '../model/jwt';
import { User } from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.apiHost + 'user/getAll', {headers: this.headers});
  }

  registerUser(user: any): Observable<any> {
    return this.http.post<any>(this.apiHost + 'user/register', user);
  }

  login(login:any): Observable<any> {
    return this.http.post<any>(this.apiHost + 'user/login', login);
  }

  getByEmail(email : any): Observable<any> {
    return this.http.get<any>(this.apiHost + 'user/by-email/' + email);
  }
}
