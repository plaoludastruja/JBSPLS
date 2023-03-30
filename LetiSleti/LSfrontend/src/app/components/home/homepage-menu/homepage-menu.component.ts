import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/app/shared/services/auth.service';
import { UserService } from 'src/app/shared/services/user.service';

@Component({
  selector: 'app-homepage-menu',
  templateUrl: './homepage-menu.component.html',
  styleUrls: ['./homepage-menu.component.css']
})
export class HomepageMenuComponent implements OnInit {

  userRole: String='';
  constructor(public authService: AuthService, private userService: UserService) { }

  ngOnInit(): void {
    this.userRole=String(this.authService.decodeToken()?.role);
    if (this.userRole === 'undefined') {
      this.userRole = ''
    }
    console.log(this.userRole);
  }

  logout() {
    this.userService.logout()
  }

}
