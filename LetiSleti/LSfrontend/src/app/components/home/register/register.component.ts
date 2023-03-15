import { Component, OnInit } from '@angular/core';
import { User } from 'src/app/shared/model/user';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  user: User = new User;
  constructor(private userService:UserService) { }

  ngOnInit(): void {
    this.getAllUsers()
  }

  onRegister():void{
    this.user.role = "USERROLE"
    this.userService.registerUser(this.user).subscribe(res => {
      console.log(res)
    });
  }

  getAllUsers():void{
    this.userService.getAllUsers().subscribe(res => {
      console.log(res)
    });
  }

}
