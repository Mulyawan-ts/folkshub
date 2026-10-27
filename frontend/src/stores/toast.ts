import { defineStore } from 'pinia'

export type ToastType = 'success' | 'error' | 'warning' | 'info'

export interface ToastAction {
  label: string
}

export interface ToastOptions {
  title?: string
  duration?: number
  action?: ToastAction | null
}

export interface ToastItem {
  id: number
  type: ToastType
  title: string
  message: string
  duration: number
  action: ToastAction | null
}

const DEFAULT_DURATION = 4000

let counter = 0

export const useToastStore = defineStore('toast', {
  state: () => ({
    items: [] as ToastItem[]
  }),

  getters: {
    hasItems: (state) => state.items.length > 0
  },

  actions: {
    show(options: Partial<ToastItem> = {}): number {
      const {
        type = 'info',
        title = '',
        message = '',
        duration = DEFAULT_DURATION,
        action = null
      } = options

      const id = ++counter

      this.items.push({ id, type, title, message, duration, action })

      if (duration > 0) {
        setTimeout(() => this.dismiss(id), duration)
      }

      return id
    },

    success(message: string, options: ToastOptions = {}): number {
      return this.show({ ...options, type: 'success', message })
    },

    error(message: string, options: ToastOptions = {}): number {
      return this.show({ ...options, type: 'error', message })
    },

    warning(message: string, options: ToastOptions = {}): number {
      return this.show({ ...options, type: 'warning', message })
    },

    info(message: string, options: ToastOptions = {}): number {
      return this.show({ ...options, type: 'info', message })
    },

    dismiss(id: number): void {
      const index = this.items.findIndex((item) => item.id === id)
      if (index !== -1) this.items.splice(index, 1)
    },

    clear(): void {
      this.items.splice(0, this.items.length)
    }
  }
})
