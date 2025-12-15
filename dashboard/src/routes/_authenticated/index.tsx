import { createFileRoute } from '@tanstack/react-router'
import { ClusterOverview } from '@/views/overview'

export const Route = createFileRoute('/_authenticated/')({
  component: ClusterOverview
})
