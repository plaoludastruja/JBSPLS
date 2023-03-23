import { Component, OnInit } from '@angular/core';
import { NgbDateStruct, NgbCalendar, NgbDatepickerModule } from '@ng-bootstrap/ng-bootstrap';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { IFlight } from 'src/app/shared/material/Flight';
import { FlightService } from 'src/app/shared/services/flight.service';

@Component({
  selector: 'app-create-flight',
  templateUrl: './create-flight.component.html',
  styleUrls: ['./create-flight.component.css']
})
export class CreateFlightComponent implements OnInit {

  public dateValue: Date = new Date("06/06/2022 15:30");
  public form: FormGroup;
  public errorMessage: String = "";

  constructor(
    fb: FormBuilder,
    public flightService: FlightService) {
      this.form = fb.group({
        startDateTime: ['', [Validators.required]],
        startPlace: ['', [Validators.required, Validators.pattern("^[A-Z].*$")]],
        endDateTime: ['', [Validators.required]],
        endPlace: ['', [Validators.required, Validators.pattern("^[A-Z].*$")]],
        maxNumberOfPlaces: ['', [Validators.required]],
        pricePerPlace: ['', [Validators.required]]
      });
     }
     
  get startDateTime(){
    return this.form.get('startDateTime');
  }
  get startPlace(){
    return this.form.get('startPlace');
  }
  get endDateTime(){
    return this.form.get('endDateTime');
  }
  get endPlace(){
    return this.form.get('endPlace');
  }
  get maxNumberOfPlaces(){
    return this.form.get('maxNumberOfPlaces');
  }
  get pricePerPlace(){
    return this.form.get('pricePerPlace');
  }

  ngOnInit(): void {
  }

  register(){
    
    var startConverted = new Date(this.startDateTime?.value);
    var newStart = new Date(startConverted.getFullYear(), startConverted.getMonth(), startConverted.getDate(), startConverted.getHours() + 2,
      startConverted.getMinutes(), startConverted.getSeconds(), startConverted.getMilliseconds()); 
    //startConverted.setHours(startConverted.getHours() + 2);
    var endConverted = new Date(this.endDateTime?.value);
    var newEnd = new Date(endConverted.getFullYear(), endConverted.getMonth(), endConverted.getDate(), endConverted.getHours() + 2,
      endConverted.getMinutes(), endConverted.getSeconds(), endConverted.getMilliseconds()); 

    let flight: IFlight = {
      id: '',
      start: newStart.toISOString(), //new Date(this.startDateTime?.value).toISOString(),
      startPlace: String(this.startPlace?.value),
      end: newEnd.toISOString(), //new Date(this.endDateTime?.value).toISOString(),
      endPlace: String(this.endPlace?.value),
      maxNumberOfPlaces: this.maxNumberOfPlaces?.value,
      pricePerPlace: this.pricePerPlace?.value
    }
    console.log(flight)
    if(this.form.valid && this.compareDates(this.startDateTime, this.endDateTime)){
      this.errorMessage=""
      this.flightService.register(flight).subscribe();
    }else{
      if(this.errorMessage=="") this.errorMessage="All fields must be entered"
    }
  }

  compareDates(start: any, end: any): boolean{
    const startDate = new Date(start.value).valueOf();
    const endDate = new Date(end.value).valueOf();
    const now = new Date().valueOf();
    if(startDate <= now){
      this.errorMessage="Start date must be in future";
      return false;
    }
    if(endDate <= now){
      this.errorMessage="End date must be in future";
      return false;
    }
    if(startDate - endDate > 0){
      this.errorMessage="End date must be after start date"
      return false;
    }
    return true;
  }

}
