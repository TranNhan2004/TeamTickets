import { Component, signal } from '@angular/core';
import { Button } from "./common/button/button";
import { Chip } from "./common/chip/chip";

@Component({
  selector: 'app-root',
  imports: [Button, Chip],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App {
  protected readonly title = signal('Team Tickets');
}
