import { createFileRoute } from '@tanstack/react-router'
import { ConnectView } from '@/views/connect'

export const Route = createFileRoute('/connect')({
  component: ConnectView
})
