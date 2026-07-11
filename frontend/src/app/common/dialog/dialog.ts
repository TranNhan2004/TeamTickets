import { NgStyle } from '@angular/common';
import {
  ChangeDetectionStrategy,
  Component,
  HostListener,
  computed,
  input,
  linkedSignal,
  output,
} from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';

@Component({
  selector: 'app-dialog',
  imports: [MatButtonModule, MatDialogModule, MatIconModule, NgStyle],
  templateUrl: './dialog.html',
  styleUrl: './dialog.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Dialog {
  open = input(false);
  title = input('Dialog');
  ariaLabel = input<string | null>(null);
  color = input('#ffffff');
  textColor = input('#0f172a');
  borderColor = input('#e2e8f0');
  width = input('min(90vw, 560px)');
  maxWidth = input('560px');
  height = input('auto');
  maxHeight = input('90vh');
  closeOnBackdrop = input(true);
  closeOnEscape = input(true);

  onOpen = output<void>();
  onClose = output<void>();

  protected readonly isOpen = linkedSignal(() => this.open());

  protected readonly dialogStyle = computed(() => ({
    '--tt-dialog-color': this.color(),
    '--tt-dialog-text-color': this.textColor(),
    '--tt-dialog-border-color': this.borderColor(),
    width: this.width(),
    maxWidth: this.maxWidth(),
    height: this.height(),
    maxHeight: this.maxHeight(),
  }));

  openDialog(): void {
    if (this.isOpen()) return;

    this.isOpen.set(true);
    this.onOpen.emit();
  }

  closeDialog(): void {
    if (!this.isOpen()) return;

    this.isOpen.set(false);
    this.onClose.emit();
  }

  protected handleBackdropClick(event: MouseEvent): void {
    if (this.closeOnBackdrop() && event.target === event.currentTarget) {
      this.closeDialog();
    }
  }

  @HostListener('document:keydown.escape')
  protected handleEscape(): void {
    if (this.isOpen() && this.closeOnEscape()) {
      this.closeDialog();
    }
  }
}
