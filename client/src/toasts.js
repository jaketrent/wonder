import nanoid from 'nanoid'
import { writable } from 'svelte/store'

export const toasts = writable([])

export function add({
  text = 'Success!',
  status = 'success',
  autodismiss = true
}) {
  const id = nanoid()
  const toast = { id, text, status }
  toasts.update(t => [...t, toast])

  if (autodismiss)
    setTimeout(_ => toasts.update(ts => ts.filter(t => t.id !== id)), 2000)
}

export function remove(id) {
  toasts.update(ts => ts.filter(t => t.id !== id))
}
