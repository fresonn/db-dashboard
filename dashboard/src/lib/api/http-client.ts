const API_BASE_URL = import.meta.env.VITE_API_URL

export type RequestConfig<TData = unknown> = {
  url?: string
  method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE'
  params?: Record<string, any>
  data?: TData | FormData | unknown
  headers?: Record<string, string>
  signal?: AbortSignal
}

export type ResponseConfig<TData = unknown> = {
  data: TData
  status: number
  statusText: string
  headers: Headers
}

export type ResponseErrorConfig<T = unknown> = T

const buildQueryParams = (params?: Record<string, any>): string => {
  if (!params || Object.keys(params).length === 0) return ''
  const searchParams = new URLSearchParams(
    Object.entries(params).flatMap(([k, v]) =>
      v == null ? [] : Array.isArray(v) ? v.map((val) => [k, String(val)]) : [[k, String(v)]]
    )
  )
  return `?${searchParams.toString()}`
}

const parseResponseBody = async (response: Response): Promise<unknown> => {
  const contentType = response.headers.get('content-type') ?? ''
  if (contentType.includes('application/json')) return response.json()
  if (contentType.includes('text/')) return response.text()
  return response.blob()
}

const client = async <TResponse = unknown, TData = unknown, TError = unknown>(
  config: RequestConfig<TData>
): Promise<ResponseConfig<TResponse>> => {
  const { method, url, params, data, headers = {}, signal } = config

  const fullUrl = `${API_BASE_URL}${url}${buildQueryParams(params)}`

  const fetchHeaders: Record<string, string> = {
    ...headers,
    ...(data instanceof FormData ? {} : { 'Content-Type': 'application/json' })
  }

  const response = await fetch(fullUrl, {
    method,
    headers: fetchHeaders,
    body: data instanceof FormData ? data : data != null ? JSON.stringify(data) : undefined,
    signal
  })

  const body = await parseResponseBody(response)

  if (!response.ok) {
    throw body as ResponseErrorConfig<TError>
  }

  return {
    data: body as TResponse,
    status: response.status,
    statusText: response.statusText,
    headers: response.headers
  }
}

export default client
