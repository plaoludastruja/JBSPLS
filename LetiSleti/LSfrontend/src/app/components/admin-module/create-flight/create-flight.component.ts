import { Component, OnInit } from '@angular/core';
import { NgbDateStruct, NgbCalendar, NgbDatepickerModule } from '@ng-bootstrap/ng-bootstrap';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-create-flight',
  templateUrl: './create-flight.component.html',
  styleUrls: ['./create-flight.component.css']
})
export class CreateFlightComponent implements OnInit {

  public dateValue: Date = new Date("06/06/2022 15:30");
  public form: FormGroup;

  constructor(
    fb: FormBuilder) {
      this.form = fb.group({
        startDateTime: ['', [Validators.required]],
        startPlace: ['', [Validators.required]],
        endDateTime: ['', [Validators.required]],
        endPlace: ['', [Validators.required]],
        maxNumberOfPlaces: ['', [Validators.required]],
        pricePerPlace: ['', [Validators.required]]
      });
     }

  ngOnInit(): void {
  }

}
