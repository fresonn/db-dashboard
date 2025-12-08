import { createFileRoute } from '@tanstack/react-router'
import { Overview } from '@/views/overview'

export const Route = createFileRoute('/overview')({
  component: Overview
})
