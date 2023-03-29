import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/shared/model/user';
import { UserService } from 'src/app/shared/services/user.service';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

    user: User = new User;
    errorMessage: string = ''
    constructor(private userService: UserService, private toastr: ToastrService, private router: Router) { }

    ngOnInit(): void {
        this.getAllUsers()
    }

    onRegister(): void {
        if(!this.CheckFields()){
            this.errorMessage ="All fields must be entered"
        }else{
            this.user.role = "USERROLE"
            this.userService.registerUser(this.user).subscribe({
                next: (res) => {
                    this.toastr.success('User successfully created', '', { closeButton: true, timeOut : 1500  });
                    this.router.navigate(['login'])
                },
                error: () => {
                    this.toastr.error('Username taken', '', { closeButton: true, timeOut : 1500  });
                }
            });
        }
        
    }

    getAllUsers(): void {
        this.userService.getAllUsers().subscribe(res => {
            console.log(res)
        });
    }

    CheckFields(): boolean {
        if ( this.user.firstName == '' || this.user.lastName == '' || this.user.email == '' || this.user.password == '' ){
            return false
        }
        return true
    }
}
