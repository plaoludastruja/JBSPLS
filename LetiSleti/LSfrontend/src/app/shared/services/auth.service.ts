import { Injectable } from '@angular/core';
import jwtDecode from 'jwt-decode';
import { Jwt } from '../model/jwt';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private jwtHelper: JwtHelperService) { }

  decodeToken(): Jwt | undefined{
    const token = localStorage.getItem('token')
    if(token) 
      return jwtDecode<Jwt>(token)
    return undefined
  }

  public isAuthenticated(): boolean {
    const token = localStorage.getItem('token');
    if(token)
    {
      return !this.jwtHelper.isTokenExpired(token!);
    }
    return false;
  }
}