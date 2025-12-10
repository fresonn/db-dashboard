import { useEffect, useState } from 'react'

export function TimeTicker({ startedAtIso }: { startedAtIso: string }) {
  const [uptime, setUptime] = useState('0 sec')

  useEffect(() => {
    const startMs = Date.parse(startedAtIso)
    if (isNaN(startMs)) {
      setUptime('time parse error')
      return
    }

    const update = () => {
      const diff = Math.floor((Date.now() - startMs) / 1000)

      const days = Math.floor(diff / 86400)
      const hours = Math.floor((diff % 86400) / 3600)
      const minutes = Math.floor((diff % 3600) / 60)
      const seconds = diff % 60

      const parts: string[] = []

      if (days > 0) parts.push(`${days} days`)
      if (hours > 0) parts.push(`${hours}h`)
      if (minutes > 0 || hours > 0 || days > 0) {
        parts.push(`${minutes} min`)
      }
      parts.push(`${seconds} sec`)

      setUptime(parts.join(' '))
    }

    update()
    const id = setInterval(update, 1000)
    return () => clearInterval(id)
  }, [startedAtIso])

  return <span>{uptime}</span>
}
