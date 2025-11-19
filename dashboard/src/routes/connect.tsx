import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/connect')({
  component: Connect
})

function Connect() {
  return (
    <div className="p-2">
      <h3>Connect to postgres</h3>
    </div>
  )
}
