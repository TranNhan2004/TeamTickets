import { Component } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { MatTooltip } from '@angular/material/tooltip';

import { Chip } from './chip';

@Component({
  imports: [Chip],
  template: '<app-chip><span class="projected-content">Active</span></app-chip>',
})
class ChipTestHost {}

describe('Chip', () => {
  let component: Chip;
  let fixture: ComponentFixture<Chip>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Chip],
    }).compileComponents();

    fixture = TestBed.createComponent(Chip);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should project content into the chip', () => {
    const hostFixture = TestBed.createComponent(ChipTestHost);
    hostFixture.detectChanges();

    const projectedContent: HTMLElement = hostFixture.nativeElement.querySelector(
      'mat-chip .projected-content',
    );

    expect(projectedContent.textContent).toBe('Active');
  });

  it('should use the filled variant styles', async () => {
    fixture.componentRef.setInput('color', 'danger');
    fixture.detectChanges();
    await fixture.whenStable();

    const chip: HTMLElement = fixture.nativeElement.querySelector('mat-chip');

    expect(chip.style.getPropertyValue('--tt-chip-color')).toBe('#dc2626');
    expect(chip.style.getPropertyValue('--tt-chip-label-color')).toBe('#ffffff');
    expect(chip.classList.contains('tt-chip--filled')).toBe(true);
  });

  it('should use the tonal variant styles', async () => {
    fixture.componentRef.setInput('variant', 'tonal');
    fixture.componentRef.setInput('color', 'success');
    fixture.detectChanges();
    await fixture.whenStable();

    const chip: HTMLElement = fixture.nativeElement.querySelector('mat-chip');

    expect(chip.classList.contains('tt-chip--tonal')).toBe(true);
    expect(chip.style.getPropertyValue('--tt-chip-tonal-bg')).toBe('#16a34a14');
  });

  it('should keep ripple disabled to avoid press behavior', () => {
    const chip = fixture.debugElement.query(By.css('mat-chip')).componentInstance as {
      disableRipple: boolean;
    };

    expect(chip.disableRipple).toBe(true);
  });

  it('should keep the tooltip disabled when tooltipTitle is not provided', () => {
    const tooltip = fixture.debugElement.query(By.css('mat-chip')).injector.get(MatTooltip);

    expect(tooltip.message).toBe('');
    expect(tooltip.disabled).toBe(true);
  });

  it('should enable the tooltip when tooltipTitle is provided', async () => {
    fixture.componentRef.setInput('tooltipTitle', 'More details');
    fixture.detectChanges();
    await fixture.whenStable();

    const tooltip = fixture.debugElement.query(By.css('mat-chip')).injector.get(MatTooltip);

    expect(tooltip.message).toBe('More details');
    expect(tooltip.disabled).toBe(false);
  });
});
