import { defineConfig } from '@kubb/core'
import { pluginTs } from '@kubb/plugin-ts'
import { pluginOas } from '@kubb/plugin-oas'
import { pluginClient } from '@kubb/plugin-client'
import { pluginReactQuery } from '@kubb/plugin-react-query'

export default defineConfig({
  name: 'dashboard',
  root: '.',
  input: {
    path: '../openapi/api.yaml'
  },
  output: {
    path: './src/lib/api/gen',
    clean: true,
    format: 'prettier'
  },
  plugins: [
    pluginOas(),
    pluginTs({
      output: { path: 'models.ts' }
    }),
    pluginClient({
      importPath: '@/lib/api/http-client.ts'
    }),
    pluginReactQuery({
      client: {
        importPath: '@/lib/api/http-client.ts'
      }
    })
  ]
})
