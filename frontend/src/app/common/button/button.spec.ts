import { Component } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { MatTooltip } from '@angular/material/tooltip';

import { Button } from './button';

@Component({
  imports: [Button],
  template: '<app-button><span class="projected-content">Save</span></app-button>',
})
class ButtonTestHost {}

describe('Button', () => {
  let component: Button;
  let fixture: ComponentFixture<Button>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Button],
    }).compileComponents();

    fixture = TestBed.createComponent(Button);
    component = fixture.componentInstance;
    fixture.detectChanges();
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should project content into the button', () => {
    const hostFixture = TestBed.createComponent(ButtonTestHost);
    hostFixture.detectChanges();

    const projectedContent: HTMLElement = hostFixture.nativeElement.querySelector(
      'button .projected-content',
    );

    expect(projectedContent.textContent).toBe('Save');
  });

  it('should use the filled button color tokens', async () => {
    fixture.componentRef.setInput('color', 'danger');
    fixture.detectChanges();
    await fixture.whenStable();

    const button: HTMLButtonElement = fixture.nativeElement.querySelector('button');

    expect(button.style.getPropertyValue('--mat-button-filled-container-color')).toBe('#dc2626');
    expect(button.style.getPropertyValue('--mat-button-filled-label-text-color')).toBe('#ffffff');
    expect(button.style.getPropertyValue('--mat-button-filled-state-layer-color')).toBe('#ffffff');
    expect(button.style.getPropertyValue('--mat-button-filled-ripple-color')).toBe('#ffffff1f');
  });

  it('should use the outlined button color tokens when variant is outlined', async () => {
    fixture.componentRef.setInput('variant', 'outlined');
    fixture.componentRef.setInput('color', 'danger');
    fixture.detectChanges();
    await fixture.whenStable();

    const button: HTMLButtonElement = fixture.nativeElement.querySelector('button');

    expect(button.classList.contains('mat-mdc-outlined-button')).toBe(true);
    expect(button.style.getPropertyValue('--mat-button-outlined-label-text-color')).toBe('#dc2626');
    expect(button.style.getPropertyValue('--mat-button-outlined-outline-color')).toBe('#dc2626');
    expect(button.style.getPropertyValue('--mat-button-outlined-state-layer-color')).toBe('#dc2626');
    expect(button.style.getPropertyValue('--mat-button-outlined-ripple-color')).toBe('#dc26261f');
  });

  it('should use the tonal variant with an outlined shape and tinted background', async () => {
    fixture.componentRef.setInput('variant', 'tonal');
    fixture.componentRef.setInput('color', 'success');
    fixture.detectChanges();
    await fixture.whenStable();

    const button: HTMLButtonElement = fixture.nativeElement.querySelector('button');

    expect(button.classList.contains('mat-mdc-outlined-button')).toBe(true);
    expect(button.classList.contains('tt-button--tonal')).toBe(true);
    expect(button.style.getPropertyValue('--tt-button-tonal-bg')).toBe('#16a34a14');
    expect(button.style.getPropertyValue('--mat-button-outlined-label-text-color')).toBe('#16a34a');
    expect(button.style.getPropertyValue('--mat-button-outlined-outline-color')).toBe('#16a34a');
  });

  it('should use the text variant tokens and underline press class', async () => {
    fixture.componentRef.setInput('variant', 'text');
    fixture.componentRef.setInput('color', 'info');
    fixture.componentRef.setInput('afterPressEffect', 'underline');
    fixture.detectChanges();
    await fixture.whenStable();

    const button: HTMLButtonElement = fixture.nativeElement.querySelector('button');

    expect(button.classList.contains('mat-mdc-button')).toBe(true);
    expect(button.classList.contains('tt-button--text-effect-underline')).toBe(true);
    expect(button.style.getPropertyValue('--mat-button-text-label-text-color')).toBe('#0891b2');
    expect(button.style.getPropertyValue('--mat-button-text-state-layer-color')).toBe('#60a5fa');
    expect(button.style.getPropertyValue('--mat-button-text-ripple-color')).toBe('#60a5fa1f');
  });

  it('should keep the tooltip disabled when tooltipTitle is not provided', () => {
    const tooltip = fixture.debugElement.query(By.css('button')).injector.get(MatTooltip);

    expect(tooltip.message).toBe('');
    expect(tooltip.disabled).toBe(true);
  });

  it('should enable the tooltip when tooltipTitle is provided', async () => {
    fixture.componentRef.setInput('tooltipTitle', 'More details');
    fixture.detectChanges();
    await fixture.whenStable();

    const tooltip = fixture.debugElement.query(By.css('button')).injector.get(MatTooltip);

    expect(tooltip.message).toBe('More details');
    expect(tooltip.disabled).toBe(false);
  });
});
