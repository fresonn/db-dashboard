import { TriangleAlert } from 'lucide-react'

export function OfflineApiGuard({ open }: { open: boolean }) {
  if (!open) return

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 backdrop-blur-sm">
      <div className="w-full max-w-md rounded-2xl bg-white p-8 shadow-2xl">
        <div className="mb-6 text-center">
          <div className="mx-auto mb-4 flex size-16 items-center justify-center rounded-full bg-red-100 text-red-600">
            <TriangleAlert className="size-10" />
          </div>
          <h2 className="text-2xl font-bold text-gray-900">No connection to the api server</h2>
          <p className="mt-3 text-gray-600">
            Failed to connect to the backend. Check your internet connection and server status.
          </p>
        </div>
        <button
          onClick={() => window.location.reload()}
          className="w-full rounded-lg bg-red-600 px-6 py-3 font-medium text-white transition hover:bg-red-700"
        >
          Try again
        </button>
      </div>
    </div>
  )
}
