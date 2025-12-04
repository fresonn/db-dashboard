import * as v from 'valibot'

export const validationSchema = v.object({
  database: v.optional(v.string()),
  host: v.pipe(v.string(), v.trim(), v.nonEmpty('required')),
  port: v.pipe(
    v.string(),
    v.trim(),
    v.regex(/^\d+$/),
    v.transform(Number),
    v.integer(),
    v.gtValue(0, 'The port must be greater than zero.')
  ),
  user: v.pipe(v.string(), v.trim(), v.nonEmpty('required')),
  password: v.pipe(v.string(), v.trim(), v.nonEmpty('required')),
  sslmode: v.optional(v.string(), '')
})

export type FormSubmitData = v.InferOutput<typeof validationSchema>
