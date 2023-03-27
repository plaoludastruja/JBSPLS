import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import jwtDecode from 'jwt-decode';
import { Jwt } from '../model/jwt';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });

  constructor(private http: HttpClient) { }

  decodeToken(): Jwt | undefined{
    const token = localStorage.getItem('token')
    if(token) 
      return jwtDecode<Jwt>(token)
    return undefined
  }
}