import { createFileRoute } from '@tanstack/react-router'
import { ThemeToggle } from '@/components/ui/theme-toggle'
import { useState } from 'react'
import { Button } from '@/components/ui/button'
import { AlarmClockCheck, Database } from 'lucide-react'
import { Input } from '@/components/ui/input'

export const Route = createFileRoute('/ui')({
  component: RouteComponent
})

function RouteComponent() {
  const [loading, setLoading] = useState(false)

  const hanleClick = () => {
    setLoading((prev) => !prev)
  }

  return (
    <div className="p-2">
      <h3>Connect to postgres</h3>

      <div>
        <ThemeToggle />
      </div>

      <Button>Click me!</Button>
      <Button variant="destructive">Click me!</Button>
      <Button variant="outline">Click me!</Button>
      <Button variant="secondary">Click me!</Button>
      <Button variant="ghost">Click me!</Button>
      <Button variant="link">Click me!</Button>
      <div className="mt-10">
        <Button loading={loading} onClick={hanleClick}>
          <Database />
          Loading
        </Button>
        <Button loading={loading} onClick={hanleClick} variant="secondary">
          <AlarmClockCheck />
          Loading
        </Button>

        <div className="max-w-42">
          <Input placeholder="Port" type="number" />
          <Input disabled placeholder="Port" type="number" />
        </div>
      </div>
    </div>
  )
}
