import { createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  beforeLoad({ context }) {
    if (!context.clusterStatus) {
      throw redirect({
        to: '/connect'
      })
    }
    throw redirect({
      to: context.clusterStatus.postgres_connection === 'connected' ? '/overview' : '/connect'
    })
  }
})
