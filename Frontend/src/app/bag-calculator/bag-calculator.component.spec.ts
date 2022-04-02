import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { BagCalculatorComponent } from './bag-calculator.component';

describe('BagCalculatorComponent', () => {
  let component: BagCalculatorComponent;
  let fixture: ComponentFixture<BagCalculatorComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ BagCalculatorComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BagCalculatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
