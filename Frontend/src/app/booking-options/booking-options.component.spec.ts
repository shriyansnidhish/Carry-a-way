import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BookingOptionsComponent } from './booking-options.component';

describe('BookingOptionsComponent', () => {
  let component: BookingOptionsComponent;
  let fixture: ComponentFixture<BookingOptionsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BookingOptionsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BookingOptionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
