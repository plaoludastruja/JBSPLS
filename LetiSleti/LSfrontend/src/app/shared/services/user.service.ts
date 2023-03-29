import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import jwtDecode from 'jwt-decode';
import { ToastrService } from 'ngx-toastr';
import { Observable } from 'rxjs';
import { Jwt } from '../model/jwt';
import { User } from '../model/user';

@Injectable({
    providedIn: 'root'
})
export class UserService {

    apiHost: string = 'http://localhost:8080/';
    headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json' });
    jwtHelper: any;

    constructor(private http: HttpClient, private router: Router, private toastr: ToastrService) { }

    getAllUsers(): Observable<User[]> {
        return this.http.get<User[]>(this.apiHost + 'user/getAll', { headers: this.headers });
    }

    registerUser(user: any): Observable<any> {
        return this.http.post<any>(this.apiHost + 'user/register', user);
    }

    getByEmail(email: any): Observable<any> {
        return this.http.get<any>(this.apiHost + 'user/by-email/' + email);
    }

    public login = (login: any): void => {
        this.http.post<any>(this.apiHost + 'user/login', login).subscribe({
            next: (res) => {
                localStorage.setItem('token', res);
                this.toastr.success('Login successfull', '', { closeButton: true, timeOut : 1500  });
                this.router.navigate(['home']);
            },
            //TODO: handle errors
            error: err => {
                this.toastr.error('Wrong username or password', '', { closeButton: true, timeOut : 1500  });
            }
        }
        );
    }

    public logout = () => {
        localStorage.removeItem('token');
        this.toastr.success('Logout successfull', '', { closeButton: true, timeOut : 1500  });
        this.router.navigate(['login']);
    }
}
