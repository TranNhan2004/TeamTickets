import { NgStyle } from '@angular/common';
import { ChangeDetectionStrategy, Component, computed, input } from '@angular/core';
import { MatCardModule } from '@angular/material/card';

@Component({
  selector: 'app-card',
  imports: [MatCardModule, NgStyle],
  templateUrl: './card.html',
  styleUrl: './card.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Card {
  color = input('#ffffff');
  textColor = input('#0f172a');
  borderColor = input('#e2e8f0');
  width = input('100%');
  maxWidth = input('none');
  height = input('auto');
  maxHeight = input('none');

  protected readonly cardStyle = computed(() => ({
    '--tt-card-color': this.color(),
    '--tt-card-text-color': this.textColor(),
    '--tt-card-border-color': this.borderColor(),
    width: this.width(),
    maxWidth: this.maxWidth(),
    height: this.height(),
    maxHeight: this.maxHeight(),
  }));
}
