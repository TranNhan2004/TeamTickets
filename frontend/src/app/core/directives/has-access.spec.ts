import { Component, signal } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HasAccess } from './has-access';

@Component({
  imports: [HasAccess],
  template: `<span *appHasAccess="hasAccess()">Visible content</span>`,
})
class TestHost {
  hasAccess = signal(false);
}

describe('HasAccess', () => {
  let fixture: ComponentFixture<TestHost>;
  let host: TestHost;

  beforeEach(async () => {
    await TestBed.configureTestingModule({ imports: [TestHost] }).compileComponents();

    fixture = TestBed.createComponent(TestHost);
    host = fixture.componentInstance;
  });

  it('does not render content when access is false', () => {
    fixture.detectChanges();

    expect(fixture.nativeElement.textContent).not.toContain('Visible content');
  });

  it('renders content when access is true', () => {
    host.hasAccess.set(true);
    fixture.detectChanges();

    expect(fixture.nativeElement.textContent).toContain('Visible content');
  });

  it('removes rendered content when access changes to false', () => {
    host.hasAccess.set(true);
    fixture.detectChanges();
    host.hasAccess.set(false);
    fixture.detectChanges();

    expect(fixture.nativeElement.textContent).not.toContain('Visible content');
  });
});
