import { Component, OnInit } from '@angular/core';
import { LoginDTO } from 'src/app/shared/model/DTO/loginDTO';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginDto: LoginDTO = new LoginDTO;
  constructor(private userService:UserService) { }

  ngOnInit(): void {
  }

  onLogin():void{
    this.userService.login(this.loginDto)
  }
}
